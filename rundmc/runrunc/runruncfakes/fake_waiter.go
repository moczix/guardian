// This file was generated by counterfeiter
package runruncfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
)

type FakeWaiter struct {
	WaitStub        func() (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWaiter) Wait() (int, error) {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	return fake.waitReturns.result1, fake.waitReturns.result2
}

func (fake *FakeWaiter) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeWaiter) WaitReturns(result1 int, result2 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeWaiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWaiter) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.Waiter = new(FakeWaiter)
