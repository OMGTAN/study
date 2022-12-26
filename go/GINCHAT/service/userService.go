package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func CreateDB(c *gin.Context) {
	utils.DB.AutoMigrate(&models.UserBasic{})
	c.JSON(http.StatusOK, gin.H{
		"message": "create db success",
	})
}

// @Tags 用户管理
// @Accept json
// @param user body models.UserBasic  true "入参"
// @Success 200 {string} Helloworld
// @Router /createUser [post]
func CreateUser(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(&user)
	utils.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
}

// @Tags 用户管理
// @Accept json
// @param id query uint  false "入参"
// @param user body models.UserBasic  true "入参"
// @Success 200 {string} Helloworld
// @Router /updateUser [post]
func UpdateUser(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(&user)
	// id := user.ID
	id, _ := c.GetQuery("id")
	fmt.Println("id ==========", id, user)
	// user.ID = uint(id)
	utils.DB.Model(&user).Where("id", id).Updates(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "update user success",
	})
}

// @Tags 用户管理
// @Accept json
// @param id query uint  false "入参"
// @Success 200 {string} Helloworld
// @Router /deleteUser [get]
func DeleteUser(c *gin.Context) {
	id, _ := c.GetQuery("id")
	// user.ID = uint(id)
	utils.DB.Delete(&models.UserBasic{}, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
}

// @Tags 用户管理
// @param id query uint  false "入参"
// @Success 200
// @Router /listUsers [get]
func ListUsers(c *gin.Context) {
	// id, _ := c.GetQuery("id")
	var data []models.UserBasic
	utils.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if nil != err {
		fmt.Println("sendMsg error: ", err)
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if nil != err {
			fmt.Println("ws close error: ", err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if nil != err {
		fmt.Println(" MsgHandler error: ", err)
	}
	tm := time.Now().Format("2022-12-12 12:12:12")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if nil != err {
		fmt.Println(" MsgHandler error: ", err)
	}
}
