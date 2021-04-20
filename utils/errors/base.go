package errors

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type content struct {
	Name         string `json:"name"`
	Endpoint     string `json:"endpoint"`
	Method       string `json:"method"`
	ErrorCode    int16  `json:"httpStatus"`
	Source       string `json:"source,omitempty"`
	Message      string `json:"message"`
	Timestamp    int64  `json:"timestamp"`
	CleanedStack string `json:"stack"`
}

// Base Error
type Base struct {
	c     content
	stack string
}

var serviceName string

func init() {
	serviceName = strings.ReplaceAll(os.Getenv("SERVICE_NAME"), "_", "-")
}

// Error getter message
func (err *Base) Error() string {
	if err.c.CleanedStack == "" {
		err.c.CleanedStack = cleanStack(err.stack)
	}

	j, _ := json.Marshal(&err.c)

	return fmt.Sprintf("%s\n", string(j))
}

// SetSource setter
func (err *Base) SetSource(source string) {
	err.c.Source = source
}

// Name getter
func (err *Base) Name() string {
	return err.c.Name
}

// ErrorCode getter
func (err *Base) ErrorCode() int16 {
	return err.c.ErrorCode
}

// Message getter
func (err *Base) Message() string {
	return err.c.Message
}

// NewError function
func NewError(
	name,
	endpoint,
	method,
	message string,
	errorCode int16,
	stack []byte) *Base {
	return &Base{
		c: content{
			Name:         name,
			Endpoint:     endpoint,
			Method:       method,
			ErrorCode:    errorCode,
			Message:      message,
			Timestamp:    time.Now().UnixNano() / 1000,
			CleanedStack: "",
		},
		stack: string(stack),
	}
}

func cleanStack(stack string) string {
	splitted := strings.Split(stack, "\n")

	splitted = splitted[5:]

	for i, str := range splitted {
		arrStr := strings.Split(str, "/"+serviceName)
		cleanedStr := arrStr[len(arrStr)-1]
		cleanedStr = strings.ReplaceAll(cleanedStr, os.Getenv("GOPATH"), "GOPATH")
		cleanedStr = strings.ReplaceAll(cleanedStr, os.Getenv("GOROOT"), "GOROOT")
		cleanedStr = strings.ReplaceAll(cleanedStr, "\t", "")

		regexFunctionArgs := regexp.MustCompile(`\(0x.*\)`)
		cleanedStr = regexFunctionArgs.ReplaceAllString(cleanedStr, "()")

		regexMemoryAddr := regexp.MustCompile(`\ \+0x.*`)
		cleanedStr = regexMemoryAddr.ReplaceAllString(cleanedStr, "")

		splitted[i] = cleanedStr
	}

	return strings.Join(splitted, "; ")
}
