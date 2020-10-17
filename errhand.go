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
