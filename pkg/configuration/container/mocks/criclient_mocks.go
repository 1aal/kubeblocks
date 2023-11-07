// /*
// Copyright (C) 2022-2023 ApeCloud Co., Ltd
//
// This file is part of KubeBlocks project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/cri-api/pkg/apis/runtime/v1 (interfaces: RuntimeServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// MockRuntimeServiceClient is a mock of RuntimeServiceClient interface.
type MockRuntimeServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeServiceClientMockRecorder
}

// MockRuntimeServiceClientMockRecorder is the mock recorder for MockRuntimeServiceClient.
type MockRuntimeServiceClientMockRecorder struct {
	mock *MockRuntimeServiceClient
}

// NewMockRuntimeServiceClient creates a new mock instance.
func NewMockRuntimeServiceClient(ctrl *gomock.Controller) *MockRuntimeServiceClient {
	mock := &MockRuntimeServiceClient{ctrl: ctrl}
	mock.recorder = &MockRuntimeServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeServiceClient) EXPECT() *MockRuntimeServiceClientMockRecorder {
	return m.recorder
}

// Attach mocks base method.
func (m *MockRuntimeServiceClient) Attach(arg0 context.Context, arg1 *v1.AttachRequest, arg2 ...grpc.CallOption) (*v1.AttachResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Attach", varargs...)
	ret0, _ := ret[0].(*v1.AttachResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attach indicates an expected call of Attach.
func (mr *MockRuntimeServiceClientMockRecorder) Attach(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockRuntimeServiceClient)(nil).Attach), varargs...)
}

// CheckpointContainer mocks base method.
func (m *MockRuntimeServiceClient) CheckpointContainer(arg0 context.Context, arg1 *v1.CheckpointContainerRequest, arg2 ...grpc.CallOption) (*v1.CheckpointContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckpointContainer", varargs...)
	ret0, _ := ret[0].(*v1.CheckpointContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckpointContainer indicates an expected call of CheckpointContainer.
func (mr *MockRuntimeServiceClientMockRecorder) CheckpointContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckpointContainer", reflect.TypeOf((*MockRuntimeServiceClient)(nil).CheckpointContainer), varargs...)
}

// ContainerStats mocks base method.
func (m *MockRuntimeServiceClient) ContainerStats(arg0 context.Context, arg1 *v1.ContainerStatsRequest, arg2 ...grpc.CallOption) (*v1.ContainerStatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerStats", varargs...)
	ret0, _ := ret[0].(*v1.ContainerStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerStats indicates an expected call of ContainerStats.
func (mr *MockRuntimeServiceClientMockRecorder) ContainerStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStats", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ContainerStats), varargs...)
}

// ContainerStatus mocks base method.
func (m *MockRuntimeServiceClient) ContainerStatus(arg0 context.Context, arg1 *v1.ContainerStatusRequest, arg2 ...grpc.CallOption) (*v1.ContainerStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ContainerStatus", varargs...)
	ret0, _ := ret[0].(*v1.ContainerStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerStatus indicates an expected call of ContainerStatus.
func (mr *MockRuntimeServiceClientMockRecorder) ContainerStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStatus", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ContainerStatus), varargs...)
}

// CreateContainer mocks base method.
func (m *MockRuntimeServiceClient) CreateContainer(arg0 context.Context, arg1 *v1.CreateContainerRequest, arg2 ...grpc.CallOption) (*v1.CreateContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateContainer", varargs...)
	ret0, _ := ret[0].(*v1.CreateContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockRuntimeServiceClientMockRecorder) CreateContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockRuntimeServiceClient)(nil).CreateContainer), varargs...)
}

// Exec mocks base method.
func (m *MockRuntimeServiceClient) Exec(arg0 context.Context, arg1 *v1.ExecRequest, arg2 ...grpc.CallOption) (*v1.ExecResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(*v1.ExecResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockRuntimeServiceClientMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockRuntimeServiceClient)(nil).Exec), varargs...)
}

