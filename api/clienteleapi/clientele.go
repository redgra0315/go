package clienteleapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_redgra/model/clienteles"
	"github.com/go_redgra/web/clientele"
)

type CustomerView struct {
	key             string `json:"key"`
	loop            bool   `json:"loop"`
	customerService clientele.CustomerService
}

// @Description  list  customer
// @Tags clientele
// @Accept  json
// @Produce  json
// @Success 200 {string} clienteles.Customers	"ok"
// @Failure 400 {object} clienteles.Customers "We need ID!!"
// @Failure 404 {object} clienteles.Customers "Can not find ID"
// @Router /list [get]
func (cu *CustomerView) ListApi(c *gin.Context) {
	customers := cu.customerService.List()
	c.JSON(200, gin.H{
		"message": customers,
	})

}

// @Description  add  customer
// @Tags clientele
// @Accept  json
// @Produce  json
// @Param   Id      path     int       true        "Id"
// @Param   Name    query    string    true        "Name"
// @Param   Gender  query    string    true        "Gender"
// @Param   Age     query    int       true        "Age"
// @Param   Phone   query    string    true        "Phone"
// @Param   Email   query    string    true        "Email"
// @Success 200 {string} clienteles.Customers	"ok"
// @Failure 400 {object} clienteles.Customers "We need ID!!"
// @Failure 404 {object} clienteles.Customers "Can not find ID"
// @Router /add [post]
func (cu *CustomerView) AddApi(c *gin.Context) {
	var mer clienteles.Customers
	c.ShouldBind(&mer)
	customer := clientele.NewCustomer(mer.Id, mer.Name, mer.Gender, mer.Age, mer.Phone, mer.Email)
	if cu.customerService.Add(customer) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": &mer,
		})
	}
}
