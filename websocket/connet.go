package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chat(c *gin.Context) {
	// 建立连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		println("建立连接失败：" + err.Error())
		return
	}
	defer conn.Close()
	for {
		// 读取客户端发送的数据
		mt, message, err := conn.ReadMessage()
		if err != nil {
			println("读取客户端发送的数据失败：" + err.Error())
			return
		}
		// 写回客户端数据
		err = conn.WriteMessage(mt, message)
	}
}
