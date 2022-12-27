package models

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/set"
	"github.com/gorilla/websocket"
)

// swagger:model Message
type Message struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	media       int
	sender_id   string
	receiver_id string
	content     string
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var ClientMap map[string]*Node = make(map[string]*Node, 0)

var rwLock sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	//获取参数
	query := request.URL.Query()
	senderId := query.Get("userId")
	// media := query.Get("media")
	// content := query.Get("content")
	// receiverId := query.Get("receiverId")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
	}
	//连接信息
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwLock.Lock()
	ClientMap[senderId] = node
	rwLock.Unlock()
	//发送
	go sendProc()
	//接收
	go recvProc()
}

func sendProc() {

}

func recvProc() {

}
