package cli

import "flag"

func DeclareClientFlags(config *Config, clientCommand *flag.FlagSet) {
	clientCommand.BoolVar(&config.Tls, "tls", false, "Enable TLS encryption")
	clientCommand.StringVar(&config.PublicKey, "pubkey", "", "Client public key path")
	clientCommand.StringVar(&config.PrivateKey, "key", "", "Client private key path")
	clientCommand.StringVar(&config.CertificationAuthority, "ca", "", "CA authority")
	clientCommand.IntVar(&config.Port, "dport", 5058, "Port of the server in the remote container")
	clientCommand.StringVar(&config.Script, "script", "", "Script to execute")
	clientCommand.StringVar(&config.Hostname, "host", "localhost", "Hostname of the remote container")
}

func DeclareServerFlags(config *Config, serverCommand *flag.FlagSet) {
	serverCommand.BoolVar(&config.Tls, "tls", false, "Enable TLS encryption")
	serverCommand.StringVar(&config.PublicKey, "pubkey", "", "Client public key path")
	serverCommand.StringVar(&config.PrivateKey, "key", "", "Client private key path")
	serverCommand.StringVar(&config.CertificationAuthority, "ca", "", "CA authority")
	serverCommand.IntVar(&config.Port, "sport", 5058, "Port to listen")
	serverCommand.StringVar(&config.ScriptsDir, "dir", "", "Directory to look for the executable scripts")
}
