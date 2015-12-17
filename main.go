package main

import (
	"fmt"
	"github.com/nmonterroso/cowsay/lib"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	message, _ := lib.Cowsay(args)
	fmt.Printf("%s", message)
}
