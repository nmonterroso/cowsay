package lib

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"strings"
	"unicode/utf8"
)

const (
	minTongueSize  = 2
	defaultMessage = "moo"
)

type options struct {
	Eyes     string `short:"e" long:"eyes" description:"the eyes to use" default:"oo"`
	Tongue   string `short:"T" long:"tongue" description:"the tongue to use" default:"  "`
	Cow      string `short:"f" long:"file" description:"the file to use" default:"default"`
	MaxWidth int    `short:"W" long:"max-width" description:"specify where the word should be wrapped" default:"40"`
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

	if opts.List {
		return list(), nil
	}

	forceMode(opts)
	normalize(opts)

	balloon, trail := buildBalloon(getMessage(messageArgs), opts)
	cow, err := buildCow(opts, trail)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n%s", balloon, cow), nil
}

func forceMode(opts *options) {
	switch {
	case opts.Borg:
		opts.Eyes = "=="
		opts.Tongue = "  "
	case opts.Dead:
		opts.Eyes = "xx"
		opts.Tongue = "U "
	case opts.Greedy:
		opts.Eyes = "$$"
		opts.Tongue = "  "
	case opts.Paranoia:
		opts.Eyes = "@@"
		opts.Tongue = "  "
	case opts.Stoned:
		opts.Eyes = "**"
		opts.Tongue = "U "
	case opts.Tired:
		opts.Eyes = "--"
		opts.Tongue = "  "
	case opts.Wired:
		opts.Eyes = "OO"
		opts.Tongue = "  "
	case opts.Youthful:
		opts.Eyes = ".."
		opts.Tongue = "  "
	}
}

func normalize(opts *options) {
	normalizeTongue(opts)
}

func normalizeTongue(opts *options) {
	for utf8.RuneCountInString(opts.Tongue) < minTongueSize {
		opts.Tongue += " "
	}
}

func getMessage(args []string) string {
	if len(args) == 0 {
		return defaultMessage
	}

	return strings.Join(args, " ")
}

func list() string {
	return ""
}
