package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Read() []TodoList
}

type TodoList struct {
	ID    string `json:"id"`
	TITLE string `json:"title"`
	TASKS string `json:"task"`
}

// func inmemory() MemoryStorage {
var list = []TodoList{{ID: "0", TITLE: "Work", TASKS: "do some shit"},
	{ID: "1", TITLE: "Sleep", TASKS: "go sleep"},
	{ID: "2", TITLE: "Eat", TASKS: "it's ok"},
}

// 	return MemoryStorage{TodoLists: list}
// }

// var data = inmemory()

//	func (s Storage) Read() []TodoList {
//		return s.
//	}
func GetToDoLists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, list)
}

func PgConnection() {

}

func getRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", GetToDoLists)
	return router
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")
}
