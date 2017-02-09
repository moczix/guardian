// This file was generated by counterfeiter
package runruncfakes

import (
	"os/exec"
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/runrunc"
)

type FakeProcessTracker struct {
	RunStub        func(processID string, cmd *exec.Cmd, io garden.ProcessIO, tty *garden.TTYSpec, pidFile string) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		processID string
		cmd       *exec.Cmd
		io        garden.ProcessIO
		tty       *garden.TTYSpec
		pidFile   string
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	AttachStub        func(processID string, io garden.ProcessIO, pidFilePath string) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		processID   string
		io          garden.ProcessIO
		pidFilePath string
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessTracker) Run(processID string, cmd *exec.Cmd, io garden.ProcessIO, tty *garden.TTYSpec, pidFile string) (garden.Process, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		processID string
		cmd       *exec.Cmd
		io        garden.ProcessIO
		tty       *garden.TTYSpec
		pidFile   string
	}{processID, cmd, io, tty, pidFile})
	fake.recordInvocation("Run", []interface{}{processID, cmd, io, tty, pidFile})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(processID, cmd, io, tty, pidFile)
	}
	return fake.runReturns.result1, fake.runReturns.result2
}

func (fake *FakeProcessTracker) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeProcessTracker) RunArgsForCall(i int) (string, *exec.Cmd, garden.ProcessIO, *garden.TTYSpec, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].processID, fake.runArgsForCall[i].cmd, fake.runArgsForCall[i].io, fake.runArgsForCall[i].tty, fake.runArgsForCall[i].pidFile
}

func (fake *FakeProcessTracker) RunReturns(result1 garden.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessTracker) Attach(processID string, io garden.ProcessIO, pidFilePath string) (garden.Process, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		processID   string
		io          garden.ProcessIO
		pidFilePath string
	}{processID, io, pidFilePath})
	fake.recordInvocation("Attach", []interface{}{processID, io, pidFilePath})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(processID, io, pidFilePath)
	}
	return fake.attachReturns.result1, fake.attachReturns.result2
}

func (fake *FakeProcessTracker) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeProcessTracker) AttachArgsForCall(i int) (string, garden.ProcessIO, string) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].processID, fake.attachArgsForCall[i].io, fake.attachArgsForCall[i].pidFilePath
}

func (fake *FakeProcessTracker) AttachReturns(result1 garden.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessTracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProcessTracker) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.ProcessTracker = new(FakeProcessTracker)
