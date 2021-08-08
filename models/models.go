package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"website/pkg/logging"
	"website/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // 添加 tag
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Setup() {

	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		logging.Info(err)
	}

	// 数据库名字前面加上前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)

	// 注册回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

/**
 * @Description: set `CreatedOn`, `ModifiedOn` when creating
 * @param scope
 */
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// scope.FieldByName 通过 scope.Fields() 获取所有字段，判断当前是否包含所需字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifiTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiTimeField.IsBlank {
				modifiTimeField.Set("nowTime")
			}
		}
	}
}

/**
 * @Description: set `ModifyTime` when updating
 * @param scope
 */
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
