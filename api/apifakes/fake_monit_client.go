// This file was generated by counterfeiter
package apifakes

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/galera-healthcheck/api"
)

type FakeMonitClient struct {
	StartServiceBootstrapStub        func(req *http.Request) (string, error)
	startServiceBootstrapMutex       sync.RWMutex
	startServiceBootstrapArgsForCall []struct {
		req *http.Request
	}
	startServiceBootstrapReturns struct {
		result1 string
		result2 error
	}
	StartServiceJoinStub        func(req *http.Request) (string, error)
	startServiceJoinMutex       sync.RWMutex
	startServiceJoinArgsForCall []struct {
		req *http.Request
	}
	startServiceJoinReturns struct {
		result1 string
		result2 error
	}
	StartServiceSingleNodeStub        func(req *http.Request) (string, error)
	startServiceSingleNodeMutex       sync.RWMutex
	startServiceSingleNodeArgsForCall []struct {
		req *http.Request
	}
	startServiceSingleNodeReturns struct {
		result1 string
		result2 error
	}
	StopServiceStub        func(req *http.Request) (string, error)
	stopServiceMutex       sync.RWMutex
	stopServiceArgsForCall []struct {
		req *http.Request
	}
	stopServiceReturns struct {
		result1 string
		result2 error
	}
	GetStatusStub        func(req *http.Request) (string, error)
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		req *http.Request
	}
	getStatusReturns struct {
		result1 string
		result2 error
	}
	GetLoggerStub        func(req *http.Request) lager.Logger
	getLoggerMutex       sync.RWMutex
	getLoggerArgsForCall []struct {
		req *http.Request
	}
	getLoggerReturns struct {
		result1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMonitClient) StartServiceBootstrap(req *http.Request) (string, error) {
	fake.startServiceBootstrapMutex.Lock()
	fake.startServiceBootstrapArgsForCall = append(fake.startServiceBootstrapArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("StartServiceBootstrap", []interface{}{req})
	fake.startServiceBootstrapMutex.Unlock()
	if fake.StartServiceBootstrapStub != nil {
		return fake.StartServiceBootstrapStub(req)
	} else {
		return fake.startServiceBootstrapReturns.result1, fake.startServiceBootstrapReturns.result2
	}
}

func (fake *FakeMonitClient) StartServiceBootstrapCallCount() int {
	fake.startServiceBootstrapMutex.RLock()
	defer fake.startServiceBootstrapMutex.RUnlock()
	return len(fake.startServiceBootstrapArgsForCall)
}

func (fake *FakeMonitClient) StartServiceBootstrapArgsForCall(i int) *http.Request {
	fake.startServiceBootstrapMutex.RLock()
	defer fake.startServiceBootstrapMutex.RUnlock()
	return fake.startServiceBootstrapArgsForCall[i].req
}

func (fake *FakeMonitClient) StartServiceBootstrapReturns(result1 string, result2 error) {
	fake.StartServiceBootstrapStub = nil
	fake.startServiceBootstrapReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMonitClient) StartServiceJoin(req *http.Request) (string, error) {
	fake.startServiceJoinMutex.Lock()
	fake.startServiceJoinArgsForCall = append(fake.startServiceJoinArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("StartServiceJoin", []interface{}{req})
	fake.startServiceJoinMutex.Unlock()
	if fake.StartServiceJoinStub != nil {
		return fake.StartServiceJoinStub(req)
	} else {
		return fake.startServiceJoinReturns.result1, fake.startServiceJoinReturns.result2
	}
}

func (fake *FakeMonitClient) StartServiceJoinCallCount() int {
	fake.startServiceJoinMutex.RLock()
	defer fake.startServiceJoinMutex.RUnlock()
	return len(fake.startServiceJoinArgsForCall)
}

func (fake *FakeMonitClient) StartServiceJoinArgsForCall(i int) *http.Request {
	fake.startServiceJoinMutex.RLock()
	defer fake.startServiceJoinMutex.RUnlock()
	return fake.startServiceJoinArgsForCall[i].req
}

func (fake *FakeMonitClient) StartServiceJoinReturns(result1 string, result2 error) {
	fake.StartServiceJoinStub = nil
	fake.startServiceJoinReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMonitClient) StartServiceSingleNode(req *http.Request) (string, error) {
	fake.startServiceSingleNodeMutex.Lock()
	fake.startServiceSingleNodeArgsForCall = append(fake.startServiceSingleNodeArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("StartServiceSingleNode", []interface{}{req})
	fake.startServiceSingleNodeMutex.Unlock()
	if fake.StartServiceSingleNodeStub != nil {
		return fake.StartServiceSingleNodeStub(req)
	} else {
		return fake.startServiceSingleNodeReturns.result1, fake.startServiceSingleNodeReturns.result2
	}
}

func (fake *FakeMonitClient) StartServiceSingleNodeCallCount() int {
	fake.startServiceSingleNodeMutex.RLock()
	defer fake.startServiceSingleNodeMutex.RUnlock()
	return len(fake.startServiceSingleNodeArgsForCall)
}

func (fake *FakeMonitClient) StartServiceSingleNodeArgsForCall(i int) *http.Request {
	fake.startServiceSingleNodeMutex.RLock()
	defer fake.startServiceSingleNodeMutex.RUnlock()
	return fake.startServiceSingleNodeArgsForCall[i].req
}

func (fake *FakeMonitClient) StartServiceSingleNodeReturns(result1 string, result2 error) {
	fake.StartServiceSingleNodeStub = nil
	fake.startServiceSingleNodeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMonitClient) StopService(req *http.Request) (string, error) {
	fake.stopServiceMutex.Lock()
	fake.stopServiceArgsForCall = append(fake.stopServiceArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("StopService", []interface{}{req})
	fake.stopServiceMutex.Unlock()
	if fake.StopServiceStub != nil {
		return fake.StopServiceStub(req)
	} else {
		return fake.stopServiceReturns.result1, fake.stopServiceReturns.result2
	}
}

func (fake *FakeMonitClient) StopServiceCallCount() int {
	fake.stopServiceMutex.RLock()
	defer fake.stopServiceMutex.RUnlock()
	return len(fake.stopServiceArgsForCall)
}

func (fake *FakeMonitClient) StopServiceArgsForCall(i int) *http.Request {
	fake.stopServiceMutex.RLock()
	defer fake.stopServiceMutex.RUnlock()
	return fake.stopServiceArgsForCall[i].req
}

func (fake *FakeMonitClient) StopServiceReturns(result1 string, result2 error) {
	fake.StopServiceStub = nil
	fake.stopServiceReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMonitClient) GetStatus(req *http.Request) (string, error) {
	fake.getStatusMutex.Lock()
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("GetStatus", []interface{}{req})
	fake.getStatusMutex.Unlock()
	if fake.GetStatusStub != nil {
		return fake.GetStatusStub(req)
	} else {
		return fake.getStatusReturns.result1, fake.getStatusReturns.result2
	}
}

func (fake *FakeMonitClient) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

func (fake *FakeMonitClient) GetStatusArgsForCall(i int) *http.Request {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return fake.getStatusArgsForCall[i].req
}

func (fake *FakeMonitClient) GetStatusReturns(result1 string, result2 error) {
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeMonitClient) GetLogger(req *http.Request) lager.Logger {
	fake.getLoggerMutex.Lock()
	fake.getLoggerArgsForCall = append(fake.getLoggerArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("GetLogger", []interface{}{req})
	fake.getLoggerMutex.Unlock()
	if fake.GetLoggerStub != nil {
		return fake.GetLoggerStub(req)
	} else {
		return fake.getLoggerReturns.result1
	}
}

func (fake *FakeMonitClient) GetLoggerCallCount() int {
	fake.getLoggerMutex.RLock()
	defer fake.getLoggerMutex.RUnlock()
	return len(fake.getLoggerArgsForCall)
}

func (fake *FakeMonitClient) GetLoggerArgsForCall(i int) *http.Request {
	fake.getLoggerMutex.RLock()
	defer fake.getLoggerMutex.RUnlock()
	return fake.getLoggerArgsForCall[i].req
}

func (fake *FakeMonitClient) GetLoggerReturns(result1 lager.Logger) {
	fake.GetLoggerStub = nil
	fake.getLoggerReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeMonitClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startServiceBootstrapMutex.RLock()
	defer fake.startServiceBootstrapMutex.RUnlock()
	fake.startServiceJoinMutex.RLock()
	defer fake.startServiceJoinMutex.RUnlock()
	fake.startServiceSingleNodeMutex.RLock()
	defer fake.startServiceSingleNodeMutex.RUnlock()
	fake.stopServiceMutex.RLock()
	defer fake.stopServiceMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	fake.getLoggerMutex.RLock()
	defer fake.getLoggerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMonitClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.MonitClient = new(FakeMonitClient)
