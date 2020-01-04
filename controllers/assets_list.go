package controllers

import "assets-api/models"

type ListResult struct {
	Id          int	`json:"id"`
	TypeName    string `json:"type_name"`
	Cost        float64 `json:"cost"`
	MarketValue float64 `json:"market_value"`
	Earnings    float64 `json:"earnings"`
	Roa         float64 `json:"roa"`
	Ratio       float64 `json:"ratio"`
	Details     string `json:"details"`
	Ctime       int64 `json:"ctime"`
	Mtime       int64 `json:"mtime"`
	Remark      string `json:"remark"`
}
// @Title 创建卡片
// @Description 创建卡片
// @Param
func (c *BaseController) List() {
	resp := NewResponse()
	defer func() {
		c.ResponseDone(resp)
	}()
	page ,_:= c.GetInt("page")
	pageSize ,_:= c.GetInt("page_size")
	assetsList , _:= models.PageList(page,pageSize)
	var list []ListResult
	for _,v := range assetsList{
		var listResult ListResult
		listResult.Id = v.Id
		listResult.TypeName = v.TypeName
		listResult.Cost = v.Cost
		listResult.MarketValue = v.MarketValue
		listResult.Earnings = v.Earnings
		listResult.Roa = v.Roa
		listResult.Ratio = v.Ratio
		listResult.Details = v.Details
		listResult.Remark = v.Remark
		listResult.Ctime = v.Ctime
		listResult.Mtime = v.Mtime
		list = append(list,listResult)
	}
	/*result := make(map[string]interface{})
	result["list"] = list
	result["total"] = total*/
	resp.Result = list
}

