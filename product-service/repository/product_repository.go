package repository

import (
	"be-golang-chapter-56/product-service/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}
