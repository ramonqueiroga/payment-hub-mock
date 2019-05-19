package repository

import (
	"fmt"
	"math/rand"
	"payment-hub-mock/business"
	"payment-hub-mock/model"
	"time"

	"github.com/jinzhu/gorm"
)

//TransactionRepository implements repository database logic for transaction model
type TransactionRepository struct {
	Db *gorm.DB
}

//Save transcation model implementation
func (tr TransactionRepository) Save(p business.Payments) (bool, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	paymentID := fmt.Sprintf("%s%d", "10000", r1)
	transModel := model.TransactionModel{
		PaymentID:    paymentID,
		Amount:       p.Payments[0].Amount,
		Status:       "AUTHORIZE",
		Installments: p.Payments[0].Installments,
	}

	tr.Db.AutoMigrate(&model.TransactionModel{})

	fmt.Println("saving transaction model", transModel)
	tr.Db.Save(&transModel)

	return true, nil
}

//FindOne transcation model implementation
func (tr TransactionRepository) FindOne(id uint64) (interface{}, error) {
	var findTransModel model.TransactionModel
	tr.Db.Where("ID = ?", id).First(&findTransModel)

	fmt.Println("finding", findTransModel)
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
