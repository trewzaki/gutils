package gutils

import (
	"encoding/json"
)

const (
	// StringType : string type symbol
	StringType = 1

	// IntType : int type symbol
	IntType = 2

	// Int8Type : int8 type symbol
	Int8Type = 3

	// Int16Type : int16 type symbol
	Int16Type = 4

	// Int32Type : int32 type symbol
	Int32Type = 5

	// Int64Type : int64 type symbol
	Int64Type = 6

	// UIntType : uint type symbol
	UIntType = 7

	// UInt8Type : uint8 type symbol
	UInt8Type = 8

	// UInt16Type : uint16 type symbol
	UInt16Type = 9

	// UInt32Type : uint32 type symbol
	UInt32Type = 10

	// UInt64Type : uint64 type symbol
	UInt64Type = 11

	// Float32Type : float32 type symbol
	Float32Type = 12

	// Float64Type : float64 type symbol
	Float64Type = 13

	// ObjectType : map[string]interface{} type symbol
	ObjectType = 14

	// ArrayType : []interface{} type symbol
	ArrayType = 15

	// ArrayObjectType : []map[string]interface{} type symbol
	ArrayObjectType = 16
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

// IsString : Deprecated!!
func IsString(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(string); !ok {
			return false
		}
	}

	return true
}

// IsInt : Deprecated!!
func IsInt(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(int); !ok {
			return false
		}
	}

	return true
}

// IsUInt32 : Deprecated!!
func IsUInt32(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(uint32); !ok {
			return false
		}
	}

	return true
}

// IsFloat : Deprecated!!
func IsFloat(dataMap map[string]interface{}, varNames []string) bool {
	for _, vn := range varNames {
		if _, ok := dataMap[vn].(float64); !ok {
			return false
		}
	}

	return true
}

// IsBool : Deprecated!!
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

// IsObject : Deprecated!!
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

// IsArray : Deprecated!!
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

// IsArrayObject : Deprecated!!
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
