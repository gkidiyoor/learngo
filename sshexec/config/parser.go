package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	yaml "gopkg.in/yaml.v2"
)

func ReadConfig(file string) (*Config, error) {
	filename, _ := filepath.Abs(file)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read config file: ", filename)
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Unable to unmarshall config file: ", filename)
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	return &config, nil
}

func getCurrentUser() (*user.User, error) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get current user info")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	return usr, nil
}

func getSigner(file string) (ssh.Signer, error) {
	filename, _ := filepath.Abs(file)
	keyFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read private key file: ", filename)
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(keyFile)
	if err != nil {
		fmt.Println("Unable to parse private key file")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	return key, nil
}

func ConfigureAuth(config *Config) (*ssh.ClientConfig, error) {
	authType := config.AuthType

	username := strings.TrimSpace(config.Username)
	if len(username) == 0 {
		user, err := getCurrentUser()
		if err != nil {
			fmt.Println("Unable to get username")
			return nil, err
		}

		username = user.Username
	}

	var clientConfig *ssh.ClientConfig
	var err error

	switch {
	case strings.EqualFold(authType, AuthTypePassword):
		password := strings.TrimSpace(config.Password)
		clientConfig, err = createPasswordAuth(username, password)

	case strings.EqualFold(authType, AuthTypePublicKey):
		clientConfig, err = createPrivateKeyAuth(username, config.PrivateKeyFile)

	case strings.EqualFold(authType, AuthTypeSSHAgent):
		clientConfig, err = createSSHAgentAuth(username)
	}

	return clientConfig, err
}

func createPasswordAuth(username string, password string) (*ssh.ClientConfig, error) {
	if len(password) == 0 {
		err := errors.New("No password specified")
		fmt.Println("Unable to configure authorization")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	return clientConfig, nil
}

func createPrivateKeyAuth(username string, keyFile string) (*ssh.ClientConfig, error) {
	key, err := getSigner(keyFile)
	if err != nil {
		fmt.Println("Unable to configure authorization")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	return clientConfig, nil
}

func createSSHAgentAuth(username string) (*ssh.ClientConfig, error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		fmt.Println("Unable to configure authorization")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers),
		},
	}

	return clientConfig, nil
}
