// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PreWorkflowHookRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockPreWorkflowHookRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPreWorkflowHookRunner(options ...pegomock.Option) *MockPreWorkflowHookRunner {
	mock := &MockPreWorkflowHookRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPreWorkflowHookRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPreWorkflowHookRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPreWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, shell string, shellArgs string, path string) (string, string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPreWorkflowHookRunner().")
	}
	_params := []pegomock.Param{ctx, command, shell, shellArgs, path}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Run", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 string
	var _ret1 string
	var _ret2 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(string)
		}
		if _result[2] != nil {
			_ret2 = _result[2].(error)
		}
	}
	return _ret0, _ret1, _ret2
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledOnce() *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPreWorkflowHookRunner struct {
	mock                   *MockPreWorkflowHookRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPreWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, shell string, shellArgs string, path string) *MockPreWorkflowHookRunner_Run_OngoingVerification {
	_params := []pegomock.Param{ctx, command, shell, shellArgs, path}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", _params, verifier.timeout)
	return &MockPreWorkflowHookRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPreWorkflowHookRunner_Run_OngoingVerification struct {
	mock              *MockPreWorkflowHookRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPreWorkflowHookRunner_Run_OngoingVerification) GetCapturedArguments() (models.WorkflowHookCommandContext, string, string, string, string) {
	ctx, command, shell, shellArgs, path := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], command[len(command)-1], shell[len(shell)-1], shellArgs[len(shellArgs)-1], path[len(path)-1]
}

func (c *MockPreWorkflowHookRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.WorkflowHookCommandContext, _param1 []string, _param2 []string, _param3 []string, _param4 []string) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]models.WorkflowHookCommandContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(models.WorkflowHookCommandContext)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]string, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(string)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]string, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(string)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]string, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(string)
			}
		}
		if len(_params) > 4 {
			_param4 = make([]string, len(c.methodInvocations))
			for u, param := range _params[4] {
				_param4[u] = param.(string)
			}
		}
	}
	return
}
