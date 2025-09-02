package main

import (
    "net/mail"
)

// IsValidEmailSyntax checks if the email syntax is valid
func IsValidEmailSyntax(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}
