package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
)

// TableName 设置表名
func (c *WorkDay) TableName() string {
	return WorkDayTBName()
}

type WorkDay struct {
	Id        int   `json:"id"`
	ProjectId int   `json:"project_id"`
	GroupId   int   `json:"group_id"`
	Amount1   int   `json:"amount1"`
	Amount2   int   `json:"amount2"`
	Amount3   int   `json:"amount3"`
	Amount4   int   `json:"amount4"`
	Amount5   int   `json:"amount5"`
	Amount6   int   `json:"amount6"`
	Amount7   int   `json:"amount7"`
	Ctime     int64 `json:"ctime"`
}

func WorkDayList(day string, endDay string, group_id int) ([]*WorkDay, int64) {
	list := make([]*WorkDay, 0)
	query := orm.NewOrm().QueryTable(WorkDayTBName())

	total, _ := query.Count()
	num, err := query.OrderBy("-id").All(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, total
}
func WorkDayTotal(startTime int64, endTime int64, groupId int) ([]*WorkDay, int64) {
	list := make([]*WorkDay, 0)
	sql := "SELECT `project_id`, sum(`amount1`) as amount1 , " +
		"sum(`amount2`) as amount2, sum(`amount3`) as amount3, sum(`amount4`) as amount4, sum(`amount5`) as amount5, sum(`amount6`) as amount6, sum(`amount7`) as amount7 " +
		"FROM `work_day` WHERE "
	if groupId > 0 {
		sql += fmt.Sprintf(" group_id = %d AND ", groupId)
	}

	sql += " `ctime` > %d  AND `ctime` <= %d  GROUP BY `project_id` "

	sql = fmt.Sprintf(sql, startTime, endTime)
	num, err := orm.NewOrm().Raw(sql).QueryRows(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, num
}
func WorkDayGroupList(startTime int, endTime int, group_id int) ([]*WorkDay, int64) {
	list := make([]*WorkDay, 0)
	sql := "SELECT `id`, `project_id`, `group_id`, sum(`amount1`), " +
		"sum(`amount2`), sum(`amount3`), sum(`amount4`), sum(`amount5`), sum(`amount6`), sum(`amount7`) " +
		"FROM `work_day` " +
		"WHERE ctime > startDay AND ctime <= endTime GROUP BY project_id  ORDER BY `id` DESC  "
	num, err := orm.NewOrm().Raw(sql).QueryRows(&list)
	if err != nil {
		beeLogger.Log.Errorf("PageList err", num, err)
	}
	return list, num
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
