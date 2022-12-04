package lcu

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/slowptr/uggo/utils"
)

type LCUConnection struct {
	HttpClient http.Client
	Url        string
	AuthToken  string
}

func fetchLCUConnection() (connection LCUConnection) {
	processCMDLine, err := utils.GetCMDLineOfProcess("LeagueClientUx.exe")
	if err != nil {
		log.Fatal(err)
	}

	for _, argument := range processCMDLine {
		argument = strings.ReplaceAll(argument, "\"", "")
		if strings.HasPrefix(argument, "--app-port=") {
			port := strings.Split(argument, "=")[1]
			connection.Url = fmt.Sprintf("https://127.0.0.1:%s", port)
			continue
		}
		if strings.HasPrefix(argument, "--remoting-auth-token=") {
			token := strings.Split(argument, "=")[1]
			connection.AuthToken = token
			continue
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	connection.HttpClient = http.Client{Timeout: time.Second * 2, Transport: tr}
	return connection
}

func Init() LCUConnection {
	return fetchLCUConnection()
}

func (lcu *LCUConnection) NewRequest(method string, body io.Reader, subUrl string, f func(r *http.Request)) {
	url := fmt.Sprintf("%s%s", lcu.Url, subUrl)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("riot", lcu.AuthToken)

	f(req)
}
