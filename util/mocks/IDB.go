// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import util "github.com/Amazeful/Amazeful-Backend/util"

// IDB is an autogenerated mock type for the IDB type
type IDB struct {
	mock.Mock
}

// Collection provides a mock function with given fields: collection
func (_m *IDB) Collection(collection string) util.ICollection {
	ret := _m.Called(collection)

	var r0 util.ICollection
	if rf, ok := ret.Get(0).(func(string) util.ICollection); ok {
		r0 = rf(collection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(util.ICollection)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *IDB) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
