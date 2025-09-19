// result.go
package Methods

type Syntax struct {
	Username string
	Domain   string
	Valid    bool
}

type Mx struct {
	HasMXRecord bool
}

type SMTP struct {
	Deliverable bool
	CatchAll    bool
}

type Result struct {
	Email        string
	Syntax       Syntax
	Disposable   bool
	HasMxRecords bool
	SMTP         *SMTP
	Reachable    string
}