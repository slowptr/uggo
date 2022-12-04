package lcu

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
	"github.com/slowptr/uggo/ddragon"
	"github.com/slowptr/uggo/ugg"
)

type LCUPage struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	PrimaryStyleId  int    `json:"primaryStyleId"`
	SelectedPerkIds []int  `json:"selectedPerkIds"`
	SubStyleId      int    `json:"subStyleId"`
}

func MakeLCUPage(pageID int, pageName string, primaryID int, subID int, itemIDs []int) LCUPage {
	return LCUPage{
		Id:              pageID,
		Name:            pageName,
		PrimaryStyleId:  primaryID,
		SelectedPerkIds: itemIDs,
		SubStyleId:      subID,
	}
}
func UGGRequestToLCUPage(c *colly.Collector, dd *ddragon.DDragonClient, pageID int, pageName string, f func()) LCUPage {
	var primaryID int
	var subID int
	var itemIDs []int

	ugg.OnHTML_Iterate(c, func(idx int, runeTree *ugg.RuneTree) {
		ddRuneTree := ddragon.GetRuneTree(dd, runeTree.Name)
		switch idx {
		case 0:
			primaryID = ddRuneTree.Id
		case 1:
			subID = ddRuneTree.Id
		}

		for _, item := range ddragon.GetItemsOfRuneTree(dd, runeTree) {
			itemIDs = append(itemIDs, item.Id)
		}
	})

	// ugg visit
	f()

	return MakeLCUPage(pageID, pageName, primaryID, subID, itemIDs)
}
func LCUPageToJSON(page LCUPage) []byte {
	json, err := json.Marshal(page)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
