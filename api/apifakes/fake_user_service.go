// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"context"
	"sync"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
)

type FakeUserService struct {
	CreateUserStub        func(context.Context, model.User) (*model.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
	}
	createUserReturns struct {
		result1 *model.User
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	DeleteUserStub        func(context.Context, model.User) error
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
	}
	deleteUserReturns struct {
		result1 error
	}
	deleteUserReturnsOnCall map[int]struct {
		result1 error
	}
	FindUserByIDStub        func(context.Context, uint) (*model.User, error)
	findUserByIDMutex       sync.RWMutex
	findUserByIDArgsForCall []struct {
		arg1 context.Context
		arg2 uint
	}
	findUserByIDReturns struct {
		result1 *model.User
		result2 error
	}
	findUserByIDReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	FindUserByUserNameStub        func(context.Context, string) (*model.User, error)
	findUserByUserNameMutex       sync.RWMutex
	findUserByUserNameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	findUserByUserNameReturns struct {
		result1 *model.User
		result2 error
	}
	findUserByUserNameReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	LoginStub        func(context.Context, string, string) (*model.User, error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	loginReturns struct {
		result1 *model.User
		result2 error
	}
	loginReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	UpdatePasswordStub        func(context.Context, model.User, string) (*model.User, error)
	updatePasswordMutex       sync.RWMutex
	updatePasswordArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
		arg3 string
	}
	updatePasswordReturns struct {
		result1 *model.User
		result2 error
	}
	updatePasswordReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	UpdateUserStub        func(context.Context, model.User) (*model.User, error)
	updateUserMutex       sync.RWMutex
	updateUserArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
	}
	updateUserReturns struct {
		result1 *model.User
		result2 error
	}
	updateUserReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserService) CreateUser(arg1 context.Context, arg2 model.User) (*model.User, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
	}{arg1, arg2})
	stub := fake.CreateUserStub
	fakeReturns := fake.createUserReturns
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2})
	fake.createUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUserService) CreateUserCalls(stub func(context.Context, model.User) (*model.User, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeUserService) CreateUserArgsForCall(i int) (context.Context, model.User) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) CreateUserReturns(result1 *model.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) CreateUserReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) DeleteUser(arg1 context.Context, arg2 model.User) error {
	fake.deleteUserMutex.Lock()
	ret, specificReturn := fake.deleteUserReturnsOnCall[len(fake.deleteUserArgsForCall)]
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
	}{arg1, arg2})
	stub := fake.DeleteUserStub
	fakeReturns := fake.deleteUserReturns
	fake.recordInvocation("DeleteUser", []interface{}{arg1, arg2})
	fake.deleteUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserService) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeUserService) DeleteUserCalls(stub func(context.Context, model.User) error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = stub
}

