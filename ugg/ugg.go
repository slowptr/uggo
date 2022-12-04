package ugg

import (
	"github.com/gocolly/colly"
)

const topContainerSelector = "div.rune-trees-container-2.media-query.media-query_MOBILE_LARGE__DESKTOP_LARGE"
const runeTreeSelector = "div .rune-tree_v2"

func Init() *colly.Collector {
	return colly.NewCollector(colly.AllowedDomains("u.gg"))
}
func Visit(c *colly.Collector, url string) {
	c.Visit(url)
}
func VisitChamp(c *colly.Collector, championName string) {
	Visit(c, "https://u.gg/lol/champions/"+championName+"/build")
}
func OnHTML_Iterate(c *colly.Collector, f func(i int, r *RuneTree)) {
	c.OnHTML(topContainerSelector, func(e *colly.HTMLElement) {
		e.ForEach(runeTreeSelector, func(idx int, treeElem *colly.HTMLElement) {
			f(idx, FetchRuneTree(treeElem))
		})
	})
}
