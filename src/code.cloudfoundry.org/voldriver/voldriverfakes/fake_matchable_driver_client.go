// This file was generated by counterfeiter
package voldriverfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/voldriver"
)

type FakeMatchableDriver struct {
	MatchesStub        func(lager.Logger, string, *voldriver.TLSConfig) bool
	matchesMutex       sync.RWMutex
	matchesArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 *voldriver.TLSConfig
	}
	matchesReturns struct {
		result1 bool
	}
	ActivateStub        func(env voldriver.Env) voldriver.ActivateResponse
	activateMutex       sync.RWMutex
	activateArgsForCall []struct {
		env voldriver.Env
	}
	activateReturns struct {
		result1 voldriver.ActivateResponse
	}
	GetStub        func(env voldriver.Env, getRequest voldriver.GetRequest) voldriver.GetResponse
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		env        voldriver.Env
		getRequest voldriver.GetRequest
	}
	getReturns struct {
		result1 voldriver.GetResponse
	}
	ListStub        func(env voldriver.Env) voldriver.ListResponse
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		env voldriver.Env
	}
	listReturns struct {
		result1 voldriver.ListResponse
	}
	MountStub        func(env voldriver.Env, mountRequest voldriver.MountRequest) voldriver.MountResponse
	mountMutex       sync.RWMutex
	mountArgsForCall []struct {
		env          voldriver.Env
		mountRequest voldriver.MountRequest
	}
	mountReturns struct {
		result1 voldriver.MountResponse
	}
	PathStub        func(env voldriver.Env, pathRequest voldriver.PathRequest) voldriver.PathResponse
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
		env         voldriver.Env
		pathRequest voldriver.PathRequest
	}
	pathReturns struct {
		result1 voldriver.PathResponse
	}
	UnmountStub        func(env voldriver.Env, unmountRequest voldriver.UnmountRequest) voldriver.ErrorResponse
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		env            voldriver.Env
		unmountRequest voldriver.UnmountRequest
	}
	unmountReturns struct {
		result1 voldriver.ErrorResponse
	}
	CapabilitiesStub        func(env voldriver.Env) voldriver.CapabilitiesResponse
	capabilitiesMutex       sync.RWMutex
	capabilitiesArgsForCall []struct {
		env voldriver.Env
	}
	capabilitiesReturns struct {
		result1 voldriver.CapabilitiesResponse
	}
	CreateStub        func(env voldriver.Env, createRequest voldriver.CreateRequest) voldriver.ErrorResponse
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		env           voldriver.Env
		createRequest voldriver.CreateRequest
	}
	createReturns struct {
		result1 voldriver.ErrorResponse
	}
	RemoveStub        func(env voldriver.Env, removeRequest voldriver.RemoveRequest) voldriver.ErrorResponse
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		env           voldriver.Env
		removeRequest voldriver.RemoveRequest
	}
	removeReturns struct {
		result1 voldriver.ErrorResponse
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMatchableDriver) Matches(arg1 lager.Logger, arg2 string, arg3 *voldriver.TLSConfig) bool {
	fake.matchesMutex.Lock()
	fake.matchesArgsForCall = append(fake.matchesArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 *voldriver.TLSConfig
	}{arg1, arg2, arg3})
	fake.recordInvocation("Matches", []interface{}{arg1, arg2, arg3})
	fake.matchesMutex.Unlock()
	if fake.MatchesStub != nil {
		return fake.MatchesStub(arg1, arg2, arg3)
	}
	return fake.matchesReturns.result1
}

func (fake *FakeMatchableDriver) MatchesCallCount() int {
	fake.matchesMutex.RLock()
	defer fake.matchesMutex.RUnlock()
	return len(fake.matchesArgsForCall)
}

func (fake *FakeMatchableDriver) MatchesArgsForCall(i int) (lager.Logger, string, *voldriver.TLSConfig) {
	fake.matchesMutex.RLock()
	defer fake.matchesMutex.RUnlock()
	return fake.matchesArgsForCall[i].arg1, fake.matchesArgsForCall[i].arg2, fake.matchesArgsForCall[i].arg3
}

