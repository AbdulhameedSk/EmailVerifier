// mx.go
package Methods

import "net"

// CheckMX checks for MX records for the given domain.
func CheckMX(domain string) (*Mx, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil && len(mxRecords) == 0 {
		return &Mx{HasMXRecord: false}, err
	}
	return &Mx{
		HasMXRecord: len(mxRecords) > 0,
	}, nil
}