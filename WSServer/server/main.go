package main

import (
	"flag"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"
	"log"
	"net/http"

	"ws-server/proto/wsserver"
)

var addr = flag.String("addr", "0.0.0.0:80", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Debug("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Debug("Read Error: ", err)
			break
		}

		var ping wsserver.CommandPing
		proto.Unmarshal(message, &ping)
		logger.Debug("Ping Code: %v", ping.Code)

		data := wsserver.CommandPong{
			Code: 200,
		}
		dataBuffer, _ := proto.Marshal(&data)
		err = ws.WriteMessage(mt, dataBuffer)

		if err != nil {
			logger.Debug("Write Error: ", err)
			break
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Debug("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Debug("Read Error: ", err)
			break
		}

		var req wsserver.LoginRequest
		proto.Unmarshal(message, &req)
		logger.Debug("Login UserName: %v", req.UserName)

		var Orders []*wsserver.Order
		Orders = append(Orders, &wsserver.Order{
			OrderId:   1,
			OrderName: "Test1",
		})
		Orders = append(Orders, &wsserver.Order{
			OrderId:   2,
			OrderName: "Test2",
		})
		data := wsserver.User{
			UserName: req.UserName,
			Orders:   Orders,
		}
		dataBuffer, _ := proto.Marshal(&data)
		err = ws.WriteMessage(mt, dataBuffer)

		if err != nil {
			logger.Debug("Write Error: ", err)
			break
		}
	}
}

func main() {
	logger.Debug("Running")

	http.HandleFunc("/echo", echo)

	http.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
