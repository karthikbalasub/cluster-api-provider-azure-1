/*
Copyright The Kubernetes Authors.

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
// Source: ../roleassignments.go

// Package mock_roleassignments is a generated GoMock package.
package mock_roleassignments

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockRoleAssignmentScope is a mock of RoleAssignmentScope interface.
type MockRoleAssignmentScope struct {
	ctrl     *gomock.Controller
	recorder *MockRoleAssignmentScopeMockRecorder
}

// MockRoleAssignmentScopeMockRecorder is the mock recorder for MockRoleAssignmentScope.
type MockRoleAssignmentScopeMockRecorder struct {
	mock *MockRoleAssignmentScope
}

// NewMockRoleAssignmentScope creates a new mock instance.
func NewMockRoleAssignmentScope(ctrl *gomock.Controller) *MockRoleAssignmentScope {
	mock := &MockRoleAssignmentScope{ctrl: ctrl}
	mock.recorder = &MockRoleAssignmentScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleAssignmentScope) EXPECT() *MockRoleAssignmentScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockRoleAssignmentScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockRoleAssignmentScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockRoleAssignmentScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockRoleAssignmentScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockRoleAssignmentScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockRoleAssignmentScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockRoleAssignmentScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockRoleAssignmentScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockRoleAssignmentScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockRoleAssignmentScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockRoleAssignmentScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockRoleAssignmentScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockRoleAssignmentScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockRoleAssignmentScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockRoleAssignmentScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockRoleAssignmentScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockRoleAssignmentScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockRoleAssignmentScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockRoleAssignmentScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockRoleAssignmentScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockRoleAssignmentScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockRoleAssignmentScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockRoleAssignmentScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockRoleAssignmentScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockRoleAssignmentScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockRoleAssignmentScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockRoleAssignmentScope)(nil).ClusterName))
}

// Enabled mocks base method.
func (m *MockRoleAssignmentScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockRoleAssignmentScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockRoleAssignmentScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockRoleAssignmentScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockRoleAssignmentScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockRoleAssignmentScope)(nil).Error), varargs...)
}

// FailureDomains mocks base method.
func (m *MockRoleAssignmentScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockRoleAssignmentScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockRoleAssignmentScope)(nil).FailureDomains))
}

// HashKey mocks base method.
func (m *MockRoleAssignmentScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockRoleAssignmentScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockRoleAssignmentScope)(nil).HashKey))
}

// Info mocks base method.
func (m *MockRoleAssignmentScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockRoleAssignmentScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockRoleAssignmentScope)(nil).Info), varargs...)
}

// Location mocks base method.
func (m *MockRoleAssignmentScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockRoleAssignmentScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockRoleAssignmentScope)(nil).Location))
}

// ResourceGroup mocks base method.
func (m *MockRoleAssignmentScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockRoleAssignmentScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockRoleAssignmentScope)(nil).ResourceGroup))
}

// RoleAssignmentSpecs mocks base method.
func (m *MockRoleAssignmentScope) RoleAssignmentSpecs() []azure.RoleAssignmentSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleAssignmentSpecs")
	ret0, _ := ret[0].([]azure.RoleAssignmentSpec)
	return ret0
}

// RoleAssignmentSpecs indicates an expected call of RoleAssignmentSpecs.
func (mr *MockRoleAssignmentScopeMockRecorder) RoleAssignmentSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleAssignmentSpecs", reflect.TypeOf((*MockRoleAssignmentScope)(nil).RoleAssignmentSpecs))
}

// SubscriptionID mocks base method.
func (m *MockRoleAssignmentScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockRoleAssignmentScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockRoleAssignmentScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockRoleAssignmentScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockRoleAssignmentScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockRoleAssignmentScope)(nil).TenantID))
}

// V mocks base method.
func (m *MockRoleAssignmentScope) V(level int) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockRoleAssignmentScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockRoleAssignmentScope)(nil).V), level)
}

// WithName mocks base method.
func (m *MockRoleAssignmentScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockRoleAssignmentScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockRoleAssignmentScope)(nil).WithName), name)
}

// WithValues mocks base method.
func (m *MockRoleAssignmentScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockRoleAssignmentScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockRoleAssignmentScope)(nil).WithValues), keysAndValues...)
}
