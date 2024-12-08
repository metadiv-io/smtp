package smtp

import (
	"strings"

	"gopkg.in/mail.v2"
)

// SendTextEmail sends a text email.
// valuesMap is a map of key-value pairs that will be used to replace the placeholders in the subject and body.
func (d *Dialer) SendTextEmail(to, cc, bcc []string, subject, body string, valuesMap ...map[string]string) error {

	if len(valuesMap) > 0 && valuesMap[0] != nil {
		valueMap := valuesMap[0]
		for key, value := range valueMap {
			subject = strings.ReplaceAll(subject, key, value)
			body = strings.ReplaceAll(body, key, value)
		}
	}

	m := mail.NewMessage()

	m.SetHeader("From", d.sender())
	m.SetHeader("To", to...)
	m.SetHeader("Cc", cc...)
	m.SetHeader("Bcc", bcc...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	mailD := mail.NewDialer(d.Host, d.Port, d.Username, d.Password)
	if err := mailD.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
