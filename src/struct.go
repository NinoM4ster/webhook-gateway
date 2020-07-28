package main

// Bots .
type Bots struct {
	Bot []Bot `json:"bots"`
}

// Bot .
type Bot struct {
	URI  string `json:"uri"`
	Host string `json:"host"`
	// Cert string `json:"cert_file"`
	// Key  string `json:"key_file"`
}

// Config .
type Config struct {
	ListenPort string `json:"listen_port"`
	CertFile   string `json:"cert_file"`
	KeyFile    string `json:"key_file"`
}
