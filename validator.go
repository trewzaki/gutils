package gutils

import (
	"encoding/json"
)

// >> Validation example:
// if ok := gutils.IsString(dataMap, []string{"variable1", "variable2"}); !ok {
// 	return gutils.ReturnValidationError()
// }

// IsString : Common Validator with string type value
func IsString(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(string); !ok {
			return false
		}
	}

	return true
}

// IsInt : Common Validator with int type value
func IsInt(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(uint32); !ok {
			return false
		}
	}

	return true
}

// IsFloat : Common Validator with float type value
func IsFloat(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(float64); !ok {
			return false
		}
	}

	return true
}

// IsBool : Common Validator with boolean type value
func IsBool(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(bool); !ok {
			return false
		}
	}

	return true
}

// IsInterface : Common Validator with interface type value
func IsInterface(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(interface{}); !ok {
			return false
		}
	}

	return true
}

// IsArrayInterface : Common Validator with array interface type value
func IsArrayInterface(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].([]interface{}); !ok {
			return false
		} else if ok {
			tmpMap := dataMap[vn].([]interface{})
			if len(tmpMap) == 0 {
				return false
			}
		}
	}

	return true
}

// ReturnValidationError : Return error response when somethong wrong about validate variable
func ReturnValidationError() []byte {
	errRes := errorResponse{}
	errRes.Success = false
	errRes.Message = "Invalid data type or variable name."

	resByte, _ := json.Marshal(errRes)

	return resByte
}
