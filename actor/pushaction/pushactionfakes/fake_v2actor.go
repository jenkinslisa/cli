// This file was generated by counterfeiter
package pushactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/v2action"
)

type FakeV2Actor struct {
	CheckRouteStub        func(route v2action.Route) (bool, v2action.Warnings, error)
	checkRouteMutex       sync.RWMutex
	checkRouteArgsForCall []struct {
		route v2action.Route
	}
	checkRouteReturns struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}
	checkRouteReturnsOnCall map[int]struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}
	CreateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	createApplicationMutex       sync.RWMutex
	createApplicationArgsForCall []struct {
		application v2action.Application
	}
	createApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	createApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationRoutesStub        func(applicationGUID string) ([]v2action.Route, v2action.Warnings, error)
	getApplicationRoutesMutex       sync.RWMutex
	getApplicationRoutesArgsForCall []struct {
		applicationGUID string
	}
	getApplicationRoutesReturns struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	getApplicationRoutesReturnsOnCall map[int]struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationDomainsStub        func(orgGUID string) ([]v2action.Domain, v2action.Warnings, error)
	getOrganizationDomainsMutex       sync.RWMutex
	getOrganizationDomainsArgsForCall []struct {
		orgGUID string
	}
	getOrganizationDomainsReturns struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationDomainsReturnsOnCall map[int]struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}
	GetRouteByHostAndDomainStub        func(host string, domainGUID string) (v2action.Route, v2action.Warnings, error)
	getRouteByHostAndDomainMutex       sync.RWMutex
	getRouteByHostAndDomainArgsForCall []struct {
		host       string
		domainGUID string
	}
	getRouteByHostAndDomainReturns struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	getRouteByHostAndDomainReturnsOnCall map[int]struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	UpdateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		application v2action.Application
	}
	updateApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2Actor) CheckRoute(route v2action.Route) (bool, v2action.Warnings, error) {
	fake.checkRouteMutex.Lock()
	ret, specificReturn := fake.checkRouteReturnsOnCall[len(fake.checkRouteArgsForCall)]
	fake.checkRouteArgsForCall = append(fake.checkRouteArgsForCall, struct {
		route v2action.Route
	}{route})
	fake.recordInvocation("CheckRoute", []interface{}{route})
	fake.checkRouteMutex.Unlock()
	if fake.CheckRouteStub != nil {
		return fake.CheckRouteStub(route)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.checkRouteReturns.result1, fake.checkRouteReturns.result2, fake.checkRouteReturns.result3
}

func (fake *FakeV2Actor) CheckRouteCallCount() int {
	fake.checkRouteMutex.RLock()
	defer fake.checkRouteMutex.RUnlock()
	return len(fake.checkRouteArgsForCall)
}

func (fake *FakeV2Actor) CheckRouteArgsForCall(i int) v2action.Route {
	fake.checkRouteMutex.RLock()
	defer fake.checkRouteMutex.RUnlock()
	return fake.checkRouteArgsForCall[i].route
}

func (fake *FakeV2Actor) CheckRouteReturns(result1 bool, result2 v2action.Warnings, result3 error) {
	fake.CheckRouteStub = nil
	fake.checkRouteReturns = struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CheckRouteReturnsOnCall(i int, result1 bool, result2 v2action.Warnings, result3 error) {
	fake.CheckRouteStub = nil
	if fake.checkRouteReturnsOnCall == nil {
		fake.checkRouteReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.checkRouteReturnsOnCall[i] = struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.createApplicationMutex.Lock()
	ret, specificReturn := fake.createApplicationReturnsOnCall[len(fake.createApplicationArgsForCall)]
	fake.createApplicationArgsForCall = append(fake.createApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("CreateApplication", []interface{}{application})
	fake.createApplicationMutex.Unlock()
	if fake.CreateApplicationStub != nil {
		return fake.CreateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationReturns.result1, fake.createApplicationReturns.result2, fake.createApplicationReturns.result3
}

func (fake *FakeV2Actor) CreateApplicationCallCount() int {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return len(fake.createApplicationArgsForCall)
}

func (fake *FakeV2Actor) CreateApplicationArgsForCall(i int) v2action.Application {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return fake.createApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) CreateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	fake.createApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	if fake.createApplicationReturnsOnCall == nil {
		fake.createApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].name, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutes(applicationGUID string) ([]v2action.Route, v2action.Warnings, error) {
	fake.getApplicationRoutesMutex.Lock()
	ret, specificReturn := fake.getApplicationRoutesReturnsOnCall[len(fake.getApplicationRoutesArgsForCall)]
	fake.getApplicationRoutesArgsForCall = append(fake.getApplicationRoutesArgsForCall, struct {
		applicationGUID string
	}{applicationGUID})
	fake.recordInvocation("GetApplicationRoutes", []interface{}{applicationGUID})
	fake.getApplicationRoutesMutex.Unlock()
	if fake.GetApplicationRoutesStub != nil {
		return fake.GetApplicationRoutesStub(applicationGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationRoutesReturns.result1, fake.getApplicationRoutesReturns.result2, fake.getApplicationRoutesReturns.result3
}

func (fake *FakeV2Actor) GetApplicationRoutesCallCount() int {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return len(fake.getApplicationRoutesArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationRoutesArgsForCall(i int) string {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return fake.getApplicationRoutesArgsForCall[i].applicationGUID
}

func (fake *FakeV2Actor) GetApplicationRoutesReturns(result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	fake.getApplicationRoutesReturns = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutesReturnsOnCall(i int, result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	if fake.getApplicationRoutesReturnsOnCall == nil {
		fake.getApplicationRoutesReturnsOnCall = make(map[int]struct {
			result1 []v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationRoutesReturnsOnCall[i] = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetOrganizationDomains(orgGUID string) ([]v2action.Domain, v2action.Warnings, error) {
	fake.getOrganizationDomainsMutex.Lock()
	ret, specificReturn := fake.getOrganizationDomainsReturnsOnCall[len(fake.getOrganizationDomainsArgsForCall)]
	fake.getOrganizationDomainsArgsForCall = append(fake.getOrganizationDomainsArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationDomains", []interface{}{orgGUID})
	fake.getOrganizationDomainsMutex.Unlock()
	if fake.GetOrganizationDomainsStub != nil {
		return fake.GetOrganizationDomainsStub(orgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationDomainsReturns.result1, fake.getOrganizationDomainsReturns.result2, fake.getOrganizationDomainsReturns.result3
}

func (fake *FakeV2Actor) GetOrganizationDomainsCallCount() int {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return len(fake.getOrganizationDomainsArgsForCall)
}

func (fake *FakeV2Actor) GetOrganizationDomainsArgsForCall(i int) string {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return fake.getOrganizationDomainsArgsForCall[i].orgGUID
}

func (fake *FakeV2Actor) GetOrganizationDomainsReturns(result1 []v2action.Domain, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationDomainsStub = nil
	fake.getOrganizationDomainsReturns = struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetOrganizationDomainsReturnsOnCall(i int, result1 []v2action.Domain, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationDomainsStub = nil
	if fake.getOrganizationDomainsReturnsOnCall == nil {
		fake.getOrganizationDomainsReturnsOnCall = make(map[int]struct {
			result1 []v2action.Domain
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationDomainsReturnsOnCall[i] = struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetRouteByHostAndDomain(host string, domainGUID string) (v2action.Route, v2action.Warnings, error) {
	fake.getRouteByHostAndDomainMutex.Lock()
	ret, specificReturn := fake.getRouteByHostAndDomainReturnsOnCall[len(fake.getRouteByHostAndDomainArgsForCall)]
	fake.getRouteByHostAndDomainArgsForCall = append(fake.getRouteByHostAndDomainArgsForCall, struct {
		host       string
		domainGUID string
	}{host, domainGUID})
	fake.recordInvocation("GetRouteByHostAndDomain", []interface{}{host, domainGUID})
	fake.getRouteByHostAndDomainMutex.Unlock()
	if fake.GetRouteByHostAndDomainStub != nil {
		return fake.GetRouteByHostAndDomainStub(host, domainGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getRouteByHostAndDomainReturns.result1, fake.getRouteByHostAndDomainReturns.result2, fake.getRouteByHostAndDomainReturns.result3
}

func (fake *FakeV2Actor) GetRouteByHostAndDomainCallCount() int {
	fake.getRouteByHostAndDomainMutex.RLock()
	defer fake.getRouteByHostAndDomainMutex.RUnlock()
	return len(fake.getRouteByHostAndDomainArgsForCall)
}

func (fake *FakeV2Actor) GetRouteByHostAndDomainArgsForCall(i int) (string, string) {
	fake.getRouteByHostAndDomainMutex.RLock()
	defer fake.getRouteByHostAndDomainMutex.RUnlock()
	return fake.getRouteByHostAndDomainArgsForCall[i].host, fake.getRouteByHostAndDomainArgsForCall[i].domainGUID
}

func (fake *FakeV2Actor) GetRouteByHostAndDomainReturns(result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetRouteByHostAndDomainStub = nil
	fake.getRouteByHostAndDomainReturns = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetRouteByHostAndDomainReturnsOnCall(i int, result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetRouteByHostAndDomainStub = nil
	if fake.getRouteByHostAndDomainReturnsOnCall == nil {
		fake.getRouteByHostAndDomainReturnsOnCall = make(map[int]struct {
			result1 v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getRouteByHostAndDomainReturnsOnCall[i] = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UpdateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("UpdateApplication", []interface{}{application})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateApplicationReturns.result1, fake.updateApplicationReturns.result2, fake.updateApplicationReturns.result3
}

func (fake *FakeV2Actor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV2Actor) UpdateApplicationArgsForCall(i int) v2action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.updateApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) UpdateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UpdateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkRouteMutex.RLock()
	defer fake.checkRouteMutex.RUnlock()
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	fake.getRouteByHostAndDomainMutex.RLock()
	defer fake.getRouteByHostAndDomainMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeV2Actor) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.V2Actor = new(FakeV2Actor)
