// This file was generated by counterfeiter
package volmanfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/volman/voldriver"
	"github.com/cloudfoundry-incubator/volman/voldriver.old/driverhttp"
)

type FakeRemoteClientFactory struct {
	NewRemoteClientStub        func(url string) (voldriver.Driver, error)
	newRemoteClientMutex       sync.RWMutex
	newRemoteClientArgsForCall []struct {
		url string
	}
	newRemoteClientReturns struct {
		result1 voldriver.Driver
		result2 error
	}
}

func (fake *FakeRemoteClientFactory) NewRemoteClient(url string) (voldriver.Driver, error) {
	fake.newRemoteClientMutex.Lock()
	fake.newRemoteClientArgsForCall = append(fake.newRemoteClientArgsForCall, struct {
		url string
	}{url})
	fake.newRemoteClientMutex.Unlock()
	if fake.NewRemoteClientStub != nil {
		return fake.NewRemoteClientStub(url)
	} else {
		return fake.newRemoteClientReturns.result1, fake.newRemoteClientReturns.result2
	}
}

func (fake *FakeRemoteClientFactory) NewRemoteClientCallCount() int {
	fake.newRemoteClientMutex.RLock()
	defer fake.newRemoteClientMutex.RUnlock()
	return len(fake.newRemoteClientArgsForCall)
}

func (fake *FakeRemoteClientFactory) NewRemoteClientArgsForCall(i int) string {
	fake.newRemoteClientMutex.RLock()
	defer fake.newRemoteClientMutex.RUnlock()
	return fake.newRemoteClientArgsForCall[i].url
}

func (fake *FakeRemoteClientFactory) NewRemoteClientReturns(result1 voldriver.Driver, result2 error) {
	fake.NewRemoteClientStub = nil
	fake.newRemoteClientReturns = struct {
		result1 voldriver.Driver
		result2 error
	}{result1, result2}
}

var _ driverhttp.RemoteClientFactory = new(FakeRemoteClientFactory)