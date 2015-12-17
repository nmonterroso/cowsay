package lib

import (
	"github.com/jessevdk/go-flags"
	"strings"
)

type options struct {
	Eyes     string `short:"e" long:"eyes" description:"the eyes to use" default:"oo"`
	Tongue   string `short:"T" long:"tongue" description:"the tongue to use" default:"  "`
	Cow      string `short:"f" long:"file" description:"the file to use" default:"default"`
	Wrap     int    `short:"W" long:"wrap-width" description:"specify where the word should be wrapped" default:"40"`
	List     bool   `short:"l" long:"list" description:"list available cows" default:"false"`
	Borg     bool   `short:"b" long:"borg" description:"borg mode" default:"false"`
	Dead     bool   `short:"d" long:"dead" description:"dead mode" default:"false"`
	Greedy   bool   `short:"g" long:"greedy" description:"greedy mode" default:"false"`
	Paranoia bool   `short:"p" long:"paranoia" description:"paranoia mode" default:"false"`
	Stoned   bool   `short:"s" long:"stoned" description:"stoned mode" default:"false"`
	Tired    bool   `short:"t" long:"tired" description:"tired mode" default:"false"`
	Wired    bool   `short:"w" long:"wired" description:"wired mode" default:"false"`
	Youthful bool   `short:"y" long:"youthful" description:"youthful mode" default:"false"`
	Think    bool   `long:"think" description:"thinking cow" default:"false"`
}

func Cowsay(args string) (string, error) {
	opts := &options{}
	messageArgs, err := flags.ParseArgs(opts, strings.Split(args, " "))

	if err != nil {
		return "", err
	}

	message := strings.Join(messageArgs, " ")
	return message + opts.Eyes, nil
}