func (fake *FakeMatchableDriver) MatchesReturns(result1 bool) {
	fake.MatchesStub = nil
	fake.matchesReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMatchableDriver) Activate(env voldriver.Env) voldriver.ActivateResponse {
	fake.activateMutex.Lock()
	fake.activateArgsForCall = append(fake.activateArgsForCall, struct {
		env voldriver.Env
	}{env})
	fake.recordInvocation("Activate", []interface{}{env})
	fake.activateMutex.Unlock()
	if fake.ActivateStub != nil {
		return fake.ActivateStub(env)
	}
	return fake.activateReturns.result1
}

func (fake *FakeMatchableDriver) ActivateCallCount() int {
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	return len(fake.activateArgsForCall)
}

func (fake *FakeMatchableDriver) ActivateArgsForCall(i int) voldriver.Env {
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	return fake.activateArgsForCall[i].env
}

func (fake *FakeMatchableDriver) ActivateReturns(result1 voldriver.ActivateResponse) {
	fake.ActivateStub = nil
	fake.activateReturns = struct {
		result1 voldriver.ActivateResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Get(env voldriver.Env, getRequest voldriver.GetRequest) voldriver.GetResponse {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		env        voldriver.Env
		getRequest voldriver.GetRequest
	}{env, getRequest})
	fake.recordInvocation("Get", []interface{}{env, getRequest})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(env, getRequest)
	}
	return fake.getReturns.result1
}

func (fake *FakeMatchableDriver) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeMatchableDriver) GetArgsForCall(i int) (voldriver.Env, voldriver.GetRequest) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].env, fake.getArgsForCall[i].getRequest
}

func (fake *FakeMatchableDriver) GetReturns(result1 voldriver.GetResponse) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 voldriver.GetResponse
	}{result1}
}

func (fake *FakeMatchableDriver) List(env voldriver.Env) voldriver.ListResponse {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		env voldriver.Env
	}{env})
	fake.recordInvocation("List", []interface{}{env})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(env)
	}
	return fake.listReturns.result1
}

func (fake *FakeMatchableDriver) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeMatchableDriver) ListArgsForCall(i int) voldriver.Env {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].env
}

func (fake *FakeMatchableDriver) ListReturns(result1 voldriver.ListResponse) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 voldriver.ListResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Mount(env voldriver.Env, mountRequest voldriver.MountRequest) voldriver.MountResponse {
	fake.mountMutex.Lock()
	fake.mountArgsForCall = append(fake.mountArgsForCall, struct {
		env          voldriver.Env
		mountRequest voldriver.MountRequest
	}{env, mountRequest})
	fake.recordInvocation("Mount", []interface{}{env, mountRequest})
	fake.mountMutex.Unlock()
	if fake.MountStub != nil {
		return fake.MountStub(env, mountRequest)
	}
	return fake.mountReturns.result1
}

func (fake *FakeMatchableDriver) MountCallCount() int {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return len(fake.mountArgsForCall)
}

func (fake *FakeMatchableDriver) MountArgsForCall(i int) (voldriver.Env, voldriver.MountRequest) {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return fake.mountArgsForCall[i].env, fake.mountArgsForCall[i].mountRequest
}

func (fake *FakeMatchableDriver) MountReturns(result1 voldriver.MountResponse) {
	fake.MountStub = nil
	fake.mountReturns = struct {
		result1 voldriver.MountResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Path(env voldriver.Env, pathRequest voldriver.PathRequest) voldriver.PathResponse {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
		env         voldriver.Env
		pathRequest voldriver.PathRequest
	}{env, pathRequest})
	fake.recordInvocation("Path", []interface{}{env, pathRequest})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub(env, pathRequest)
	}
	return fake.pathReturns.result1
}

