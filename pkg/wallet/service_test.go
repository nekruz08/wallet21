package wallet

import (
	// "cmd/go/internal/fmtcmd"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/nekruz08/wallet/pkg/types"
)
// func TestService_Reject_success(t *testing.T) {
// 	// создаём сервис
// 	s := &Service{}

// 	// регистрируем там пользователья 
// 	phone:=types.Phone("+992000000001")
// 	account,err:=s.RegisterAccount(phone)
// 	if err != nil {
// 		t.Errorf("Reject(): can't register account, error = %v",err)
// 		return
// 	}		

// 	// пополняем его счёт
// 	err=s.Deposit(account.ID,10_000_00)
// 	if err != nil {
// 		t.Errorf("Reject(): can't deposit account, error = %v",err)
// 		return
// 	}

// 	// осуществляем платеж на его счет
// 	payment,err:=s.Pay(account.ID,1000_00,"auto")
// 	if err != nil {
// 		t.Errorf("Reject(): can't create payment, error = %v",err)
// 		return
// 	}

// 	// попробуем отменить платёж
// 	err=s.Reject(payment.ID)
// 	if err != nil {
// 		t.Errorf("Reject(): error = %v",err)
// 		return
// 	}
	
// 	// А как проверить статус платежа?
// 	// И баланс аккаунта?
// }
//**********************************************************************************
// func TestService_FindPaymentByID_success(t *testing.T) {
// 	// создаём сервис
// 	s:=&Service{}

// 	// регистрируем там ползователья 
// 	phone:=types.Phone("+992000000001")
// 	account,err:=s.RegisterAccount(phone)
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): can't register account, error = %v", err)
// 		return
// 	}

// 	// пополняем его счёт
// 	err=s.Deposit(account.ID,10_000_00)
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): can't deposit account, error=%v",err)
// 		return
// 	}

// 	// осуществляем платеж на его счёт
// 	payment, err:=s.Pay(account.ID,1000_00,"auto")
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): can't create payment, error = %v", err)
// 		return
// 	}

// 	// попробуем найти платёж
// 	got, err:=s.FindPaymentByID(payment.ID)
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): error = %v",err)
// 	}

// 	// сравниваем платежи
// 	if reflect.DeepEqual(payment,got){
// 		t.Errorf("FindPaymetByID(): wrong payment returned = %v",err)
// 	}
// }
//**********************************************************************************
// func TestService_FindPaymentByID_success(t *testing.T) {
// 	// создаём сервис
// 	s:=testService{}
// 	account,err:=s.addAccountWithBalance("+99200000000",10_000_00)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	// осуществляем платеж на его счёт
// 	payment, err:=s.Pay(account.ID,1000_00,"auto")
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): can't create payment, error = %v", err)
// 		return
// 	}

// 	// попробуем найти платёж
// 	got, err:=s.FindPaymentByID(payment.ID)
// 	if err != nil {
// 		t.Errorf("FindPaymetByID(): error = %v",err)
// 	}

// 	// сравниваем платежи
// 	if reflect.DeepEqual(payment,got){
// 		t.Errorf("FindPaymetByID(): wrong payment returned = %v",err)
// 	}
// }
//**********************************************************************************
// func TestService_FindPaymentByID_fail(t *testing.T) {
// 	// создаем сервис
// 	s:=newTestService()
// 	account,err:=s.addAccountWithBalance("+992000000001",10_000_00)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
	
// 	// осуществляем платеж на его счёт 
// 	payment, err :=s.Pay(account.ID,1000_00,"auto")
// 	if err != nil {
// 		t.Errorf("FindPaymentByID(): can't create payment, error = %v", err)
// 		return
// 	}

// 	// попробуем найти несуществующий платеж
// 	got, err:=s.FindPaymentByID(uuid.New().String())
// 	if err == nil {
// 		t.Error("FindPaymentByID(): must return error, returned nil")
// 		return
// 	}
// 	if err != ErrPaymentNotFound {
// 		t.Errorf("FindPaymentByID(): must return ErrPaymentNotFound, returned = %v", err )
// 		return
// 	}

