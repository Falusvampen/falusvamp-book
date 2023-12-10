package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader      = websocket.Upgrader{}
	wsConnections = &sync.Map{}
)

type wsInput struct {
	Identifier string                 `json:"identifier"`
	Payload    map[string]interface{} `json:"payload"`
}

type wsPost struct {
	Identifier  string            `json:"identifier"`
	PayloadPost PostDTOoutElement `json:"payload_post"`
}

type wsComment struct {
	Identifier     string  `json:"identifier"`
	PostID         int     `json:"postID"`
	PayloadComment Comment `json:"payload_comment"`
}

type wsStatus struct {
	Identifier    string `json:"identifier"`
	UserId        int    `json:"userid"`
	PayloadOnline bool   `json:"payload_online"`
}

type wsMessage struct {
	Identifier     string  `json:"identifier"`
	PayloadMessage Message `json:"payload_message"`
}

type wsTyping struct {
	Type         string `json:"type"`
	UsernameFrom string `json:"usernameFrom"`
	Typing       bool   `json:"typing"`
}

type wsError struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

type wsGroupPosts struct {
	Type  string              `json:"type"`
	Posts []PostDTOoutElement `json:"posts"`
}

type wsGroupPost struct {
	Type string `json:"type"`
	Post Post   `json:"post"`
}

type wsGroupPostComment struct {
	Type    string       `json:"type"`
	PostID  int          `json:"postID"`
	Comment GroupComment `json:"group_comment"`
}

type wsGroupPostComments struct {
	Type          string         `json:"type"`
	PostID        int            `json:"postID"`
	GroupComments []GroupComment `json:"group_comments"`
}

type wsNotification struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func wsConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Get the UUID from the cookie.
	cookie, err := r.Cookie("user_uuid")
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, wsError{"error", "no cookie"})
		return
	}
	uuid := cookie.Value

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	handleConnection(uuid, conn)
}

func handleConnection(uuid string, conn *websocket.Conn) {
	defer wsConnections.Delete(conn)
	defer conn.Close()
	for {
		messageType, incoming, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if messageType == websocket.TextMessage {
			var data wsInput
			if err := json.Unmarshal([]byte(incoming), &data); err != nil {
				log.Println(err)
				return
			}
			switch data.Identifier {
			case "login":
				wsConnections.Store(uuid, conn)
				sendWSStatus(data.Payload["id"].(int), true)
				defer sendWSStatus(data.Payload["id"].(int), false)
			case "logout":
				conn.Close()
				wsConnections.Delete(conn)
				sendWSStatus(data.Payload["id"].(int), false)
			}
		}
		if messageType == websocket.BinaryMessage {
			log.Println("Binary message received")
		}
	}
}

func sendWSALLPosts(posts []PostDTOoutElement) {
	data := wsGroupPosts{"posts", posts}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if c, ok := key.(*websocket.Conn); ok {
			err = c.WriteMessage(websocket.TextMessage, output)
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}

func sendWSPost(post PostDTOoutElement) {
	data := wsPost{"post", post}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if c, ok := key.(*websocket.Conn); ok {
			err = c.WriteMessage(websocket.TextMessage, output)
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}

func sendWSComment(postID int, comment Comment) {
	data := wsComment{"comment", postID, comment}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if c, ok := key.(*websocket.Conn); ok {
			err = c.WriteMessage(websocket.TextMessage, output)
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}

func sendWSStatus(id int, online bool) {
	data := wsStatus{"status", id, online}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if value.(string) != "" {
			err = key.(*websocket.Conn).WriteMessage(websocket.TextMessage, output) // todo: CHECK IT! err was added, not sure it is correct
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}

func sendWSMessage(message Message) {
	data := wsMessage{"message", message}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if value.(string) == message.UsernameFrom || value.(string) == message.UsernameTo {
			err = key.(*websocket.Conn).WriteMessage(websocket.TextMessage, output) // todo: CHECK IT! err was added, not sure it is correct
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}

func sendWSTyping(typing Typing) {
	data := wsTyping{"typing", typing.UsernameFrom, typing.Typing}
	output, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	wsConnections.Range(func(key, value interface{}) bool {
		if value.(string) == typing.UsernameTo {
			err = key.(*websocket.Conn).WriteMessage(websocket.TextMessage, output) // todo: CHECK IT! err was added, not sure it is correct
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
}
