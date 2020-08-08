//

package GroupWebApp

import (
	"errors"
	"io/ioutil"
	"os"
)

// Default "config.json" file path (when not specified as running argument).
const CONFIG_DEFAULTPATH = "config.json"

type AppConfig struct {
	dbDriver         string
	dbDataSourceName string
	setup            bool
}

var config AppConfig

func InitConfig(filePath string) (bool, error) {

	if filePath == "" {
		filePath = CONFIG_DEFAULTPATH
	}

	fileData, err := ioutil.ReadFile(filePath)
	config = nil

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
