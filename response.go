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

// ErrorResponse : Error response model
type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// >> Send response example:
// return SendResposne(true, &message, map[string]interface{}{"foo": "bar"})

// SendResponse :  Standard response form in my projects
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
		fmt.Println("[x] Send Response :>> ", string(resByte))
	}
	responseLogger = true

	return resByte
}

// SetResponseLogger : Set logger status to print or no print in SendResponse function
func SetResponseLogger(printStatus bool) {
	responseLogger = printStatus
}
