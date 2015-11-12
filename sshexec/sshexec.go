package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abhigupta912/learngo/sshexec/config"
	"github.com/abhigupta912/learngo/sshexec/executor"
)

func main() {
	configFile := flag.String("cfg", "cfg.yml", "Config File")
	flag.Parse()

	fmt.Println("Reading config file : ", *configFile)
	execConfig, err := config.ReadConfig(*configFile)
	if err != nil {
		fmt.Println("Unable to proceed with execution")
		os.Exit(1)
	}

	clientConfig, err := config.ConfigureAuth(execConfig)
	if err != nil {
		fmt.Println("Unable to proceed with execution")
		os.Exit(1)
	}

	serverFile := execConfig.HostsFile
	port := execConfig.Port
	envFile := execConfig.EnvFile
	commandsFile := execConfig.CommandsFile

	servers, err := config.GetServers(serverFile, port)
	if err != nil {
		fmt.Println("Unable to proceed with execution")
		os.Exit(1)
	}

	var env map[string]string
	if envFile != "" {
		env, err = config.GetEnv(envFile)
		if err != nil {
			fmt.Println("Unable to proceed with execution")
			os.Exit(1)
		}
	}

	commands, err := config.ReadFileLines(commandsFile)
	if err != nil {
		fmt.Println("Unable to proceed with execution")
		os.Exit(1)
	}

	for _, server := range servers {
		executor.Execute(clientConfig, server, env, commands)
	}
}
