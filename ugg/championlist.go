package ugg

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type ChampionEntry struct {
	Name string
	Url  string
}

const championSelector = "div.champions-container > a"

func getChampionList() (championList []ChampionEntry) {
	c := colly.NewCollector()
	c.OnHTML(championSelector, func(e *colly.HTMLElement) {
		champ := ChampionEntry{
			Name: e.ChildText(".champion-name"),
			Url:  fmt.Sprintf("https://u.gg%s", e.Attr("href")),
		}
		championList = append(championList, champ)
	})
	c.Visit("https://u.gg/lol/champions/")
	return championList
}

func SearchChampion(c *colly.Collector, searchQuery string) (selectedChamp ChampionEntry) {
	champList := getChampionList()
	championNameList := []string{}
	for _, champ := range champList {
		championNameList = append(championNameList, champ.Name)
	}

	selected := fuzzy.RankFindNormalizedFold(searchQuery, championNameList)
	if selected.Len() < 1 {
		log.Fatal(championNameList)
		return selectedChamp
	}

	selectedChamp.Name = selected[0].Target
	for _, champ := range champList {
		if champ.Name == selectedChamp.Name {
			selectedChamp.Url = champ.Url
			break
		}
	}

	return selectedChamp
}
