// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/volman/voldriver"
	"github.com/pivotal-golang/lager"
)

type FakeClient struct {
	CreateStub        func(logger lager.Logger, name string, opts voldriver.Opts) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		logger lager.Logger
		name   string
		opts   voldriver.Opts
	}
	createReturns struct {
		result1 error
	}
	RemoveStub        func(logger lager.Logger, name string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	removeReturns struct {
		result1 error
	}
	MountStub        func(logger lager.Logger, name string) (string, error)
	mountMutex       sync.RWMutex
	mountArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	mountReturns struct {
		result1 string
		result2 error
	}
	PathStub        func(logger lager.Logger, name string) (string, error)
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	pathReturns struct {
		result1 string
		result2 error
	}
	UnmountStub        func(logger lager.Logger, name string) error
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	unmountReturns struct {
		result1 error
	}
	GetStub        func(logger lager.Logger, name string) (voldriver.Volume, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	getReturns struct {
		result1 voldriver.Volume
		result2 error
	}
	ListStub        func(logger lager.Logger, name string) (voldriver.Volumes, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		logger lager.Logger
		name   string
	}
	listReturns struct {
		result1 voldriver.Volumes
		result2 error
	}
}

func (fake *FakeClient) Create(logger lager.Logger, name string, opts voldriver.Opts) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		logger lager.Logger
		name   string
		opts   voldriver.Opts
	}{logger, name, opts})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(logger, name, opts)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeClient) CreateArgsForCall(i int) (lager.Logger, string, voldriver.Opts) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].logger, fake.createArgsForCall[i].name, fake.createArgsForCall[i].opts
}

func (fake *FakeClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Remove(logger lager.Logger, name string) error {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(logger, name)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeClient) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeClient) RemoveArgsForCall(i int) (lager.Logger, string) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].logger, fake.removeArgsForCall[i].name
}

func (fake *FakeClient) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Mount(logger lager.Logger, name string) (string, error) {
	fake.mountMutex.Lock()
	fake.mountArgsForCall = append(fake.mountArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.mountMutex.Unlock()
	if fake.MountStub != nil {
		return fake.MountStub(logger, name)
	} else {
		return fake.mountReturns.result1, fake.mountReturns.result2
	}
}

func (fake *FakeClient) MountCallCount() int {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return len(fake.mountArgsForCall)
}

func (fake *FakeClient) MountArgsForCall(i int) (lager.Logger, string) {
	fake.mountMutex.RLock()
	defer fake.mountMutex.RUnlock()
	return fake.mountArgsForCall[i].logger, fake.mountArgsForCall[i].name
}

func (fake *FakeClient) MountReturns(result1 string, result2 error) {
	fake.MountStub = nil
	fake.mountReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Path(logger lager.Logger, name string) (string, error) {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub(logger, name)
	} else {
		return fake.pathReturns.result1, fake.pathReturns.result2
	}
}

func (fake *FakeClient) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeClient) PathArgsForCall(i int) (lager.Logger, string) {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return fake.pathArgsForCall[i].logger, fake.pathArgsForCall[i].name
}

func (fake *FakeClient) PathReturns(result1 string, result2 error) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Unmount(logger lager.Logger, name string) error {
	fake.unmountMutex.Lock()
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(logger, name)
	} else {
		return fake.unmountReturns.result1
	}
}

func (fake *FakeClient) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeClient) UnmountArgsForCall(i int) (lager.Logger, string) {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].logger, fake.unmountArgsForCall[i].name
}

func (fake *FakeClient) UnmountReturns(result1 error) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Get(logger lager.Logger, name string) (voldriver.Volume, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(logger, name)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeClient) GetArgsForCall(i int) (lager.Logger, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].logger, fake.getArgsForCall[i].name
}

func (fake *FakeClient) GetReturns(result1 voldriver.Volume, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 voldriver.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) List(logger lager.Logger, name string) (voldriver.Volumes, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		logger lager.Logger
		name   string
	}{logger, name})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(logger, name)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeClient) ListArgsForCall(i int) (lager.Logger, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].logger, fake.listArgsForCall[i].name
}

func (fake *FakeClient) ListReturns(result1 voldriver.Volumes, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 voldriver.Volumes
		result2 error
	}{result1, result2}
}

var _ voldriver.Client = new(FakeClient)