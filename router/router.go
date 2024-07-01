package router

import (
    "github.com/gin-gonic/gin"
    "github.com/intern/controllers"
    "github.com/intern/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/users/register", controllers.Register)
    r.POST("/photos", middlewares.AuthMiddleware(), controllers.UploadPhoto)

    return r
}
