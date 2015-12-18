package main

import (
	"fmt"
	"github.com/nmonterroso/cowsay"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	message, _ := cowsay.Cowsay(args)
	fmt.Printf("%s", message)
}
