package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"simpleChat/database"
	"simpleChat/entities"
)

func hashMd5(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func checkUserName(username string) bool {
	db := database.DBConn()

	var user entities.User
	result := db.Where("user_name = ?", username).First(&user)

	// records count
	if result.Error != nil && result.RowsAffected == 0 {
		return false
	}

	return true
}

func CreateUser(c *gin.Context) {
	db := database.DBConn()

	// Struct parameter request from client
	var json entities.User
	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		// Check user name is existed ?
		if !checkUserName(json.UserName) {
			var userCreate entities.User

			userCreate.UserName = json.UserName
			userCreate.Password = hashMd5(json.Password)
			userCreate.Name = json.Name

			result := db.Create(&userCreate)

			if result.Error != nil {
				c.JSON(500, gin.H{
					"messages": result.Error,
				})
			} else {
				c.JSON(200, gin.H{
					"messages": userCreate.Id,
				})
			}
		} else {
			c.JSON(500, gin.H{
				"messages": "User name already exists",
			})
		}
	}

	defer db.Close()
}

func GetUser(c *gin.Context) {

	db := database.DBConn()
	var listUser []entities.User
	err := db.Find(&listUser).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "List user is empty",
		})
	} else {
		c.JSON(200, listUser)
	}

	defer db.Close()
}

func GetUserByListId(c *gin.Context) {
	db := database.DBConn()

	var listId entities.ListId
	var listUser []entities.User

	// Fetch value from request body
	err := c.ShouldBindJSON(&listId)

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	} else {
		result := db.Find(&listUser, listId.ListUserId)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
		} else {
			c.JSON(200, listUser)
		}
	}

	defer db.Close()
}
