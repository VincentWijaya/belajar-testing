package repository

import "belajar-testing/entity"

type ProductRepository interface {
	FindByID(id string) *entity.Product
}
