package main

import (
	"fmt"
	"github.com/nmonterroso/cowsay"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	message, _ := cowsay.Say(args)
	fmt.Printf("%s", message)
}
