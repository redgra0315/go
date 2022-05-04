package clientele

import "github.com/go_redgra/model/clienteles"

type Customes interface {
	List() []clienteles.Customers
	Add(cu clienteles.Customers) bool
	Delete(id int) bool
	FindById(id int) int
}

type CustomerService struct {
	Customers   []clienteles.Customers `json:"customers"`
	CustomerNum int                    `json:"customer_num"`
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customer := NewCustomer(1, "张三", "男", 20, "19876543243", "zhangsan@163.com")
	customerService.Customers = append(customerService.Customers, customer)
	return customerService
}

func (c *CustomerService) List() []clienteles.Customers {
	return c.Customers
}

func (c *CustomerService) Add(cu clienteles.Customers) bool {
	c.CustomerNum++
	cu.Id = c.CustomerNum
	c.Customers = append(c.Customers, cu)
	return true
}

// 根据id删除客户
func (c CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	// 如果index == -1，说明没有这个客户
	if index == -1 {
		return false
	}
	// 从切片中删除一个元素
	c.Customers = append(c.Customers[:index], c.Customers[index+1:]...)
	return true
}

// 根据id来查找客户在切片中对应的下标，如果没有客户就返回-1
func (c CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.Customers); i++ {
		if c.Customers[i].Id == id {
			index = 1
		}
	}
	return index
}
