package result

import "Nft-Go/gateway/internal/types"

func OperateSuccess(data string, message string) (*types.CommonResponse, error) {
	return &types.CommonResponse{
		Code:    200,
		Data:    data,
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
