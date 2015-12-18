package cowsay

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"strings"
	"unicode/utf8"
)

const (
	minTongueSize  = 2
	defaultMessage = "moo"
	HelpEyes       = "oo"
	HelpTongue     = "  "
	HelpCow        = "dragon-and-cow"
)

var (
	CowNotFound = errors.New("cow not found")
)

type Options struct {
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

type Cow struct {
	Options  *Options
	Message  string
	IsHelper bool
}

func (c *Cow) Speak() (string, error) {
	if c.Options.List {
		return cowList, nil
	}

	balloon, trail := buildBalloon(c.Message, c.Options)
	cow, err := buildCow(c.Options, trail)

	if err != nil {
		return "", CowNotFound
	}

	return fmt.Sprintf("%s\n%s", balloon, cow), nil
}

func Say(args string) (string, error) {
	cow, err := ParseArgs(args)

	if err != nil {
		return "", err
	}

	return cow.Speak()
}

func ParseArgs(args string) (*Cow, error) {
	opts := &Options{}
	messageArgs, err := flags.ParseArgs(opts, strings.Split(args, " "))

	if err != nil {
		flagError, ok := err.(*flags.Error)
		if ok && flagError.Type == flags.ErrHelp {
			opts.Eyes = HelpEyes
			opts.Cow = HelpCow
			opts.Tongue = HelpTongue

			return &Cow{opts, flagError.Message, true}, nil
		} else {
			return nil, err
		}
	}

	forceMode(opts)
	normalize(opts)

	return &Cow{opts, getMessage(messageArgs), false}, nil
}

func forceMode(opts *Options) {
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

func normalize(opts *Options) {
	normalizeTongue(opts)
}

func normalizeTongue(opts *Options) {
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
