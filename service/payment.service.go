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
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentService struct {
	repository         portRepo.PaymentInterface
	orderRepo portRepo.OrderRepository
	envMidtrans        midtrans.EnvironmentType
	serverKey          string
	s                  snap.Client
}

func NewPaymentService(repository portRepo.PaymentInterface, env midtrans.EnvironmentType, serverkey string, snaps snap.Client, orderRepo portRepo.OrderRepository) portService.PaymentInterface {
	return &paymentService{
		repository:         repository,
		envMidtrans:        env,
		serverKey:          serverkey,
		s:                  snaps,
		orderRepo: orderRepo,
	}
}

func (payment *paymentService) Create(req domain.Payment) (*web.ResponsePayment, error) {
	data, err := payment.repository.Create(req)
	if err != nil {
		return nil, err
	}

	transaction := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: data.OrderID,
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

	req, errRequest := http.NewRequest(http.MethodGet, payment.s.Env.BaseUrl() + "/v2/" + id + "/status", nil)
	if errRequest != nil {
		return nil, web.Error{
			Message: "Tidak bisa request untuk mendapatkan status",
			Code: 400,
		}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(payment.s.ServerKey + ":")))
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
			Message: errDecode.Error(),
			Code: 500,
		}
	}

	return &status, nil
}

func (payment *paymentService) NotificationStream(orderID string) (bool, error) {
	var client coreapi.Client
	client.New(payment.serverKey, payment.envMidtrans)
	var order domain.Order

	transactionResp, err := client.CheckTransaction(orderID)
	if err != nil {
		return false, web.Error{
			Message: "Order id tidak valid",
			Code: 400,
		}
	} else {
		if transactionResp != nil {
			if transactionResp.StatusCode == "200" {
				findOrder, _ := payment.orderRepo.FindByUUID(orderID)
				fmt.Println(findOrder)
				order.StatusPayment = transactionResp.FraudStatus
				order.Name = "Order" + orderID
				order.ID = findOrder.ID
				order.UserID = findOrder.UserID
				order.EventID = findOrder.EventID
				order.UUID = orderID

				
				_, errs := payment.orderRepo.Update(orderID, order)
				if errs != nil {
					return false, web.Error{
						Message: "Cant update status payment",
						Code: 500,
					}
				}
				return true, nil
			} else if transactionResp.TransactionStatus == "deny" {
				return false, web.Error{
					Message: "Pembayaran anda ditolak",
					Code: 400,
				}
			}
		}
	}

	return false, nil
}