package service

import (
	"belajar-testing/entity"
	"belajar-testing/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (s ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := s.Repository.FindByID(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
