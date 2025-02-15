package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"go-ec2-instance-connect/internal/aws"
	"go-ec2-instance-connect/internal/configure"
	"go-ec2-instance-connect/internal/ssh"
)

func main() {
	cfg := configure.ParseArgs()

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.AWSRegion))
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	availabilityZone, publicIP, err := aws.GetInstanceDetails(awsCfg, cfg.InstanceID)
	if err != nil {
		log.Fatalf("Failed to get instance details: %v", err)
	}

	publicKey, err := ssh.ReadPublicKey()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = aws.SendPublicKey(awsCfg, cfg.InstanceID, availabilityZone, cfg.SSHUser, publicKey)
	if err != nil {
		log.Fatalf("Failed to send SSH public key: %v", err)
	}

	fmt.Printf("Public key sent successfully, connecting to %s@%s \n", cfg.SSHUser, publicIP)
	ssh.ConnectSSH(cfg.SSHUser, publicIP)
}
