package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *WorkGroup) TableName() string {
	return WorkGroupTBName()
}

type WorkGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GroupList() ([]*WorkGroup, int64) {
	list := make([]*WorkGroup, 0)
	query := orm.NewOrm().QueryTable(WorkGroupTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-id").All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func GroupInsert(data *WorkGroup) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err == nil {
		return id, nil
	}
	return id, nil
}
func UpdateGroupById(data *WorkGroup) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Update(data)
	if err == nil {
		return num, nil
	}
	return num, nil
}

//获取单条
func GetOneGroupByType(id int) (*WorkGroup, error) {
	o := orm.NewOrm()
	m := WorkGroup{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
