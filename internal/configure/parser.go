package configure

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	AWSRegion  string
	SSHUser    string
	InstanceID string
	Debug      bool
}

func ParseArgs() Config {
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	if defaultRegion == "" {
		defaultRegion = "us-east-1"
	}

	awsRegion := flag.String("region", defaultRegion, "AWS region")
	sshUser := flag.String("user", "centos", "SSH user")
	instanceID := flag.String("instance", "", "EC2 Instance ID")
	debug := flag.Bool("debug", false, "Turn on debug messages")

	flag.Parse()
	if *instanceID == "" {
		fmt.Println("Instance ID is required")
		os.Exit(1)
	}

	return Config{
		AWSRegion:  *awsRegion,
		SSHUser:    *sshUser,
		InstanceID: *instanceID,
		Debug:      *debug,
	}
}
