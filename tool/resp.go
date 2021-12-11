package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)//处理因写入产生的错误
func RespErrorWithDate(ctx *gin.Context, date interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": date,
	})
}//服务器内部错误（注意不能将错误展现在服务器页面上，防止sql注入）
func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}
//回复成功响应
func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}
//
func RespSuccessfulWithDate(ctx *gin.Context, date interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"date": date,
	})
}
