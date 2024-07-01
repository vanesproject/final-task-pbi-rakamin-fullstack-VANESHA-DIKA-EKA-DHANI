package main

import (
    "github.com/intern/router"
    "github.com/intern/database"
)

func main() {
    database.ConnectDatabase()

    r := router.SetupRouter()
    r.Run()
}
