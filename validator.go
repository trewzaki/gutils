package gutils

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// >> Validation example:
// if ok := utils.IsString(dataMap, []string{"variable1", "variable2"}); !ok {
// 	return utils.ReturnValidationError()
// }

func IsString(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(string); !ok {
			return false
		}
	}

	return true
}

func IsInt(dataMap map[string]interface{}, varNames []string, c amqp.Delivery) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(uint32); !ok {
			return false
		}
	}

	return true
}

func IsFloat(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(float64); !ok {
			return false
		}
	}

	return true
}

func IsBool(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(bool); !ok {
			return false
		}
	}

	return true
}

func IsInterface(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(interface{}); !ok {
			return false
		}
	}

	return true
}

func IsArrayInterface(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].([]interface{}); !ok {
			return false
		} else {
			tmpMap := dataMap[vn].([]interface{})
			if len(tmpMap) == 0 {
				return false
			}
		}
	}

	return true
}

func ReturnValidationError() []byte {
	errRes := ErrorResponse{}
	errRes.Success = false
	errRes.Message = "Invalid data type or variable name."

	resByte, _ := json.Marshal(errRes)

	return resByte
}
