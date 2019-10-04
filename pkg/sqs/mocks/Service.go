// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import sqs "github.com/applike/gosoline/pkg/sqs"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateQueue provides a mock function with given fields: _a0
func (_m *Service) CreateQueue(_a0 sqs.Settings) (*sqs.Properties, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.Properties
	if rf, ok := ret.Get(0).(func(sqs.Settings) *sqs.Properties); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.Properties)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(sqs.Settings) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArn provides a mock function with given fields: _a0
func (_m *Service) GetArn(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPropertiesByArn provides a mock function with given fields: _a0
func (_m *Service) GetPropertiesByArn(_a0 string) (*sqs.Properties, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.Properties
	if rf, ok := ret.Get(0).(func(string) *sqs.Properties); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.Properties)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPropertiesByName provides a mock function with given fields: _a0
func (_m *Service) GetPropertiesByName(_a0 string) (*sqs.Properties, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.Properties
	if rf, ok := ret.Get(0).(func(string) *sqs.Properties); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.Properties)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrl provides a mock function with given fields: _a0
func (_m *Service) GetUrl(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Purge provides a mock function with given fields: _a0
func (_m *Service) Purge(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueExists provides a mock function with given fields: _a0
func (_m *Service) QueueExists(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
