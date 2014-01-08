package keiko

import (
	"encoding/json"
	"os"
)

type keikoConf struct {
	KeikoIP string
}

var conf keikoConf

func readConfig() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&conf)
	return
}
