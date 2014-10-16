// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/actors/service_builder"
	"github.com/cloudfoundry/cli/cf/models"
	"sync"
)

type FakeServiceBuilder struct {
	GetServiceByNameStub        func(string) (models.ServiceOffering, error)
	getServiceByNameMutex       sync.RWMutex
	getServiceByNameArgsForCall []struct {
		arg1 string
	}
	getServiceByNameReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServiceByNameForOrgStub        func(string, string) (models.ServiceOffering, error)
	getServiceByNameForOrgMutex       sync.RWMutex
	getServiceByNameForOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceByNameForOrgReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServicesForBrokerStub        func(string) ([]models.ServiceOffering, error)
	getServicesForBrokerMutex       sync.RWMutex
	getServicesForBrokerArgsForCall []struct {
		arg1 string
	}
	getServicesForBrokerReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServiceVisibleToOrgStub        func(string, string) (models.ServiceOffering, error)
	getServiceVisibleToOrgMutex       sync.RWMutex
	getServiceVisibleToOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceVisibleToOrgReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	GetServicesVisibleToOrgStub        func(string) ([]models.ServiceOffering, error)
	getServicesVisibleToOrgMutex       sync.RWMutex
	getServicesVisibleToOrgArgsForCall []struct {
		arg1 string
	}
	getServicesVisibleToOrgReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
}

func (fake *FakeServiceBuilder) GetServiceByName(arg1 string) (models.ServiceOffering, error) {
	fake.getServiceByNameMutex.Lock()
	defer fake.getServiceByNameMutex.Unlock()
	fake.getServiceByNameArgsForCall = append(fake.getServiceByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetServiceByNameStub != nil {
		return fake.GetServiceByNameStub(arg1)
	} else {
		return fake.getServiceByNameReturns.result1, fake.getServiceByNameReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameCallCount() int {
	fake.getServiceByNameMutex.RLock()
	defer fake.getServiceByNameMutex.RUnlock()
	return len(fake.getServiceByNameArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameArgsForCall(i int) string {
	fake.getServiceByNameMutex.RLock()
	defer fake.getServiceByNameMutex.RUnlock()
	return fake.getServiceByNameArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServiceByNameReturns(result1 models.ServiceOffering, result2 error) {
	fake.getServiceByNameReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrg(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceByNameForOrgMutex.Lock()
	defer fake.getServiceByNameForOrgMutex.Unlock()
	fake.getServiceByNameForOrgArgsForCall = append(fake.getServiceByNameForOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.GetServiceByNameForOrgStub != nil {
		return fake.GetServiceByNameForOrgStub(arg1, arg2)
	} else {
		return fake.getServiceByNameForOrgReturns.result1, fake.getServiceByNameForOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgCallCount() int {
	fake.getServiceByNameForOrgMutex.RLock()
	defer fake.getServiceByNameForOrgMutex.RUnlock()
	return len(fake.getServiceByNameForOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgArgsForCall(i int) (string, string) {
	fake.getServiceByNameForOrgMutex.RLock()
	defer fake.getServiceByNameForOrgMutex.RUnlock()
	return fake.getServiceByNameForOrgArgsForCall[i].arg1, fake.getServiceByNameForOrgArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceByNameForOrgReturns(result1 models.ServiceOffering, result2 error) {
	fake.getServiceByNameForOrgReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesForBroker(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesForBrokerMutex.Lock()
	defer fake.getServicesForBrokerMutex.Unlock()
	fake.getServicesForBrokerArgsForCall = append(fake.getServicesForBrokerArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetServicesForBrokerStub != nil {
		return fake.GetServicesForBrokerStub(arg1)
	} else {
		return fake.getServicesForBrokerReturns.result1, fake.getServicesForBrokerReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesForBrokerCallCount() int {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return len(fake.getServicesForBrokerArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesForBrokerArgsForCall(i int) string {
	fake.getServicesForBrokerMutex.RLock()
	defer fake.getServicesForBrokerMutex.RUnlock()
	return fake.getServicesForBrokerArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesForBrokerReturns(result1 []models.ServiceOffering, result2 error) {
	fake.getServicesForBrokerReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrg(arg1 string, arg2 string) (models.ServiceOffering, error) {
	fake.getServiceVisibleToOrgMutex.Lock()
	defer fake.getServiceVisibleToOrgMutex.Unlock()
	fake.getServiceVisibleToOrgArgsForCall = append(fake.getServiceVisibleToOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	if fake.GetServiceVisibleToOrgStub != nil {
		return fake.GetServiceVisibleToOrgStub(arg1, arg2)
	} else {
		return fake.getServiceVisibleToOrgReturns.result1, fake.getServiceVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgCallCount() int {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return len(fake.getServiceVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgArgsForCall(i int) (string, string) {
	fake.getServiceVisibleToOrgMutex.RLock()
	defer fake.getServiceVisibleToOrgMutex.RUnlock()
	return fake.getServiceVisibleToOrgArgsForCall[i].arg1, fake.getServiceVisibleToOrgArgsForCall[i].arg2
}

func (fake *FakeServiceBuilder) GetServiceVisibleToOrgReturns(result1 models.ServiceOffering, result2 error) {
	fake.getServiceVisibleToOrgReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrg(arg1 string) ([]models.ServiceOffering, error) {
	fake.getServicesVisibleToOrgMutex.Lock()
	defer fake.getServicesVisibleToOrgMutex.Unlock()
	fake.getServicesVisibleToOrgArgsForCall = append(fake.getServicesVisibleToOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.GetServicesVisibleToOrgStub != nil {
		return fake.GetServicesVisibleToOrgStub(arg1)
	} else {
		return fake.getServicesVisibleToOrgReturns.result1, fake.getServicesVisibleToOrgReturns.result2
	}
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgCallCount() int {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return len(fake.getServicesVisibleToOrgArgsForCall)
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgArgsForCall(i int) string {
	fake.getServicesVisibleToOrgMutex.RLock()
	defer fake.getServicesVisibleToOrgMutex.RUnlock()
	return fake.getServicesVisibleToOrgArgsForCall[i].arg1
}

func (fake *FakeServiceBuilder) GetServicesVisibleToOrgReturns(result1 []models.ServiceOffering, result2 error) {
	fake.getServicesVisibleToOrgReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

var _ ServiceBuilder = new(FakeServiceBuilder)
