// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type PreDeployCheckService struct {
	ListPendingDirectorChangesStub        func() (api.PendingDirectorChangesOutput, error)
	listPendingDirectorChangesMutex       sync.RWMutex
	listPendingDirectorChangesArgsForCall []struct {
	}
	listPendingDirectorChangesReturns struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}
	listPendingDirectorChangesReturnsOnCall map[int]struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PreDeployCheckService) ListPendingDirectorChanges() (api.PendingDirectorChangesOutput, error) {
	fake.listPendingDirectorChangesMutex.Lock()
	ret, specificReturn := fake.listPendingDirectorChangesReturnsOnCall[len(fake.listPendingDirectorChangesArgsForCall)]
	fake.listPendingDirectorChangesArgsForCall = append(fake.listPendingDirectorChangesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListPendingDirectorChanges", []interface{}{})
	fake.listPendingDirectorChangesMutex.Unlock()
	if fake.ListPendingDirectorChangesStub != nil {
		return fake.ListPendingDirectorChangesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listPendingDirectorChangesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesCallCount() int {
	fake.listPendingDirectorChangesMutex.RLock()
	defer fake.listPendingDirectorChangesMutex.RUnlock()
	return len(fake.listPendingDirectorChangesArgsForCall)
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesCalls(stub func() (api.PendingDirectorChangesOutput, error)) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = stub
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesReturns(result1 api.PendingDirectorChangesOutput, result2 error) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = nil
	fake.listPendingDirectorChangesReturns = struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) ListPendingDirectorChangesReturnsOnCall(i int, result1 api.PendingDirectorChangesOutput, result2 error) {
	fake.listPendingDirectorChangesMutex.Lock()
	defer fake.listPendingDirectorChangesMutex.Unlock()
	fake.ListPendingDirectorChangesStub = nil
	if fake.listPendingDirectorChangesReturnsOnCall == nil {
		fake.listPendingDirectorChangesReturnsOnCall = make(map[int]struct {
			result1 api.PendingDirectorChangesOutput
			result2 error
		})
	}
	fake.listPendingDirectorChangesReturnsOnCall[i] = struct {
		result1 api.PendingDirectorChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *PreDeployCheckService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listPendingDirectorChangesMutex.RLock()
	defer fake.listPendingDirectorChangesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PreDeployCheckService) recordInvocation(key string, args []interface{}) {
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
