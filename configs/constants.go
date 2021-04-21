package configs

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

func toInt(v interface{}) (int, error) {
	switch v.(type) {
	case int:
		return v.(int), nil
	case int32:
		return int(v.(int32)), nil
	case uint32:
		return int(v.(uint32)), nil
	case float32:
		return int(v.(float32)), nil
	case float64:
		return int(v.(float64)), nil
	case string:
		t, err := strconv.Atoi(v.(string))
		if err != nil {
			return 0, fmt.Errorf("toInt: string %s cannot be converted to int", v.(string))
		}
		return t, nil
	default:
		return 0, fmt.Errorf("toInt: %s cannot be converted to int", reflect.ValueOf(v).Kind().String())
	}
}

type constant struct {
	TimeoutOnSeconds       time.Duration
	OperationOnEachContext int
}

func setupConstant() *constant {
	timeoutOnSecondsInt, _ := toInt(os.Getenv("TIMEOUT_ON_SECONDS"))
	timeoutOnSeconds := time.Duration(timeoutOnSecondsInt)
	operationOnEachContext, _ := toInt(os.Getenv("OPERATION_ON_EACH_CONTEXT"))

	v := &constant{
		TimeoutOnSeconds:       timeoutOnSeconds,
		OperationOnEachContext: operationOnEachContext,
	}

	return v
}
