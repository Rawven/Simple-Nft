package middleware

import (
	"Nft-Go/common/util"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
	"net/http"
)

type JwtMiddleware struct {
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		userId, err := util.ParseToken(token)
		// 如果报错则不合法直接返回， 否则将信息塞入context
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorized"))
			if err != nil {
				return
			}
			return
		}
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "userId", convertor.ToString(userId))
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
