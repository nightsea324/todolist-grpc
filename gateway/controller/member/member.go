package member

import (
	"log"
	"net/http"
	"strings"

	"main/gateway/model/member"
	pb "main/gateway/protobuf/member"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var MemberClient pb.MemberClient

// New -
func New() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("初始化微服務錯誤")
	}

	// 啟動會員服務
	MemberClient = pb.NewMemberClient(conn)
}

// Register - 註冊會員
func Register(ctx *gin.Context) {

	var status string
	var msg string

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status": status,
			"msg":    msg,
		})
	}()

	// 取得註冊資料
	req := new(member.RegisterMemberRequest)
	ctx.BindJSON(&req)

	// 呼叫會員註冊方法
	in := member.RegisterMemberRequestToPb(req)
	_, err := MemberClient.Register(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	status = "ok"
	msg = "註冊成功"

	return
}

// SignIn - 登入會員
func SignIn(ctx *gin.Context) {

	var status string
	var msg string

	defer func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得登入資訊
	req := new(member.SignInMemberRequest)
	ctx.BindJSON(&req)

	// 呼叫會員登入方法
	in := member.SignInMemberRequestToPb(req)
	res, err := MemberClient.SignIn(ctx, in)
	if err != nil {
		errCode := strings.Split(err.Error(), " ")
		status = "failed"
		msg = errCode[len(errCode)-1]
		return
	}

	// 設置cookie
	ctx.SetCookie("token", res.Token, int(res.TokenLife), "/", "localhost", false, true)

	status = "ok"
	msg = "登入成功"

	return
}
