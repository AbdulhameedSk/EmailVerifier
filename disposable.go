// disposable.go
package emailverifier

import "sync"

var disposableDomains = map[string]bool{
	"10minutemail.com": true,
	"yopmail.com":      true,
	// ... add more domains here
}

var disposableSyncDomains sync.Map

func init() {
	for d := range disposableDomains {
		disposableSyncDomains.Store(d, struct{}{})
	}
}

// IsDisposable checks if a domain is a known disposable email provider
func IsDisposable(domain string) bool {
	_, ok := disposableSyncDomains.Load(domain)
	return ok
}