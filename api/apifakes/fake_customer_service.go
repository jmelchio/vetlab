// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"context"
	"sync"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
)

type FakeCustomerService struct {
	CreateCustomerStub        func(ctx context.Context, user model.Customer) (*model.Customer, error)
	createCustomerMutex       sync.RWMutex
	createCustomerArgsForCall []struct {
		ctx  context.Context
		user model.Customer
	}
	createCustomerReturns struct {
		result1 *model.Customer
		result2 error
	}
	createCustomerReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	UpdateCustomerStub        func(ctx context.Context, user model.Customer) (*model.Customer, error)
	updateCustomerMutex       sync.RWMutex
	updateCustomerArgsForCall []struct {
		ctx  context.Context
		user model.Customer
	}
	updateCustomerReturns struct {
		result1 *model.Customer
		result2 error
	}
	updateCustomerReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	DeleteCustomerStub        func(ctx context.Context, user model.Customer) error
	deleteCustomerMutex       sync.RWMutex
	deleteCustomerArgsForCall []struct {
		ctx  context.Context
		user model.Customer
	}
	deleteCustomerReturns struct {
		result1 error
	}
	deleteCustomerReturnsOnCall map[int]struct {
		result1 error
	}
	UpdatePasswordStub        func(ctx context.Context, user model.Customer, password string) (*model.Customer, error)
	updatePasswordMutex       sync.RWMutex
	updatePasswordArgsForCall []struct {
		ctx      context.Context
		user     model.Customer
		password string
	}
	updatePasswordReturns struct {
		result1 *model.Customer
		result2 error
	}
	updatePasswordReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	LoginStub        func(ctx context.Context, userName string, password string) (*model.Customer, error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		ctx      context.Context
		userName string
		password string
	}
	loginReturns struct {
		result1 *model.Customer
		result2 error
	}
	loginReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	FindCustomerByUserNameStub        func(ctx context.Context, userName string) (*model.Customer, error)
	findCustomerByUserNameMutex       sync.RWMutex
	findCustomerByUserNameArgsForCall []struct {
		ctx      context.Context
		userName string
	}
	findCustomerByUserNameReturns struct {
		result1 *model.Customer
		result2 error
	}
	findCustomerByUserNameReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	FindCustomerByIDStub        func(ctx context.Context, userID uint) (*model.Customer, error)
	findCustomerByIDMutex       sync.RWMutex
	findCustomerByIDArgsForCall []struct {
		ctx    context.Context
		userID uint
	}
	findCustomerByIDReturns struct {
		result1 *model.Customer
		result2 error
	}
	findCustomerByIDReturnsOnCall map[int]struct {
		result1 *model.Customer
		result2 error
	}
	FindCustomerByVetOrgStub        func(ctx context.Context, vetOrg model.VetOrg) ([]model.Customer, error)
	findCustomerByVetOrgMutex       sync.RWMutex
	findCustomerByVetOrgArgsForCall []struct {
		ctx    context.Context
		vetOrg model.VetOrg
	}
	findCustomerByVetOrgReturns struct {
		result1 []model.Customer
		result2 error
	}
	findCustomerByVetOrgReturnsOnCall map[int]struct {
		result1 []model.Customer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCustomerService) CreateCustomer(ctx context.Context, user model.Customer) (*model.Customer, error) {
	fake.createCustomerMutex.Lock()
	ret, specificReturn := fake.createCustomerReturnsOnCall[len(fake.createCustomerArgsForCall)]
	fake.createCustomerArgsForCall = append(fake.createCustomerArgsForCall, struct {
		ctx  context.Context
		user model.Customer
	}{ctx, user})
	fake.recordInvocation("CreateCustomer", []interface{}{ctx, user})
	fake.createCustomerMutex.Unlock()
	if fake.CreateCustomerStub != nil {
		return fake.CreateCustomerStub(ctx, user)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCustomerReturns.result1, fake.createCustomerReturns.result2
}

func (fake *FakeCustomerService) CreateCustomerCallCount() int {
	fake.createCustomerMutex.RLock()
	defer fake.createCustomerMutex.RUnlock()
	return len(fake.createCustomerArgsForCall)
}

func (fake *FakeCustomerService) CreateCustomerArgsForCall(i int) (context.Context, model.Customer) {
	fake.createCustomerMutex.RLock()
	defer fake.createCustomerMutex.RUnlock()
	return fake.createCustomerArgsForCall[i].ctx, fake.createCustomerArgsForCall[i].user
}

func (fake *FakeCustomerService) CreateCustomerReturns(result1 *model.Customer, result2 error) {
	fake.CreateCustomerStub = nil
	fake.createCustomerReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) CreateCustomerReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.CreateCustomerStub = nil
	if fake.createCustomerReturnsOnCall == nil {
		fake.createCustomerReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.createCustomerReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) UpdateCustomer(ctx context.Context, user model.Customer) (*model.Customer, error) {
	fake.updateCustomerMutex.Lock()
	ret, specificReturn := fake.updateCustomerReturnsOnCall[len(fake.updateCustomerArgsForCall)]
	fake.updateCustomerArgsForCall = append(fake.updateCustomerArgsForCall, struct {
		ctx  context.Context
		user model.Customer
	}{ctx, user})
	fake.recordInvocation("UpdateCustomer", []interface{}{ctx, user})
	fake.updateCustomerMutex.Unlock()
	if fake.UpdateCustomerStub != nil {
		return fake.UpdateCustomerStub(ctx, user)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateCustomerReturns.result1, fake.updateCustomerReturns.result2
}

func (fake *FakeCustomerService) UpdateCustomerCallCount() int {
	fake.updateCustomerMutex.RLock()
	defer fake.updateCustomerMutex.RUnlock()
	return len(fake.updateCustomerArgsForCall)
}

func (fake *FakeCustomerService) UpdateCustomerArgsForCall(i int) (context.Context, model.Customer) {
	fake.updateCustomerMutex.RLock()
	defer fake.updateCustomerMutex.RUnlock()
	return fake.updateCustomerArgsForCall[i].ctx, fake.updateCustomerArgsForCall[i].user
}

func (fake *FakeCustomerService) UpdateCustomerReturns(result1 *model.Customer, result2 error) {
	fake.UpdateCustomerStub = nil
	fake.updateCustomerReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) UpdateCustomerReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.UpdateCustomerStub = nil
	if fake.updateCustomerReturnsOnCall == nil {
		fake.updateCustomerReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.updateCustomerReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) DeleteCustomer(ctx context.Context, user model.Customer) error {
	fake.deleteCustomerMutex.Lock()
	ret, specificReturn := fake.deleteCustomerReturnsOnCall[len(fake.deleteCustomerArgsForCall)]
	fake.deleteCustomerArgsForCall = append(fake.deleteCustomerArgsForCall, struct {
		ctx  context.Context
		user model.Customer
	}{ctx, user})
	fake.recordInvocation("DeleteCustomer", []interface{}{ctx, user})
	fake.deleteCustomerMutex.Unlock()
	if fake.DeleteCustomerStub != nil {
		return fake.DeleteCustomerStub(ctx, user)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteCustomerReturns.result1
}

func (fake *FakeCustomerService) DeleteCustomerCallCount() int {
	fake.deleteCustomerMutex.RLock()
	defer fake.deleteCustomerMutex.RUnlock()
	return len(fake.deleteCustomerArgsForCall)
}

func (fake *FakeCustomerService) DeleteCustomerArgsForCall(i int) (context.Context, model.Customer) {
	fake.deleteCustomerMutex.RLock()
	defer fake.deleteCustomerMutex.RUnlock()
	return fake.deleteCustomerArgsForCall[i].ctx, fake.deleteCustomerArgsForCall[i].user
}

func (fake *FakeCustomerService) DeleteCustomerReturns(result1 error) {
	fake.DeleteCustomerStub = nil
	fake.deleteCustomerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCustomerService) DeleteCustomerReturnsOnCall(i int, result1 error) {
	fake.DeleteCustomerStub = nil
	if fake.deleteCustomerReturnsOnCall == nil {
		fake.deleteCustomerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCustomerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCustomerService) UpdatePassword(ctx context.Context, user model.Customer, password string) (*model.Customer, error) {
	fake.updatePasswordMutex.Lock()
	ret, specificReturn := fake.updatePasswordReturnsOnCall[len(fake.updatePasswordArgsForCall)]
	fake.updatePasswordArgsForCall = append(fake.updatePasswordArgsForCall, struct {
		ctx      context.Context
		user     model.Customer
		password string
	}{ctx, user, password})
	fake.recordInvocation("UpdatePassword", []interface{}{ctx, user, password})
	fake.updatePasswordMutex.Unlock()
	if fake.UpdatePasswordStub != nil {
		return fake.UpdatePasswordStub(ctx, user, password)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updatePasswordReturns.result1, fake.updatePasswordReturns.result2
}

func (fake *FakeCustomerService) UpdatePasswordCallCount() int {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	return len(fake.updatePasswordArgsForCall)
}

func (fake *FakeCustomerService) UpdatePasswordArgsForCall(i int) (context.Context, model.Customer, string) {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	return fake.updatePasswordArgsForCall[i].ctx, fake.updatePasswordArgsForCall[i].user, fake.updatePasswordArgsForCall[i].password
}

func (fake *FakeCustomerService) UpdatePasswordReturns(result1 *model.Customer, result2 error) {
	fake.UpdatePasswordStub = nil
	fake.updatePasswordReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) UpdatePasswordReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.UpdatePasswordStub = nil
	if fake.updatePasswordReturnsOnCall == nil {
		fake.updatePasswordReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.updatePasswordReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) Login(ctx context.Context, userName string, password string) (*model.Customer, error) {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		ctx      context.Context
		userName string
		password string
	}{ctx, userName, password})
	fake.recordInvocation("Login", []interface{}{ctx, userName, password})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(ctx, userName, password)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loginReturns.result1, fake.loginReturns.result2
}

func (fake *FakeCustomerService) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeCustomerService) LoginArgsForCall(i int) (context.Context, string, string) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return fake.loginArgsForCall[i].ctx, fake.loginArgsForCall[i].userName, fake.loginArgsForCall[i].password
}

func (fake *FakeCustomerService) LoginReturns(result1 *model.Customer, result2 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) LoginReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByUserName(ctx context.Context, userName string) (*model.Customer, error) {
	fake.findCustomerByUserNameMutex.Lock()
	ret, specificReturn := fake.findCustomerByUserNameReturnsOnCall[len(fake.findCustomerByUserNameArgsForCall)]
	fake.findCustomerByUserNameArgsForCall = append(fake.findCustomerByUserNameArgsForCall, struct {
		ctx      context.Context
		userName string
	}{ctx, userName})
	fake.recordInvocation("FindCustomerByUserName", []interface{}{ctx, userName})
	fake.findCustomerByUserNameMutex.Unlock()
	if fake.FindCustomerByUserNameStub != nil {
		return fake.FindCustomerByUserNameStub(ctx, userName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCustomerByUserNameReturns.result1, fake.findCustomerByUserNameReturns.result2
}

func (fake *FakeCustomerService) FindCustomerByUserNameCallCount() int {
	fake.findCustomerByUserNameMutex.RLock()
	defer fake.findCustomerByUserNameMutex.RUnlock()
	return len(fake.findCustomerByUserNameArgsForCall)
}

func (fake *FakeCustomerService) FindCustomerByUserNameArgsForCall(i int) (context.Context, string) {
	fake.findCustomerByUserNameMutex.RLock()
	defer fake.findCustomerByUserNameMutex.RUnlock()
	return fake.findCustomerByUserNameArgsForCall[i].ctx, fake.findCustomerByUserNameArgsForCall[i].userName
}

func (fake *FakeCustomerService) FindCustomerByUserNameReturns(result1 *model.Customer, result2 error) {
	fake.FindCustomerByUserNameStub = nil
	fake.findCustomerByUserNameReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByUserNameReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.FindCustomerByUserNameStub = nil
	if fake.findCustomerByUserNameReturnsOnCall == nil {
		fake.findCustomerByUserNameReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.findCustomerByUserNameReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByID(ctx context.Context, userID uint) (*model.Customer, error) {
	fake.findCustomerByIDMutex.Lock()
	ret, specificReturn := fake.findCustomerByIDReturnsOnCall[len(fake.findCustomerByIDArgsForCall)]
	fake.findCustomerByIDArgsForCall = append(fake.findCustomerByIDArgsForCall, struct {
		ctx    context.Context
		userID uint
	}{ctx, userID})
	fake.recordInvocation("FindCustomerByID", []interface{}{ctx, userID})
	fake.findCustomerByIDMutex.Unlock()
	if fake.FindCustomerByIDStub != nil {
		return fake.FindCustomerByIDStub(ctx, userID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCustomerByIDReturns.result1, fake.findCustomerByIDReturns.result2
}

func (fake *FakeCustomerService) FindCustomerByIDCallCount() int {
	fake.findCustomerByIDMutex.RLock()
	defer fake.findCustomerByIDMutex.RUnlock()
	return len(fake.findCustomerByIDArgsForCall)
}

func (fake *FakeCustomerService) FindCustomerByIDArgsForCall(i int) (context.Context, uint) {
	fake.findCustomerByIDMutex.RLock()
	defer fake.findCustomerByIDMutex.RUnlock()
	return fake.findCustomerByIDArgsForCall[i].ctx, fake.findCustomerByIDArgsForCall[i].userID
}

func (fake *FakeCustomerService) FindCustomerByIDReturns(result1 *model.Customer, result2 error) {
	fake.FindCustomerByIDStub = nil
	fake.findCustomerByIDReturns = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByIDReturnsOnCall(i int, result1 *model.Customer, result2 error) {
	fake.FindCustomerByIDStub = nil
	if fake.findCustomerByIDReturnsOnCall == nil {
		fake.findCustomerByIDReturnsOnCall = make(map[int]struct {
			result1 *model.Customer
			result2 error
		})
	}
	fake.findCustomerByIDReturnsOnCall[i] = struct {
		result1 *model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.Customer, error) {
	fake.findCustomerByVetOrgMutex.Lock()
	ret, specificReturn := fake.findCustomerByVetOrgReturnsOnCall[len(fake.findCustomerByVetOrgArgsForCall)]
	fake.findCustomerByVetOrgArgsForCall = append(fake.findCustomerByVetOrgArgsForCall, struct {
		ctx    context.Context
		vetOrg model.VetOrg
	}{ctx, vetOrg})
	fake.recordInvocation("FindCustomerByVetOrg", []interface{}{ctx, vetOrg})
	fake.findCustomerByVetOrgMutex.Unlock()
	if fake.FindCustomerByVetOrgStub != nil {
		return fake.FindCustomerByVetOrgStub(ctx, vetOrg)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCustomerByVetOrgReturns.result1, fake.findCustomerByVetOrgReturns.result2
}

func (fake *FakeCustomerService) FindCustomerByVetOrgCallCount() int {
	fake.findCustomerByVetOrgMutex.RLock()
	defer fake.findCustomerByVetOrgMutex.RUnlock()
	return len(fake.findCustomerByVetOrgArgsForCall)
}

func (fake *FakeCustomerService) FindCustomerByVetOrgArgsForCall(i int) (context.Context, model.VetOrg) {
	fake.findCustomerByVetOrgMutex.RLock()
	defer fake.findCustomerByVetOrgMutex.RUnlock()
	return fake.findCustomerByVetOrgArgsForCall[i].ctx, fake.findCustomerByVetOrgArgsForCall[i].vetOrg
}

func (fake *FakeCustomerService) FindCustomerByVetOrgReturns(result1 []model.Customer, result2 error) {
	fake.FindCustomerByVetOrgStub = nil
	fake.findCustomerByVetOrgReturns = struct {
		result1 []model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) FindCustomerByVetOrgReturnsOnCall(i int, result1 []model.Customer, result2 error) {
	fake.FindCustomerByVetOrgStub = nil
	if fake.findCustomerByVetOrgReturnsOnCall == nil {
		fake.findCustomerByVetOrgReturnsOnCall = make(map[int]struct {
			result1 []model.Customer
			result2 error
		})
	}
	fake.findCustomerByVetOrgReturnsOnCall[i] = struct {
		result1 []model.Customer
		result2 error
	}{result1, result2}
}

func (fake *FakeCustomerService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCustomerMutex.RLock()
	defer fake.createCustomerMutex.RUnlock()
	fake.updateCustomerMutex.RLock()
	defer fake.updateCustomerMutex.RUnlock()
	fake.deleteCustomerMutex.RLock()
	defer fake.deleteCustomerMutex.RUnlock()
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.findCustomerByUserNameMutex.RLock()
	defer fake.findCustomerByUserNameMutex.RUnlock()
	fake.findCustomerByIDMutex.RLock()
	defer fake.findCustomerByIDMutex.RUnlock()
	fake.findCustomerByVetOrgMutex.RLock()
	defer fake.findCustomerByVetOrgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCustomerService) recordInvocation(key string, args []interface{}) {
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

var _ api.CustomerService = new(FakeCustomerService)
