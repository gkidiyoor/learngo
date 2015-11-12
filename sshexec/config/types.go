package config

const (
	AuthTypePassword  = "password"
	AuthTypePublicKey = "publickey"
	AuthTypeSSHAgent  = "sshagent"
)

type Config struct {
	AuthType       string `yaml:"authtype"`
	Username       string `yaml:"username,omitempty"`
	Password       string `yaml:"password,omitempty"`
	PrivateKeyFile string `yaml:"privatekeyfile,omitempty"`
	HostsFile      string `yaml:"hostsfile"`
	Port           string `yaml:"port,omitempty"`
	EnvFile        string `yaml:"envfile,omitempty"`
	CommandsFile   string `yaml:"commandsfile"`
}
