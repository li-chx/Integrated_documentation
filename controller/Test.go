package controller

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Cou1 struct {
	Box int64 `json:"box"`
	Bag int64 `json:"bag"`
	//Name [32]byte `json:"name"`
	Cou2 Cou2 `json:"cou2"`
}

type Cou2 struct {
	//Name [20]byte `json:"name"`
	Mn int64 `json:"mn"`
}

func Test1(c *gin.Context) {
	cm := &Cou1{}
	if err := c.ShouldBindJSON(cm); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	fmt.Println("cm  :   ", cm, "\n")
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, cm)
	if err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	fmt.Println("buf   :  ", buf.Bytes(), "\n", buf, "\n")

	cm2 := &Cou1{}
	err = binary.Read(buf, binary.BigEndian, cm2)
	if err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	fmt.Println(" cm2  :  ", cm2, "\n")
}

type stu struct {
	Name string                 `json:"name"`
	Age  int                    `json:"age"`
	Pp   map[string]interface{} `json:"pp"`
}

type info struct {
	Name string      `json:"name"`
	Om   interface{} `json:"om"`
}

type cla struct {
	Name string      `json:"name"`
	Stus []stu       `json:"stus"`
	Om   interface{} `json:"om"`
}

func Test(c *gin.Context) {
	cm := cla{}
	if err := c.ShouldBindJSON(&cm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	fmt.Println("cm  :   ", cm, "\n")

	class, err := json.Marshal(cm)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	fmt.Println("class   :  ", class, "\n", string(class), "\n")
	test := model.Test{Data: string(class)}
	if err = model.Creat(&test); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建失败", err))
		return
	}

	id, _ := strconv.Atoi(c.Query("id"))
	test2, err1 := model.GetById(id)
	if err1 != nil {
		fmt.Println(err1)
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err1))
		return
	}
	fmt.Println("  test2  :  ", test2, "\n")
	m := "[" + test2.Data + "," + string(class) + "]"
	//var cm2 cla
	var cm2 []cla
	//var cm2 []info
	err = json.Unmarshal([]byte(m), &cm2)
	if err != nil {
		fmt.Println(err.Error())
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据解析失败", err))
		return
	}
	fmt.Println(" cm2  :  ", cm2, "\n")
	//info := info(cm2)

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "测试成功", cm2))
}
