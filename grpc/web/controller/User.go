package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc/web/model"
	"net/http"
	"strconv"
	"time"
)

type response struct {
	Msg string

	Data interface{}
}

func Index(ctx *gin.Context) {
	fmt.Println("hello world")

	data := make(map[string]interface{})

	data["id"] = 1
	data["name"] = "xmcx_5235"
	data["timestamp"] = time.Now().Unix()
	resp := make(map[string]interface{})
	resp["msg"] = "ok"
	resp["data"] = data
	ctx.JSON(0, resp)
}

func PostUserAvatar(ctx *gin.Context) {
	fmt.Println("start post avatar")
	file, _ := ctx.FormFile("avatar")
	uuid := string(time.Now().Unix())

	err := ctx.SaveUploadedFile(file, "./test/avatar/"+uuid+file.Filename)

	response := new(response)
	code := 0
	if err != nil {
		code = 1
		response.Msg = "error"
		response.Data = ""
	} else {
		response.Msg = "ok"
		response.Data = ""
	}

	member := new(model.Member)
	/*member.Nickname = "fsdf"
	member.Email = "fsdfsa"
	member.Level = 1*/
	//model.DbCon.First(member)

	fmt.Println(member)
	ctx.JSON(code, response)
}

func SelectUser(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.PostForm("id"))

	member := model.SelectMember(id)

	resp := make(map[string]interface{})

	resp["errorno"] = 0
	resp["data"] = member
	resp["msg"] = "success"

	ctx.JSON(http.StatusOK, resp)
}