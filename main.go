package main

import (
	"github.com/nmonterroso/cowsay/lib"
	"strings"
	"os"
	"fmt"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	message, _ := lib.Cowsay(args)
	fmt.Printf("%s", message)
}
