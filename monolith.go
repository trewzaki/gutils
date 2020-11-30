package gutils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/labstack/echo/v4"
)

// GatewayLogger : Gateway log with physical path, function name, client IP and request timefor gin-gonic library
func GatewayLogger(c *gin.Context, funcName string) {
	fmt.Printf("\n>>>>>>>>> Path: %s => Trig %s() function\n", c.Request.URL, funcName)
	fmt.Print("========> Request From: ", c.ClientIP())
	fmt.Println(" | Request Time:", time.Now())
}

// GatewayLoggerEcho : Gateway log with physical path, function name, client IP and request time for echo library
func GatewayLoggerEcho(c echo.Context, funcName string) {
	fmt.Printf("\n>>>>>>>>> Path: %s => Trig %s() function\n", c.Request().RequestURI, funcName)
	fmt.Print("========> Request From: ", c.RealIP())
	fmt.Println(" | Request Time:", time.Now())
}

// GetRequest : Get payload request of all request method
func GetRequest(c *gin.Context, reqMap map[string]interface{}, getMethodVarNames []string) {
	reqMethod := c.Request.Method
	if reqMethod != "GET" {
		c.ShouldBindWith(&reqMap, binding.JSON)
	}

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

	reqByte, _ := json.Marshal(reqMap)
	fmt.Printf("[x] Request JSON: %s\n", string(reqByte))
}
