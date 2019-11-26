package utils

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	env "github.com/segmentio/go-env"
)

// ResolveEnv check the env with the provived name and if doesn't exist
// it will create it and set value to the provided one
func ResolveEnv(envName, def string) string {
	curr := os.Getenv(envName)

	if curr == "" {
		os.Setenv(envName, def)
	}

	return os.Getenv(envName)
}

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func EnvVariable(exportedName string) (variable string, err error) {
	variable, err = env.Get(exportedName)
	return
}

func DoesFileExist(filename string) (fileok bool) {
	fileok = false

	if _, err := os.Stat(filename); err == nil {
		fileok = true
		return
		// fmt.Println("File" + filename + " exists!")
	} else if os.IsNotExist(err) {
		// fmt.Println("File" + filename + " doesn't exist!")
	} else {
		// fmt.Println("File" + filename + " probably doesn't exist!")
		// file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	return
}

type File struct {
	Filepath string
	Name     string
	Time     time.Time
	Size     int64
	// date
}

// FilesFromDirectory returns a list of files in passed directory.
// If mustCompile variable passed, only the files that pass the regexp will be returned
func FilesFromDirectory(directory string, mustCompile string) (result []*File, err error) {

	if _, err = os.Stat(directory); os.IsNotExist(err) {
		return nil, err
	}

	filepath.Walk(directory+"/", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == true {
			return nil
		}

		if mustCompile != "" {
			r := regexp.MustCompile(mustCompile)

			if r.MatchString(path) {
				result = append(result, &File{strings.Replace(path, directory+"/", "", -1), f.Name(), f.ModTime(), f.Size()})
			}
		} else {
			result = append(result, &File{strings.Replace(path, directory+"/", "", -1), f.Name(), f.ModTime(), f.Size()})
		}

		return nil
	})

	return
}

// Untar function unpacking an archive to destination directory
// Returns nil if everything goes just fine, otherwise error is sent
func Untar(path string, destinationpath string) error {
	if destinationpath != "" {
		os.RemoveAll(destinationpath)
	}
	os.Mkdir(destinationpath, 0700)

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

func StructToMap(entry interface{}) (retval map[string]interface{}) {
	s := reflect.ValueOf(entry)

	typeOfT := s.Type()
	if typeOfT.Kind() == reflect.Ptr {
		typeOfT = typeOfT.Elem()
	}

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		// uncomment below to see what's really happening
		// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		retval[typeOfT.Field(i).Name] = f.Interface()
	}

	return
}

// BodyToString returns the pure json-string data to screen
func BodyToString(body io.Reader) (string, error) {

	data, err := ioutil.ReadAll(body)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// PrettyJson returns pretty converted json
func PrettyJson(i interface{}) (string, error) {

	// printing pretty
	b, err := json.MarshalIndent(i, "", "  ")

	if err != nil {
		return "", err
	}

	// os.Stdout.Write(b)
	// fmt.Printf("%s", b)

	return string(b), nil
}

func SliceContains(slice []string, needle string) bool {

	for _, a := range slice {

		if a == needle {

			return true
		}
	}

	return false
}

// SlicesDiff comparef two slices together and returns two slices of the difference
// The first slice contains what's been removed, opposite to the other slice
// Note: First passed slice should be the old/previous/base one, second the new/current one
func SlicesDiff(oldSlice, newSlice []string) (removed, added []string) {
	for _, o := range oldSlice {
		found := false

		for i, n := range newSlice {
			if n == o {
				found = true
				newSlice = append(newSlice[:i], newSlice[i+1:]...)
				break
			}
		}

		if !found {
			removed = append(removed, o)
		}
	}

	added = newSlice

	return
}

// ReadWholeFile opens and reads whole file and return a body of this file
// Note that this function should not be used for big files!
func ReadWholeFile(path string) (body string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("No mapping file found. Skipping: %v\n", err)
		return
	}

	body = string(file)

	return
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
