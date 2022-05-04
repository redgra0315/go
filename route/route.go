package route

import (
	"net/http"
	"time"

	"github.com/go_redgra/api/clienteleapi"

	"github.com/gin-gonic/gin"
	"github.com/go_redgra/api"
	_ "github.com/go_redgra/docs"
	"github.com/go_redgra/web"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteInit() {
	r := gin.New()
	gin.SetMode(gin.DebugMode)

	s := http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	aaa := clienteleapi.CustomerView{}

	r.GET("/v2/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.GET("/v2/testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)
	r.GET("/v2/login", web.Login)
	r.POST("/v2/login", web.Register)
	r.GET("/v2/list", aaa.ListApi)
	r.POST("/v2/add", aaa.AddApi)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.ListenAndServe()
}
