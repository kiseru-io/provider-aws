/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type ExecutionRedriveFilter string

const (
	ExecutionRedriveFilter_REDRIVEN     ExecutionRedriveFilter = "REDRIVEN"
	ExecutionRedriveFilter_NOT_REDRIVEN ExecutionRedriveFilter = "NOT_REDRIVEN"
)

type ExecutionRedriveStatus string

const (
	ExecutionRedriveStatus_REDRIVABLE            ExecutionRedriveStatus = "REDRIVABLE"
	ExecutionRedriveStatus_NOT_REDRIVABLE        ExecutionRedriveStatus = "NOT_REDRIVABLE"
	ExecutionRedriveStatus_REDRIVABLE_BY_MAP_RUN ExecutionRedriveStatus = "REDRIVABLE_BY_MAP_RUN"
)

type ExecutionStatus string

const (
	ExecutionStatus_RUNNING         ExecutionStatus = "RUNNING"
	ExecutionStatus_SUCCEEDED       ExecutionStatus = "SUCCEEDED"
	ExecutionStatus_FAILED          ExecutionStatus = "FAILED"
	ExecutionStatus_TIMED_OUT       ExecutionStatus = "TIMED_OUT"
	ExecutionStatus_ABORTED         ExecutionStatus = "ABORTED"
	ExecutionStatus_PENDING_REDRIVE ExecutionStatus = "PENDING_REDRIVE"
)

type HistoryEventType string

const (
	HistoryEventType_ActivityFailed               HistoryEventType = "ActivityFailed"
	HistoryEventType_ActivityScheduled            HistoryEventType = "ActivityScheduled"
	HistoryEventType_ActivityScheduleFailed       HistoryEventType = "ActivityScheduleFailed"
	HistoryEventType_ActivityStarted              HistoryEventType = "ActivityStarted"
	HistoryEventType_ActivitySucceeded            HistoryEventType = "ActivitySucceeded"
	HistoryEventType_ActivityTimedOut             HistoryEventType = "ActivityTimedOut"
	HistoryEventType_ChoiceStateEntered           HistoryEventType = "ChoiceStateEntered"
	HistoryEventType_ChoiceStateExited            HistoryEventType = "ChoiceStateExited"
	HistoryEventType_ExecutionAborted             HistoryEventType = "ExecutionAborted"
	HistoryEventType_ExecutionFailed              HistoryEventType = "ExecutionFailed"
	HistoryEventType_ExecutionStarted             HistoryEventType = "ExecutionStarted"
	HistoryEventType_ExecutionSucceeded           HistoryEventType = "ExecutionSucceeded"
	HistoryEventType_ExecutionTimedOut            HistoryEventType = "ExecutionTimedOut"
	HistoryEventType_FailStateEntered             HistoryEventType = "FailStateEntered"
	HistoryEventType_LambdaFunctionFailed         HistoryEventType = "LambdaFunctionFailed"
	HistoryEventType_LambdaFunctionScheduled      HistoryEventType = "LambdaFunctionScheduled"
	HistoryEventType_LambdaFunctionScheduleFailed HistoryEventType = "LambdaFunctionScheduleFailed"
	HistoryEventType_LambdaFunctionStarted        HistoryEventType = "LambdaFunctionStarted"
	HistoryEventType_LambdaFunctionStartFailed    HistoryEventType = "LambdaFunctionStartFailed"
	HistoryEventType_LambdaFunctionSucceeded      HistoryEventType = "LambdaFunctionSucceeded"
	HistoryEventType_LambdaFunctionTimedOut       HistoryEventType = "LambdaFunctionTimedOut"
	HistoryEventType_MapIterationAborted          HistoryEventType = "MapIterationAborted"
	HistoryEventType_MapIterationFailed           HistoryEventType = "MapIterationFailed"
	HistoryEventType_MapIterationStarted          HistoryEventType = "MapIterationStarted"
	HistoryEventType_MapIterationSucceeded        HistoryEventType = "MapIterationSucceeded"
	HistoryEventType_MapStateAborted              HistoryEventType = "MapStateAborted"
	HistoryEventType_MapStateEntered              HistoryEventType = "MapStateEntered"
	HistoryEventType_MapStateExited               HistoryEventType = "MapStateExited"
	HistoryEventType_MapStateFailed               HistoryEventType = "MapStateFailed"
	HistoryEventType_MapStateStarted              HistoryEventType = "MapStateStarted"
	HistoryEventType_MapStateSucceeded            HistoryEventType = "MapStateSucceeded"
	HistoryEventType_ParallelStateAborted         HistoryEventType = "ParallelStateAborted"
	HistoryEventType_ParallelStateEntered         HistoryEventType = "ParallelStateEntered"
	HistoryEventType_ParallelStateExited          HistoryEventType = "ParallelStateExited"
	HistoryEventType_ParallelStateFailed          HistoryEventType = "ParallelStateFailed"
	HistoryEventType_ParallelStateStarted         HistoryEventType = "ParallelStateStarted"
	HistoryEventType_ParallelStateSucceeded       HistoryEventType = "ParallelStateSucceeded"
	HistoryEventType_PassStateEntered             HistoryEventType = "PassStateEntered"
	HistoryEventType_PassStateExited              HistoryEventType = "PassStateExited"
	HistoryEventType_SucceedStateEntered          HistoryEventType = "SucceedStateEntered"
	HistoryEventType_SucceedStateExited           HistoryEventType = "SucceedStateExited"
	HistoryEventType_TaskFailed                   HistoryEventType = "TaskFailed"
	HistoryEventType_TaskScheduled                HistoryEventType = "TaskScheduled"
	HistoryEventType_TaskStarted                  HistoryEventType = "TaskStarted"
	HistoryEventType_TaskStartFailed              HistoryEventType = "TaskStartFailed"
	HistoryEventType_TaskStateAborted             HistoryEventType = "TaskStateAborted"
	HistoryEventType_TaskStateEntered             HistoryEventType = "TaskStateEntered"
	HistoryEventType_TaskStateExited              HistoryEventType = "TaskStateExited"
	HistoryEventType_TaskSubmitFailed             HistoryEventType = "TaskSubmitFailed"
	HistoryEventType_TaskSubmitted                HistoryEventType = "TaskSubmitted"
	HistoryEventType_TaskSucceeded                HistoryEventType = "TaskSucceeded"
	HistoryEventType_TaskTimedOut                 HistoryEventType = "TaskTimedOut"
	HistoryEventType_WaitStateAborted             HistoryEventType = "WaitStateAborted"
	HistoryEventType_WaitStateEntered             HistoryEventType = "WaitStateEntered"
	HistoryEventType_WaitStateExited              HistoryEventType = "WaitStateExited"
	HistoryEventType_MapRunAborted                HistoryEventType = "MapRunAborted"
	HistoryEventType_MapRunFailed                 HistoryEventType = "MapRunFailed"
	HistoryEventType_MapRunStarted                HistoryEventType = "MapRunStarted"
	HistoryEventType_MapRunSucceeded              HistoryEventType = "MapRunSucceeded"
	HistoryEventType_ExecutionRedriven            HistoryEventType = "ExecutionRedriven"
	HistoryEventType_MapRunRedriven               HistoryEventType = "MapRunRedriven"
)

