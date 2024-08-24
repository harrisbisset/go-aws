package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/harrisbisset/go-aws/pkg/models"
)

func GetInstanceList() ([]types.Instance, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
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

func CheckInstanceConnection(id string) (models.ConnectionReport, error) {
	return models.ConnectionReport{}, nil
}

func CheckDocker() {

}
