package main

import (	
	"net"
	"strings"
)

func HasMXRecord(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	mxRecords, err := net.LookupMX(domain)
	return err == nil && len(mxRecords) > 0
}