// 	// сравниваем платежи
// 	if reflect.DeepEqual(payment,got){
// 		t.Errorf("FindPaymetByID(): wrong payment returned = %v",err)
// 	}
// }
//**********************************************************************************1
type testService struct{
	*Service	//embedding (встраивание)
}

func newTestService() *testService{		//функция-конструктор
	return &testService{Service:&Service{}}
}
// func (s *testService) addAccountWithBalance(phone types.Phone, balance types.Money)	(*types.Account,error)  {
// 	// регистрируем там пользователя
// 	account, err:=s.RegisterAccount(phone)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't register account, error = %v",err)
// 	}

// 	// пополняем его счёт
// 	err=s.Deposit(account.ID,balance)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't deposit account, error = %v",err)
// 	}
// 	return account,nil
// }
// //**********************************************************************************
type testAccount struct{
	phone types.Phone
	balance types.Money
	payments []struct{
		amount types.Money
		category types.PaymentCategory
	}
}
// Метод, который генерирует тестовы аккаунт и платеж  
func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	// регистрируем там пользователя
	account,err:=s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	// пополняем его счёт 
	err=s.Deposit(account.ID,data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposity account, error = %v", err)
	}

	// выполняем платежи
	// можем создать слайс нужной длины, поскольку знаем размер
	payments:=make([]*types.Payment,len(data.payments))
	for i, payment := range data.payments {
		//тогда здесь работаем через index, а не через append
		payments[i],err=s.Pay(account.ID,payment.amount,payment.category)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v",err)
		}
	}
	return account, payments,nil
}
// //**********************************************************************************
var	defaultTestAccount=testAccount{
	phone: "+992000000001",
	balance: 10_000_00,
	payments: []struct{
		amount types.Money
		category types.PaymentCategory
	}{
	{amount:10_000_00, category: "auto" },
	},
}
func TestService_FindPaymentByID_success(t *testing.T)  {
	// создаём сервис
	s:=newTestService()
	_,payments,err:=s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	
	// попробуем найти платеж
	payment:=payments[0]
	got,err:=s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error=%v",err)
		return
	}

	// сравниваем платежи
	if !reflect.DeepEqual(payment,got) {
		t.Errorf("FindPaymentByID(): error=%v",err)
		return
	}
}
// //**********************************************************************************
func TestService_FindPaymentByID_fail(t *testing.T) {
	// создаём сервис
	s:=newTestService()
	_,_,err:=s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// попробуем найти несуществующий платёж
	_,err=s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Error("FindPaymentByID(): must return error, returned nil")
		return
	}

	if err != ErrPaymentNotFound {
		t.Errorf("FindPaymentByID(): must return ErrPaymentNotFound, returned = %v",err)
		return
	}
}
//**********************************************************************************
func TestService_Reject_success(t *testing.T) {
	// создаём сервис
	s := newTestService()
	_,payments,err:=s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// попробуем отменить платёж
	payment:=payments[0]
	err=s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v",err)
		return
	}

	savedPayment, err:=s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("Reject(): can't find payment by id, error = %v",err)
		return
	}
	if savedPayment.Status!=types.PaymentStatusFail{
		t.Errorf("Reject(): status didn't changed, payment = %v", savedPayment)
		return
	}
	savedAccount,err:=s.FindAccountByID(payment.AccountID)
	if err != nil {
		t.Errorf("Reject(): can't find account by id, error = %v", err)
		return
	}
	if savedAccount.Balance != defaultTestAccount.balance{
		t.Errorf("Reject(): balance didn't changed, account = %v", savedAccount)
	}
}
//**********************************************************************************
func TestService_Repeat_success(t *testing.T) {
	svc:=newTestService()
	_,payments,err:=svc.addAccount(defaultTestAccount)
	if err != nil {
		t.Errorf("Repeat error = %v",err)
	}
	payment:=payments[0]
	_,err=svc.Repeat(payment.ID)
	if err != nil {
		t.Errorf("Repeat(), cant repeat payment, error = %v",err)
	}
}