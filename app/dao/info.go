package dao

import (
	"yugod-backend/app/model"
	"yugod-backend/app/openapi"
)

type IInfoDao interface {
	GetClickVolume() (data openapi.GetClickVolumeData, err error)
}

type InfoDao struct{}

var Info IInfoDao = InfoDao{}

func (InfoDao) GetClickVolume() (data openapi.GetClickVolumeData, err error) {
	d := openapi.GetClickVolumeData{}

	DB.Model(&model.ClickVolume{})
	if err := DB.First(&d).Where("id = ?", 0).Error; err != nil {
		return d, err
	}

	return d, nil
}
