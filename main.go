package main

import (
	"fmt"
	"time"

	"github.com/slowptr/uggo/ddragon"
	"github.com/slowptr/uggo/lcu"
	"github.com/slowptr/uggo/ugg"
)

func main() {
	c := ugg.Init()
	dd := ddragon.Init()
	lol := lcu.Init()

	var lastChampName string
	for true {
		time.Sleep(time.Second * 1)
		if lol.RequestGameFlow() != "ChampSelect" {
			time.Sleep(time.Second * 5)
			continue
		}

		// var buffer string
		// fmt.Print("press enter")
		// fmt.Scan(&buffer)

		champID := lol.RequestCurrentChampionID()
		if champID == "" {
			continue
		}

		champName := ddragon.GetChampionNameFromID(dd, champID)
		selectedChamp := ugg.SearchChampion(c, champName)
		if selectedChamp.Url == "" {
			continue
		}

		if lastChampName == champName {
			continue
		}
		lastChampName = champName

		fmt.Printf("[%s ~ %s] ", champID, champName)
		fmt.Printf("%s selected..\n", selectedChamp.Name)

		pageID := lol.RequestCurrentPage().Id
		page := lcu.UGGRequestToLCUPage(c, dd, pageID, selectedChamp.Name, func() {
			ugg.Visit(c, selectedChamp.Url)
		})
		lol.RequestSetPage(page.Id, page)
	}
}
