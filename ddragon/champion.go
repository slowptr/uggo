package ddragon

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/slowptr/uggo/utils"
)

type ChampionDataJSON struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
type ChampionJSON struct {
	Data map[string]ChampionDataJSON `json:"data"`
}

func fetchChampionJSON(cl http.Client, version string) (champResult ChampionJSON) {
	championURL := "https://ddragon.leagueoflegends.com/cdn/" + version + "/data/en_US/champion.json"
	body := utils.GetHTMLBody(cl, championURL)
	err := json.Unmarshal(body, &champResult)
	if err != nil {
		log.Fatal(err)
	}
	return champResult
}
func GetChampionNameFromID(dd *DDragonClient, id string) (name string) {
	for _, v := range dd.Champion.Data {
		if v.Key == id {
			name = v.Name
		}
	}
	return name
}
