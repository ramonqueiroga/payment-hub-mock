package repository

import (
	"fmt"
	"payment-hub-mock/business"
	"payment-hub-mock/model"

	"github.com/jinzhu/gorm"
)

//TransactionRepository implements repository database logic for transaction model
type TransactionRepository struct {
	Db *gorm.DB
}

//Save transcation model implementation
func (tr TransactionRepository) Save(p business.Payments) (bool, error) {
	fmt.Print("p", p)

	transModel := model.TransactionModel{
		PaymentID: "1",
		Amount:    1.0,
		Status:    "CAPTURADO",
	}

	tr.Db.AutoMigrate(&model.TransactionModel{})
	tr.Db.Create(&transModel)

	var findTransModel model.TransactionModel
	tr.Db.First(&findTransModel, 1)

	fmt.Print("paymentId", findTransModel.PaymentID)
	return true, nil
}

//FindOne transcation model implementation
func (tr TransactionRepository) FindOne(id uint64) (interface{}, error) {
	var findTransModel model.TransactionModel
	tr.Db.First(&findTransModel, 1)
	return findTransModel, nil
}

//FindAll transcation model implementation
func (tr TransactionRepository) FindAll() ([]interface{}, error) {
	return nil, nil
}

//Delete transcation model implementation
func (tr TransactionRepository) Delete(id uint64) (bool, error) {
	return true, nil
}
