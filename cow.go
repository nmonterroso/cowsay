package cowsay

import (
	"fmt"
	"github.com/nmonterroso/cowsay/templates"
	"math/rand"
	"regexp"
	"sort"
	"strings"
)

const (
	cowSuffix = ".cow"
)

var (
	commentFilterRegex = regexp.MustCompile("##.*\n")
	animalList         = generateAnimalList()
)

func buildAnimal(opts *Options, trail balloonTrail) (string, error) {
	if opts.Random {
		opts.Animal = randomAnimalFile()
	}

	cowFile := fmt.Sprintf("%s%s", opts.Animal, cowSuffix)
	cowBytes, err := templates.Asset(cowFile)

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

func randomAnimalFile() string {
	index := rand.Intn(len(animalList))
	return animalList[index]
}

func generateAnimalList() []string {
	list := make([]string, 0)

	for _, cow := range templates.AssetNames() {
		list = append(list, strings.TrimSuffix(cow, cowSuffix))
	}

	sort.Strings(list)
	return list
}
