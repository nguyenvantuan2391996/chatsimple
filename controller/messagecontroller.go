package controller

import (
	"github.com/gin-gonic/gin"
	"simpleChat/database"
	"simpleChat/entities"
)

func UpdateMessageNoRead(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.SenderReceive
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		var listMessage []entities.Message
		result := db.Table("message").Joins("JOIN sender_receive ON message.id = sender_receive.id AND sender_receive.sender_id = ? AND sender_receive.receive_id = ?", json.ReceiveId, json.SenderId).Find(&listMessage)

		// Update message
		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		} else {
			for i := 0; i < len(listMessage); i++ {
				listMessage[i].IsRead = "1"
				db.Model(&listMessage[i]).Update("is_read", "1")
			}

			c.JSON(200, gin.H{
				"messages": "Success",
			})
		}
	}

	defer db.Close()
}

func CountMessageNoRead(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var listCountMessageNoRead []int64
	var json entities.ListId
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		for i := 0; i < len(json.ListUserId); i++ {
			var count int64
			result := db.Table("message").Joins("JOIN sender_receive ON message.id = sender_receive.id AND sender_receive.sender_id = ? AND sender_receive.receive_id = ? AND message.is_read = ?", json.ListUserId[i], json.SenderId, "0").Count(&count)

			if result.Error != nil {
				c.JSON(500, gin.H{
					"messages": result.Error,
				})
			} else {
				listCountMessageNoRead = append(listCountMessageNoRead, count)
			}
		}

		c.JSON(200, listCountMessageNoRead)
	}

	defer db.Close()
}

func GetMessage(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.SenderReceive
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		var listMessage []entities.Message
		result := db.Table("message").Joins("JOIN sender_receive ON message.id = sender_receive.id").Where("sender_receive.sender_id = ? AND sender_receive.receive_id = ?", json.SenderId, json.ReceiveId).Or("sender_receive.sender_id = ? AND sender_receive.receive_id = ?", json.ReceiveId, json.SenderId).Find(&listMessage)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		} else {
			c.JSON(200, listMessage)
		}
	}

	defer db.Close()
}

func AddFirstMessage(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.DataChat
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		var _senderReceive entities.SenderReceive
		result := db.Where("sender_id = ? AND receive_id = ?", json.SenderReceive.SenderId, json.SenderReceive.ReceiveId).Find(&_senderReceive)

		if result.Error != nil && result.RowsAffected != 0 {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		} else {
			if result.RowsAffected == 0 {
				// add table message
				db.Create(&json.MessageChat)

				// add table sender_receive
				json.SenderReceive.MessageId = json.MessageChat.Id
				db.Create(&json.SenderReceive)
			}
			c.JSON(200, "Success")
		}
	}

	defer db.Close()
}

func AddMessage(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.DataChat

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		// add table message
		resultMessage := db.Create(&json.MessageChat)

		if resultMessage.Error != nil {
			c.JSON(500, gin.H{
				"messages": resultMessage.Error,
			})
		}

		// add table sender_receive
		json.SenderReceive.MessageId = json.MessageChat.Id
		result := db.Create(&json.SenderReceive)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		}

		c.JSON(200, "Success")
	}

	defer db.Close()
}

func GetReceiver(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.SenderReceive

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		var listSenderReceive []entities.SenderReceive
		var listId []int
		result := db.Where("sender_id = ?", json.SenderId).Order("id desc").Find(&listSenderReceive)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		} else {
			for i := 0; i < len(listSenderReceive); i++ {
				listId = append(listId,listSenderReceive[i].ReceiveId)
			}
			c.JSON(200, listId)
		}
	}

	defer db.Close()
}
