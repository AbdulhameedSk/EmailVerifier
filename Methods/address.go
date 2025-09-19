// address.go
package emailverifier

import (
	"strings"
)

// ParseAddress attempts to parse an email address
func ParseAddress(email string) Syntax {
	isAddressValid := emailRegex.MatchString(email)
	if !isAddressValid {
		return Syntax{Valid: false}
	}

	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := strings.ToLower(email[index+1:])

	return Syntax{
		Username: username,
		Domain:   domain,
		Valid:    isAddressValid,
	}
}