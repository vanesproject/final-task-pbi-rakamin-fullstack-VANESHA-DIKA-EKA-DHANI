package router

import (
    "github.com/gin-gonic/gin"
    "github.com/intern/controllers"
    "github.com/intern/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    r.GET("/register", func(c *gin.Context) {
        c.HTML(http.StatusOK, "register.html", nil)
    })

    r.GET("/upload", middlewares.AuthMiddleware(), func(c *gin.Context) {
        c.HTML(http.StatusOK, "upload.html", nil)

    r.POST("/users/register", controllers.Register)
    r.POST("/users/login", controllers.Login)
    r.POST("/photos", middlewares.AuthMiddleware(), controllers.UploadPhoto)

    return r
}
