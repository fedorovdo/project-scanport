package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	target := "127.0.0.1" // Replace with the IP address or hostname of the target computer
	startPort := 1        // Start port number
	endPort := 1024       // End port number

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting host name:", err)
	} else {
		fmt.Printf("Host Name: %s\n", hostName)
	}

	ipAddress, err := getLocalIP()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
	} else {
		fmt.Printf("IP Address: %s\n", ipAddress)
	}

	fmt.Printf("Scanning ports on %s...\n", target)

	openPorts := scanPorts(target, startPort, endPort)

	if len(openPorts) == 0 {
		fmt.Println("No open ports found.")
	} else {
		fmt.Println("Open ports:")
		for _, port := range openPorts {
			fmt.Printf("Port %d is open\n", port)
		}
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}

func scanPorts(target string, startPort, endPort int) []int {
	openPorts := []int{}

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			continue // Port is closed
		}
		conn.Close()
		openPorts = append(openPorts, port)
	}

	return openPorts
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("No IPv4 address found")
}
