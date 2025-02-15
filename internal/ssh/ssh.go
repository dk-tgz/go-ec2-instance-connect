package ssh

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ReadPublicKey() (string, error) {
	pubKeyPath := os.ExpandEnv("$HOME/.ssh/id_rsa.pub")
	publicKeyBytes, err := os.ReadFile(pubKeyPath)
	if err != nil {
		return "", fmt.Errorf("Failed to read SSH public key: %v", err)
	}
	return strings.TrimSpace(string(publicKeyBytes)), nil
}

func ConnectSSH(user, ip string) {
	privateKeyPath := os.ExpandEnv("$HOME/.ssh/id_rsa")
	sshCmd := exec.Command("ssh", "-i", privateKeyPath, fmt.Sprintf("%s@%s", user, ip))
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	sshCmd.Stdin = os.Stdin

	if err := sshCmd.Run(); err != nil {
		log.Fatalf("Failed to connect via SSH: %v", err)
	}
}
