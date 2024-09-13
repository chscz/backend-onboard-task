package router

import (
	"github.com/chscz/backend-onboard-task/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(uh *handler.UserHandler, ph *handler.PostHandler, localDebugMode bool) *gin.Engine {
	r := gin.Default()

	if localDebugMode {
		r.LoadHTMLGlob("templates/*")
	} else {
		r.LoadHTMLGlob("/app/templates/*")
	}

	// templateDir := os.Getenv("TEMPLATE_DIR")
	// if templateDir == "" {
	// 	templateDir = "templates/*"
	// }

	r.Use(AuthMiddleware(uh.Auth))

	r.GET("/", ph.GetPosts)

	r.GET("/login", uh.LoginPage)
	r.POST("/login", uh.Login)

	r.POST("/logout", uh.Logout)

	r.GET("/register", uh.RegisterPage)
	r.POST("/register", uh.Register)

	postGroup := r.Group("/post")
	{
		postGroup.GET("/create", ph.CreatePostPage)
		postGroup.POST("/create", ph.CreatePost)

		postGroup.GET("/update/:id", ph.UpdatePostPage)
		postGroup.POST("/update/:id", ph.UpdatePost)

		postGroup.GET("/delete/:id", ph.DeletePost)

		postGroup.GET("/detail/:id", ph.GetPostDetail)
	}

	return r
}
