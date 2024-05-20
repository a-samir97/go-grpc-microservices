package api

import (
	"Github.com/a-samir97/microservices/payment/internals/application/core/domain"
	"Github.com/a-samir97/microservices/payment/internals/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) Charge(payement domain.Payment) (domain.Payment, error) {
	err := a.db.Save(&payement)

	if err != nil {
		return domain.Payment{}, err
	}
	return payement, nil
}
