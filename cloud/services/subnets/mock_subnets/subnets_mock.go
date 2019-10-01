/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi (interfaces: SubnetsClientAPI)

// Package mock_subnets is a generated GoMock package.
package mock_subnets

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockSubnetsClientAPI is a mock of SubnetsClientAPI interface
type MockSubnetsClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSubnetsClientAPIMockRecorder
}

// MockSubnetsClientAPIMockRecorder is the mock recorder for MockSubnetsClientAPI
type MockSubnetsClientAPIMockRecorder struct {
	mock *MockSubnetsClientAPI
}

// NewMockSubnetsClientAPI creates a new mock instance
func NewMockSubnetsClientAPI(ctrl *gomock.Controller) *MockSubnetsClientAPI {
	mock := &MockSubnetsClientAPI{ctrl: ctrl}
	mock.recorder = &MockSubnetsClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubnetsClientAPI) EXPECT() *MockSubnetsClientAPIMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockSubnetsClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 network.Subnet) (network.SubnetsCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(network.SubnetsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockSubnetsClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockSubnetsClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4)
}

// Delete mocks base method
func (m *MockSubnetsClientAPI) Delete(arg0 context.Context, arg1, arg2, arg3 string) (network.SubnetsDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.SubnetsDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockSubnetsClientAPIMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubnetsClientAPI)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockSubnetsClientAPI) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSubnetsClientAPIMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubnetsClientAPI)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// List mocks base method
func (m *MockSubnetsClientAPI) List(arg0 context.Context, arg1, arg2 string) (network.SubnetListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.SubnetListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSubnetsClientAPIMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSubnetsClientAPI)(nil).List), arg0, arg1, arg2)
}

// PrepareNetworkPolicies mocks base method
func (m *MockSubnetsClientAPI) PrepareNetworkPolicies(arg0 context.Context, arg1, arg2, arg3 string, arg4 network.PrepareNetworkPoliciesRequest) (network.SubnetsPrepareNetworkPoliciesFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareNetworkPolicies", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(network.SubnetsPrepareNetworkPoliciesFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareNetworkPolicies indicates an expected call of PrepareNetworkPolicies
func (mr *MockSubnetsClientAPIMockRecorder) PrepareNetworkPolicies(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareNetworkPolicies", reflect.TypeOf((*MockSubnetsClientAPI)(nil).PrepareNetworkPolicies), arg0, arg1, arg2, arg3, arg4)
}