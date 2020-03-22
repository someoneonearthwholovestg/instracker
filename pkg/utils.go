package pkg

import (
	"encoding/json"
	"io/ioutil"
)

// GetMapFromJSON gets a map[string]string from JSON file
func GetMapFromJSON(pathToJSONFile string) (
	result map[string]string,
	err error,
) {
	dataJSON, err := ioutil.ReadFile(pathToJSONFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataJSON, &result)
	return
}
