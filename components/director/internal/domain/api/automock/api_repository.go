// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// APIRepository is an autogenerated mock type for the APIRepository type
type APIRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *APIRepository) Create(item *model.APIDefinition) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.APIDefinition) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: item
func (_m *APIRepository) CreateMany(item []*model.APIDefinition) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*model.APIDefinition) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: item
func (_m *APIRepository) Delete(item *model.APIDefinition) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.APIDefinition) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByApplicationID provides a mock function with given fields: id
func (_m *APIRepository) DeleteAllByApplicationID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *APIRepository) GetByID(id string) (*model.APIDefinition, error) {
	ret := _m.Called(id)

	var r0 *model.APIDefinition
	if rf, ok := ret.Get(0).(func(string) *model.APIDefinition); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.APIDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByApplicationID provides a mock function with given fields: applicationID, pageSize, cursor
func (_m *APIRepository) ListByApplicationID(applicationID string, pageSize *int, cursor *string) (*model.APIDefinitionPage, error) {
	ret := _m.Called(applicationID, pageSize, cursor)

	var r0 *model.APIDefinitionPage
	if rf, ok := ret.Get(0).(func(string, *int, *string) *model.APIDefinitionPage); ok {
		r0 = rf(applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.APIDefinitionPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *int, *string) error); ok {
		r1 = rf(applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: item
func (_m *APIRepository) Update(item *model.APIDefinition) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.APIDefinition) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
