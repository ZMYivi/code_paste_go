package logic

import (
	"code_paste/biz/dal"
	"code_paste/biz/model"
	"code_paste/biz/utils"
)

func SaveCode(code string) (string, error) {
	codeInfo := model.CodeInfo{Code: code}
	err := dal.CodePasteDB.Create(&codeInfo).Error
	if err != nil {
		return "", err
	}
	key := utils.ShaEncrypt(codeInfo.Id)
	codeInfo.Key = key
	err = dal.CodePasteDB.Save(&codeInfo).Error
	if err != nil {
		return "", err
	}
	return key, nil
}

func GetCode(key string) (string, error) {
	codeInfo := model.CodeInfo{}
	err := dal.CodePasteDB.Where(&model.CodeInfo{Key: key}).First(&codeInfo).Error
	if err != nil {
		return "", err
	}
	return codeInfo.Code, nil
}
