package main

import (
	"github.com/gin-gonic/gin"
	"simpleChat/controller"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		// API user
		api.GET("/user", controller.GetUser)
		api.POST("/user/getbylistid", controller.GetUserByListId)
		api.POST("/user", controller.CreateUser)

		// API message
		api.PUT("/message/updatemessagenoread", controller.UpdateMessageNoRead)
		api.POST("/message/countmessagenoread", controller.CountMessageNoRead)
		api.POST("/message/getmessage", controller.GetMessage)
		api.POST("/message/addfirstmessage", controller.AddFirstMessage)
		api.POST("/message/addmessage", controller.AddMessage)
		api.POST("/message/getreceiver", controller.GetReceiver)
	}
	r.Run()
}
