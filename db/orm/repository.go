package orm

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

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

type QueryOption struct {
	Preload []string
	Order   []string
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
	t := reflect.TypeOf(m.Model) // Get the element type of the slice
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, limit)

	slicePtr := reflect.New(slice.Type())

	return slicePtr.Interface()
}

func (m *Instance) convertSingleData(data interface{}) (interface{}, error) {
	obj := m.newObject()
	t := reflect.TypeOf(m.Model)
	listObj := reflect.MakeSlice(reflect.SliceOf(t), 0, 1).Interface()

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

func (m *Instance) Create(entity interface{}) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}
	err := m.DB.WithContext(context.TODO()).Table(m.TableName).Create(entity).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
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

func (m *Instance) QueryOne(query interface{},  option *QueryOption) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	entity := m.newObject()

	db := m.DB.WithContext(context.TODO()).Table(m.TableName).Model(m.Model)

	if option != nil {
		if option.Preload != nil {
			for _, preload := range option.Preload {
				db.Preload(preload)
			}
		}

		if option.Order != nil {
			orders := strings.Join(option.Order, ", ")
			db.Order(orders)
		}
	}

	err := db.Where(query).First(entity).Error

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

func (m *Instance) Query(query interface{}, offset int32, limit int32, option *QueryOption) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	entities := m.newListObject(int(limit))
	var total int64

	db := m.DB.WithContext(context.TODO()).Table(m.TableName).Model(m.Model)

	if option != nil {
		if option.Preload != nil {
			for _, preload := range option.Preload {
				db.Preload(preload)
			}
		}

		if option.Order != nil {
			orders := strings.Join(option.Order, ", ")
			db.Order(orders)
		}
	}

	err := db.Where(query).Count(&total). // count
		Offset(int((offset - 1) * limit)).Limit(int(limit)). // paginate
		Where(query).Find(entities).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot find item in table " + ". Error detail: " + err.Error(),
		}
	}

	var data interface{}
	v := reflect.ValueOf(entities)

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem() // Dereference the pointer
		// Now 'v' holds the value that the pointer points to
		data = v.Interface()
	} else {
		data = nil
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    data,
		Total:   total,
		Message: "Query " + m.TableName + " successfully.",
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

	err := m.DB.WithContext(context.TODO()).Table(m.TableName).
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

	err := m.DB.WithContext(context.TODO()).Table(m.TableName).
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

	var count int64 = 0
	m.DB.WithContext(context.TODO()).Table(m.TableName).Where(params).Count(&count)
	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Total:  count,
	}
}
