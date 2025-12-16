package storage

import (
	"encoding/json"
	"os"
)

type State struct {
	SentMessages int `json:"sent_messages"`
}

func SaveState(state State) {
	file, _ := os.Create("state.json")
	defer file.Close()
	json.NewEncoder(file).Encode(state)
}
