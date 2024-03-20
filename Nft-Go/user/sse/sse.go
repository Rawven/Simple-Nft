package sse

import (
	"Nft-Go/common/util"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"log"
	"net/http"
	"sync"

	"gopkg.in/antage/eventsource.v1"
)

var (
	sessions map[int32]eventsource.EventSource
	mu       sync.Mutex
)

func InitSse() {
	sessions = make(map[int32]eventsource.EventSource)
	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		// 获取用户ID，这里假设用户ID在URL参数中
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}
		userId, err := util.ParseToken(token)
		if err != nil {
			return
		}
		// 创建一个新的EventSource实例并保存到sessions中
		userEs := eventsource.New(nil, nil)
		mu.Lock()
		sessions[*userId] = userEs
		mu.Unlock()

		// 处理连接
		userEs.ServeHTTP(w, r)
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	logger.Info("Open URL http://localhost:8080/events?userID=<your-user-id> in your browser.")
}

func SendNotificationToSingleUser(userID int32, message string) {
	mu.Lock()
	userEs, ok := sessions[userID]
	mu.Unlock()
	if !ok {
		logger.Info("No session found for user %s", userID)
		return
	}
	userEs.SendEventMessage(message, "", "")
}
func SendNotificationToAllUser(message string) {
	mu.Lock()
	for _, userEs := range sessions {
		userEs.SendEventMessage(message, "", "")
	}
	mu.Unlock()
}
