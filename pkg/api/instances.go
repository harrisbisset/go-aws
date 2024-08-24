package api

import (
	"bytes"
	"context"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/charmbracelet/keygen"
	"github.com/harrisbisset/go-aws/pkg/models"
	"golang.org/x/crypto/ssh"
)

func GetInstanceList() ([]types.Instance, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return []types.Instance{}, err
	}

	client := ec2.NewFromConfig(cfg)
	out, err := client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		// Filters: []types.Filter{
		// {Name: aws.String("architecture")},
		// {Name: aws.String("instance-id")},
		// {Name: aws.String("instance-type")},
		// {Name: aws.String("monitoring-state")},
		// {Name: aws.String("dns-name")},
		// {Name: aws.String("ip-address")},
		// {Name: aws.String("monitoring-state")},
		// {Name: aws.String("instance.group-name")},
		// {Name: aws.String("iam-instance-profile.arn")},
		// {Name: aws.String("iam-instance-profile.name")},
		// },
		MaxResults: aws.Int32(10),
	})

	return out.Reservations[0].Instances, err
}

func CheckInstanceConnection(id *string, zone *string) (models.ConnectionReport, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return models.ConnectionReport{}, err
	}

	//generate key
	kp, err := keygen.New("my_temp_key", keygen.WithKeyType(keygen.Ed25519))
	if err != nil {
		return models.ConnectionReport{}, err
	}

	//create 60s connection
	client := ec2instanceconnect.NewFromConfig(cfg)
	_, err = client.SendSSHPublicKey(
		context.Background(),
		&ec2instanceconnect.SendSSHPublicKeyInput{
			InstanceId:       id,
			InstanceOSUser:   aws.String("ubuntu"),
			SSHPublicKey:     aws.String(kp.AuthorizedKey()),
			AvailabilityZone: zone,
		},
	)
	if err != nil {
		return models.ConnectionReport{}, err
	}

	return models.ConnectionReport{}, nil
}

// e.g. output, err := remoteRun("root", "MY_IP", "PRIVATE_KEY", "ls")
func remoteRun(user string, addr string, privateKey string, cmd string) (string, error) {
	// privateKey could be read from a file, or retrieved from another storage
	// source, such as the Secret Service / GNOME Keyring
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		// https://github.com/golang/go/issues/19767
		// as clientConfig is non-permissive by default
		// you can set ssh.InsercureIgnoreHostKey to allow any host
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		//alternatively, you could use a password
		/*
		   Auth: []ssh.AuthMethod{
		       ssh.Password("PASSWORD"),
		   },
		*/
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// you can also pass what gets input to the stdin, allowing you to pipe
	// content from client to server
	//      session.Stdin = bytes.NewBufferString("My input")

	// Finally, run the command
	err = session.Run(cmd)
	return b.String(), err
}

func CheckDocker() {

}
