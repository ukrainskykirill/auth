// Code generated by http://github.com/gojuno/minimock (v3.3.14). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/ukrainskykirill/auth/internal/cache.UserCache -o user_cache_minimock.go -n UserCacheMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/ukrainskykirill/auth/internal/model"
)

// UserCacheMock implements cache.UserCache
type UserCacheMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, userIn *model.User) (err error)
	inspectFuncCreate   func(ctx context.Context, userIn *model.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserCacheMockCreate

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mUserCacheMockDelete

	funcGet          func(ctx context.Context, id int64) (up1 *model.User, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserCacheMockGet
}

// NewUserCacheMock returns a mock for cache.UserCache
func NewUserCacheMock(t minimock.Tester) *UserCacheMock {
	m := &UserCacheMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserCacheMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserCacheMockCreateParams{}

	m.DeleteMock = mUserCacheMockDelete{mock: m}
	m.DeleteMock.callArgs = []*UserCacheMockDeleteParams{}

	m.GetMock = mUserCacheMockGet{mock: m}
	m.GetMock.callArgs = []*UserCacheMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserCacheMockCreate struct {
	optional           bool
	mock               *UserCacheMock
	defaultExpectation *UserCacheMockCreateExpectation
	expectations       []*UserCacheMockCreateExpectation

	callArgs []*UserCacheMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCacheMockCreateExpectation specifies expectation struct of the UserCache.Create
type UserCacheMockCreateExpectation struct {
	mock      *UserCacheMock
	params    *UserCacheMockCreateParams
	paramPtrs *UserCacheMockCreateParamPtrs
	results   *UserCacheMockCreateResults
	Counter   uint64
}

// UserCacheMockCreateParams contains parameters of the UserCache.Create
type UserCacheMockCreateParams struct {
	ctx    context.Context
	userIn *model.User
}

// UserCacheMockCreateParamPtrs contains pointers to parameters of the UserCache.Create
type UserCacheMockCreateParamPtrs struct {
	ctx    *context.Context
	userIn **model.User
}

// UserCacheMockCreateResults contains results of the UserCache.Create
type UserCacheMockCreateResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mUserCacheMockCreate) Optional() *mUserCacheMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for UserCache.Create
func (mmCreate *mUserCacheMockCreate) Expect(ctx context.Context, userIn *model.User) *mUserCacheMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserCacheMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &UserCacheMockCreateParams{ctx, userIn}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for UserCache.Create
func (mmCreate *mUserCacheMockCreate) ExpectCtxParam1(ctx context.Context) *mUserCacheMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserCacheMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UserCacheMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx

	return mmCreate
}

// ExpectUserInParam2 sets up expected param userIn for UserCache.Create
func (mmCreate *mUserCacheMockCreate) ExpectUserInParam2(userIn *model.User) *mUserCacheMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserCacheMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UserCacheMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.userIn = &userIn

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserCache.Create
func (mmCreate *mUserCacheMockCreate) Inspect(f func(ctx context.Context, userIn *model.User)) *mUserCacheMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserCacheMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserCache.Create
func (mmCreate *mUserCacheMockCreate) Return(err error) *UserCacheMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserCacheMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserCacheMockCreateResults{err}
	return mmCreate.mock
}

// Set uses given function f to mock the UserCache.Create method
func (mmCreate *mUserCacheMockCreate) Set(f func(ctx context.Context, userIn *model.User) (err error)) *UserCacheMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserCache.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserCache.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UserCache.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserCacheMockCreate) When(ctx context.Context, userIn *model.User) *UserCacheMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserCacheMock.Create mock is already set by Set")
	}

	expectation := &UserCacheMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UserCacheMockCreateParams{ctx, userIn},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserCache.Create return parameters for the expectation previously defined by the When method
func (e *UserCacheMockCreateExpectation) Then(err error) *UserCacheMock {
	e.results = &UserCacheMockCreateResults{err}
	return e.mock
}

// Times sets number of times UserCache.Create should be invoked
func (mmCreate *mUserCacheMockCreate) Times(n uint64) *mUserCacheMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of UserCacheMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	return mmCreate
}

func (mmCreate *mUserCacheMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements cache.UserCache
func (mmCreate *UserCacheMock) Create(ctx context.Context, userIn *model.User) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, userIn)
	}

	mm_params := UserCacheMockCreateParams{ctx, userIn}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := UserCacheMockCreateParams{ctx, userIn}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("UserCacheMock.Create got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.userIn != nil && !minimock.Equal(*mm_want_ptrs.userIn, mm_got.userIn) {
				mmCreate.t.Errorf("UserCacheMock.Create got unexpected parameter userIn, want: %#v, got: %#v%s\n", *mm_want_ptrs.userIn, mm_got.userIn, minimock.Diff(*mm_want_ptrs.userIn, mm_got.userIn))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserCacheMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserCacheMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, userIn)
	}
	mmCreate.t.Fatalf("Unexpected call to UserCacheMock.Create. %v %v", ctx, userIn)
	return
}

