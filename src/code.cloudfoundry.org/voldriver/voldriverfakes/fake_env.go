// This file was generated by counterfeiter
package voldriverfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/voldriver"
)

type FakeEnv struct {
	LoggerStub        func() lager.Logger
	loggerMutex       sync.RWMutex
	loggerArgsForCall []struct{}
	loggerReturns     struct {
		result1 lager.Logger
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct{}
	contextReturns     struct {
		result1 context.Context
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnv) Logger() lager.Logger {
	fake.loggerMutex.Lock()
	fake.loggerArgsForCall = append(fake.loggerArgsForCall, struct{}{})
	fake.recordInvocation("Logger", []interface{}{})
	fake.loggerMutex.Unlock()
	if fake.LoggerStub != nil {
		return fake.LoggerStub()
	}
	return fake.loggerReturns.result1
}

func (fake *FakeEnv) LoggerCallCount() int {
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	return len(fake.loggerArgsForCall)
}

func (fake *FakeEnv) LoggerReturns(result1 lager.Logger) {
	fake.LoggerStub = nil
	fake.loggerReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeEnv) Context() context.Context {
	fake.contextMutex.Lock()
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct{}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	return fake.contextReturns.result1
}

func (fake *FakeEnv) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeEnv) ContextReturns(result1 context.Context) {
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeEnv) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEnv) recordInvocation(key string, args []interface{}) {
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

var _ voldriver.Env = new(FakeEnv)