type LogLevel string

const (
	LogLevel_ALL   LogLevel = "ALL"
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_FATAL LogLevel = "FATAL"
	LogLevel_OFF   LogLevel = "OFF"
)

type MapRunStatus string

const (
	MapRunStatus_RUNNING   MapRunStatus = "RUNNING"
	MapRunStatus_SUCCEEDED MapRunStatus = "SUCCEEDED"
	MapRunStatus_FAILED    MapRunStatus = "FAILED"
	MapRunStatus_ABORTED   MapRunStatus = "ABORTED"
)

type StateMachineStatus_SDK string

const (
	StateMachineStatus_SDK_ACTIVE   StateMachineStatus_SDK = "ACTIVE"
	StateMachineStatus_SDK_DELETING StateMachineStatus_SDK = "DELETING"
)

type StateMachineType string

const (
	StateMachineType_STANDARD StateMachineType = "STANDARD"
	StateMachineType_EXPRESS  StateMachineType = "EXPRESS"
)

type SyncExecutionStatus string

const (
	SyncExecutionStatus_SUCCEEDED SyncExecutionStatus = "SUCCEEDED"
	SyncExecutionStatus_FAILED    SyncExecutionStatus = "FAILED"
	SyncExecutionStatus_TIMED_OUT SyncExecutionStatus = "TIMED_OUT"
)

type ValidationExceptionReason string

const (
	ValidationExceptionReason_API_DOES_NOT_SUPPORT_LABELED_ARNS ValidationExceptionReason = "API_DOES_NOT_SUPPORT_LABELED_ARNS"
	ValidationExceptionReason_MISSING_REQUIRED_PARAMETER        ValidationExceptionReason = "MISSING_REQUIRED_PARAMETER"
	ValidationExceptionReason_CANNOT_UPDATE_COMPLETED_MAP_RUN   ValidationExceptionReason = "CANNOT_UPDATE_COMPLETED_MAP_RUN"
	ValidationExceptionReason_INVALID_ROUTING_CONFIGURATION     ValidationExceptionReason = "INVALID_ROUTING_CONFIGURATION"
)
