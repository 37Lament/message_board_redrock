package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆
	engine.POST("ping",Ping)//检查网络连接
	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码
	}

	postGroup := engine.Group("/post")
	{
		postGroup.Use(auth)
		postGroup.POST("/", addPost) //发布新留言
		postGroup.POST("/:post_id")  //修改留言
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/", addComment)  //发送评论
		commentGroup.DELETE("/:comment_id") //删除评论
	}
	viewGroup := engine.Group("/view")//游客模式可使用的
	{
		viewGroup.GET("/", briefPosts)         //查看全部留言概略
		viewGroup.GET("/:post_id", postDetail) //查看一条留言详细信息和其下属评论

	}
	adGroup:=engine.Group("ad")
	{
		adGroup.GET("/",Admin)
	}

	engine.Run()
}
