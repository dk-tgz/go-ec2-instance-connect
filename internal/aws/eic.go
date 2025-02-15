package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
)

func SendPublicKey(cfg aws.Config, instanceID, availabilityZone, sshUser, publicKey string) error {
	eicClient := ec2instanceconnect.NewFromConfig(cfg)

	_, err := eicClient.SendSSHPublicKey(context.TODO(), &ec2instanceconnect.SendSSHPublicKeyInput{
		InstanceId:       aws.String(instanceID),
		InstanceOSUser:   aws.String(sshUser),
		SSHPublicKey:     aws.String(publicKey),
		AvailabilityZone: aws.String(availabilityZone),
	})
	return err
}
