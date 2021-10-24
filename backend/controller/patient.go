package controller

import (
	"github.com/nitaxxix/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Patient

func CreatePatient(c *gin.Context) {

	var Patient entity.Patient

	if err := c.ShouldBindJSON(&Patient); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&Patient).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": Patient})

}

func GetPatient(c *gin.Context) {
	var Patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&Patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Patient})
}

// GET /Patient

func ListPatient(c *gin.Context) {

	var patients []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})

}
