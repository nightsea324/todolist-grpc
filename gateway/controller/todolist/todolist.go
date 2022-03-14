package todolist

import (
	"log"
	"net/http"
	"strings"

	"main/gateway/model/todolist"
	pb "main/gateway/protobuf/todolist"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var TodoClient pb.TodoClient

// New -
func New() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("初始化微服務錯誤")
	}

	// 啟動會員服務
	TodoClient = pb.NewTodoClient(conn)
}

// Create - 新增待辦事項
func Create(ctx *gin.Context) {

	var status string
	var msg string

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得資料
	req := new(todolist.CreateRequest)
	ctx.BindJSON(&req)
	req.MemberId = ctx.GetString("memberId")

	// 呼叫新增事項方法
	in := todolist.CreateRequestToPb(req)
	_, err := TodoClient.Create(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	status = "ok"
	msg = "已成功新增" + req.Name + "待辦事項"

	return
}

// Delete -
func Delete(ctx *gin.Context) {
	var status string
	var msg string

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得資料
	req := new(todolist.DeleteRequest)
	ctx.BindJSON(&req)
	req.ID = ctx.Param("id")
	req.MemberId = ctx.GetString("memberId")

	in := todolist.DeleteRequestToPb(req)
	_, err := TodoClient.Delete(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	status = "ok"
	msg = "已成功刪除待辦事項"

	return
}

// Update -
func Update(ctx *gin.Context) {

	var status string
	var msg string

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得資料
	req := new(todolist.UpdateRequest)
	ctx.BindJSON(&req)
	req.ID = ctx.Param("id")
	req.MemberId = ctx.GetString("memberId")

	in := todolist.UpdateRequestToPb(req)
	_, err := TodoClient.Update(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	status = "ok"
	msg = "已成功完成待辦事項"

	return
}

// Get -
func Get(ctx *gin.Context) {

	var status string
	var msg string
	var results []*todolist.Todolist

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"results": results,
			"status":  status,
			"message": msg,
		})
	}()

	// 取得使用者資料
	req := new(todolist.GetRequest)
	ctx.BindJSON(&req)
	req.MemberId = ctx.GetString("memberId")

	in := todolist.GetRequestToPb(req)
	out, err := TodoClient.Get(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	results = todolist.ConvertTodoListToModel(out.GetTodoList())
	status = "ok"
	msg = "已查詢使用者待辦事項"

	return
}
