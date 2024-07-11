package nft

import (
	"net/http"

	"Nft-Go/nft/http/internal/logic/nft"
	"Nft-Go/nft/http/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMyDcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := nft.NewGetMyDcLogic(r.Context(), svcCtx)
		resp, err := l.GetMyDc()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
