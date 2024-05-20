package ports

import (
	"Github.com/a-samir97/microservices/payment/internals/application/core/domain"
)

type APIPort interface {
	Charge(payment domain.Payment) (domain.Payment, error)
}
