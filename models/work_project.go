package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *WorkProject) TableName() string {
	return WorkProjectTBName()
}

type WorkProject struct {
	Id          int
	Name    string
}

func ProjectList(page, pageSize int) ([]*WorkProject, int64) {
	offset := (page - 1) * pageSize
	list := make([]*WorkProject, 0)
	query := orm.NewOrm().QueryTable(AssetsTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-cost").Limit(pageSize, offset).All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func ProjectInsert(data *WorkProject) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err == nil {
		return id, nil
	}
	return id, nil
}
func UpdateProjectById(data *WorkProject) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Update(data)
	if err == nil {
		return num, nil
	}
	return num, nil
}

//获取单条
func GetOneProjectByType(id int) (*WorkProject, error) {
	o := orm.NewOrm()
	m := WorkProject{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
