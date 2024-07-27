package main

import (
	"flag"
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

const (
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
	listenAddress := flag.String("listen-address", ":63104", "Address and port to listen on (e.g., :63104 or 0.0.0.0:63104)")
	flag.Parse()

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", *listenAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}
