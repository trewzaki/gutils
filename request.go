package gutils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

// EchoBind : Bind all request payload and addition data in echo context
func EchoBind(c echo.Context, req interface{}, contextList []string) {
	reqMap := map[string]interface{}{}
	c.Bind(&reqMap)

	for _, v := range contextList {
		if c.Get(v) != nil {
			reqMap[v] = c.Get(v).(string)
		}
	}

	tmpByte, _ := json.Marshal(reqMap)
	json.Unmarshal(tmpByte, &req)
}
