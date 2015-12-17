package lib
import (
	"github.com/nmonterroso/cowsay/cows"
	"fmt"
	"regexp"
	"strings"
)

var (
	commentFilterRegex = regexp.MustCompile("##.*\n")
)

func buildCow(opts *options, trail balloonTrail) (string, error) {
	cowFile := fmt.Sprintf("%s.%s", opts.Cow, "cow")
	cowBytes, err := cows.Asset(cowFile)

	if err != nil {
		return "", err
	}

	cow := string(cowBytes[:])
	cow = commentFilterRegex.ReplaceAllString(cow, "")

	cow = strings.Replace(cow, "$the_cow = <<EOC;\n", "", 1)
	cow = strings.Replace(cow, "$the_cow = <<\"EOC\";\n", "", 1)
	cow = strings.Replace(cow, "\\\\", "\\", -1)
	cow = strings.Replace(cow, "\\@", "@", -1)
	cow = strings.Replace(cow, "$eyes", opts.Eyes, -1)
	cow = strings.Replace(cow, "$tongue", opts.Tongue, -1)
	cow = strings.Replace(cow, "$thoughts", string(trail), -1)
	cow = strings.Replace(cow, "\nEOC", "", 1)

	return cow, nil
}
