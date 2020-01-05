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
func (c *BaseController) Save() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	assets :=  new(models.Assets)
	assets.Id, _ = c.GetInt("id")
	assets.TypeName = c.GetString("type_name")
	assets.Cost,_ = c.GetFloat("cost")
	assets.MarketValue,_ = c.GetFloat("market_value")
	assets.Earnings,_ = c.GetFloat("earnings")
	assets.Roa,_ = c.GetFloat("roa")
	assets.Ratio,_ = c.GetFloat("ratio")
	assets.Details = c.GetString("details")
	assets.Remark = c.GetString("remark")
	fmt.Println(assets)
	now := time.Now().Unix()
	assets.Ctime = now
	if assets.Id == 0 {
		_, err := models.AssetsInsert(assets)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	}else{
		assets.Mtime = now
		_, err := models.UpdateById(assets)
		if err != nil {
			resp.Error = &errors.SaveError
			return
		}
	}
	saveLog(assets)
}
func saveLog(assets *models.Assets)  {
	assetsLog :=  new(models.AssetsLog)
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
