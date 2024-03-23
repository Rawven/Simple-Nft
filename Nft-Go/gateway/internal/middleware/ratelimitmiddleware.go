package middleware

import (
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type RateLimitMiddleware struct {
}

func NewRateLimitMiddleware() *RateLimitMiddleware {
	return &RateLimitMiddleware{}
}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := []string{r.URL.Path}
		limiter := util.GetLimiter()
		maxTokens := viper.GetInt("rateLimit.maxTokens")
		tokenRate := viper.GetInt("rateLimit.tokenRate")
		currentTime := time.Now().Unix()
		scriptArgs := []interface{}{maxTokens, tokenRate, currentTime}
		result, err := limiter.Run(r.Context(), db.GetRedis(), keys, scriptArgs).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.(int64) == 0 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
