package utils

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/v3/process"
)

func GetHTMLBody(cl http.Client, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := cl.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
func GetPIDOfProcess(processName string) (pid int) {
	ps, _ := ps.Processes()
	for i := range ps {
		if ps[i].Executable() == processName {
			pid = ps[i].Pid()
		}
	}
	return pid
}
func GetCMDLineOfProcess(processName string) (cmd []string, err error) {
	var proc *process.Process
	pid := GetPIDOfProcess(processName)
	proc, err = process.NewProcess(int32(pid))
	cmd, err = proc.CmdlineSlice()
	return cmd, err
}
