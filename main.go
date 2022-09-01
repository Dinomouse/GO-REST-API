package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID 			string	`json:"id"` 
	Item 		string	`json:"title`
	Completed 	bool	`json:"completed`
}

var todos = []todo{
	{ID:"1", Item:"Clean Room", Completed: false},
	{ID:"2", Item:"Feed the dog", Completed: false},
	{ID:"3", Item:"Read a book", Completed: false},
}
//context contains information about the incoming http request
func getTodos(context *gin.Context){

}

func main(){
	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:9090")
}