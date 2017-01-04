package cli

type Config struct {
	Port                   int
	Script                 string
	ScriptsDir             string
	Hostname               string
	Path                   string
	Tls                    bool
	PublicKey              string
	PrivateKey             string
	CertificationAuthority string
}
