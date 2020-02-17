package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *WorkDay) TableName() string {
	return WorkDayTBName()
}

type WorkDay struct {
	Id        int
	Name      string
	ProjectId int
	GroupId   int
	Amount1   int
	Amount2   int
	Amount3   int
	Amount4   int
	Amount5   int
	Amount6   int
	Amount7   int
	Ctime     string
}

func WorkDayList(page, pageSize int) ([]*WorkDay, int64) {
	offset := (page - 1) * pageSize
	list := make([]*WorkDay, 0)
	query := orm.NewOrm().QueryTable(AssetsTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-cost").Limit(pageSize, offset).All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func WorkDayInsert(data *WorkDay) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err == nil {
		return id, nil
	}
	return id, nil
}
func UpdateWorkDayById(data *WorkDay) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Update(data)
	if err == nil {
		return num, nil
	}
	return num, nil
}

//获取单条
func GetOneWorkDayByType(id int) (*WorkDay, error) {
	o := orm.NewOrm()
	m := WorkDay{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