func (fake *FakeMatchableDriver) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeMatchableDriver) PathArgsForCall(i int) (voldriver.Env, voldriver.PathRequest) {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return fake.pathArgsForCall[i].env, fake.pathArgsForCall[i].pathRequest
}

func (fake *FakeMatchableDriver) PathReturns(result1 voldriver.PathResponse) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 voldriver.PathResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Unmount(env voldriver.Env, unmountRequest voldriver.UnmountRequest) voldriver.ErrorResponse {
	fake.unmountMutex.Lock()
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		env            voldriver.Env
		unmountRequest voldriver.UnmountRequest
	}{env, unmountRequest})
	fake.recordInvocation("Unmount", []interface{}{env, unmountRequest})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(env, unmountRequest)
	}
	return fake.unmountReturns.result1
}

func (fake *FakeMatchableDriver) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeMatchableDriver) UnmountArgsForCall(i int) (voldriver.Env, voldriver.UnmountRequest) {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].env, fake.unmountArgsForCall[i].unmountRequest
}

func (fake *FakeMatchableDriver) UnmountReturns(result1 voldriver.ErrorResponse) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 voldriver.ErrorResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Capabilities(env voldriver.Env) voldriver.CapabilitiesResponse {
	fake.capabilitiesMutex.Lock()
	fake.capabilitiesArgsForCall = append(fake.capabilitiesArgsForCall, struct {
		env voldriver.Env
	}{env})
	fake.recordInvocation("Capabilities", []interface{}{env})
	fake.capabilitiesMutex.Unlock()
	if fake.CapabilitiesStub != nil {
		return fake.CapabilitiesStub(env)
	}
	return fake.capabilitiesReturns.result1
}

func (fake *FakeMatchableDriver) CapabilitiesCallCount() int {
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	return len(fake.capabilitiesArgsForCall)
}

func (fake *FakeMatchableDriver) CapabilitiesArgsForCall(i int) voldriver.Env {
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	return fake.capabilitiesArgsForCall[i].env
}

func (fake *FakeMatchableDriver) CapabilitiesReturns(result1 voldriver.CapabilitiesResponse) {
	fake.CapabilitiesStub = nil
	fake.capabilitiesReturns = struct {
		result1 voldriver.CapabilitiesResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Create(env voldriver.Env, createRequest voldriver.CreateRequest) voldriver.ErrorResponse {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		env           voldriver.Env
		createRequest voldriver.CreateRequest
	}{env, createRequest})
	fake.recordInvocation("Create", []interface{}{env, createRequest})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(env, createRequest)
	}
	return fake.createReturns.result1
}

func (fake *FakeMatchableDriver) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeMatchableDriver) CreateArgsForCall(i int) (voldriver.Env, voldriver.CreateRequest) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].env, fake.createArgsForCall[i].createRequest
}

func (fake *FakeMatchableDriver) CreateReturns(result1 voldriver.ErrorResponse) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 voldriver.ErrorResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Remove(env voldriver.Env, removeRequest voldriver.RemoveRequest) voldriver.ErrorResponse {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		env           voldriver.Env
		removeRequest voldriver.RemoveRequest
	}{env, removeRequest})
	fake.recordInvocation("Remove", []interface{}{env, removeRequest})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(env, removeRequest)
	}
	return fake.removeReturns.result1
}

func (fake *FakeMatchableDriver) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeMatchableDriver) RemoveArgsForCall(i int) (voldriver.Env, voldriver.RemoveRequest) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].env, fake.removeArgsForCall[i].removeRequest
}

func (fake *FakeMatchableDriver) RemoveReturns(result1 voldriver.ErrorResponse) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 voldriver.ErrorResponse
	}{result1}
}

func (fake *FakeMatchableDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.matchesMutex.RLock()
	defer fake.matchesMutex.RUnlock()
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMatchableDriver) recordInvocation(key string, args []interface{}) {
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

var _ voldriver.MatchableDriver = new(FakeMatchableDriver)
