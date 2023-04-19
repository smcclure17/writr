// websockets.go
package main

import (
	"fmt"
	"net/http"

	"github.com/smcclure17/writr/pkg/cache"
	"github.com/smcclure17/writr/pkg/database"
	"github.com/smcclure17/writr/pkg/models"
	"github.com/smcclure17/writr/pkg/server"
	"golang.org/x/net/websocket"
)

func main() {

	db := database.NewDatabase()
	cache := cache.NewCache()
	server := server.NewServer(*cache, *db)
	fmt.Println("Created database", db)

	http.Handle("/ws", websocket.Handler(server.HandleConnections))

	// // TODO: think about when/how to save to disk to keep data up to date without too many writes
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.InsertData("documents", models.CreateMessage("the he he yup!e", "doc-1"))
		if err != nil {
			fmt.Println("error inserting data: ", err)
		}

		w.Write([]byte("success"))
	})

	http.ListenAndServe(":8081", nil)
}
