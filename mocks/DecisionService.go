// Code generated by mockery 2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/mikhailbolshakov/decision/domain/decision"
	mock "github.com/stretchr/testify/mock"
)

// DecisionService is an autogenerated mock type for the DecisionService type
type DecisionService struct {
	mock.Mock
}

// MakeDecision provides a mock function with given fields: ctx, userId, problem
func (_m *DecisionService) MakeDecision(ctx context.Context, userId string, problem *domain.Problem) (*domain.Decision, error) {
	ret := _m.Called(ctx, userId, problem)

	var r0 *domain.Decision
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Problem) *domain.Decision); ok {
		r0 = rf(ctx, userId, problem)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Decision)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *domain.Problem) error); ok {
		r1 = rf(ctx, userId, problem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDecisionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDecisionService creates a new instance of DecisionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDecisionService(t mockConstructorTestingTNewDecisionService) *DecisionService {
	mock := &DecisionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}