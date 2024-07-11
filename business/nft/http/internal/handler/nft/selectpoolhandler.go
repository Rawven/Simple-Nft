package nft

import (
	"net/http"

	"Nft-Go/nft/http/internal/logic/nft"
	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SelectPoolHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectPoolRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := nft.NewSelectPoolLogic(r.Context(), svcCtx)
		resp, err := l.SelectPool(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
