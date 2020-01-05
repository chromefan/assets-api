package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Assets), new(AssetsLog))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func AssetsTBName() string {
	return TableName("assets")
}
func AssetsLogTBName() string {
	return TableName("assets_log")
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func ProjectTBName() string {
	return TableName("project")
}
