package utils

import (
	
	"bookApp/errs"
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

func StructToMap(obj interface{}) map[string]interface{} {

	var myMap map[string]interface{}
	res, err := json.Marshal(obj)
	if err != nil {

		return nil
	}
	err = json.Unmarshal(res, &myMap)
	if err != nil {

		return nil
	}
	return myMap
}

func MapToStruct(mp interface{}, obj interface{}) (*errs.AppError) {

	err := mapstructure.Decode(mp, obj)
	if err != nil {

		return errs.ErrorData()
	}

	return nil
}