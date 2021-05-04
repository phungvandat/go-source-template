package helper

import (
	"encoding/json"
	"fmt"

	"github.com/phungvandat/source-template/utils/logger"
)

// Size constants
const (
	MB = 1 << 20
)

func MBToBytes(val int) int64 {
	return int64(val) * MB
}

func PrintJSON(val interface{}) {
	b, err := json.MarshalIndent(val, "  ", "  ")
	if err != nil {
		logger.Error("failed print json by error: %v", err)
		return
	}
	fmt.Println(string(b))
}
