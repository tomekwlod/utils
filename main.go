package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// FilesFromDirectory returns a list of files in passed directory.
// If mustCompile variable passed, only the files that pass the regexp will be returned
func FilesFromDirectory(directory string, mustCompile string) []string {
	files := []string{}

	filepath.Walk(directory+"/", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == true {
			return nil
		}

		if mustCompile != "" {
			r := regexp.MustCompile(mustCompile)
			if r.MatchString(path) {
				files = append(files, strings.Replace(path, directory+"/", "", -1))
			}
		} else {
			files = append(files, strings.Replace(path, directory+"/", "", -1))
		}

		return nil
	})

	return files
}

// Untar function unpacking an archive to destination directory
// Returns nil if everything goes just fine, otherwise error is sent
func Untar(path string, destinationpath string) error {
	if destinationpath != "" {
		os.RemoveAll(destinationpath)
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tr := tar.NewReader(r)

	for {
		header, err := tr.Next()

		// if no more files are found within an archive, kill the loop
		if err == io.EOF {
			break
		}

		// return any other error
		if err != nil {
			return err
		}

		// if the header is nil, just skip it (not sure how this can happen)
		if header == nil {
			continue
		}

		target := filepath.Join(destinationpath, header.Name)

		// check the file type
		switch header.Typeflag {

		case tar.TypeDir: // if directory
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}

				// fmt.Println("Directory:", target)
			}
		case tar.TypeReg: // if regular file
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer f.Close()

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// fmt.Println("Regular file:", target)
		}
	}

	return nil
}

// Below example shows how to execute many goroutines and wait until they are all done
// We don't have to use the channels here because we have no values to pass from the goroutines
func wgExample() {
	// this is just to later check if all the goroutines are done
	var wg sync.WaitGroup
	// adding as many waitgroups as necessary

	wg.Add(1)
	go function1(&wg, "variable")

	wg.Add(1)
	go function1(&wg, "variable")

	fmt.Println("Waiting the groups to be all finished")
	wg.Wait()
	fmt.Println("All done now!")
}
func function1(wg *sync.WaitGroup, variable string) {
	defer wg.Done()
	// function body
}

func channelsExample() {
	// initializing the channel
	ch := make(chan string)

	go function2(ch, "variable")
	go function2(ch, "variable")

	x, y := <-ch, <-ch

	fmt.Println("All done", x, y)
}
func function2(c chan string, variable string) {
	// function body
	c <- "value to be returned"
}
