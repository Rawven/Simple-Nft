package middleware

import (
	"Nft-Go/common/util"
	"context"
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
			return
		}
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "userId", userId)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
