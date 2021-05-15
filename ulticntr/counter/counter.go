package counter

import "os"

const (
	BACKEND_ENV = "ULTICNTR_BACKEND"
)

// Counter is an interface that has a Count() method that increments the counter
// and returns the total count and an error
type Counter interface {
	Count() (uint64, error)
}

// New returns a Counter and an error. It defaults to a dynamoDB counter unless
// the BACKEND_ENV is set to tmp or mem.
func New() (Counter, error) {
	switch os.Getenv(BACKEND_ENV) {
	case "tmp":
		return newTmp()
	case "mem":
		return newMem()
	default:
		return newDynamo()
	}
}
