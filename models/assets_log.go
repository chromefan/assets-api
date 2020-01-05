package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *AssetsLog) TableName() string {
	return AssetsLogTBName()
}

type AssetsLog struct {
	Id          int
	TypeName    string
	Cost        float64
	MarketValue float64
	Earnings    float64
	Roa         float64
	Ratio       float64
	Details     string
	Remark      string
	Ctime       int64
	Uid         int64
}

func LogPageList(page, pageSize int) ([]*AssetsLog, int64) {
	offset := (page - 1) * pageSize
	list := make([]*AssetsLog, 0)
	query := orm.NewOrm().QueryTable(AssetsTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func LogInsert(assets *AssetsLog) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(assets)
	if err == nil {
		return id, nil
	}
	return id, nil
}
