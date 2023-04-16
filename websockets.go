// websockets.go
package main

import (
	"net/http"

	"github.com/smcclure17/writr/pkg/database"
	"github.com/smcclure17/writr/pkg/models"
	"github.com/smcclure17/writr/pkg/server"
	"golang.org/x/net/websocket"
)

func main() {

	server := server.NewServer()
	db := database.NewDatabase()

	http.Handle("/ws", websocket.Handler(server.HandleConnections))

	// HACK: serve temp html file for testing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	// TODO: think about when/how to save to disk to keep data up to date without too many writes
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		db.InsertData("Messages", models.CreateMessage("this is a placeholder message", "doc-1"))

		w.Write([]byte("success"))
	})

	http.ListenAndServe(":8081", nil)
}
