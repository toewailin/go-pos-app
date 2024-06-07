package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/toewailin/pos-app/models"
	"github.com/toewailin/pos-app/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

var DB *gorm.DB

func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_pos?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Item{})
}

func main() {
	initDB()
	r := routes.SetupRouter(DB)
	r.Run() // listen and serve on 0.0.0.0:8080
}
