package dbOrm

import (
	"context"

	"github.com/hadanhtuan/go-sdk/common"
	"gorm.io/gorm"
)

// Each model will define an Instance
type Instance struct {
	TableName string
	DBName    string
	Model     interface{}

	db *gorm.DB
}

// Apply database connection for each model
func (m *Instance) ApplyDatabase(database *gorm.DB) *Instance {
	m.db = database
	m.DBName = database.Name()
	return m
}

func (m *Instance) Create(entity interface{}) *common.APIResponse {
	// check table
	if m.db == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}
	err := m.db.WithContext(context.TODO()).Create(&entity).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.Created,
			Message: "Cannot create item in table " + m.TableName + ". Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status: common.APIStatus.Created,
	}
}

func (m *Instance) CreateMany(entity *[]interface{}) *common.APIResponse {
	// check table
	if m.db == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	err := m.db.WithContext(context.TODO()).Create(&entity).Error
	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.Created,
			Message: "Cannot create item in table " + m.TableName + ". Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status: common.APIStatus.Created,
	}
}

func (m *Instance) QueryOne(params interface{}) *common.APIResponse {

	// check table
	if m.db == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	var entity interface{}
	err := m.db.WithContext(context.TODO()).Where(&params).First(&entity).Error

	if entity == nil || err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Not found any matched. Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    []interface{}{entity},
		Message: "Query " + m.TableName + " successfully.",
	}
}

func (m *Instance) Query(params interface{}, offset int, limit int, sortFields interface{}) *common.APIResponse {
	var entities []interface{}
	var total int64

	err := m.db.WithContext(context.TODO()).Model(&entities).Where(&params).Count(&total).Error
	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot count item in table " + m.TableName + ". Error detail: " + err.Error(),
		}
	}
	err = m.db.WithContext(context.TODO()).Offset((offset - 1) * limit).Limit(limit).Where(&params).Find(&entities).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot find item in table " + m.TableName + ". Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    entities,
		Total:   total,
		Message: "Query " + m.TableName + " successfully.",
	}
}

func (m *Instance) UpdateOne(entity interface{}) *common.APIResponse {
	err := m.db.WithContext(context.TODO()).Save(&entity).Error

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

func (m *Instance) UpdateMany(entities []interface{}) *common.APIResponse {
	err := m.db.WithContext(context.TODO()).Save(&entities).Error

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

func (m *Instance) DeleteOne(entity interface{}) *common.APIResponse {
	err := m.db.WithContext(context.TODO()).Delete(&entity).Error

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
	var entity interface{}
	var count int64
	m.db.WithContext(context.TODO()).Model(&entity).Where(&params).Count(&count)
	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Total:  count,
	}
}
