package controllers

import (
	"sagara-msib-test/models"
	"sagara-msib-test/services"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ClothController struct {
	service *services.ClothService
}

func NewClothController(service *services.ClothService) *ClothController {
	return &ClothController{service: service}
}

func (c *ClothController) CreateCloth(ctx *gin.Context) {
	var cloth models.Cloth
	cloth.CreatedAt = time.Now()

	if err := ctx.ShouldBindJSON(&cloth); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCloth(cloth); err != nil {	
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Cloth created successfully"})
}

func (c *ClothController) GetCloth(ctx *gin.Context) {
	clothID, _ := strconv.Atoi(ctx.Param("id"))

	cloth, err := c.service.GetCloth(uint(clothID))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": cloth})
}

func (c *ClothController) GetClothes(ctx *gin.Context) {
	color := ctx.Query("color")
	size := ctx.Query("size")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	clothes, err := c.service.GetClothes(color, size, page, pageSize)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": clothes})
}

func (c *ClothController) UpdateCloth(ctx *gin.Context) {
	clothID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid cloth ID"})
		return
	}

	_, err = c.service.GetCloth(uint(clothID))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cloth not found!"})
		return
	}

	var updatedCloth models.Cloth
	if err := ctx.ShouldBindJSON(&updatedCloth); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCloth.ID = uint(clothID)
	updatedCloth.UpdatedAt = time.Now()

	err = c.service.UpdateCloth(updatedCloth)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cloth", "details": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Cloth updated successfully"})
}

func (c *ClothController) IncrementClothStock(ctx *gin.Context) {
    clothID, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid cloth ID"})
        return
    }

    cloth, err := c.service.GetCloth(uint(clothID))
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cloth not found!"})
        return
    }

    cloth.Stock++

    err = c.service.UpdateCloth(cloth)
    if err != nil {
        ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment cloth", "details": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Cloth incremented successfully"})
}

func (c *ClothController) DecrementClothStock(ctx *gin.Context) {
    clothID, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid cloth ID"})
        return
    }

    cloth, err := c.service.GetCloth(uint(clothID))
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cloth not found!"})
        return
    }

	if cloth.Stock <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock is already 0"})
		return
	}

    cloth.Stock--

    err = c.service.UpdateCloth(cloth)
    if err != nil {
        ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrement cloth", "details": err.Error()})
        return
    }

    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Cloth decremented successfully"})
}

func (c *ClothController) DeleteCloth(ctx *gin.Context) {
	clothID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid cloth ID"})
		return
	}

	_, err = c.service.GetCloth(uint(clothID))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cloth not found!"})
		return
	}

	if err := c.service.DeleteCloth(uint(clothID)); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Cloth deleted successfully"})
}

func (c *ClothController) FindOutOfStockClothes(ctx *gin.Context) {
	color := ctx.Query("color")
	size := ctx.Query("size")

	clothes, err := c.service.FindOutOfStockClothes(color, size)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": clothes})
}

func (c *ClothController) FindLowStockClothes(ctx *gin.Context) {
	color := ctx.Query("color")
	size := ctx.Query("size")

	clothes, err := c.service.FindLowStockClothes(color, size)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": clothes})
}
