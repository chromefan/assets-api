package controllers

import (
	"assets-api/errors"
	"assets-api/models"
	"fmt"
	"time"
)

// @Title 创建卡片
// @Description 创建卡片
// @Param
func (c *BaseController) WorkDaySave() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	work := new(models.WorkDay)
	work.Id, _ = c.GetInt("id")
	work.ProjectId,_ = c.GetInt("project_id")
	work.GroupId, _ = c.GetInt("group_id")
	work.Amount1, _ = c.GetInt("amount1")
	work.Amount2, _ = c.GetInt("amount2")
	work.Amount3, _ = c.GetInt("amount3")
	work.Amount4, _ = c.GetInt("amount4")
	work.Amount5, _ = c.GetInt("amount5")
	work.Amount6, _= c.GetInt("amount6")
	work.Amount7, _= c.GetInt("amount7")
	work.Ctime = time.Now().Unix()
	fmt.Println(work)
	if work.Id == 0 {
		_, err := models.WorkDayInsert(work)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	} else {
		_, err := models.UpdateWorkDayById(work)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	}
}
func (c *BaseController) WorkGroupSave() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	fmt.Println(c.GetString("name"))
	group := new(models.WorkGroup)
	group.Id, _ = c.GetInt("id")
	group.Name = c.GetString("name")
	if group.Name == "" {
		resp.Error = &errors.SaveError
		return
	}
	if group.Id == 0 {
		_, err := models.GroupInsert(group)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	} else {
		_, err := models.UpdateGroupById(group)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	}
}
func (c *BaseController) WorkProjectSave() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	project := new(models.WorkProject)
	project.Id, _ = c.GetInt("id")
	project.Name = c.GetString("name")
	fmt.Println(project)
	if project.Name == "" {
		resp.Error = &errors.SaveError
		return
	}
	if project.Id == 0 {
		_, err := models.ProjectInsert(project)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	} else {
		_, err := models.UpdateProjectById(project)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	}
}
func saveGroup(assets *models.Assets) {
	assetsLog := new(models.AssetsLog)
	assetsLog.Id = assets.Id
	assetsLog.TypeName = assets.TypeName
	assetsLog.Cost = assets.Cost
	assetsLog.MarketValue = assets.MarketValue
	assetsLog.Earnings = assets.Earnings
	assetsLog.Roa = assets.Roa
	assetsLog.Ratio = assets.Ratio
	assetsLog.Details = assets.Details
	assetsLog.Remark = assets.Remark
	assetsLog.Ctime = assets.Ctime
	assetsLog.Uid = assets.Uid
	models.LogInsert(assetsLog)
}
