package logic

import (
	"Nft-Go/nft/http/internal/types"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
)

func OperateSuccess(data interface{}, message string) (*types.CommonResponse, error) {
	return &types.CommonResponse{
		Code:    200,
		Data:    util.ToJsonString(data),
		Message: message,
	}, nil
}

func OperateSuccessWithoutData(message string) (*types.CommonResponse, error) {
	return &types.CommonResponse{
		Code:    200,
		Data:    "",
		Message: message,
	}, nil
}

func OperateFailed(message string) (*types.CommonResponse, error) {
	return &types.CommonResponse{
		Code:    500,
		Data:    "",
		Message: message,
	}, nil
}
