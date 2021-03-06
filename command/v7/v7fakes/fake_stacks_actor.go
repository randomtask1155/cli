// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeStacksActor struct {
	GetStacksStub        func() ([]v7action.Stack, v7action.Warnings, error)
	getStacksMutex       sync.RWMutex
	getStacksArgsForCall []struct {
	}
	getStacksReturns struct {
		result1 []v7action.Stack
		result2 v7action.Warnings
		result3 error
	}
	getStacksReturnsOnCall map[int]struct {
		result1 []v7action.Stack
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStacksActor) GetStacks() ([]v7action.Stack, v7action.Warnings, error) {
	fake.getStacksMutex.Lock()
	ret, specificReturn := fake.getStacksReturnsOnCall[len(fake.getStacksArgsForCall)]
	fake.getStacksArgsForCall = append(fake.getStacksArgsForCall, struct {
	}{})
	fake.recordInvocation("GetStacks", []interface{}{})
	fake.getStacksMutex.Unlock()
	if fake.GetStacksStub != nil {
		return fake.GetStacksStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getStacksReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStacksActor) GetStacksCallCount() int {
	fake.getStacksMutex.RLock()
	defer fake.getStacksMutex.RUnlock()
	return len(fake.getStacksArgsForCall)
}

func (fake *FakeStacksActor) GetStacksCalls(stub func() ([]v7action.Stack, v7action.Warnings, error)) {
	fake.getStacksMutex.Lock()
	defer fake.getStacksMutex.Unlock()
	fake.GetStacksStub = stub
}

func (fake *FakeStacksActor) GetStacksReturns(result1 []v7action.Stack, result2 v7action.Warnings, result3 error) {
	fake.getStacksMutex.Lock()
	defer fake.getStacksMutex.Unlock()
	fake.GetStacksStub = nil
	fake.getStacksReturns = struct {
		result1 []v7action.Stack
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStacksActor) GetStacksReturnsOnCall(i int, result1 []v7action.Stack, result2 v7action.Warnings, result3 error) {
	fake.getStacksMutex.Lock()
	defer fake.getStacksMutex.Unlock()
	fake.GetStacksStub = nil
	if fake.getStacksReturnsOnCall == nil {
		fake.getStacksReturnsOnCall = make(map[int]struct {
			result1 []v7action.Stack
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getStacksReturnsOnCall[i] = struct {
		result1 []v7action.Stack
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStacksActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStacksMutex.RLock()
	defer fake.getStacksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStacksActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.StacksActor = new(FakeStacksActor)
