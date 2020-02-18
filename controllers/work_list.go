package controllers

import (
	"assets-api/models"
	"time"
)

type workListResult struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ProjectId int    `json:"project_id"`
	GroupId   int    `json:"group_id"`
	Amount1   int    `json:"amount_1"`
	Amount2   int    `json:"amount_2"`
	Amount3   int    `json:"amount_3"`
	Amount4   int    `json:"amount_4"`
	Amount5   int    `json:"amount_5"`
	Amount6   int    `json:"amount_6"`
	Amount7   int    `json:"amount_7"`
	Ctime     string `json:"ctime"`
}
type workGroup struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	WorkDayList []*models.WorkDay `json:"workDaytList"`
}

// @Title 创建卡片
// @Description 创建卡片
// @Param
func (c *BaseController) WorkList() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	//startDay := c.GetString("start_day")
	//endDay := c.GetString("end_day")

	startTime := int64(1)
	endTime := time.Now().Unix()
	workList, _ := models.WorkDayTotal(startTime, endTime, 0)

	groupList, _ := models.GroupList()
	var workGroupList []workGroup
	for _, v := range groupList {
		var wg workGroup
		wg.WorkDayList, _ = models.WorkDayTotal(startTime, endTime, v.Id)
		wg.Id = v.Id
		wg.Name = v.Name
		workGroupList = append(workGroupList, wg)
	}
	projectList, _ := models.ProjectList()
	result := make(map[string]interface{})
	result["workDayList"] = workList
	result["workGroupList"] = workGroupList
	result["groupList"] = groupList
	result["projectList"] = projectList
	resp.Result = result
}
