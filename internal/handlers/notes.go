package handlers

import (
	"net/http"

	"github.com/ItsMeSouptik/insightly/internal/db"
	"github.com/ItsMeSouptik/insightly/internal/models"
	"github.com/gin-gonic/gin"
)

func ListNotes(c *gin.Context) {
	var notes []models.Note
	db.DB.Find(&notes)
	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if note.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	if note.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}
	db.DB.Create(&note)
	c.JSON(http.StatusCreated, note)
}

func GetNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := db.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := db.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	var update models.Note
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	note.Title = update.Title
	note.Content = update.Content
	db.DB.Save(&note)
	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	db.DB.Delete(&models.Note{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
