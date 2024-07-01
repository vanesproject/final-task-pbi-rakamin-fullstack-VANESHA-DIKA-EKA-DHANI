package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/intern/models"
    "github.com/intern/database"
    "github.com/intern/helpers"
)

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := helpers.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user := models.User{Username: input.Username, Email: input.Email, Password: hashedPassword}
    database.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}
