// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"sync"

	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/service"
)

type FakeDiagnosticRequestRepo struct {
	CreateStub        func(*model.DiagnosticRequest) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *model.DiagnosticRequest
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(*model.DiagnosticRequest) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 *model.DiagnosticRequest
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetByCustomerIDStub        func(uint) ([]model.DiagnosticRequest, error)
	getByCustomerIDMutex       sync.RWMutex
	getByCustomerIDArgsForCall []struct {
		arg1 uint
	}
	getByCustomerIDReturns struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	getByCustomerIDReturnsOnCall map[int]struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	GetByIDStub        func(uint) (*model.DiagnosticRequest, error)
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		arg1 uint
	}
	getByIDReturns struct {
		result1 *model.DiagnosticRequest
		result2 error
	}
	getByIDReturnsOnCall map[int]struct {
		result1 *model.DiagnosticRequest
		result2 error
	}
	GetByUserIDStub        func(uint) ([]model.DiagnosticRequest, error)
	getByUserIDMutex       sync.RWMutex
	getByUserIDArgsForCall []struct {
		arg1 uint
	}
	getByUserIDReturns struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	getByUserIDReturnsOnCall map[int]struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	GetByVetOrgIDStub        func(uint) ([]model.DiagnosticRequest, error)
	getByVetOrgIDMutex       sync.RWMutex
	getByVetOrgIDArgsForCall []struct {
		arg1 uint
	}
	getByVetOrgIDReturns struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	getByVetOrgIDReturnsOnCall map[int]struct {
		result1 []model.DiagnosticRequest
		result2 error
	}
	UpdateStub        func(*model.DiagnosticRequest) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 *model.DiagnosticRequest
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDiagnosticRequestRepo) Create(arg1 *model.DiagnosticRequest) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *model.DiagnosticRequest
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDiagnosticRequestRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) CreateCalls(stub func(*model.DiagnosticRequest) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeDiagnosticRequestRepo) CreateArgsForCall(i int) *model.DiagnosticRequest {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) Delete(arg1 *model.DiagnosticRequest) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 *model.DiagnosticRequest
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDiagnosticRequestRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) DeleteCalls(stub func(*model.DiagnosticRequest) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeDiagnosticRequestRepo) DeleteArgsForCall(i int) *model.DiagnosticRequest {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerID(arg1 uint) ([]model.DiagnosticRequest, error) {
	fake.getByCustomerIDMutex.Lock()
	ret, specificReturn := fake.getByCustomerIDReturnsOnCall[len(fake.getByCustomerIDArgsForCall)]
	fake.getByCustomerIDArgsForCall = append(fake.getByCustomerIDArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.GetByCustomerIDStub
	fakeReturns := fake.getByCustomerIDReturns
	fake.recordInvocation("GetByCustomerID", []interface{}{arg1})
	fake.getByCustomerIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerIDCallCount() int {
	fake.getByCustomerIDMutex.RLock()
	defer fake.getByCustomerIDMutex.RUnlock()
	return len(fake.getByCustomerIDArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerIDCalls(stub func(uint) ([]model.DiagnosticRequest, error)) {
	fake.getByCustomerIDMutex.Lock()
	defer fake.getByCustomerIDMutex.Unlock()
	fake.GetByCustomerIDStub = stub
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerIDArgsForCall(i int) uint {
	fake.getByCustomerIDMutex.RLock()
	defer fake.getByCustomerIDMutex.RUnlock()
	argsForCall := fake.getByCustomerIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerIDReturns(result1 []model.DiagnosticRequest, result2 error) {
	fake.getByCustomerIDMutex.Lock()
	defer fake.getByCustomerIDMutex.Unlock()
	fake.GetByCustomerIDStub = nil
	fake.getByCustomerIDReturns = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByCustomerIDReturnsOnCall(i int, result1 []model.DiagnosticRequest, result2 error) {
	fake.getByCustomerIDMutex.Lock()
	defer fake.getByCustomerIDMutex.Unlock()
	fake.GetByCustomerIDStub = nil
	if fake.getByCustomerIDReturnsOnCall == nil {
		fake.getByCustomerIDReturnsOnCall = make(map[int]struct {
			result1 []model.DiagnosticRequest
			result2 error
		})
	}
	fake.getByCustomerIDReturnsOnCall[i] = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByID(arg1 uint) (*model.DiagnosticRequest, error) {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.GetByIDStub
	fakeReturns := fake.getByIDReturns
	fake.recordInvocation("GetByID", []interface{}{arg1})
	fake.getByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiagnosticRequestRepo) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) GetByIDCalls(stub func(uint) (*model.DiagnosticRequest, error)) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = stub
}

func (fake *FakeDiagnosticRequestRepo) GetByIDArgsForCall(i int) uint {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	argsForCall := fake.getByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) GetByIDReturns(result1 *model.DiagnosticRequest, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 *model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByIDReturnsOnCall(i int, result1 *model.DiagnosticRequest, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 *model.DiagnosticRequest
			result2 error
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 *model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByUserID(arg1 uint) ([]model.DiagnosticRequest, error) {
	fake.getByUserIDMutex.Lock()
	ret, specificReturn := fake.getByUserIDReturnsOnCall[len(fake.getByUserIDArgsForCall)]
	fake.getByUserIDArgsForCall = append(fake.getByUserIDArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.GetByUserIDStub
	fakeReturns := fake.getByUserIDReturns
	fake.recordInvocation("GetByUserID", []interface{}{arg1})
	fake.getByUserIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiagnosticRequestRepo) GetByUserIDCallCount() int {
	fake.getByUserIDMutex.RLock()
	defer fake.getByUserIDMutex.RUnlock()
	return len(fake.getByUserIDArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) GetByUserIDCalls(stub func(uint) ([]model.DiagnosticRequest, error)) {
	fake.getByUserIDMutex.Lock()
	defer fake.getByUserIDMutex.Unlock()
	fake.GetByUserIDStub = stub
}

func (fake *FakeDiagnosticRequestRepo) GetByUserIDArgsForCall(i int) uint {
	fake.getByUserIDMutex.RLock()
	defer fake.getByUserIDMutex.RUnlock()
	argsForCall := fake.getByUserIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) GetByUserIDReturns(result1 []model.DiagnosticRequest, result2 error) {
	fake.getByUserIDMutex.Lock()
	defer fake.getByUserIDMutex.Unlock()
	fake.GetByUserIDStub = nil
	fake.getByUserIDReturns = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByUserIDReturnsOnCall(i int, result1 []model.DiagnosticRequest, result2 error) {
	fake.getByUserIDMutex.Lock()
	defer fake.getByUserIDMutex.Unlock()
	fake.GetByUserIDStub = nil
	if fake.getByUserIDReturnsOnCall == nil {
		fake.getByUserIDReturnsOnCall = make(map[int]struct {
			result1 []model.DiagnosticRequest
			result2 error
		})
	}
	fake.getByUserIDReturnsOnCall[i] = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgID(arg1 uint) ([]model.DiagnosticRequest, error) {
	fake.getByVetOrgIDMutex.Lock()
	ret, specificReturn := fake.getByVetOrgIDReturnsOnCall[len(fake.getByVetOrgIDArgsForCall)]
	fake.getByVetOrgIDArgsForCall = append(fake.getByVetOrgIDArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.GetByVetOrgIDStub
	fakeReturns := fake.getByVetOrgIDReturns
	fake.recordInvocation("GetByVetOrgID", []interface{}{arg1})
	fake.getByVetOrgIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgIDCallCount() int {
	fake.getByVetOrgIDMutex.RLock()
	defer fake.getByVetOrgIDMutex.RUnlock()
	return len(fake.getByVetOrgIDArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgIDCalls(stub func(uint) ([]model.DiagnosticRequest, error)) {
	fake.getByVetOrgIDMutex.Lock()
	defer fake.getByVetOrgIDMutex.Unlock()
	fake.GetByVetOrgIDStub = stub
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgIDArgsForCall(i int) uint {
	fake.getByVetOrgIDMutex.RLock()
	defer fake.getByVetOrgIDMutex.RUnlock()
	argsForCall := fake.getByVetOrgIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgIDReturns(result1 []model.DiagnosticRequest, result2 error) {
	fake.getByVetOrgIDMutex.Lock()
	defer fake.getByVetOrgIDMutex.Unlock()
	fake.GetByVetOrgIDStub = nil
	fake.getByVetOrgIDReturns = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) GetByVetOrgIDReturnsOnCall(i int, result1 []model.DiagnosticRequest, result2 error) {
	fake.getByVetOrgIDMutex.Lock()
	defer fake.getByVetOrgIDMutex.Unlock()
	fake.GetByVetOrgIDStub = nil
	if fake.getByVetOrgIDReturnsOnCall == nil {
		fake.getByVetOrgIDReturnsOnCall = make(map[int]struct {
			result1 []model.DiagnosticRequest
			result2 error
		})
	}
	fake.getByVetOrgIDReturnsOnCall[i] = struct {
		result1 []model.DiagnosticRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeDiagnosticRequestRepo) Update(arg1 *model.DiagnosticRequest) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 *model.DiagnosticRequest
	}{arg1})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDiagnosticRequestRepo) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeDiagnosticRequestRepo) UpdateCalls(stub func(*model.DiagnosticRequest) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeDiagnosticRequestRepo) UpdateArgsForCall(i int) *model.DiagnosticRequest {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiagnosticRequestRepo) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDiagnosticRequestRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getByCustomerIDMutex.RLock()
	defer fake.getByCustomerIDMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.getByUserIDMutex.RLock()
	defer fake.getByUserIDMutex.RUnlock()
	fake.getByVetOrgIDMutex.RLock()
	defer fake.getByVetOrgIDMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDiagnosticRequestRepo) recordInvocation(key string, args []interface{}) {
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

var _ service.DiagnosticRequestRepo = new(FakeDiagnosticRequestRepo)
