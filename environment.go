package smtp

import (
	"github.com/metadiv-io/env"
)

var (
	EnvHost        = env.New("SMTP_HOST", false)
	EnvPort        = env.New("SMTP_PORT", false)
	EnvUsername    = env.New("SMTP_USERNAME", false)
	EnvPassword    = env.New("SMTP_PASSWORD", false)
	EnvDisplayName = env.New("SMTP_DISPLAY_NAME", false)
)
