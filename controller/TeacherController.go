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

/*
// 将 map[string]interface{} 转换成字符串
func ChangeToString(num map[string]interface{}) (string, error) {
	cjson, err := json.Marshal(num)
	if err != nil {
		return "", err
	}
	return string(cjson), nil
}
*/

// 对 map[string]interface{} 转换为 存储时 的字符串
// 以及 按指定顺序 转换为 观看时为占比形式 的字符串

func ChangeAndDeal(Model map[string]interface{}, order map[int]string) (string, string, error) {
	num := Model //err := json.Unmarshal([]byte(data), &num) //if err != nil { //	return "", err //} //fmt.Println("ChangeAndDeal start  model : ", model.Callname) //fmt.Println("num : ", num) //fmt.Printf("%T %v\n", num["sum"])
	gojson, err := json.Marshal(num)
	if err != nil {
		return "", "", err
	}
	store := string(gojson)

	rate := make(map[string]float64) //rate["test"] = 6.6 //fmt.Println("\nrate : ", rate)
	tmp := num["sum"].(string)
	rate["sum"], _ = strconv.ParseFloat(tmp, 32)
	for k, v := range num {
		if k == "sum" {
			continue
		}
		tmp = v.(string)
		rate[k], _ = strconv.ParseFloat(tmp, 16)
		rate[k] /= rate["sum"]
	} //fmt.Println("rate : ", rate)

	view := ""
	for i := 1; i <= len(order); i++ {
		view += fmt.Sprintf("%s:%3.f%% ", order[i], rate[order[i]]*100)
	} //fmt.Println("view : ", view) //fmt.Println("ChangeAndDeal end  model : ", model.Callname)
	return store, view, nil
}

// 将 map[string]interface{} 指定 字段 对应的 数值 加一
// map 和 slice 存储的 是地址，不能直接赋值

func NumIncrease(name string, store string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(store), &data)
	if err != nil {
		return nil, err
	}
	fmt.Println("NumIncrease start : ", data)
	NumString := ""
	NumInt := 0
	for i := 1; i <= 2; i++ {
		NumString = data[name].(string)
		NumInt, err = strconv.Atoi(NumString)
		if err != nil {
			return nil, err
		}
		data[name] = strconv.Itoa(NumInt + 1)
		name = "sum" // 再来一次 使总和 增加
	}
	fmt.Println("NumIncrease end : ", data)
	return data, nil
}

// 对 某字段 的 观看数据 和 存储数据 进行处理

func Deal(view *string, store *string, Model map[string]interface{}, order map[int]string, op int) error {
	if *store == "" { // 新添加时初始化
		temp, err := json.Marshal(Model)
		if err != nil {
			return err
		}
		*store = string(temp)
	}
	if op == 0 {
		temp, err := NumIncrease(*view, *store)
		if err != nil {
			return err
		}
		if *store, *view, err = ChangeAndDeal(temp, order); err != nil {
			return err
		}
	} else if op == 1 {

	}

	return nil
}

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
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建成功", Teacher))

}

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

func GetAllTeacher(c *gin.Context) {
	Teachers, count, err := model.GetAllTeacher()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Teachers))
}

func GetTeacherByName(c *gin.Context) {
	name := c.Query("name")
	Teachers, count, err := model.GetTeacherByNAndC(name, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Teachers))

}

func GetTeacherByCourse(c *gin.Context) {
	course := c.Query("course")
	Teachers, count, err := model.GetTeacherByNAndC("", course)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}
	msg := fmt.Sprintf("已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Teachers))

}

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
	}
	fmt.Println("teacher: ", Teacher)
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
	fmt.Println("s:  ", s)
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
	}
	fmt.Println("\nmp:  ", mp)
	// 更新
	if err = model.UpdateTeacher(&Teacher, mp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "更新失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "更新成功", Teacher))

}

func DeleteTeacher(c *gin.Context) {
	Id := c.Query("teacherid")
	Teacherid, _ := strconv.Atoi(Id)
	Teacher, err := model.GetTeacherById(Teacherid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	//fmt.Println(Teacher)
	if Teacher.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "无删除对象", nil))
		return
	}
	if err = model.DeleteTeacher(Teacher); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))

}
