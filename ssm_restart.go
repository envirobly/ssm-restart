package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

const (
	PORT           = ":9009" // Port to listen on
	SECRET_MESSAGE = "restart_ssm_agent"
	SERVICE_NAME   = "amazon-ssm-agent"
)

func restartService() {
	cmd := exec.Command("systemctl", "restart", "--no-block", SERVICE_NAME)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to restart service %s: %v\n", SERVICE_NAME, err)
	} else {
		fmt.Printf("Service %s restart scheduled successfully.\n", SERVICE_NAME)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading message: %v\n", err)
		return
	}
	message = strings.TrimSpace(message)
	if message == SECRET_MESSAGE {
		fmt.Println("Received correct message, restarting service...")
		restartService()
	} else {
		fmt.Printf("Received incorrect message: %s\n", message)
	}
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}
