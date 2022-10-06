package api

import (
	"appMagic/internal/server/api/methods"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service *Handler
}

type Handler interface {
	HandelFunc()
}

func (service *Service) HandelFunc() {
	method := methods.Service{}
	router := gin.Default()
	router.GET("api/getData", method.GetData)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
