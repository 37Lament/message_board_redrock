package dao//注释完毕

import "message-board/model"
//插入评论
func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(username, txt, comment_time, post_id) "+"values(?, ?, ?, ?);", comment.Username, comment.Txt, comment.CommentTime, comment.PostId)
	return err
}
//对选择的ID进行查询并评论
func SelectCommentByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment
	//在数据库中查找
	rows, err := dB.Query("SELECT id, post_id, txt, username, comment_time FROM comment WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	//关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Txt, &comment.Username, &comment.CommentTime)
		if err != nil {
			return nil, err
		}
	//没有问题就加入切片
		comments = append(comments, comment)
	}
	//返回评论和nil
	return comments, nil
}
