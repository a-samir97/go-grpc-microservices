package ports

import "Github.com/a-samir97/microservices/payment/internals/application/core/domain"

type DBPort interface {
	Get(paymentId string) (domain.Payment, error)
	Save(payment *domain.Payment) error
}
