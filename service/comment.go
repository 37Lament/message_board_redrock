package service//服务包，注释中

import (
	"message-board/dao"
	"message-board/model"
)
//加评论
func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}
//发送评论
func GetPostComments(postId int) ([]model.Comment, error) {
	return dao.SelectCommentByPostId(postId)
}
