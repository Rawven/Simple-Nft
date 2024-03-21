package nft

import (
	"net/http"

	"Nft-Go/gateway/internal/logic/nft"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDigitalCollectionHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDigitalCollectionHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := nft.NewGetDigitalCollectionHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetDigitalCollectionHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