// CreateAfterCounter returns a count of finished UserCacheMock.Create invocations
func (mmCreate *UserCacheMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserCacheMock.Create invocations
func (mmCreate *UserCacheMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserCacheMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserCacheMockCreate) Calls() []*UserCacheMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserCacheMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserCacheMock) MinimockCreateDone() bool {
	if m.CreateMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMock.invocationsDone()
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserCacheMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCacheMock.Create with params: %#v", *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCacheMock.Create")
		} else {
			m.t.Errorf("Expected call to UserCacheMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Error("Expected call to UserCacheMock.Create")
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCacheMock.Create but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), afterCreateCounter)
	}
}

type mUserCacheMockDelete struct {
	optional           bool
	mock               *UserCacheMock
	defaultExpectation *UserCacheMockDeleteExpectation
	expectations       []*UserCacheMockDeleteExpectation

	callArgs []*UserCacheMockDeleteParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCacheMockDeleteExpectation specifies expectation struct of the UserCache.Delete
type UserCacheMockDeleteExpectation struct {
	mock      *UserCacheMock
	params    *UserCacheMockDeleteParams
	paramPtrs *UserCacheMockDeleteParamPtrs
	results   *UserCacheMockDeleteResults
	Counter   uint64
}

// UserCacheMockDeleteParams contains parameters of the UserCache.Delete
type UserCacheMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// UserCacheMockDeleteParamPtrs contains pointers to parameters of the UserCache.Delete
type UserCacheMockDeleteParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UserCacheMockDeleteResults contains results of the UserCache.Delete
type UserCacheMockDeleteResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDelete *mUserCacheMockDelete) Optional() *mUserCacheMockDelete {
	mmDelete.optional = true
	return mmDelete
}

// Expect sets up expected params for UserCache.Delete
func (mmDelete *mUserCacheMockDelete) Expect(ctx context.Context, id int64) *mUserCacheMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserCacheMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.paramPtrs != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by ExpectParams functions")
	}

	mmDelete.defaultExpectation.params = &UserCacheMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// ExpectCtxParam1 sets up expected param ctx for UserCache.Delete
func (mmDelete *mUserCacheMockDelete) ExpectCtxParam1(ctx context.Context) *mUserCacheMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserCacheMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &UserCacheMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.ctx = &ctx

	return mmDelete
}

// ExpectIdParam2 sets up expected param id for UserCache.Delete
func (mmDelete *mUserCacheMockDelete) ExpectIdParam2(id int64) *mUserCacheMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserCacheMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &UserCacheMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.id = &id

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the UserCache.Delete
func (mmDelete *mUserCacheMockDelete) Inspect(f func(ctx context.Context, id int64)) *mUserCacheMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for UserCacheMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by UserCache.Delete
func (mmDelete *mUserCacheMockDelete) Return(err error) *UserCacheMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserCacheMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &UserCacheMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the UserCache.Delete method
func (mmDelete *mUserCacheMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *UserCacheMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the UserCache.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the UserCache.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the UserCache.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mUserCacheMockDelete) When(ctx context.Context, id int64) *UserCacheMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserCacheMock.Delete mock is already set by Set")
	}

	expectation := &UserCacheMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &UserCacheMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up UserCache.Delete return parameters for the expectation previously defined by the When method
func (e *UserCacheMockDeleteExpectation) Then(err error) *UserCacheMock {
	e.results = &UserCacheMockDeleteResults{err}
	return e.mock
}

