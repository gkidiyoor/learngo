package config

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFileLines(file string) ([]string, error) {
	filename, _ := filepath.Abs(file)
	inFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file: ", filename)
		fmt.Println("Error : ", err.Error())
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents, nil
}

func GetServers(file string, port string) ([]string, error) {
	fmt.Println("Reading servers file")
	contents, err := ReadFileLines(file)
	if err != nil {
		fmt.Println("Unable to read servers file")
		return nil, err
	}

	servers := make([]string, 0, len(contents))

	for _, server := range contents {
		if strings.Contains(server, ":") {
			servers = append(servers, strings.TrimSpace(server))
		} else {
			if port == "" {
				fmt.Println("No port specified for server: ", server)
				continue
			}

			serverBytes := bytes.NewBufferString(strings.TrimSpace(server))
			serverBytes.WriteString(":")
			serverBytes.WriteString(port)
			servers = append(servers, serverBytes.String())
		}
	}

	fmt.Println("Done reading servers file")
	return servers, nil
}

func GetEnv(file string) (map[string]string, error) {
	fmt.Println("Reading env file")
	contents, err := ReadFileLines(file)
	if err != nil {
		fmt.Println("Unable to read env file")
		return nil, err
	}

	env := make(map[string]string, len(contents))

	for _, envline := range contents {
		kvpair := strings.Split(envline, "=")
		if len(kvpair) != 2 {
			log.Println("Ignoring invalid env line : ", envline)
			continue
		}

		key := strings.TrimSpace(kvpair[0])
		value := strings.TrimSpace(kvpair[1])
		env[key] = value
	}

	fmt.Println("Done reading env file")
	return env, nil
}
