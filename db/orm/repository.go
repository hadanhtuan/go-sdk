package orm

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hadanhtuan/go-sdk/common"
	"gorm.io/gorm"
)

// Each model will define an Instance
type Instance struct {
	TableName string
	DBName    string
	Model     interface{}

	DB *gorm.DB
}

// Apply database connection for each model
func (m *Instance) ApplyDatabase(database *gorm.DB) *Instance {
	m.DB = database
	m.DBName = database.Name()
	return m
}

// TODO: return same type with m.Model
func (m *Instance) newObject() interface{} {
	t := reflect.TypeOf(m.Model)
	return reflect.New(t).Interface()
}

func (m *Instance) newListObject(limit int) interface{} {
	t := reflect.TypeOf(m.Model)
	return reflect.MakeSlice(reflect.SliceOf(t), 0, limit).Interface()
}

func (m *Instance) convertSingleData(data interface{}) (interface{}, error) {
	obj := m.newObject()
	listObj := m.newListObject(1)

	encodeData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(encodeData, &obj)

	if err != nil {
		return nil, err
	}

	listValue := reflect.Append(reflect.ValueOf(listObj),
		reflect.Indirect(reflect.ValueOf(obj)))

	return listValue.Interface(), nil
}

func (m *Instance) convertMultiData(data []map[string]interface{}) (interface{}, error) {
	listObj := m.newListObject(len(data))
	var listValue reflect.Value = reflect.ValueOf(listObj)
	// TODO: Don't need call Elem() in here: listValue = listValue.Elem()
	for _, mapData := range data {
		obj := m.newObject()

		encodeData, err := json.Marshal(mapData)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(encodeData, &obj)
		if err != nil {
			return nil, err
		}
		listValue = reflect.Append(listValue,
			reflect.Indirect(reflect.ValueOf(obj)))

	}

	return listValue.Interface(), nil
}

func (m *Instance) Create(entity interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}
	err := m.DB.WithContext(context.TODO()).Create(entity).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.Created,
			Message: "Cannot create item in table " + m.TableName + ". Error detail: " + err.Error(),
		}
	}

	data, _ := m.convertSingleData(entity)

	return &common.APIResponse{
		Status: common.APIStatus.Created,
		Data:   data,
		Total:  1,
	}
}

func (m *Instance) QueryOne(query interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	entity := m.newObject()
	err := m.DB.WithContext(context.TODO()).Where(query).First(entity).Error

	if entity == nil || err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Not found any matched. Error detail: " + err.Error(),
			Total:   0,
		}
	}

	data, _ := m.convertSingleData(entity)

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    data,
		Message: "Query " + m.TableName + " successfully.",
		Total:   1,
	}
}

func (m *Instance) Query(query interface{}, offset int, limit int) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	entities := []map[string]interface{}{} // TODO: use for dynamic result
	var total int64

	err := m.DB.WithContext(context.TODO()).
		Model(m.Model).Where(query).Count(&total).
		Offset((offset - 1) * limit).Limit(limit).Where(query).
		Find(&entities).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot find item in table " + ". Error detail: " + err.Error(),
		}
	}

	data, _ := m.convertMultiData(entities)

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    data,
		Total:   total,
		Message: "Query " + " successfully.",
	}
}

func (m *Instance) Update(query interface{}, payload interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	err := m.DB.WithContext(context.TODO()).Model(m.Model).
		Where(query).Updates(payload).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Cannot update. Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Item in table " + m.TableName + " updated.",
	}
}

func (m *Instance) Delete(payload interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	err := m.DB.WithContext(context.TODO()).
		Where(payload).Delete(m.Model).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Cannot delete. Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Item in table " + m.TableName + " deleted.",
	}
}

func (m *Instance) Count(params interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	var entity interface{}
	var count int64
	m.DB.WithContext(context.TODO()).Model(&entity).Where(&params).Count(&count)
	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Total:  count,
	}
}
