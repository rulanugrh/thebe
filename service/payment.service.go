package service

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	portRepo "be-project/repository/port"
	portService "be-project/service/port"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentService struct {
	repository         portRepo.PaymentInterface
	orderRepo portRepo.OrderRepository
	envMidtrans        midtrans.EnvironmentType
	serverKey          string
	s                  snap.Client
}

func NewPaymentService(repository portRepo.PaymentInterface, env midtrans.EnvironmentType, serverkey string, snaps snap.Client) portService.PaymentInterface {
	return &paymentService{
		repository:         repository,
		envMidtrans:        env,
		serverKey:          serverkey,
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
			FName: data.Orders.UserDetail.Name,
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
		Name:    data.Orders.UserDetail.Name,
		Event:   data.Orders.Events.Name,
		Price:   data.Orders.Events.Price,
		SnapURL: createTransaction.RedirectURL,
		Token:   createTransaction.Token,
	}

	save := domain.Transaction{
		Name: data.Orders.UserDetail.Name,
		Event: data.Orders.Events.Name,
		Price: data.Orders.Events.Price,
		SnapURL: createTransaction.RedirectURL,
		Token: createTransaction.Token,
	}

	errSave := payment.repository.Save(save)
	if errSave != nil {
		return nil, errSave
	}
	history, errHistory := payment.HandlingStatus(strconv.Itoa(req.OrderID))
	if errHistory != nil {
		return nil, web.Error{
			Message: "Tidak bisa mendapatkan status payment",
			Code: 400,
		}
	}

	switch history.StatusCode {
	case "200":
		req.Orders.StatusPayment = "success"
		err := payment.orderRepo.AppendToEvents(req)
		if err != nil {
			return nil, web.Error{
				Message: "Tidak bisa menambahkanmu ke dalam events",
				Code: 500,
			}
		}
	case "201":
		req.Orders.StatusPayment = "pending"
		return nil, web.Error{
			Message: "Pembayaran Pending, tunggu sebentar",
			Code: 201,
		}
	case "202":
		req.Orders.StatusPayment = "denied"
		return nil, web.Error{
			Message: "Pembayaran kamu ditolak",
			Code: 202,
		}

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
		Name:   data.Orders.UserDetail.Name,
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
			Name:   result.Orders.UserDetail.Name,
			Event:  result.Orders.Events.Name,
			Price:  result.Orders.Events.Price,
			Status: result.Orders.StatusPayment,
		}

		responsePayment = append(responsePayment, response)
	}

	return responsePayment, nil
}

func (payment *paymentService) HandlingStatus( id string ) (*web.StatusPayment, error){

	url := fmt.Sprintf("%s/%s/%s", payment.s.Env.BaseUrl(), id, "status")
	req, errRequest := http.NewRequest(http.MethodGet, url, nil)
	if errRequest != nil {
		return nil, web.Error{
			Message: "Tidak bisa request untuk mendapatkan status",
			Code: 400,
		}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(payment.s.ServerKey)))
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, web.Error{
			Message: "cant get status",
			Code: 400,
		}
	}

	defer response.Body.Close()

	status := web.StatusPayment{}
	errDecode := json.NewDecoder(response.Body).Decode(&status)
	if errDecode != nil {
		return nil, web.Error{
			Message: "Tidak bisa membaca response body",
			Code: 500,
		}
	}

	return &status, nil
}