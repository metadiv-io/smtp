package smtp

import "fmt"

// EmailList returns a list of emails.
func EmailList(emails ...string) []string {
	if len(emails) == 0 {
		return nil
	}
	return emails
}

// sender returns the sender of the email.
func (d *Dialer) sender() string {
	if d.DisplayName == "" {
		return d.Username
	}
	return fmt.Sprintf("%s <%s>", d.DisplayName, d.Username)
}
