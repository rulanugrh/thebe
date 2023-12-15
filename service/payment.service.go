package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"log"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentService struct {
	repository         portRepo.PaymentInterface
	envMidtrans        midtrans.EnvironmentType
	serverKey          string
	// paymentAppendUrl   string
	// paymentOverrideUrl string
	s                  snap.Client
}

func NewPaymentService(repository portRepo.PaymentInterface, env midtrans.EnvironmentType, serverkey string, snaps snap.Client) portService.PaymentInterface {
// midtrans.SetPaymentAppendNotification(paymentAppendUrl)
	// midtrans.SetPaymentOverrideNotification(paymentOverride)

	return &paymentService{
		repository:         repository,
		envMidtrans:        env,
		serverKey:          serverkey,
		// paymentAppendUrl:   paymentAppendUrl,
		// paymentOverrideUrl: paymentOverride,
		s:                  snaps,
	}
}

func (payment *paymentService) Create(req domain.Payment) (*web.ResponsePayment, error) {
	data, err := payment.repository.Create(req)
	if err != nil {
		return nil, err
	}

	transaction := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(data.OrderID),
			GrossAmt: int64(data.Orders.Events.Price),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.Orders.UserDetail.FName,
			LName: data.Orders.UserDetail.LName,
			Email: data.Orders.UserDetail.Email,
			Phone: data.Orders.UserDetail.Telephone,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.Itoa(int(data.Orders.Events.ID)),
				Price: int64(data.Orders.Events.Price),
				Qty:   1,
				Name:  data.Orders.Events.Name,
			},
		},
	}

	// payment.s.Options.SetPaymentOverrideNotification(payment.paymentOverrideUrl)
	// payment.s.Options.SetPaymentAppendNotification(payment.paymentAppendUrl)
	payment.s.Env = payment.envMidtrans
	payment.s.ServerKey = payment.serverKey

	createTransaction, errTransaction := payment.s.CreateTransaction(&transaction)
	if errTransaction != nil {
		return nil, web.Error{
			Message: errTransaction.GetMessage(),
			Code: 400,
		}
	}

	responseData := web.ResponsePayment{
		Name:    data.Orders.UserDetail.FName + " " + data.Orders.UserDetail.LName,
		Event:   data.Orders.Events.Name,
		Price:   data.Orders.Events.Price,
		SnapURL: createTransaction.RedirectURL,
		Token:   createTransaction.Token,
	}

	save := domain.Transaction{
		Name: data.Orders.UserDetail.FName + " " + data.Orders.UserDetail.LName,
		Event: data.Orders.Events.Name,
		Price: data.Orders.Events.Price,
		SnapURL: createTransaction.RedirectURL,
		Token: createTransaction.Token,
	}

	errSave := payment.repository.Save(save)
	if errSave != nil {
		return nil, errSave
	}

	return &responseData, nil
}
func (payment *paymentService) FindByID(id string) (*web.ResponseForPayment, error) {
	data, err := payment.repository.FindByID(id)
	if err != nil {
		log.Printf("Cannot find the order")
		return nil, err
	}

	response := web.ResponseForPayment{
		Name:   data.Orders.UserDetail.FName + " " + data.Orders.UserDetail.FName,
		Event:  data.Orders.Events.Name,
		Price:  data.Orders.Events.Price,
		Status: data.Orders.StatusPayment,
	}

	return &response, nil
}

func (payment *paymentService) FindAll() ([]web.ResponseForPayment, error) {
	data, err := payment.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responsePayment []web.ResponseForPayment
	for _, result := range data {
		response := web.ResponseForPayment{
			Name:   result.Orders.UserDetail.FName + " " + result.Orders.UserDetail.LName,
			Event:  result.Orders.Events.Name,
			Price:  result.Orders.Events.Price,
			Status: result.Orders.StatusPayment,
		}

		responsePayment = append(responsePayment, response)
	}

	return responsePayment, nil
}
