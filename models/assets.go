package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *Assets) TableName() string {
	return AssetsTBName()
}

type Assets struct {
	Id          int
	TypeName    string
	Cost        float64
	MarketValue float64
	Earnings    float64
	Roa         float64
	Ratio       float64
	Details     string
	Ctime       int64
	Mtime       int64
	Remark      string
	Uid      int64
}

func PageList(page, pageSize int) ([]*Assets, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Assets, 0)
	query := orm.NewOrm().QueryTable(AssetsTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-cost").Limit(pageSize, offset).All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func AssetsInsert(assets *Assets) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(assets)
	if err == nil {
		return id, nil
	}
	return id, nil
}
func UpdateById(assets *Assets) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Update(assets)
	if err == nil {
		return num, nil
	}
	return num, nil
}

//获取单条
func GetOneByType(id int) (*Assets, error) {
	o := orm.NewOrm()
	m := Assets{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
