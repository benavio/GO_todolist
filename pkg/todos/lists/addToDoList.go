package todos

import (
	"github.com/gin-gonic/gin"
)

type lists struct {
	ListsNames []list `json:"listsnames"`
}
type list struct {
	ID       string `json:"id"`
	ListName string `json:"listname"`
	// тот кто дал задание
	// тот кто выполняет задание
}

func postList(c *gin.Context) {
	var newList list
	// id := c.Param("id")
	c.BindJSON(&newList)
	// album := CreateList(id, newList)
	// c.IndentedJSON(http.StatusCreated, album)
}

func CreateList(id string, list string) {

}
