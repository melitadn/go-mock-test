package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ParentHandler interface {
	Parent(ctx context.Context, req string) (res string, err error)
}

type ParentHandlerImpl struct {
	service ParentService
}

func (handler *ParentHandlerImpl) Parent(ctx context.Context, req string) (res string, err error) {
	res, err = handler.service.StartService()
	return res, err
}
func (handler *ParentHandlerImpl) ParentReq(c *gin.Context) {
	log.Debug().Msg("Receiving parent request")

}
func (handler *ParentHandlerImpl) Register(engine gin.IRouter) {
	engine.POST("/test", handler.ParentReq)
}
