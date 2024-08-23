package api_test

import (
	"fmt"
	"testing"

	"github.com/harrisbisset/go-aws/pkg/api"
)

func TestGetInstanceList(t *testing.T) {
	list, err := api.GetInstanceList()
	if err != nil {
		panic(err)
	}

	for _, k := range list {
		fmt.Println(*k.InstanceId)
		fmt.Println(k.Architecture)
		fmt.Println(k.InstanceType)
		fmt.Println(*k.PublicDnsName)
		fmt.Println(k.State.Name)
		fmt.Println(*k.PublicIpAddress)
		fmt.Println(k.Monitoring.State)
		for _, j := range k.SecurityGroups {
			fmt.Println(*j.GroupName)
		}
		fmt.Println(*k.SubnetId)
		fmt.Println(*k.IamInstanceProfile.Arn)
	}
}
