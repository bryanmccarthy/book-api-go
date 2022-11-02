package books

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "example.com/book-api-go/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {
    var books []models.Book

    if result := h.DB.Find(&books); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &books)
}
