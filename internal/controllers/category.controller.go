package controllers

import (
	"net/http"
	"stockhive-server/internal/models"
	"stockhive-server/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service}
}

func (ctl *CategoryController) GetAll(c *gin.Context) {
	Categorys, err := ctl.service.GetAllCategorys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Categorys"})
		return
	}
	c.JSON(http.StatusOK, Categorys)
}

func (ctl *CategoryController) GetByID(c *gin.Context) {
	id := c.Param("id")
	Category, err := ctl.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, Category)
}

func (ctl *CategoryController) Create(c *gin.Context) {
	var Category models.Category
	if err := c.ShouldBindJSON(&Category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := ctl.service.CreateCategory(&Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (ctl *CategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	Category, err := ctl.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&Category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := ctl.service.UpdateCategory(&Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (ctl *CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	Category, err := ctl.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := ctl.service.DeleteCategory(&Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}


