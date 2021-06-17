package Util

import (
	"bufio"
	"os"
	"regexp"
)

type Preference struct {
	Name   string
	Config map[string]string
}

func ParsePreferenceMap(path string) map[string]*Preference {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parser := preferenceParser{
		scanner:     scanner,
		preferences: make(map[string]*Preference),
	}

	ok := parser.parse()
	if !ok {
		return nil
	}

	return parser.preferences
}

type preferenceParser struct {
	scanner     *bufio.Scanner
	preferences map[string]*Preference
}

func (parser *preferenceParser) parse() bool {
	s := parser.scanner

	regConfigName, _ := regexp.Compile("^\\[(.*)\\]$")
	regConfigItem, _ := regexp.Compile("^(.*?)=(.*)$")

	var currentName = ""
	for s.Scan() {
		line := s.Text()
		if regConfigName.MatchString(line) {
			name := regConfigName.FindStringSubmatch(line)[1]
			currentName = name
			p := Preference{
				Name:   name,
				Config: make(map[string]string),
			}
			parser.preferences[name] = &p

		} else if regConfigItem.MatchString(line) {
			subStrings := regConfigItem.FindStringSubmatch(line)
			key := subStrings[1]
			value := subStrings[2]
			parser.preferences[currentName].Config[key] = value
		}
	}
	return true
}
