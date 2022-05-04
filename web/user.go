package web

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int8   `json:"phone"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int8   `json:"phone"`
	Email    string `json:"email"`
}

// Login @Summary User Login
// @Description get string by Userinfo
// @Tags user
// @Accept   json
// @Produce  json
// @Param   username     path    string     true       "username"
// @Param   password     query   string     true       "password"
// @Success 200 {object} User	"ok"
// @Failure 400 {object} UserInfo "We need Username!!"
// @Failure 404 {object} UserInfo "Can not find Username"
// @Router /login [get]
func Login(c *gin.Context) {
	var user UserInfo
	c.JSON(200, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"password": user.Password,
		"phone":    user.Phone,
		"email":    user.Email,
	})
}

// @Summary User Register
// @Description get string by Userinfo
// @Tags user
// @Accept   json
// @Produce  json
// @Param   id           path    int        true       "id"
// @Param   username     path    string     true       "username"
// @Param   password     query   string     true       "password"
// @Param   phone        query   int8       false           "phone"
// @Param   email        query   string     true       "email"
// @Success 200 {object} User	"ok"
// @Failure 400 {object} UserInfo "We need Username!!"
// @Failure 404 {object} UserInfo "Can not find Username"
// @Router /login [post]
func Register(c *gin.Context) {
	var u User
	c.JSON(200, gin.H{
		"username": u.Username,
		"password": u.Password,
	})
}
