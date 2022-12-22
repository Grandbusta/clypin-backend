package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Clypin running...")
}

var wc []*websocket.Conn

type socket struct {
	conn []*websocket.Conn
	msgs []string
}

var sockets = make(map[string]socket, 0)

func watcher() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	env := os.Environ()
	fmt.Println(env[0], env[1])
	changed := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range changed {
		println(string(data))
		if ent, ok := sockets["123"]; ok {
			ent.msgs = append(sockets["123"].msgs, string(data))
			sockets["123"] = ent
		}
		curr := sockets["123"]
		msg := curr.msgs[len(curr.msgs)-1]
		for _, soc := range curr.conn {
			err := soc.WriteJSON(map[string]string{"msg": msg})
			if err != nil {
				// log.Println("write:", err)
				break
			}
		}
	}
}

func main() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/send/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed")) // true
		log.Println(c.Params("id"))      // 123
		log.Println(c.Query("v"))
		if ent, ok := sockets["123"]; ok {
			ent.conn = append(sockets["123"].conn, c)
			ent.msgs = append(ent.msgs)
			sockets["123"] = ent
		} else {
			sockets["123"] = socket{
				conn: []*websocket.Conn{c},
				msgs: []string{},
			}

		}
		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			if ent, ok := sockets["123"]; ok {
				ent.msgs = append(sockets["123"].msgs, string(msg))
				sockets["123"] = ent
			}
			curr := sockets["123"]
			fmt.Println(curr.msgs)
			msgd := curr.msgs[len(curr.msgs)-1]
			for _, soc := range curr.conn {
				err := soc.WriteJSON(map[string]string{"msg": msgd})
				if err != nil {
					// log.Println("write:", err)
					break
				}
			}
		}

	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	go watcher()
	app.Listen(":3000")
}
