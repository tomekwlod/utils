package sftp

import (
	"fmt"
	"io"
	"log"
	"os"

	_sftp "github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTP struct {
	*_sftp.Client
}
type ClientConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewClient(c *ClientConfig) (*SFTP, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            c.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(c.Password)},
	}

	conn, err := ssh.Dial("tcp", c.Host+":"+c.Port, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	fmt.Println("Successfully connected to ssh server.")

	// open an SFTP session over an existing ssh connection.
	client, err := _sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}

	srv := &SFTP{client}

	return srv, nil
}

// Example
// http://networkbit.ch/golang-sftp-client/

func (s *SFTP) SendFile(localpath, localfilename, remotepath, remotefilename string) (int64, error) {

	// Create the destination file
	dstFile, err := s.Create(remotepath + remotefilename)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	// Open the source file
	srcFile, err := os.Open(localpath + localfilename)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, err
	}

	return bytes, nil
}

func (s *SFTP) GetFile(remotepath, remotefilename, localpath, localfilename string) (int64, error) {

	// Create the destination file
	dstFile, err := os.Create(localpath + localfilename)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	// Open the source file
	srcFile, err := s.Open(remotepath + remotefilename)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, err
	}

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}
