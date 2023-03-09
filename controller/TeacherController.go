package controller

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"

	//"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// 将 map[string]interface{} 指定 字段 对应的 数值 加一
// map 和 slice 存储的 是地址，不能直接赋值

func DataDeal(view *string, store *string, order map[int]string) error {
	data := make(map[string]int)
	if err := json.Unmarshal([]byte(*store), &data); err != nil {
		return err
	} //fmt.Println("NumIncrease start : ", data)
	// 增加
	data[*view]++
	data["sum"]++
	// 求占比
	rate := make(map[string]float64) //rate["test"] = 6.6 //fmt.Println("\nrate : ", rate)
	rate["sum"] = float64(data["sum"])
	for k, v := range data {
		if k == "sum" {
			continue
		}
		rate[k] = float64(v) / rate["sum"]
	} //fmt.Println("rate : ", rate)
	// 个数数据 转换为字符串存储
	gojson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	Store := string(gojson)
	*store = Store
	// 占比数据 按指定顺序 转换为字符串存储
	View := ""
	for i := 1; i <= len(order); i++ {
		View += fmt.Sprintf("%s:%3.f%% ", order[i], rate[order[i]]*100)
	} //fmt.Println("view : ", view) //fmt.Println("ChangeAndDeal end  model : ", model.Callname)
	//fmt.Println("NumIncrease end : ", data)
	*view = View
	return nil
}

// 对分数 进行求和处理

func ScoreDeal(view *string, store *string) error {
	data := make(map[string]int)
	if err := json.Unmarshal([]byte(*store), &data); err != nil {
		return err
	} //fmt.Println("\ndata :", data)
	Add, _ := strconv.Atoi(*view)
	data["score"] += Add
	data["sum"] += 1
	gojson, err := json.Marshal(data)
	if err != nil {
		return err
	} //fmt.Println("\ndata :", data)
	Store := string(gojson)
	*store = Store
	View := fmt.Sprintf("%3.2f", float64(data["score"])/float64(data["sum"]))
	*view = View
	return nil
}

// 对 某字段 的 观看数据 和 存储数据 进行处理

func Deal(view *string, store *string, Model map[string]interface{}, order map[int]string, op int) error {
	if *store == "" { // 新添加时初始化
		temp, err := json.Marshal(Model)
		if err != nil {
			return err
		}
		*store = string(temp)
	} //fmt.Println("/n store :", *store)
	if op == 0 {
		if err := DataDeal(view, store, order); err != nil {
			return err
		}
	} else if op == 1 {
		if err := ScoreDeal(view, store); err != nil {
			return err
		}
	} //fmt.Println("/n store :", *store)
	return nil
}

// 数据处理的总函数

func TeacherDeal(t *model.Teacher) error {
	//fmt.Println("   TeacherDeal : ") //fmt.Println("\ncallname: ", model.Callname)
	err := Deal(&t.Callname, &t.Cn, model.Callname, model.CNorder, 0)
	if err != nil {
		return err
	} //fmt.Println("\nhomework: ", model.Homework)
	err = Deal(&t.Homework, &t.Hw, model.Homework, model.HWorder, 0)
	if err != nil {
		return err
	} //fmt.Println("\nmark: ", model.Mark)
	err = Deal(&t.Mark, &t.Mk, model.Mark, model.MKorder, 0)
	if err != nil {
		return err
	} //fmt.Println("\nfinishway: ", model.Finishway)
	err = Deal(&t.Finishway, &t.Fw, model.Finishway, model.FWorder, 0)
	if err != nil {
		return err
	}
	err = Deal(&t.Score, &t.Sc, model.Score, nil, 1)
	if err != nil {
		return err
	}
	return nil
}

//-------------------------------------

// 添加

func AddTeacher(c *gin.Context) {
	Teacher := &model.Teacher{}
	// 获取数据
	if err := c.ShouldBindBodyWith(Teacher, binding.JSON); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	// 检验存在性
	_, count, err := model.GetTeacherByNAndC(Teacher.Name, Teacher.Course)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "查询失败", err))
		return
	}
	if count != 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "该老师的这门课已存在", nil))
		return
	}
	// 处理多种可能的数据
	if err = TeacherDeal(Teacher); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "初始化失败", err))
		return
	}
	// 创建
	if err = model.CreateTeacher(Teacher); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "创建成功", nil))

}

// 获取 ，通过 id

func GetTeacherById(c *gin.Context) {
	Id := c.Query("teacherid") //fmt.Printf("\nid: %T %v\n", Id, Id) // string
	Teacherid, _ := strconv.Atoi(Id)
	// 判断存在
	Teacher, err := model.GetTeacherById(Teacherid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if Teacher.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "该老师不存在", nil))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取成功", Teacher))
}

// 获取全部

func GetAllTeacher(c *gin.Context) {
	Teachers, count, err := model.GetAllTeacher()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Teachers))
}

// 获取，通过 姓名 课程

func GetTeacherByNameAndCourse(c *gin.Context) {
	name := c.Query("name")
	course := c.Query("course")
	Teachers, count, err := model.GetTeacherByNAndC(name, course)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Teachers))
}

// 更新

func UpdateTeacher(c *gin.Context) {
	Id := c.Query("teacherid")
	Teacherid, _ := strconv.Atoi(Id)
	// 判断存在性
	Teacher, err := model.GetTeacherById(Teacherid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if Teacher.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "用户不存在", nil))
		return
	} //fmt.Println("teacher: ", Teacher)
	// 获取 更新 字断
	s := &model.Teacher{}
	if err = c.ShouldBindJSON(s); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}
	// 获取 原来 的 存储值
	s.Cn = Teacher.Cn
	s.Hw = Teacher.Hw
	s.Mk = Teacher.Mk
	s.Fw = Teacher.Fw
	s.Sc = Teacher.Sc
	//fmt.Println("s:  ", s)
	// 修改 存储值 和 观看值
	if err = TeacherDeal(s); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "初始化失败", err))
		return
	}
	// 去掉 不变 字段 ， 转为 map 进行 更新
	mp := structs.Map(s)
	for k, v := range mp {
		if v == "" || v == 0 {
			delete(mp, k)
		}
	} //fmt.Println("\nmp:  ", mp)
	// 更新
	if err = model.UpdateTeacher(&Teacher, mp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "更新失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "更新成功", Teacher))
}

// 删除

func DeleteTeacher(c *gin.Context) {
	Id := c.Query("teacherid")
	Teacherid, _ := strconv.Atoi(Id)
	// 验证存在性
	Teacher, err := model.GetTeacherById(Teacherid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //fmt.Println(Teacher)
	if Teacher.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "无删除对象", nil))
		return
	}
	// 删除
	if err = model.DeleteTeacher(Teacher); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))
}
