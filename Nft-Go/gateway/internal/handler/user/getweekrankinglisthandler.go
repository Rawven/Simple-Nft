package user

import (
	"net/http"

	"Nft-Go/gateway/internal/logic/user"
	"Nft-Go/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetWeekRankingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetWeekRankingListLogic(r.Context(), svcCtx)
		resp, err := l.GetWeekRankingList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
