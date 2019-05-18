package logging

import (
	"payment-hub-mock/business"
	"time"

	"github.com/go-kit/kit/log"
)

//LoggerMiddleware logs the business logic information
type LoggerMiddleware struct {
	Logger log.Logger
	Next   business.PaymentService
}

//Authorize logs the business logic of the authorize payment service
func (lm LoggerMiddleware) Authorize(payments business.Payments) (transaction business.Transaction, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "authorize",
			"input", payments,
			"output", transaction.PaymentID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Authorize(payments)
	return
}

//Capture logs the business logic of the capture payment service
func (lm LoggerMiddleware) Capture(PaymentID uint64) (transaction business.Transaction, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "capture",
			"input", PaymentID,
			"output", transaction.PaymentID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Capture(PaymentID)
	return
}

//Cancel logs the business logic of the cancel payment service
func (lm LoggerMiddleware) Cancel(p uint64) (transaction business.Transaction, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "cancel",
			"input", p,
			"output", transaction,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Cancel(p)
	return
}

//Search logs the business logic of the search payment service
func (lm LoggerMiddleware) Search(p uint64) (transaction business.Transaction, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "search",
			"input", p,
			"output", transaction,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Search(p)
	return
}
