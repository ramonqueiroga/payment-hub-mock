package repository

import (
	"payment-hub-mock/model"

	"github.com/jinzhu/gorm"
)

//Repository interface that represents the database logic
type Repository interface {
	Save(interface{}) (interface{}, error)
	FindOne(uint) (interface{}, error)
	FindAll() ([]interface{}, error)
	Delete(uint) (bool, error)
}

//TransactionRepository implements repository database logic for transaction model
type TransactionRepository struct {
	db *gorm.DB
}

//Save transcation model implementation
func (tr TransactionRepository) Save(trm model.TransactionModel) (model.TransactionModel, error) {
	return model.TransactionModel{}, nil
}

//FindOne transcation model implementation
func (tr TransactionRepository) FindOne(id uint) (model.TransactionModel, error) {
	return model.TransactionModel{}, nil
}

//FindAll transcation model implementation
func (tr TransactionRepository) FindAll() (model.TransactionModel, error) {
	return model.TransactionModel{}, nil
}

//Delete transcation model implementation
func (tr TransactionRepository) Delete(id uint) (model.TransactionModel, error) {
	return model.TransactionModel{}, nil
}