// ExecSync mocks base method.
func (m *MockRuntimeServiceClient) ExecSync(arg0 context.Context, arg1 *v1.ExecSyncRequest, arg2 ...grpc.CallOption) (*v1.ExecSyncResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecSync", varargs...)
	ret0, _ := ret[0].(*v1.ExecSyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecSync indicates an expected call of ExecSync.
func (mr *MockRuntimeServiceClientMockRecorder) ExecSync(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecSync", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ExecSync), varargs...)
}

// GetContainerEvents mocks base method.
func (m *MockRuntimeServiceClient) GetContainerEvents(arg0 context.Context, arg1 *v1.GetEventsRequest, arg2 ...grpc.CallOption) (v1.RuntimeService_GetContainerEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainerEvents", varargs...)
	ret0, _ := ret[0].(v1.RuntimeService_GetContainerEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerEvents indicates an expected call of GetContainerEvents.
func (mr *MockRuntimeServiceClientMockRecorder) GetContainerEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerEvents", reflect.TypeOf((*MockRuntimeServiceClient)(nil).GetContainerEvents), varargs...)
}

// ListContainerStats mocks base method.
func (m *MockRuntimeServiceClient) ListContainerStats(arg0 context.Context, arg1 *v1.ListContainerStatsRequest, arg2 ...grpc.CallOption) (*v1.ListContainerStatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainerStats", varargs...)
	ret0, _ := ret[0].(*v1.ListContainerStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainerStats indicates an expected call of ListContainerStats.
func (mr *MockRuntimeServiceClientMockRecorder) ListContainerStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainerStats", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListContainerStats), varargs...)
}

// ListContainers mocks base method.
func (m *MockRuntimeServiceClient) ListContainers(arg0 context.Context, arg1 *v1.ListContainersRequest, arg2 ...grpc.CallOption) (*v1.ListContainersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainers", varargs...)
	ret0, _ := ret[0].(*v1.ListContainersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockRuntimeServiceClientMockRecorder) ListContainers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListContainers), varargs...)
}

// ListMetricDescriptors mocks base method.
func (m *MockRuntimeServiceClient) ListMetricDescriptors(arg0 context.Context, arg1 *v1.ListMetricDescriptorsRequest, arg2 ...grpc.CallOption) (*v1.ListMetricDescriptorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMetricDescriptors", varargs...)
	ret0, _ := ret[0].(*v1.ListMetricDescriptorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMetricDescriptors indicates an expected call of ListMetricDescriptors.
func (mr *MockRuntimeServiceClientMockRecorder) ListMetricDescriptors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMetricDescriptors", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListMetricDescriptors), varargs...)
}

// ListPodSandbox mocks base method.
func (m *MockRuntimeServiceClient) ListPodSandbox(arg0 context.Context, arg1 *v1.ListPodSandboxRequest, arg2 ...grpc.CallOption) (*v1.ListPodSandboxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPodSandbox", varargs...)
	ret0, _ := ret[0].(*v1.ListPodSandboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodSandbox indicates an expected call of ListPodSandbox.
func (mr *MockRuntimeServiceClientMockRecorder) ListPodSandbox(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodSandbox", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListPodSandbox), varargs...)
}

// ListPodSandboxMetrics mocks base method.
func (m *MockRuntimeServiceClient) ListPodSandboxMetrics(arg0 context.Context, arg1 *v1.ListPodSandboxMetricsRequest, arg2 ...grpc.CallOption) (*v1.ListPodSandboxMetricsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPodSandboxMetrics", varargs...)
	ret0, _ := ret[0].(*v1.ListPodSandboxMetricsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodSandboxMetrics indicates an expected call of ListPodSandboxMetrics.
func (mr *MockRuntimeServiceClientMockRecorder) ListPodSandboxMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodSandboxMetrics", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListPodSandboxMetrics), varargs...)
}

// ListPodSandboxStats mocks base method.
func (m *MockRuntimeServiceClient) ListPodSandboxStats(arg0 context.Context, arg1 *v1.ListPodSandboxStatsRequest, arg2 ...grpc.CallOption) (*v1.ListPodSandboxStatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPodSandboxStats", varargs...)
	ret0, _ := ret[0].(*v1.ListPodSandboxStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodSandboxStats indicates an expected call of ListPodSandboxStats.
func (mr *MockRuntimeServiceClientMockRecorder) ListPodSandboxStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodSandboxStats", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ListPodSandboxStats), varargs...)
}

// PodSandboxStats mocks base method.
func (m *MockRuntimeServiceClient) PodSandboxStats(arg0 context.Context, arg1 *v1.PodSandboxStatsRequest, arg2 ...grpc.CallOption) (*v1.PodSandboxStatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PodSandboxStats", varargs...)
	ret0, _ := ret[0].(*v1.PodSandboxStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PodSandboxStats indicates an expected call of PodSandboxStats.
func (mr *MockRuntimeServiceClientMockRecorder) PodSandboxStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodSandboxStats", reflect.TypeOf((*MockRuntimeServiceClient)(nil).PodSandboxStats), varargs...)
}

// PodSandboxStatus mocks base method.
func (m *MockRuntimeServiceClient) PodSandboxStatus(arg0 context.Context, arg1 *v1.PodSandboxStatusRequest, arg2 ...grpc.CallOption) (*v1.PodSandboxStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PodSandboxStatus", varargs...)
	ret0, _ := ret[0].(*v1.PodSandboxStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PodSandboxStatus indicates an expected call of PodSandboxStatus.
func (mr *MockRuntimeServiceClientMockRecorder) PodSandboxStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodSandboxStatus", reflect.TypeOf((*MockRuntimeServiceClient)(nil).PodSandboxStatus), varargs...)
}

// PortForward mocks base method.
func (m *MockRuntimeServiceClient) PortForward(arg0 context.Context, arg1 *v1.PortForwardRequest, arg2 ...grpc.CallOption) (*v1.PortForwardResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PortForward", varargs...)
	ret0, _ := ret[0].(*v1.PortForwardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PortForward indicates an expected call of PortForward.
func (mr *MockRuntimeServiceClientMockRecorder) PortForward(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortForward", reflect.TypeOf((*MockRuntimeServiceClient)(nil).PortForward), varargs...)
}

// RemoveContainer mocks base method.
func (m *MockRuntimeServiceClient) RemoveContainer(arg0 context.Context, arg1 *v1.RemoveContainerRequest, arg2 ...grpc.CallOption) (*v1.RemoveContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveContainer", varargs...)
	ret0, _ := ret[0].(*v1.RemoveContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveContainer indicates an expected call of RemoveContainer.
func (mr *MockRuntimeServiceClientMockRecorder) RemoveContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockRuntimeServiceClient)(nil).RemoveContainer), varargs...)
}

// RemovePodSandbox mocks base method.
func (m *MockRuntimeServiceClient) RemovePodSandbox(arg0 context.Context, arg1 *v1.RemovePodSandboxRequest, arg2 ...grpc.CallOption) (*v1.RemovePodSandboxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemovePodSandbox", varargs...)
	ret0, _ := ret[0].(*v1.RemovePodSandboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePodSandbox indicates an expected call of RemovePodSandbox.
func (mr *MockRuntimeServiceClientMockRecorder) RemovePodSandbox(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePodSandbox", reflect.TypeOf((*MockRuntimeServiceClient)(nil).RemovePodSandbox), varargs...)
}

// ReopenContainerLog mocks base method.
func (m *MockRuntimeServiceClient) ReopenContainerLog(arg0 context.Context, arg1 *v1.ReopenContainerLogRequest, arg2 ...grpc.CallOption) (*v1.ReopenContainerLogResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReopenContainerLog", varargs...)
	ret0, _ := ret[0].(*v1.ReopenContainerLogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReopenContainerLog indicates an expected call of ReopenContainerLog.
func (mr *MockRuntimeServiceClientMockRecorder) ReopenContainerLog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReopenContainerLog", reflect.TypeOf((*MockRuntimeServiceClient)(nil).ReopenContainerLog), varargs...)
}

// RunPodSandbox mocks base method.
func (m *MockRuntimeServiceClient) RunPodSandbox(arg0 context.Context, arg1 *v1.RunPodSandboxRequest, arg2 ...grpc.CallOption) (*v1.RunPodSandboxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunPodSandbox", varargs...)
	ret0, _ := ret[0].(*v1.RunPodSandboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunPodSandbox indicates an expected call of RunPodSandbox.
func (mr *MockRuntimeServiceClientMockRecorder) RunPodSandbox(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPodSandbox", reflect.TypeOf((*MockRuntimeServiceClient)(nil).RunPodSandbox), varargs...)
}

// StartContainer mocks base method.
func (m *MockRuntimeServiceClient) StartContainer(arg0 context.Context, arg1 *v1.StartContainerRequest, arg2 ...grpc.CallOption) (*v1.StartContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartContainer", varargs...)
	ret0, _ := ret[0].(*v1.StartContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartContainer indicates an expected call of StartContainer.
func (mr *MockRuntimeServiceClientMockRecorder) StartContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockRuntimeServiceClient)(nil).StartContainer), varargs...)
}

// Status mocks base method.
func (m *MockRuntimeServiceClient) Status(arg0 context.Context, arg1 *v1.StatusRequest, arg2 ...grpc.CallOption) (*v1.StatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v1.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockRuntimeServiceClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRuntimeServiceClient)(nil).Status), varargs...)
}

// StopContainer mocks base method.
func (m *MockRuntimeServiceClient) StopContainer(arg0 context.Context, arg1 *v1.StopContainerRequest, arg2 ...grpc.CallOption) (*v1.StopContainerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopContainer", varargs...)
	ret0, _ := ret[0].(*v1.StopContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopContainer indicates an expected call of StopContainer.
func (mr *MockRuntimeServiceClientMockRecorder) StopContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockRuntimeServiceClient)(nil).StopContainer), varargs...)
}

// StopPodSandbox mocks base method.
func (m *MockRuntimeServiceClient) StopPodSandbox(arg0 context.Context, arg1 *v1.StopPodSandboxRequest, arg2 ...grpc.CallOption) (*v1.StopPodSandboxResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopPodSandbox", varargs...)
	ret0, _ := ret[0].(*v1.StopPodSandboxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPodSandbox indicates an expected call of StopPodSandbox.
func (mr *MockRuntimeServiceClientMockRecorder) StopPodSandbox(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPodSandbox", reflect.TypeOf((*MockRuntimeServiceClient)(nil).StopPodSandbox), varargs...)
}

// UpdateContainerResources mocks base method.
func (m *MockRuntimeServiceClient) UpdateContainerResources(arg0 context.Context, arg1 *v1.UpdateContainerResourcesRequest, arg2 ...grpc.CallOption) (*v1.UpdateContainerResourcesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateContainerResources", varargs...)
	ret0, _ := ret[0].(*v1.UpdateContainerResourcesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateContainerResources indicates an expected call of UpdateContainerResources.
func (mr *MockRuntimeServiceClientMockRecorder) UpdateContainerResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainerResources", reflect.TypeOf((*MockRuntimeServiceClient)(nil).UpdateContainerResources), varargs...)
}

// UpdateRuntimeConfig mocks base method.
func (m *MockRuntimeServiceClient) UpdateRuntimeConfig(arg0 context.Context, arg1 *v1.UpdateRuntimeConfigRequest, arg2 ...grpc.CallOption) (*v1.UpdateRuntimeConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRuntimeConfig", varargs...)
	ret0, _ := ret[0].(*v1.UpdateRuntimeConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRuntimeConfig indicates an expected call of UpdateRuntimeConfig.
func (mr *MockRuntimeServiceClientMockRecorder) UpdateRuntimeConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRuntimeConfig", reflect.TypeOf((*MockRuntimeServiceClient)(nil).UpdateRuntimeConfig), varargs...)
}

// Version mocks base method.
func (m *MockRuntimeServiceClient) Version(arg0 context.Context, arg1 *v1.VersionRequest, arg2 ...grpc.CallOption) (*v1.VersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Version", varargs...)
	ret0, _ := ret[0].(*v1.VersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockRuntimeServiceClientMockRecorder) Version(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockRuntimeServiceClient)(nil).Version), varargs...)
}