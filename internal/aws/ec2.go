package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func GetInstanceDetails(cfg aws.Config, instanceID string) (string, string, error) {
	ec2Client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}
	result, err := ec2Client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return "", "", err
	}

	if len(result.Reservations) == 0 || len(result.Reservations[0].Instances) == 0 {
		return "", "", fmt.Errorf("instance not found")
	}

	instance := result.Reservations[0].Instances[0]
	return *instance.Placement.AvailabilityZone, *instance.PublicIpAddress, nil
}
