package main

import (
    "github.com/gin-gonic/gin"
    "github.com/intern/database"
    "github.com/intern/router"
)

func main() {
    database.ConnectDatabase()
    r := router.SetupRouter()
    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/*")
    r.Run(":8080")
}
