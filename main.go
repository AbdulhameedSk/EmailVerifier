package main

import (
	"log"
	"emailverifier/methods"
)

func main() {
	email := "test@yopmail.com"
	syntaxResult := Methods.ParseAddress(email)
	if !syntaxResult.Valid {
		log.Printf("Email syntax is invalid: %s", email)
		return
	}

	isDisposable := Methods.IsDisposable(syntaxResult.Domain)
	if isDisposable {
		log.Printf("Email is from a disposable domain: %s", email)
		return
	}

	mxResult, err := Methods.CheckMX(syntaxResult.Domain)
	if err != nil {
		log.Printf("Failed to check MX records: %v", err)
		return
	}
	if !mxResult.HasMXRecord {
		log.Printf("Domain has no MX records: %s", syntaxResult.Domain)
		return
	}

	smtpResult, err := Methods.CheckSMTP(syntaxResult.Domain, syntaxResult.Username)
	if err != nil {
		log.Printf("SMTP check failed: %v", err)
		return
	}

	if smtpResult.Deliverable {
		log.Printf("Email address is deliverable: %s", email)
	} else {
		log.Printf("Email address is NOT deliverable: %s", email)
	}
}