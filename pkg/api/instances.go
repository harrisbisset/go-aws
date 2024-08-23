package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func GetInstanceList() (ec2.DescribeInstancesOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := ec2.NewFromConfig(cfg)
	out, err := client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		// Filters:    []types.Filter{},
		MaxResults: aws.Int32(10),
	})

	return *out, err
}
