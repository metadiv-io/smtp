package smtp

// NewDialer creates a new SMTP dialer.
func NewDialer(displayName, host string, port int, username, password string) *Dialer {
	return &Dialer{
		DisplayName: displayName,
		Host:        host,
		Port:        port,
		Username:    username,
		Password:    password,
	}
}

// NewDefaultDialer creates a new SMTP dialer with the default environment variables.
func NewDefaultDialer() *Dialer {
	return NewDialer(
		EnvDisplayName.String(),
		EnvHost.String(),
		EnvPort.Int(),
		EnvUsername.String(),
		EnvPassword.String(),
	)
}

// Dialer is a SMTP dialer.
type Dialer struct {
	DisplayName string `json:"display_name"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
