package controller

import (
	"be-golang-chapter-56/product-service/models"
	"be-golang-chapter-56/product-service/repository"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	repo repository.ProductRepository
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{repo: repository.NewProductRepository(db)}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.repo.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := ctrl.repo.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
