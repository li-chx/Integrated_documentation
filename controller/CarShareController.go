package controller

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"net/http"

	"strconv"
)

func AddCarShare(c *gin.Context) {
	Userid := c.MustGet("user_id").(int)
	carshare := &model.CarShare{}
	if err := c.ShouldBindBodyWith(carshare, binding.JSON); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}

	carshare.Luggage = fmt.Sprintf("箱数:%d,包数:%d", carshare.Box, carshare.Bag)
	carshare.Userid = Userid
	if err := model.CreateCarShare(carshare); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建失败", err))
		return
	}
	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建成功", nil))

}

func GetCarShareById(c *gin.Context) {
	Id := c.Query("carshareid")
	Carshareid, _ := strconv.Atoi(Id)
	carshare, err := model.GetCarShareById(Carshareid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if carshare.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取成功", carshare))
}

func GetCarShareByDestination(c *gin.Context) {
	mp := struct {
		Destination string `json:"destination"`
	}{""}
	if err := c.ShouldBindJSON(&mp); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	if mp.Destination == "" {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "绑定失败", nil))
		return
	}
	//fmt.Println(mp.Destination)
	carshares, count, err := model.GetCarShareByDestination(mp.Destination)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", nil))
		return
	}
	//fmt.Println(carshares, "\n", count)
	//if count == 0 {
	//	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil))
	//	return
	//}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, carshares))

}

func GetCarShareByUser(c *gin.Context) {
	Userid := c.MustGet("user_id").(int)
	carshares, count, err := model.GetCarShareByUser(Userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	//if count == 0 {
	//	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil))
	//	return
	//}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, carshares))

}

func GetAllCarShare(c *gin.Context) {
	carshares, count, err := model.GetAllCarShare()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	//if count == 0 {
	//	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil))
	//	return
	//}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, carshares))

}

func UpdateCarShare(c *gin.Context) {
	Id := c.Query("carshareid")
	Carshareid, _ := strconv.Atoi(Id)
	carshare, err := model.GetCarShareById(Carshareid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if carshare.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	}

	s := &model.CarShare{}
	if err = c.ShouldBindJSON(s); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	carshare.Luggage = fmt.Sprintf("箱数:%d,包数:%d", carshare.Box, carshare.Bag)
	mp := structs.Map(s)
	for k, v := range mp {
		if v == "" || v == 0 {
			delete(mp, k)
		}
	}
	if err = model.UpdateCarShare(carshare, mp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "更新失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "更新成功", nil))

}

func DeleteCarShare(c *gin.Context) {
	Id := c.Query("carshareid")
	Carshareid, _ := strconv.Atoi(Id)
	carshare, err := model.GetCarShareById(Carshareid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if carshare.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "目前无订单", nil))
		return
	}
	if err = model.DeleteCarShare(carshare); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))

}
