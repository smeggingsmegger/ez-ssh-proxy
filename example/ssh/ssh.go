package main

import (
	"fmt"
	"time"

	"github.com/smeggingsmegger/ez-ssh-proxy"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "smeggingsmegger",
		Server: "example.com",
		// Optional key or Password without either we try to contact your agent SOCKET
		//Password: "password",
		// Paste your source content of private key
		// Key: `-----BEGIN RSA PRIVATE KEY-----
		// MIIEpAIBAAKCAQEA4e2D/qPN08pzTac+a8ZmlP1ziJOXk45CynMPtva0rtK/RB26
		// 7XC9wlRna4b3Ln8ew3q1ZcBjXwD4ppbTlmwAfQIaZTGJUgQbdsO9YA==
		// -----END RSA PRIVATE KEY-----
		// `,
		KeyPath: "/Users/username/.ssh/id_rsa",
		Port:    "22",
		Timeout: 60 * time.Second,
		// Parse PrivateKey With Passphrase
		Passphrase: "1234",
	}

	// Call Run method with command you want to run on remote server.
	stdout, stderr, done, err := ssh.Run("ls -al", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("don is :", done, "stdout is :", stdout, ";   stderr is :", stderr)
	}

}
