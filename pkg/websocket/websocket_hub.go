package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (hub *WebSocketHub) Run() {
	for {
		select {
		case conn := <-hub.register:
			hub.mu.Lock()
			hub.clients[conn] = true
			hub.mu.Unlock()
		case conn := <-hub.unregister:
			hub.mu.Lock()
			if _, ok := hub.clients[conn]; ok {
				delete(hub.clients, conn)
				conn.Close()
			}
			hub.mu.Unlock()
		case message := <-hub.broadcast:
			hub.mu.Lock()
			for conn := range hub.clients {
				err := conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					conn.Close()
					delete(hub.clients, conn)
				}
			}
			hub.mu.Unlock()
		}
	}
}

func (hub *WebSocketHub) Broadcast(message []byte) {
	hub.broadcast <- message
}

func (hub *WebSocketHub) RegisterConnection(conn *websocket.Conn) {
	hub.register <- conn
}

func (hub *WebSocketHub) UnregisterConnection(conn *websocket.Conn) {
	hub.unregister <- conn
}

func (hub *WebSocketHub) ConnectWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Không thể mở kết nối WebSocket", http.StatusBadRequest)
		return
	}
	hub.RegisterConnection(conn)
	defer hub.UnregisterConnection(conn)

	// Đặt read deadline và xử lý tin nhắn ping/pong để giữ cho kết nối hoạt động
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Vòng lặp đọc tin nhắn từ client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break // Nếu có lỗi, thoát khỏi vòng lặp và đóng kết nối
		}
		hub.Broadcast(message) // Phát thông điệp tới tất cả các client
	}
}
