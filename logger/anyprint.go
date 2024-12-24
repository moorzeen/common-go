package logger

import (
	"encoding/json"
	"fmt"
)

func AnyPrint(obj interface{}) string {
	JSON, err := json.MarshalIndent(obj, "", "   ")

	if err != nil {
		fmt.Printf("anyprint error: %s", err)
		return ""
	}

	return "\n" + string(JSON)
}
