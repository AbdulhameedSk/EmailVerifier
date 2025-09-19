// smtp.go
package emailverifier

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

// CheckSMTP performs an email verification on the passed domain via SMTP
func CheckSMTP(domain, username string) (*SMTP, error) {
	var ret SMTP
	email := fmt.Sprintf("%s@%s", username, domain)

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return nil, errors.New("no MX records found")
	}

	client, err := smtp.Dial(mxRecords[0].Host + smtpPort)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if err = client.Hello("localhost"); err != nil {
		return nil, err
	}

	if err = client.Mail("test@example.org"); err != nil {
		return nil, err
	}

	// This is where you would check for a "catch-all" and then the specific email.
	// For simplicity, we will just check the specific email directly.
	if err = client.Rcpt(email); err == nil {
		ret.Deliverable = true
	} else if strings.Contains(err.Error(), "550") {
		ret.Deliverable = false
	} else {
		return nil, err
	}

	return &ret, nil
}