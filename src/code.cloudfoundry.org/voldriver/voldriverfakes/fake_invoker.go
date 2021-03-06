// This file was generated by counterfeiter
package voldriverfakes

import (
	"sync"

	"code.cloudfoundry.org/voldriver"
	"code.cloudfoundry.org/voldriver/invoker"
)

type FakeInvoker struct {
	InvokeStub        func(env voldriver.Env, executable string, args []string) ([]byte, error)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		env        voldriver.Env
		executable string
		args       []string
	}
	invokeReturns struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInvoker) Invoke(env voldriver.Env, executable string, args []string) ([]byte, error) {
	var argsCopy []string
	if args != nil {
		argsCopy = make([]string, len(args))
		copy(argsCopy, args)
	}
	fake.invokeMutex.Lock()
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		env        voldriver.Env
		executable string
		args       []string
	}{env, executable, argsCopy})
	fake.recordInvocation("Invoke", []interface{}{env, executable, argsCopy})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(env, executable, args)
	}
	return fake.invokeReturns.result1, fake.invokeReturns.result2
}

func (fake *FakeInvoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeInvoker) InvokeArgsForCall(i int) (voldriver.Env, string, []string) {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invokeArgsForCall[i].env, fake.invokeArgsForCall[i].executable, fake.invokeArgsForCall[i].args
}

func (fake *FakeInvoker) InvokeReturns(result1 []byte, result2 error) {
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInvoker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeInvoker) recordInvocation(key string, args []interface{}) {
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

var _ invoker.Invoker = new(FakeInvoker)
