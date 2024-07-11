package nft

import (
	"net/http"

	"Nft-Go/nft/http/internal/logic/nft"
	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDcHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDcHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := nft.NewGetDcHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetDcHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
