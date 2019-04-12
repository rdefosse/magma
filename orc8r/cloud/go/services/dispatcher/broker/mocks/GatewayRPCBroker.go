/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Code generated by mockery v1.0.0. DO NOT EDIT.
// to regenerate this file, run mockery -name=GatewayRPCBroker in /broker

package mocks

import broker "magma/orc8r/cloud/go/services/dispatcher/broker"
import mock "github.com/stretchr/testify/mock"
import protos "magma/orc8r/cloud/go/protos"

// GatewayRPCBroker is an autogenerated mock type for the GatewayRPCBroker type
type GatewayRPCBroker struct {
	mock.Mock
}

// CancelGatewayRequest provides a mock function with given fields: gwId, reqId
func (_m *GatewayRPCBroker) CancelGatewayRequest(gwId string, reqId uint32) error {
	ret := _m.Called(gwId, reqId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32) error); ok {
		r0 = rf(gwId, reqId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanupGateway provides a mock function with given fields: gwId
func (_m *GatewayRPCBroker) CleanupGateway(gwId string) error {
	ret := _m.Called(gwId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(gwId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitializeGateway provides a mock function with given fields: gwId
func (_m *GatewayRPCBroker) InitializeGateway(gwId string) chan *protos.SyncRPCRequest {
	ret := _m.Called(gwId)

	var r0 chan *protos.SyncRPCRequest
	if rf, ok := ret.Get(0).(func(string) chan *protos.SyncRPCRequest); ok {
		r0 = rf(gwId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *protos.SyncRPCRequest)
		}
	}

	return r0
}

// ProcessGatewayResponse provides a mock function with given fields: response
func (_m *GatewayRPCBroker) ProcessGatewayResponse(response *protos.SyncRPCResponse) error {
	ret := _m.Called(response)

	var r0 error
	if rf, ok := ret.Get(0).(func(*protos.SyncRPCResponse) error); ok {
		r0 = rf(response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendRequestToGateway provides a mock function with given fields: gwReq
func (_m *GatewayRPCBroker) SendRequestToGateway(gwReq *protos.GatewayRequest) (*broker.GatewayResponseChannel, error) {
	ret := _m.Called(gwReq)

	var r0 *broker.GatewayResponseChannel
	if rf, ok := ret.Get(0).(func(*protos.GatewayRequest) *broker.GatewayResponseChannel); ok {
		r0 = rf(gwReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*broker.GatewayResponseChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*protos.GatewayRequest) error); ok {
		r1 = rf(gwReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
