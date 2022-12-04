package ddragon

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/slowptr/uggo/ugg"
	"github.com/slowptr/uggo/utils"
)

type RunesReforgedJSONRune struct {
	Id   int
	Name string
}
type RunesReforgedJSONRunes struct {
	Runes []RunesReforgedJSONRune
}
type RunesReforgedJSONRuneTree struct {
	Id    int
	Key   string
	Icon  string
	Name  string
	Slots []RunesReforgedJSONRunes
}

type DDragonClient struct {
	HttpClient    http.Client
	LatestVersion string
	RunesReforged []RunesReforgedJSONRuneTree
	Champion      ChampionJSON
}

const versionURL = "http://ddragon.leagueoflegends.com/api/versions.json"

func fetchLatestVersion(cl http.Client) string {
	versionList := []string{}
	body := utils.GetHTMLBody(cl, versionURL)
	err := json.Unmarshal(body, &versionList)
	if err != nil {
		log.Fatal(err)
	}
	return versionList[0]
}

func fetchRunesReforged(cl http.Client, version string) []RunesReforgedJSONRuneTree {
	var runesReforgedJSON []RunesReforgedJSONRuneTree

	runesReforgedURL := "https://ddragon.leagueoflegends.com/cdn/" + version + "/data/en_US/runesReforged.json"
	body := utils.GetHTMLBody(cl, runesReforgedURL)
	err := json.Unmarshal(body, &runesReforgedJSON)
	if err != nil {
		log.Fatal(err)
	}

	return runesReforgedJSON
}

func GetRuneTree(dd *DDragonClient, name string) RunesReforgedJSONRuneTree {
	runeTree := RunesReforgedJSONRuneTree{}
	for _, currentRuneTree := range dd.RunesReforged {
		if currentRuneTree.Name == name {
			runeTree = currentRuneTree
			break
		}
	}
	return runeTree
}

func GetRune(dd *DDragonClient, runeTree RunesReforgedJSONRuneTree, name string) RunesReforgedJSONRune {
	rune := RunesReforgedJSONRune{}
	for _, runeSlot := range runeTree.Slots {
		for _, currentRune := range runeSlot.Runes {
			if currentRune.Name == name {
				rune = currentRune
				break
			}
		}
	}
	return rune
}
func GetShard(name string) RunesReforgedJSONRune {
	shardList := map[string]int{ // not represented in data-dragon..
		"Scaling Bonus Health": 5001,
		"Armor":                5002,
		"Magic Resist":         5003,
		"Attack Speed":         5005,
		"Scaling CDR":          5007,
		"Adaptive Force":       5008,
	}

	return RunesReforgedJSONRune{
		shardList[name],
		name,
	}
}
func GetItemsOfRuneTree(dd *DDragonClient, runeTree *ugg.RuneTree) []RunesReforgedJSONRune {
	itemList := []RunesReforgedJSONRune{}
	if runeTree.Name == "" { // shard GetItemsOfRuneTree
		for _, elem := range runeTree.Items {
			ddShard := GetShard(elem)
			itemList = append(itemList, ddShard)
		}
		return itemList
	}

	ddRuneTree := GetRuneTree(dd, runeTree.Name)
	for _, elem := range runeTree.Items {
		ddRune := GetRune(dd, ddRuneTree, elem)
		itemList = append(itemList, ddRune)
	}
	return itemList
}
func Init() *DDragonClient {
	ddragon := DDragonClient{}
	ddragon.HttpClient = http.Client{
		Timeout: time.Second * 6,
	}
	ddragon.LatestVersion = fetchLatestVersion(ddragon.HttpClient)
	ddragon.RunesReforged = fetchRunesReforged(ddragon.HttpClient, ddragon.LatestVersion)
	ddragon.Champion = fetchChampionJSON(ddragon.HttpClient, ddragon.LatestVersion)
	return &ddragon
}
