package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
	"context"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(stream user.User_UploadServer) error {
	var fileData []byte
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}

	// 将文件数据写入本地文件
	ipfs := db.GetIpfs()
	cid, err := ipfs.UploadIPFS(fileData)
	if err != nil {
		return err

	}

	// 返回上传结果
	if err := stream.SendAndClose(&user.Response{
		Message: "上传成功",
		Code:    200,
		Data:    cid,
	}); err != nil {
		return err
	}

	return nil
}
