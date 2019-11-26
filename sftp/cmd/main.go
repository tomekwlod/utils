package main

import (
	"fmt"

	"github.com/tomekwlod/utils/sftp"
)

func main() {

	config := &sftp.ClientConfig{
		Username: "backup",
		Password: "CfyQpWp8AW5fvpbEXmAF",
		Host:     "178.62.101.80",
		Port:     "21",
	}

	client, err := sftp.NewClient(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// like:  cp ./file1.txt sftp@/files/file1.txt
	bytesSent, err := client.SendFile("./", "file1.txt", "/files/", "file1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes copied\n", bytesSent)

	// like:  cp sftp@/files/file1.txt ./file1.txt
	bytesReceived, err := client.GetFile("/files/", "file1.txt", "./", "file1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes copied\n", bytesReceived)

}
