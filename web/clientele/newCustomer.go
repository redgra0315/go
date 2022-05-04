package clientele

import "github.com/go_redgra/model/clienteles"

func NewCustomer(id int, name string, gender string, age int, phone, email string) clienteles.Customers {
	return clienteles.Customers{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomers(name string, gender string, age int, phone, email string) *clienteles.Customers {
	return &clienteles.Customers{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
