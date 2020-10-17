# ErrHand - Golang error handling

ErrHand is a simple Golang package to aid writing simple
programs fast.

## Usage

Import:

	import (
		"github.com/wkoszek/errhand"
	)

API

	errhand.CrashIfErr(err)	<- will crash the program is err != nil
