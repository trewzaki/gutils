package gutils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GatewayLogger : Gateway log with physical path, function name, client IP and request time
func GatewayLogger(c *gin.Context, functionName string) {
	// tmpPath, _ := c.Get("path")
	fmt.Printf("\n>>>>>>>>> Path: %s => Trig %s() function\n", c.Request.URL, functionName)
	fmt.Print("========> Request From: ", c.ClientIP())
	fmt.Println(" | Request Time:", time.Now())
}

// GetRequest : Get both of POST and GET method request
func GetRequest(c *gin.Context, reqMap map[string]interface{}, getMethodVarNames []string) {
	c.BindJSON(&reqMap)
	reqByte, _ := json.Marshal(reqMap)

	if getMethodVarNames != nil {
		var err error
		for i := 0; i < len(getMethodVarNames); i++ {
			queryValue := c.Query(getMethodVarNames[i])
			if queryValue != "" {
				reqMap[getMethodVarNames[i]], err = strconv.Atoi(queryValue)
				if err != nil {
					reqMap[getMethodVarNames[i]] = queryValue
				}
			}

			paramValue := c.Param(getMethodVarNames[i])
			if paramValue != "" {
				reqMap[getMethodVarNames[i]], err = strconv.Atoi(paramValue)
				if err != nil {
					reqMap[getMethodVarNames[i]] = paramValue
				}
			}
		}
	}

	fmt.Printf("[x] Request JSON: %s\n", string(reqByte))
}