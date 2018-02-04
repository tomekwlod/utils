package ftp

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	f "github.com/jlaffaye/ftp"
)

var (
	// c exposed ftp connection
	c *f.ServerConn
)

type Client struct {
	Addr string
	Port int
	Auth Auth
}
type Auth struct {
	Username string
	Password string
}

type SearchInput struct {
	Path   string
	Suffix string
}

func (ftp Client) newConn() (c *f.ServerConn, err error) {
	// setting up the ftp client
	c, err = f.Dial(ftp.Addr + ":" + strconv.Itoa(ftp.Port))
	if err != nil {
		return
	}
	// defer c.Quit()

	// auth
	if ftp.Auth.Username != "" {
		err = c.Login(ftp.Auth.Username, ftp.Auth.Password)
	}

	return
}

// https://github.com/jlaffaye/ftp/blob/master/client_test.go <- simple tests
func (ftpUtils Client) FTPFilesList(in *SearchInput) (newEntries []*f.Entry, err error) {
	c, err := ftpUtils.newConn()
	if err != nil {
		return
	}
	defer c.Quit()

	// listing the files
	entries, _ := c.List(in.Path)

	// finding the only files that are important
	for _, entry := range entries {
		if in.Suffix != "" && strings.HasSuffix(entry.Name, in.Suffix) {
			newEntries = append(newEntries, entry)
		} else if in.Suffix == "" {
			newEntries = append(newEntries, entry)
		}
	}

	return
}

func (ftpUtils Client) FTPDownload(filename, targetFilename string) (err error) {
	c, err := ftpUtils.newConn()
	if err != nil {
		return
	}
	defer c.Quit()

	r, err := c.Retr(filename)
	if err != nil {
		return
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	err = c.NoOp()
	if err == nil {
		err = (errors.New("An unexpected Error"))
		return
	}

	// err = r.Close()
	// if err == nil {
	// 	err = (errors.New("Unexpected Error while closing the transfer"))
	// }

	err = ioutil.WriteFile(targetFilename, buf, 0644)

	return
}

func (ftpUtils Client) Rename(path, target string) (err error) {
	// setting up the ftp client
	c, err := ftpUtils.newConn()
	if err != nil {
		return
	}
	defer c.Quit()

	err = c.Rename(path, target)

	return
}
