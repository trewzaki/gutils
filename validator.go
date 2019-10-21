package gutils

import (
	"encoding/json"
)

const (
	// TypeString : string type symbol
	TypeString = 1

	// TypeInt : int type symbol
	TypeInt = 2

	// TypeInt8 : int8 type symbol
	TypeInt8 = 3

	// TypeInt16 : int16 type symbol
	TypeInt16 = 4

	// TypeInt32 : int32 type symbol
	TypeInt32 = 5

	// TypeInt64 : int64 type symbol
	TypeInt64 = 6

	// TypeUInt : uint type symbol
	TypeUInt = 7

	// TypeUInt8 : uint8 type symbol
	TypeUInt8 = 8

	// TypeUInt16 : uint16 type symbol
	TypeUInt16 = 9

	// TypeUInt32 : uint32 type symbol
	TypeUInt32 = 10

	// TypeUInt64 : uint64 type symbol
	TypeUInt64 = 11

	// TypeFloat32 : float32 type symbol
	TypeFloat32 = 12

	// TypeFloat64 : float64 type symbol
	TypeFloat64 = 13

	// TypeObject : map[string]interface{} type symbol
	TypeObject = 14

	// TypeArray : []interface{} type symbol
	TypeArray = 15

	// TypeArrayObject : []map[string]interface{} type symbol
	TypeArrayObject = 16
)

// TypeValidator : Common validator with custom type value
func TypeValidator(varType int, dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		switch varType {
		case 1:
			if _, ok := dataMap[vn].(string); !ok {
				return false
			}
		case 2:
			if _, ok := dataMap[vn].(int); !ok {
				return false
			}
		case 3:
			if _, ok := dataMap[vn].(int8); !ok {
				return false
			}
		case 4:
			if _, ok := dataMap[vn].(int16); !ok {
				return false
			}
		case 5:
			if _, ok := dataMap[vn].(int32); !ok {
				return false
			}
		case 6:
			if _, ok := dataMap[vn].(int64); !ok {
				return false
			}
		case 7:
			if _, ok := dataMap[vn].(uint); !ok {
				return false
			}
		case 8:
			if _, ok := dataMap[vn].(uint8); !ok {
				return false
			}
		case 9:
			if _, ok := dataMap[vn].(uint16); !ok {
				return false
			}
		case 10:
			if _, ok := dataMap[vn].(uint32); !ok {
				return false
			}
		case 11:
			if _, ok := dataMap[vn].(uint64); !ok {
				return false
			}
		case 12:
			if _, ok := dataMap[vn].(float32); !ok {
				return false
			}
		case 13:
			if _, ok := dataMap[vn].(float64); !ok {
				return false
			}
		case 14:
			for _, vn := range varNames {
				if _, ok := dataMap[vn].(map[string]interface{}); !ok {
					return false
				}
			}
		case 15:
			if _, ok := dataMap[vn].([]interface{}); !ok {
				return false
			} else if ok {
				tmpMap := dataMap[vn].([]interface{})
				if len(tmpMap) == 0 {
					return false
				}
			}
		case 16:
			if _, ok := dataMap[vn].([]map[string]interface{}); !ok {
				return false
			} else if ok {
				tmpMap := dataMap[vn].([]map[string]interface{})
				if len(tmpMap) == 0 {
					return false
				}
			}
		}

	}
	return true
}

// IsString : Common validator with string type value
func IsString(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(string); !ok {
			return false
		}
	}

	return true
}

// IsInt : Common validator with int type value
func IsInt(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(int); !ok {
			return false
		}
	}

	return true
}

// IsUInt32 : Common validator with uint32 type value
func IsUInt32(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(uint32); !ok {
			return false
		}
	}

	return true
}

// IsFloat : Common validator with float type value
func IsFloat(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(float64); !ok {
			return false
		}
	}

	return true
}

// IsBool : Common validator with boolean type value
func IsBool(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(bool); !ok {
			return false
		}
	}

	return true
}

// IsInterface : Deprecated!!
func IsInterface(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(interface{}); !ok {
			return false
		}
	}

	return true
}

// IsObject : Common validator with object type value
func IsObject(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(map[string]interface{}); !ok {
			return false
		}
	}

	return true
}

// IsArrayInterface : Deprecated!!
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

// IsArray : Common validator with array type value
func IsArray(dataMap map[string]interface{}, varNames []string) bool {
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

// IsArrayObject : Common validator with array object type value
func IsArrayObject(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].([]map[string]interface{}); !ok {
			return false
		} else if ok {
			tmpMap := dataMap[vn].([]map[string]interface{})
			if len(tmpMap) == 0 {
				return false
			}
		}
	}

	return true
}

// ReturnValidationError : Return error response when something wrong about validate variable
func ReturnValidationError() []byte {
	errRes := errorResponse{}
	errRes.Success = false
	errRes.Message = "Invalid data type or variable name."

	resByte, _ := json.Marshal(errRes)

	return resByte
}
