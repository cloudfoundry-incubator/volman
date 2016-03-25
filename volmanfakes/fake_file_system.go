// This file was generated by counterfeiter
package volmanfakes

import (
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/volman/fakedriver"
)

type FakeFileSystem struct {
	MkdirAllStub        func(string, os.FileMode) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		arg1 string
		arg2 os.FileMode
	}
	mkdirAllReturns struct {
		result1 error
	}
	TempDirStub        func() string
	tempDirMutex       sync.RWMutex
	tempDirArgsForCall []struct{}
	tempDirReturns     struct {
		result1 string
	}
	StatStub        func(string) (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		arg1 string
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	RemoveAllStub        func(string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		arg1 string
	}
	removeAllReturns struct {
		result1 error
	}
	AbsStub        func(path string) (string, error)
	absMutex       sync.RWMutex
	absArgsForCall []struct {
		path string
	}
	absReturns struct {
		result1 string
		result2 error
	}
}

func (fake *FakeFileSystem) MkdirAll(arg1 string, arg2 os.FileMode) error {
	fake.mkdirAllMutex.Lock()
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		arg1 string
		arg2 os.FileMode
	}{arg1, arg2})
	fake.mkdirAllMutex.Unlock()
	if fake.MkdirAllStub != nil {
		return fake.MkdirAllStub(arg1, arg2)
	} else {
		return fake.mkdirAllReturns.result1
	}
}

func (fake *FakeFileSystem) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *FakeFileSystem) MkdirAllArgsForCall(i int) (string, os.FileMode) {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return fake.mkdirAllArgsForCall[i].arg1, fake.mkdirAllArgsForCall[i].arg2
}

func (fake *FakeFileSystem) MkdirAllReturns(result1 error) {
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileSystem) TempDir() string {
	fake.tempDirMutex.Lock()
	fake.tempDirArgsForCall = append(fake.tempDirArgsForCall, struct{}{})
	fake.tempDirMutex.Unlock()
	if fake.TempDirStub != nil {
		return fake.TempDirStub()
	} else {
		return fake.tempDirReturns.result1
	}
}

func (fake *FakeFileSystem) TempDirCallCount() int {
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	return len(fake.tempDirArgsForCall)
}

func (fake *FakeFileSystem) TempDirReturns(result1 string) {
	fake.TempDirStub = nil
	fake.tempDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFileSystem) Stat(arg1 string) (os.FileInfo, error) {
	fake.statMutex.Lock()
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub(arg1)
	} else {
		return fake.statReturns.result1, fake.statReturns.result2
	}
}

func (fake *FakeFileSystem) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *FakeFileSystem) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.statArgsForCall[i].arg1
}

func (fake *FakeFileSystem) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFileSystem) RemoveAll(arg1 string) error {
	fake.removeAllMutex.Lock()
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.removeAllMutex.Unlock()
	if fake.RemoveAllStub != nil {
		return fake.RemoveAllStub(arg1)
	} else {
		return fake.removeAllReturns.result1
	}
}

func (fake *FakeFileSystem) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *FakeFileSystem) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return fake.removeAllArgsForCall[i].arg1
}

func (fake *FakeFileSystem) RemoveAllReturns(result1 error) {
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileSystem) Abs(path string) (string, error) {
	fake.absMutex.Lock()
	fake.absArgsForCall = append(fake.absArgsForCall, struct {
		path string
	}{path})
	fake.absMutex.Unlock()
	if fake.AbsStub != nil {
		return fake.AbsStub(path)
	} else {
		return fake.absReturns.result1, fake.absReturns.result2
	}
}

func (fake *FakeFileSystem) AbsCallCount() int {
	fake.absMutex.RLock()
	defer fake.absMutex.RUnlock()
	return len(fake.absArgsForCall)
}

func (fake *FakeFileSystem) AbsArgsForCall(i int) string {
	fake.absMutex.RLock()
	defer fake.absMutex.RUnlock()
	return fake.absArgsForCall[i].path
}

func (fake *FakeFileSystem) AbsReturns(result1 string, result2 error) {
	fake.AbsStub = nil
	fake.absReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ fakedriver.FileSystem = new(FakeFileSystem)
