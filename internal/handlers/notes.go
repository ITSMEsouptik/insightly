package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListNotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list notes - stub"})
}

func CreateNote(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "create note - stub"})
}

func GetNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get note - stub"})
}

func UpdateNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update note - stub"})
}

func DeleteNote(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "delete note - stub"})
}