// Times sets number of times UserCache.Delete should be invoked
func (mmDelete *mUserCacheMockDelete) Times(n uint64) *mUserCacheMockDelete {
	if n == 0 {
		mmDelete.mock.t.Fatalf("Times of UserCacheMock.Delete mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDelete.expectedInvocations, n)
	return mmDelete
}

func (mmDelete *mUserCacheMockDelete) invocationsDone() bool {
	if len(mmDelete.expectations) == 0 && mmDelete.defaultExpectation == nil && mmDelete.mock.funcDelete == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDelete.mock.afterDeleteCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDelete.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Delete implements cache.UserCache
func (mmDelete *UserCacheMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := UserCacheMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_want_ptrs := mmDelete.DeleteMock.defaultExpectation.paramPtrs

		mm_got := UserCacheMockDeleteParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDelete.t.Errorf("UserCacheMock.Delete got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmDelete.t.Errorf("UserCacheMock.Delete got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("UserCacheMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the UserCacheMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to UserCacheMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished UserCacheMock.Delete invocations
func (mmDelete *UserCacheMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of UserCacheMock.Delete invocations
func (mmDelete *UserCacheMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to UserCacheMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mUserCacheMockDelete) Calls() []*UserCacheMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*UserCacheMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *UserCacheMock) MinimockDeleteDone() bool {
	if m.DeleteMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DeleteMock.invocationsDone()
}

// MinimockDeleteInspect logs each unmet expectation
func (m *UserCacheMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCacheMock.Delete with params: %#v", *e.params)
		}
	}

	afterDeleteCounter := mm_atomic.LoadUint64(&m.afterDeleteCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && afterDeleteCounter < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCacheMock.Delete")
		} else {
			m.t.Errorf("Expected call to UserCacheMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && afterDeleteCounter < 1 {
		m.t.Error("Expected call to UserCacheMock.Delete")
	}

	if !m.DeleteMock.invocationsDone() && afterDeleteCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCacheMock.Delete but found %d calls",
			mm_atomic.LoadUint64(&m.DeleteMock.expectedInvocations), afterDeleteCounter)
	}
}

type mUserCacheMockGet struct {
	optional           bool
	mock               *UserCacheMock
	defaultExpectation *UserCacheMockGetExpectation
	expectations       []*UserCacheMockGetExpectation

	callArgs []*UserCacheMockGetParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserCacheMockGetExpectation specifies expectation struct of the UserCache.Get
type UserCacheMockGetExpectation struct {
	mock      *UserCacheMock
	params    *UserCacheMockGetParams
	paramPtrs *UserCacheMockGetParamPtrs
	results   *UserCacheMockGetResults
	Counter   uint64
}

// UserCacheMockGetParams contains parameters of the UserCache.Get
type UserCacheMockGetParams struct {
	ctx context.Context
	id  int64
}

// UserCacheMockGetParamPtrs contains pointers to parameters of the UserCache.Get
type UserCacheMockGetParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UserCacheMockGetResults contains results of the UserCache.Get
type UserCacheMockGetResults struct {
	up1 *model.User
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mUserCacheMockGet) Optional() *mUserCacheMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for UserCache.Get
func (mmGet *mUserCacheMockGet) Expect(ctx context.Context, id int64) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &UserCacheMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for UserCache.Get
func (mmGet *mUserCacheMockGet) ExpectCtxParam1(ctx context.Context) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserCacheMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx

	return mmGet
}

// ExpectIdParam2 sets up expected param id for UserCache.Get
func (mmGet *mUserCacheMockGet) ExpectIdParam2(id int64) *mUserCacheMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserCacheMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.id = &id

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserCache.Get
func (mmGet *mUserCacheMockGet) Inspect(f func(ctx context.Context, id int64)) *mUserCacheMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserCacheMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserCache.Get
func (mmGet *mUserCacheMockGet) Return(up1 *model.User, err error) *UserCacheMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserCacheMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserCacheMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the UserCache.Get method
func (mmGet *mUserCacheMockGet) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *UserCacheMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserCache.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserCache.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UserCache.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserCacheMockGet) When(ctx context.Context, id int64) *UserCacheMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserCacheMock.Get mock is already set by Set")
	}

	expectation := &UserCacheMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserCacheMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserCache.Get return parameters for the expectation previously defined by the When method
func (e *UserCacheMockGetExpectation) Then(up1 *model.User, err error) *UserCacheMock {
	e.results = &UserCacheMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times UserCache.Get should be invoked
func (mmGet *mUserCacheMockGet) Times(n uint64) *mUserCacheMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of UserCacheMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	return mmGet
}

func (mmGet *mUserCacheMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements cache.UserCache
func (mmGet *UserCacheMock) Get(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := UserCacheMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := UserCacheMockGetParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("UserCacheMock.Get got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmGet.t.Errorf("UserCacheMock.Get got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserCacheMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserCacheMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserCacheMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserCacheMock.Get invocations
func (mmGet *UserCacheMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserCacheMock.Get invocations
func (mmGet *UserCacheMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserCacheMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserCacheMockGet) Calls() []*UserCacheMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserCacheMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserCacheMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *UserCacheMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserCacheMock.Get with params: %#v", *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserCacheMock.Get")
		} else {
			m.t.Errorf("Expected call to UserCacheMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Error("Expected call to UserCacheMock.Get")
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to UserCacheMock.Get but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), afterGetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserCacheMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockDeleteInspect()

			m.MinimockGetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserCacheMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserCacheMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone()
}
