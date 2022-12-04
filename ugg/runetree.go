package ugg

import (
	"strings"

	"github.com/gocolly/colly"
)

type RuneTree struct {
	Name  string
	Items []string
}

func stripNames(names []string) []string {
	stripList := [...]string{
		"The ",
		"Rune ",
		"Keystone ",
		" Shard",
	}

	nameList := []string{}
	for _, elem := range names {
		for _, strip := range stripList {
			elem = strings.ReplaceAll(elem, strip, "")
		}
		nameList = append(nameList, elem)
	}
	return nameList
}

func FetchRuneTree(htmlElem *colly.HTMLElement) *RuneTree {
	runeTree := RuneTree{}

	runeTree.Name = htmlElem.Text
	if runeTree.Name == "" {
		runeTree.Items = stripNames(htmlElem.ChildAttrs("div .shard-active > img", "alt"))
	} else {
		runeTree.Items = stripNames(htmlElem.ChildAttrs("div .perk-active > img", "alt"))
	}

	return &runeTree
}
