package error

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type Error struct {
	ServiceCode int
	ErrorCode   int
	Message     string
}

func (e *Error) GRPCStatus() *status.Status {
	var code codes.Code
	code = codes.Code(uint32(e.ErrorCode))
	return status.New(code, e.GetMessage())
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	message := e.GetMessage()
	if message != "" {
		return message
	}
	return "خطای سرور"
}

func (e *Error) GetMessage() string {
	errorCode := e.ErrorCode
	errorCodeStr := strconv.Itoa(errorCode)
	if errorServiceCode != 0 {
		errorCode, _ = strconv.Atoi(errorCodeStr[len(strconv.Itoa(errorServiceCode)):])
	}

	return GetErrorMessage(errorCode)
}

func GetErrorMessage(errorCode int, params ...interface{}) string {
	if message, ok := I18nTranslates[DefaultLanguage][errorCode]; ok {
		return fmt.Sprintf(message, params...)
	}
	return ""
}
