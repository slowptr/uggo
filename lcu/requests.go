package lcu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (lol *LCUConnection) RequestCurrentPage() (page LCUPage) {
	lol.NewRequest(http.MethodGet, nil, "/lol-perks/v1/currentpage", func(req *http.Request) {
		resp, err := lol.HttpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(body, &page)
		if err != nil {
			log.Fatal(err)
		}
	})
	return page
}

func (lol *LCUConnection) RequestCurrentChampionID() (champID string) {
	lol.NewRequest(http.MethodGet, nil, "/lol-champ-select/v1/current-champion", func(req *http.Request) {
		resp, err := lol.HttpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		if resp.StatusCode != 200 {
			champID = ""
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		champID = string(body[:])
		if champID == "0" {
			champID = ""
		}
	})
	return champID
}

func (lol *LCUConnection) RequestGameFlow() (phase string) {
	lol.NewRequest(http.MethodGet, nil, "/lol-gameflow/v1/gameflow-phase", func(req *http.Request) {
		resp, err := lol.HttpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(body, &phase)
		if err != nil {
			log.Fatal(err)
		}
	})
	return phase
}

func (lol *LCUConnection) RequestSetPage(pageID int, page LCUPage) {
	json := LCUPageToJSON(page)
	lol.NewRequest(http.MethodPut, bytes.NewBuffer(json),
		fmt.Sprintf("/lol-perks/v1/pages/%d", pageID), func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			_, err := lol.HttpClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
		})
}
