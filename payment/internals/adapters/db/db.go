package db

import (
	"fmt"

	"Github.com/a-samir97/microservices/payment/internals/application/core/domain"
	"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(id string) (domain.Payment, error) {
	var paymentEntity Payment

	res := a.db.First(&paymentEntity, id)

	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderID:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnitNano(),
	}

	return payment, res.Error
}

func (a Adapter) Save(payment domain.Payment) error {
	paymentModel := Payment{
		CustomerID: payment.CustomerID,
		TotalPrice: payment.TotalPrice,
		OrderID:    payment.OrderID,
		Status:     payment.Status,
	}

	res := a.db.Create(&paymentModel)

	if res.Error == nil {
		payment.ID = int64(paymentModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	dbObject, openError := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})

	if openError != nil {
		return nil, fmt.Errorf("db connection error :%v", openError)
	}

	err := dbObject.AutoMigrate(&Payment{})

	if err != nil {
		return nil, fmt.Errorf("db migration error %v: ", err)
	}
	return &Adapter{db: dbObject}, nil
}
