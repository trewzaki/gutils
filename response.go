package gutils

import (
	"encoding/json"
	"fmt"
)

const (
	// NoPrint : Setting to no print type
	NoPrint = false

	// Print : Setting to print type
	Print = true
)

var responseLogger = true
var responseLoggerLimit = uint32(10000)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// >> Send response example:
// return SendResposne(true, &message, map[string]interface{}{"foo": "bar"})

// SendResponse : Standard response form in my projects
func SendResponse(success bool, message *string, data interface{}) []byte {
	resMap := map[string]interface{}{
		"success": success,
	}

	if message != nil {
		resMap["message"] = *message
	}

	if data != nil {
		resMap["data"] = data
	}

	resByte, _ := json.Marshal(resMap)

	if responseLogger {
		if len(resByte) < int(responseLoggerLimit) {
			fmt.Println("[x] Send Response :>> ", string(resByte))
		} else {
			fmt.Println("[x] Send Response :>> ", string(resByte[0:responseLoggerLimit]))
		}
	}
	responseLogger = true

	return resByte
}

// SendResponseV2 : Standard response form in my projects
func SendResponseV2(success bool, message *string, data interface{}, errors []string, pages *uint32) []byte {
	resMap := map[string]interface{}{
		"success": success,
	}

	if message != nil {
		resMap["message"] = *message
	}

	if data != nil {
		resMap["data"] = data
	}

	if errors != nil {
		resMap["errors"] = errors
	}

	if pages != nil {
		resMap["pages"] = pages
	}

	resByte, _ := json.Marshal(resMap)

	if responseLogger {
		if len(resByte) < int(responseLoggerLimit) {
			fmt.Println("[x] Send Response :>> ", string(resByte))
		} else {
			fmt.Println("[x] Send Response :>> ", string(resByte[0:responseLoggerLimit]))
		}
	}
	responseLogger = true

	return resByte
}

// SetResponseLogger : Set logger status to print or no print in SendResponse function
func SetResponseLogger(printStatus bool) {
	responseLogger = printStatus
}

// SetResponseLogger : Set maximum text length of log printer
func SetResponseLoggerLimit(limit uint32) {
	responseLoggerLimit = limit
}
