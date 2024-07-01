package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/intern/models"
    "github.com/intern/database"
)

func UploadPhoto(c *gin.Context) {
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    photo := models.Photo{Title: input.Title, Caption: input.Caption, PhotoUrl: input.PhotoUrl, UserID: input.UserID}
    database.DB.Create(&photo)

    c.JSON(http.StatusOK, gin.H{"data": photo})
}
