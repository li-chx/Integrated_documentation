package controller

import (
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AddCarShare(c *gin.Context) {
	Userid := c.MustGet("user_id").(int)
	carshare := &model.CarShare{}
	if err := c.ShouldBindJSON(carshare); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	carshare.Userid = Userid

}

func GetCarShareById(c *gin.Context) {

}

func GetCarShareByUser(c *gin.Context) {

}

func UpdateCarShare(c *gin.Context) {

}

func DeleteCarShare(c *gin.Context) {

}
