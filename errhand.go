package errhand

import (
	"log"
)


// CrashIfErr is expected to crash your program if `err` isn't nill
func CrashIfErr(err error) {
	if err != nil {
		log.Fatal(err, "\n")
	}
}

// DescrIfErr return description error if `err` isn't nill
func DescrIfErr(err error) string {
    if err != nil {
        return err.Error()
    }
    return ""
}
