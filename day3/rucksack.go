package main

import "unicode/utf8"

type Rucksack string

func (r Rucksack) Priority() int {
	for _, firstRune := range r.firstCompartment() {
		for _, secRune := range r.secondCompartment() {
			if firstRune == secRune {
				if int(firstRune) > 96 {
					return int(firstRune) - 96
				} else {
					return int(firstRune) - 38
				}
			}
		}
	}
	return 0
}

func (r Rucksack) firstCompartment() string {
	return string(r)[:utf8.RuneCountInString(string(r))/2]
}

func (r Rucksack) secondCompartment() string {
	return string(r)[utf8.RuneCountInString(string(r))/2:]
}
