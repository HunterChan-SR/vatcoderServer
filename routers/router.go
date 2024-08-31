package routers

import (
	"VAtcoder/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func Router() *gin.Engine {
	root := gin.Default()
	root.Use(CorsMiddleware())

	r := root.Group("/api")
	{
		admin := r.Group("/admin")
		{
			admin.POST("/reload", controllers.AdminController{}.ReLoad)
			admin.GET("/levels", controllers.AdminController{}.GetLevels)
			admin.GET("/account/:level", controllers.AdminController{}.GetAccount)
			admin.POST("/addproblem", controllers.AdminController{}.AddProblem)
			admin.POST("/adduser", controllers.AdminController{}.AddUser)
		}

		submit := r.Group("/submit")
		{
			submit.POST("/", controllers.SubmitController{}.Submit)
			submit.GET("/data/:id", controllers.SubmitController{}.Data)
		}

		user := r.Group("/user")
		{
			user.POST("/login", controllers.UserController{}.PostLogin)
			user.GET("/online", controllers.UserController{}.GetOnline)
			user.POST("/pwd", controllers.UserController{}.PostPwd)
		}
		problem := r.Group("/problem")
		{
			problem.GET("/", controllers.ProblemController{}.Get)
		}
	}

	return root
}
