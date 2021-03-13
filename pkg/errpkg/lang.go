package errpkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/phungvandat/source-template/utils/logger"
)

// Lang type
type Lang string

// List of error language
const (
	VN Lang = "vn"
	EN Lang = "en" // Default language
)

// exp: map[key][lang][message]
type langErrMap map[string]map[Lang]string

var (
	mapOfLang = make(langErrMap)
)

func init() {
	rootPath, err := os.Getwd()
	if err != nil {
		logger.Error("failed load list of error by: %v", err)
		return
	}

	// Load error language
	loadToMapOfLang(rootPath)
}

func loadToMapOfLang(rootPath string) {
	var (
		jsonPath  = fmt.Sprintf("%v/assets/error.json", rootPath)
		data, err = ioutil.ReadFile(jsonPath)
	)
	if err != nil {
		logger.Error("failed read file %v by error: %v", err)
		return
	}

	err = json.Unmarshal(data, &mapOfLang)
	if err != nil {
		logger.Error("failed unmarshal to map error by: %v", err)
	}
}
