package main

import (
	"fmt"
	"github.com/nmonterroso/cowsay"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	args := strings.Join(os.Args[1:], " ")
	message, _ := cowsay.Say(args)
	fmt.Printf("%s", message)
}
