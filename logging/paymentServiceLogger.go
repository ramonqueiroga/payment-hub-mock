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

//Capture logs the business logic of the capture payment service
func (lm LoggerMiddleware) Capture(p business.Payments) (transaction business.TransactionModel, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "capture",
			"input", p,
			"output", transaction.PaymentID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Capture(p)
	return
}

//Authorize logs the business logic of the authorize payment service
func (lm LoggerMiddleware) Authorize(s string) (transaction business.TransactionModel, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "authorize",
			"input", s,
			"output", transaction.PaymentID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Authorize(s)
	return
}

//Cancel logs the business logic of the cancel payment service
func (lm LoggerMiddleware) Cancel(s string) (transaction business.TransactionModel, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "cancel",
			"input", s,
			"output", transaction,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Cancel(s)
	return
}

//Search logs the business logic of the search payment service
func (lm LoggerMiddleware) Search(s string) (transaction business.TransactionModel, err error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"service", "search",
			"input", s,
			"output", transaction,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	transaction, err = lm.Next.Search(s)
	return
}
