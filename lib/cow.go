package lib

import (
	"fmt"
	"github.com/nmonterroso/cowsay/cows"
	"regexp"
	"sort"
	"strings"
)

const (
	cowSuffix = ".cow"
)

var (
	commentFilterRegex = regexp.MustCompile("##.*\n")
	cowList            = generateCowList()
)

func buildCow(opts *options, trail balloonTrail) (string, error) {
	cowFile := fmt.Sprintf("%s%s", opts.Cow, cowSuffix)
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

func generateCowList() string {
	list := make([]string, 0)

	for _, cow := range cows.AssetNames() {
		list = append(list, strings.TrimSuffix(cow, cowSuffix))
	}

	sort.Strings(list)
	return strings.Join(list, ", ")
}
