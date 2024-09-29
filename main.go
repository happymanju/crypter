package main

import (
	"os"

	"github.com/happymanju/crypter/crypter"
)

func main() {
	os.Exit(crypter.CLI(os.Args[1:]))
}
