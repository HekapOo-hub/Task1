// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/HekapOo-hub/Task1/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// HumanRepository is an autogenerated mock type for the HumanRepository type
type HumanRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, human
func (_m *HumanRepository) Create(ctx context.Context, human model.Human) error {
	ret := _m.Called(ctx, human)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Human) error); ok {
		r0 = rf(ctx, human)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, name
func (_m *HumanRepository) Delete(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name
func (_m *HumanRepository) Get(ctx context.Context, name string) (*model.Human, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.Human
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Human); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Human)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, name, human
func (_m *HumanRepository) Update(ctx context.Context, name string, human model.Human) error {
	ret := _m.Called(ctx, name, human)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Human) error); ok {
		r0 = rf(ctx, name, human)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}