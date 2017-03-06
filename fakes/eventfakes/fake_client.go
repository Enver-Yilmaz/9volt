// This file was generated by counterfeiter
package eventfakes

import (
	"sync"

	"github.com/9corp/9volt/event"
)

type FakeIClient struct {
	AddStub        func(string, string) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addReturns struct {
		result1 error
	}
	AddWithErrorLogStub        func(string, string) error
	addWithErrorLogMutex       sync.RWMutex
	addWithErrorLogArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addWithErrorLogReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIClient) Add(arg1 string, arg2 string) error {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1, arg2)
	} else {
		return fake.addReturns.result1
	}
}

func (fake *FakeIClient) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeIClient) AddArgsForCall(i int) (string, string) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].arg1, fake.addArgsForCall[i].arg2
}

func (fake *FakeIClient) AddReturns(result1 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) AddWithErrorLog(arg1 string, arg2 string) error {
	fake.addWithErrorLogMutex.Lock()
	fake.addWithErrorLogArgsForCall = append(fake.addWithErrorLogArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AddWithErrorLog", []interface{}{arg1, arg2})
	fake.addWithErrorLogMutex.Unlock()
	if fake.AddWithErrorLogStub != nil {
		return fake.AddWithErrorLogStub(arg1, arg2)
	} else {
		return fake.addWithErrorLogReturns.result1
	}
}

func (fake *FakeIClient) AddWithErrorLogCallCount() int {
	fake.addWithErrorLogMutex.RLock()
	defer fake.addWithErrorLogMutex.RUnlock()
	return len(fake.addWithErrorLogArgsForCall)
}

func (fake *FakeIClient) AddWithErrorLogArgsForCall(i int) (string, string) {
	fake.addWithErrorLogMutex.RLock()
	defer fake.addWithErrorLogMutex.RUnlock()
	return fake.addWithErrorLogArgsForCall[i].arg1, fake.addWithErrorLogArgsForCall[i].arg2
}

func (fake *FakeIClient) AddWithErrorLogReturns(result1 error) {
	fake.AddWithErrorLogStub = nil
	fake.addWithErrorLogReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.addWithErrorLogMutex.RLock()
	defer fake.addWithErrorLogMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIClient) recordInvocation(key string, args []interface{}) {
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

var _ event.IClient = new(FakeIClient)