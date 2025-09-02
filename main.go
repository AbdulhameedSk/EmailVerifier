package main

import "fmt"

func main() {
    emails := []string{
        "test@example.com",
        "invalid-email@",
        "user@domain.co.in",
        "user@nonexistentdomain.fake",
    }

    for _, email := range emails {
        if !IsValidEmailSyntax(email) {
            fmt.Printf("%s: Invalid Syntax\n", email)
            continue
        }
        if !HasMXRecord(email) {
            fmt.Printf("%s: No MX Record found\n", email)
            continue
        }
        fmt.Printf("%s: Syntax and MX Record OK\n", email)
    }
}
