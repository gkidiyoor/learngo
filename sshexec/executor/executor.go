package executor

import (
	"fmt"
	"io"
	"os"
	"bytes"
	"golang.org/x/crypto/ssh"
)

var modes ssh.TerminalModes = ssh.TerminalModes{
	ssh.ECHO:          0,     // disable echoing
	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
}

func Execute(config *ssh.ClientConfig, server string, env map[string]string, commands []string) error {
	client, err := connect(config, server)
	if err != nil {
		fmt.Println("Unable to execute commands on : ", server)
		return err
	}

	for _, command := range commands {
		err, output := executeCommand(client, env, command)
		fmt.Println("output : " + output)
		if err != nil {
			fmt.Println("Unable to execute commands on : ", server)
			return err
		}
	}

	return nil
}

func executeCommand(client *ssh.Client, env map[string]string, command string) (error, string) {
	//session, err := createSession(client)
	session, err := client.NewSession()
	if err != nil {
		return err, ""
	}
	defer session.Close()

	//fmt.Println("Setting up session IO")
	//err = configureSessionIO(session)
	//if err != nil {
//		fmt.Println("Unable to set up session IO")
//		return err
//	}
//	fmt.Println("Done setting up session IO")

	if len(env) > 0 {
		err = configureSessionEnv(session, env)
		if err != nil {
			return err, ""
		}
	}

	fmt.Println("Begin executing :", command)


	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(command)

	if err != nil {
		fmt.Println("Error executing command: ", command)
		fmt.Println("Error : ", err.Error())
		return err, ""
	}

	fmt.Println("End executing :", command)
	fmt.Println(b.String())
	return nil, b.String()
}

func connect(config *ssh.ClientConfig, server string) (*ssh.Client, error) {
	fmt.Println("Attempting to connect to : ", server)
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		fmt.Println("Failed to dial SSH connection to : ", server)
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	fmt.Println("Connected to : ", server)
	return client, nil
}

func createSession(client *ssh.Client) (*ssh.Session, error) {
	fmt.Println("Creating new session")
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session")
		fmt.Println("Error : ", err.Error())
		return nil, err
	}

	return session, nil
}

func configureSessionEnv(session *ssh.Session, env map[string]string) error {
	fmt.Println("Setting environment variables")

	for key, value := range env {
		if err := session.Setenv(key, value); err != nil {
			fmt.Printf("Unable to set environment %s = %s\n", key, value)
			return err
		}
	}

	fmt.Println("Done setting environment variables")
	return nil
}

func configureSessionIO(session *ssh.Session) error {
	fmt.Println("Requesting Pseudo Terminal")
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Println("Unable to request Pseudo Terminal")
		fmt.Println("Error : ", err.Error())
		return err
	}

	fmt.Println("Setting up STDIN")
	stdin, err := session.StdinPipe()
	if err != nil {
		fmt.Println("Unable to setup STDIN")
		fmt.Println("Error : ", err.Error())
		return err
	}
	go io.Copy(stdin, os.Stdin)

	fmt.Println("Setting up STDOUT")
	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Println("Unable to setup STDOUT")
		fmt.Println("Error : ", err.Error())	
		return err
	}
	go io.Copy(os.Stdout, stdout)

	fmt.Println("Setting up STDERR")
	stderr, err := session.StderrPipe()
	if err != nil {
		fmt.Println("Unable to setup STDERR")
		fmt.Println("Error : ", err.Error())
		return err
	}
	go io.Copy(os.Stderr, stderr)

	return nil
}