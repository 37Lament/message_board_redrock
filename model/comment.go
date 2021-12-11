package model

import "time"//一个评论的结构体

type Comment struct {
	Id          int
	PostId      int
	Txt         string//文本
	Username    string//用户名
	CommentTime time.Time//评论时间
}