func (fake *FakeUserService) DeleteUserArgsForCall(i int) (context.Context, model.User) {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	argsForCall := fake.deleteUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) DeleteUserReturns(result1 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) DeleteUserReturnsOnCall(i int, result1 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	if fake.deleteUserReturnsOnCall == nil {
		fake.deleteUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) FindUserByID(arg1 context.Context, arg2 uint) (*model.User, error) {
	fake.findUserByIDMutex.Lock()
	ret, specificReturn := fake.findUserByIDReturnsOnCall[len(fake.findUserByIDArgsForCall)]
	fake.findUserByIDArgsForCall = append(fake.findUserByIDArgsForCall, struct {
		arg1 context.Context
		arg2 uint
	}{arg1, arg2})
	stub := fake.FindUserByIDStub
	fakeReturns := fake.findUserByIDReturns
	fake.recordInvocation("FindUserByID", []interface{}{arg1, arg2})
	fake.findUserByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) FindUserByIDCallCount() int {
	fake.findUserByIDMutex.RLock()
	defer fake.findUserByIDMutex.RUnlock()
	return len(fake.findUserByIDArgsForCall)
}

func (fake *FakeUserService) FindUserByIDCalls(stub func(context.Context, uint) (*model.User, error)) {
	fake.findUserByIDMutex.Lock()
	defer fake.findUserByIDMutex.Unlock()
	fake.FindUserByIDStub = stub
}

func (fake *FakeUserService) FindUserByIDArgsForCall(i int) (context.Context, uint) {
	fake.findUserByIDMutex.RLock()
	defer fake.findUserByIDMutex.RUnlock()
	argsForCall := fake.findUserByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) FindUserByIDReturns(result1 *model.User, result2 error) {
	fake.findUserByIDMutex.Lock()
	defer fake.findUserByIDMutex.Unlock()
	fake.FindUserByIDStub = nil
	fake.findUserByIDReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindUserByIDReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.findUserByIDMutex.Lock()
	defer fake.findUserByIDMutex.Unlock()
	fake.FindUserByIDStub = nil
	if fake.findUserByIDReturnsOnCall == nil {
		fake.findUserByIDReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.findUserByIDReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindUserByUserName(arg1 context.Context, arg2 string) (*model.User, error) {
	fake.findUserByUserNameMutex.Lock()
	ret, specificReturn := fake.findUserByUserNameReturnsOnCall[len(fake.findUserByUserNameArgsForCall)]
	fake.findUserByUserNameArgsForCall = append(fake.findUserByUserNameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.FindUserByUserNameStub
	fakeReturns := fake.findUserByUserNameReturns
	fake.recordInvocation("FindUserByUserName", []interface{}{arg1, arg2})
	fake.findUserByUserNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) FindUserByUserNameCallCount() int {
	fake.findUserByUserNameMutex.RLock()
	defer fake.findUserByUserNameMutex.RUnlock()
	return len(fake.findUserByUserNameArgsForCall)
}

func (fake *FakeUserService) FindUserByUserNameCalls(stub func(context.Context, string) (*model.User, error)) {
	fake.findUserByUserNameMutex.Lock()
	defer fake.findUserByUserNameMutex.Unlock()
	fake.FindUserByUserNameStub = stub
}

func (fake *FakeUserService) FindUserByUserNameArgsForCall(i int) (context.Context, string) {
	fake.findUserByUserNameMutex.RLock()
	defer fake.findUserByUserNameMutex.RUnlock()
	argsForCall := fake.findUserByUserNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) FindUserByUserNameReturns(result1 *model.User, result2 error) {
	fake.findUserByUserNameMutex.Lock()
	defer fake.findUserByUserNameMutex.Unlock()
	fake.FindUserByUserNameStub = nil
	fake.findUserByUserNameReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindUserByUserNameReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.findUserByUserNameMutex.Lock()
	defer fake.findUserByUserNameMutex.Unlock()
	fake.FindUserByUserNameStub = nil
	if fake.findUserByUserNameReturnsOnCall == nil {
		fake.findUserByUserNameReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.findUserByUserNameReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Login(arg1 context.Context, arg2 string, arg3 string) (*model.User, error) {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.LoginStub
	fakeReturns := fake.loginReturns
	fake.recordInvocation("Login", []interface{}{arg1, arg2, arg3})
	fake.loginMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeUserService) LoginCalls(stub func(context.Context, string, string) (*model.User, error)) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeUserService) LoginArgsForCall(i int) (context.Context, string, string) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserService) LoginReturns(result1 *model.User, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) LoginReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) UpdatePassword(arg1 context.Context, arg2 model.User, arg3 string) (*model.User, error) {
	fake.updatePasswordMutex.Lock()
	ret, specificReturn := fake.updatePasswordReturnsOnCall[len(fake.updatePasswordArgsForCall)]
	fake.updatePasswordArgsForCall = append(fake.updatePasswordArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UpdatePasswordStub
	fakeReturns := fake.updatePasswordReturns
	fake.recordInvocation("UpdatePassword", []interface{}{arg1, arg2, arg3})
	fake.updatePasswordMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) UpdatePasswordCallCount() int {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	return len(fake.updatePasswordArgsForCall)
}

func (fake *FakeUserService) UpdatePasswordCalls(stub func(context.Context, model.User, string) (*model.User, error)) {
	fake.updatePasswordMutex.Lock()
	defer fake.updatePasswordMutex.Unlock()
	fake.UpdatePasswordStub = stub
}

func (fake *FakeUserService) UpdatePasswordArgsForCall(i int) (context.Context, model.User, string) {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	argsForCall := fake.updatePasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserService) UpdatePasswordReturns(result1 *model.User, result2 error) {
	fake.updatePasswordMutex.Lock()
	defer fake.updatePasswordMutex.Unlock()
	fake.UpdatePasswordStub = nil
	fake.updatePasswordReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) UpdatePasswordReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.updatePasswordMutex.Lock()
	defer fake.updatePasswordMutex.Unlock()
	fake.UpdatePasswordStub = nil
	if fake.updatePasswordReturnsOnCall == nil {
		fake.updatePasswordReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.updatePasswordReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) UpdateUser(arg1 context.Context, arg2 model.User) (*model.User, error) {
	fake.updateUserMutex.Lock()
	ret, specificReturn := fake.updateUserReturnsOnCall[len(fake.updateUserArgsForCall)]
	fake.updateUserArgsForCall = append(fake.updateUserArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
	}{arg1, arg2})
	stub := fake.UpdateUserStub
	fakeReturns := fake.updateUserReturns
	fake.recordInvocation("UpdateUser", []interface{}{arg1, arg2})
	fake.updateUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) UpdateUserCallCount() int {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	return len(fake.updateUserArgsForCall)
}

func (fake *FakeUserService) UpdateUserCalls(stub func(context.Context, model.User) (*model.User, error)) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = stub
}

func (fake *FakeUserService) UpdateUserArgsForCall(i int) (context.Context, model.User) {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	argsForCall := fake.updateUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) UpdateUserReturns(result1 *model.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	fake.updateUserReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) UpdateUserReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	if fake.updateUserReturnsOnCall == nil {
		fake.updateUserReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.updateUserReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	fake.findUserByIDMutex.RLock()
	defer fake.findUserByIDMutex.RUnlock()
	fake.findUserByUserNameMutex.RLock()
	defer fake.findUserByUserNameMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserService) recordInvocation(key string, args []interface{}) {
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

var _ api.UserService = new(FakeUserService)
