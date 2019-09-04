package gutils

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse : Error response model
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// >> Send response example:
// return SendResposne(true, &message, map[string]interface{}{"foo": "bar"})

// SendResponse :  Standard response form in my projects
func SendResponse(success bool, message *string, data map[string]interface{}) []byte {
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
	fmt.Println("[x] Send Response :>> ", string(resByte))

	return resByte
}
