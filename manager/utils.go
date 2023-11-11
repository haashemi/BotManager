package manager

import (
	"encoding/json"
	"os"
)

// loadBots loads the bots from json file at provided path.
// it returns no error if file not exists (it's possible situation).
//
// It must only get called once on the project startup
func loadBots(path string) (bots []*Bot, err error) {
	file, err := os.Open(BotsDataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}

		return nil, err
	}

	if err = json.NewDecoder(file).Decode(&bots); err != nil {
		return nil, err
	}

	return bots, nil
}

// save serializes the bots to json and writes them at provided path.
func save(bots []*Bot, path string) error {
	data, err := json.Marshal(bots)
	if err != nil {
		return err
	}

	return os.WriteFile(BotsDataPath, data, os.ModePerm)
}
