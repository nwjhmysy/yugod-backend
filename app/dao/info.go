package dao

import (
	"errors"
	"reflect"
	"yugod-backend/app/model"
	"yugod-backend/app/openapi"

	"gorm.io/gorm"
)

type IInfoDao interface {
	GetClickVolume() (data *openapi.GetClickVolumeData, err error)
	SetClickVolumeByTag(tag string) error
}

type InfoDao struct{}

var Info IInfoDao = InfoDao{}

func (InfoDao) GetClickVolume() (data *openapi.GetClickVolumeData, err error) {
	d := openapi.GetClickVolumeData{}
	result := model.ClickVolume{}
	DB.Model(&result)
	if err := DB.First(&result).Where("id = ?", 1).Error; err != nil {
		return nil, err
	}

	d.Web = int32(result.Web)
	d.Note = int32(result.Note)

	return &d, nil
}

func (InfoDao) SetClickVolumeByTag(tag string) error {
	result := model.ClickVolume{}

	// 判断参数是否正确
	t := reflect.TypeOf(result)
	flag := false
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == tag {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("参数不正确！")
	}

	// 参数正确
	DB = DB.Model(&result)

	err := DB.Transaction(func(tx *gorm.DB) error {
		// 查询
		if err := DB.First(&result).Where("id = ?", 1).Error; err != nil {
			return err
		}

		switch tag {
		case "Web":
			if err := DB.Update("web", (result.Web + 1)).Error; err != nil {
				return err
			}
		case "Note":
			if err := DB.Update("note", (result.Note + 1)).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
