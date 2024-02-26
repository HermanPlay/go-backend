package util

import (
	"github.com/HermanPlay/backend/internal/api/http/constant"
	"github.com/HermanPlay/backend/pkg/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatusCode(), responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](statuscode int, status, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		StatusCode:      statuscode,
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
