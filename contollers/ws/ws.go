package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

var wc []*websocket.Conn

type Data struct {
	Msg    string `json:"msg"`
	Source string `json:"source"`
}
type socket struct {
	Conn  []*websocket.Conn
	Datas []Data
}

var sockets = make(map[string]socket, 0)

func Ws() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed")) // true
		user_id := c.Params("user_id")   // 123
		token := c.Query("token")
		log.Println(token)
		if ent, ok := sockets[user_id]; ok {
			ent.Conn = append(sockets[user_id].Conn, c)
			sockets[user_id] = ent
		} else {
			sockets[user_id] = socket{
				Conn:  []*websocket.Conn{c},
				Datas: []Data{},
			}
		}
		var (
			data Data
			err  error
		)
		for {

			if err = c.ReadJSON(&data); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", data)
			if ent, ok := sockets[user_id]; ok {
				ent.Datas = append(sockets[user_id].Datas, data)
				sockets[user_id] = ent
			}
			curr := sockets[user_id]
			fmt.Println(curr.Datas)
			msgd := curr.Datas[len(curr.Datas)-1]
			for _, soc := range curr.Conn {
				err = soc.WriteJSON(msgd)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		}
	}
}
