package websocket

import (
	"Nft-Go/common/util"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	mapForConnections sync.Map
}

func NewWebsocketServer() *Server {
	return &Server{}
}

func (ws *Server) echo(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Failed to upgrade to WebSocket connection")
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	userId := getUserId(r)
	ws.mapForConnections.Store(userId, conn)
	defer func() {
		ws.mapForConnections.Delete(userId)
	}()

	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		switch message {
		//TODO add your code here

		}
	}
}

func (ws *Server) SendMessageToAll(messageType int, message []byte) {
	ws.mapForConnections.Range(func(key, value interface{}) bool {
		conn := value.(*websocket.Conn)
		err := conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			err := conn.Close()
			if err != nil {
				return false
			}
			ws.mapForConnections.Delete(key)
		}
		return true
	})
}

func (ws *Server) SendMessageToUser(userId int, messageType int, message []byte) error {
	value, ok := ws.mapForConnections.Load(userId)
	if !ok {
		return fmt.Errorf("connection not found for user %v", userId)
	}
	conn := value.(*websocket.Conn)
	return conn.WriteMessage(messageType, message)
}

func getUserId(r *http.Request) int {
	token := r.Header.Get("token")
	userId, _ := util.ParseToken(token)
	return *userId
}

func InitWebsocket() {
	ws := NewWebsocketServer()
	http.HandleFunc("/echo", ws.echo)
	log.Fatal(http.ListenAndServe(":9991", nil))
}
