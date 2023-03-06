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

type Architecture string

const (
	Architecture_x86_64 Architecture = "x86_64"
	Architecture_arm64  Architecture = "arm64"
)

type CodeSigningPolicy string

const (
	CodeSigningPolicy_Warn    CodeSigningPolicy = "Warn"
	CodeSigningPolicy_Enforce CodeSigningPolicy = "Enforce"
)

type EndPointType string

const (
	EndPointType_KAFKA_BOOTSTRAP_SERVERS EndPointType = "KAFKA_BOOTSTRAP_SERVERS"
)

type EventSourcePosition string

const (
	EventSourcePosition_TRIM_HORIZON EventSourcePosition = "TRIM_HORIZON"
	EventSourcePosition_LATEST       EventSourcePosition = "LATEST"
	EventSourcePosition_AT_TIMESTAMP EventSourcePosition = "AT_TIMESTAMP"
)

type FunctionResponseType string

const (
	FunctionResponseType_ReportBatchItemFailures FunctionResponseType = "ReportBatchItemFailures"
)

type FunctionURLAuthType string

const (
	FunctionURLAuthType_NONE    FunctionURLAuthType = "NONE"
	FunctionURLAuthType_AWS_IAM FunctionURLAuthType = "AWS_IAM"
)

type FunctionVersion string

const (
	FunctionVersion_ALL FunctionVersion = "ALL"
)

type InvocationType string

const (
	InvocationType_Event           InvocationType = "Event"
	InvocationType_RequestResponse InvocationType = "RequestResponse"
	InvocationType_DryRun          InvocationType = "DryRun"
)

type LastUpdateStatus string

const (
	LastUpdateStatus_Successful LastUpdateStatus = "Successful"
	LastUpdateStatus_Failed     LastUpdateStatus = "Failed"
	LastUpdateStatus_InProgress LastUpdateStatus = "InProgress"
)

type LastUpdateStatusReasonCode string

const (
	LastUpdateStatusReasonCode_EniLimitExceeded            LastUpdateStatusReasonCode = "EniLimitExceeded"
	LastUpdateStatusReasonCode_InsufficientRolePermissions LastUpdateStatusReasonCode = "InsufficientRolePermissions"
	LastUpdateStatusReasonCode_InvalidConfiguration        LastUpdateStatusReasonCode = "InvalidConfiguration"
	LastUpdateStatusReasonCode_InternalError               LastUpdateStatusReasonCode = "InternalError"
	LastUpdateStatusReasonCode_SubnetOutOfIPAddresses      LastUpdateStatusReasonCode = "SubnetOutOfIPAddresses"
	LastUpdateStatusReasonCode_InvalidSubnet               LastUpdateStatusReasonCode = "InvalidSubnet"
	LastUpdateStatusReasonCode_InvalidSecurityGroup        LastUpdateStatusReasonCode = "InvalidSecurityGroup"
	LastUpdateStatusReasonCode_ImageDeleted                LastUpdateStatusReasonCode = "ImageDeleted"
	LastUpdateStatusReasonCode_ImageAccessDenied           LastUpdateStatusReasonCode = "ImageAccessDenied"
	LastUpdateStatusReasonCode_InvalidImage                LastUpdateStatusReasonCode = "InvalidImage"
	LastUpdateStatusReasonCode_KMSKeyAccessDenied          LastUpdateStatusReasonCode = "KMSKeyAccessDenied"
	LastUpdateStatusReasonCode_KMSKeyNotFound              LastUpdateStatusReasonCode = "KMSKeyNotFound"
	LastUpdateStatusReasonCode_InvalidStateKMSKey          LastUpdateStatusReasonCode = "InvalidStateKMSKey"
	LastUpdateStatusReasonCode_DisabledKMSKey              LastUpdateStatusReasonCode = "DisabledKMSKey"
	LastUpdateStatusReasonCode_EFSIOError                  LastUpdateStatusReasonCode = "EFSIOError"
	LastUpdateStatusReasonCode_EFSMountConnectivityError   LastUpdateStatusReasonCode = "EFSMountConnectivityError"
	LastUpdateStatusReasonCode_EFSMountFailure             LastUpdateStatusReasonCode = "EFSMountFailure"
	LastUpdateStatusReasonCode_EFSMountTimeout             LastUpdateStatusReasonCode = "EFSMountTimeout"
	LastUpdateStatusReasonCode_InvalidRuntime              LastUpdateStatusReasonCode = "InvalidRuntime"
	LastUpdateStatusReasonCode_InvalidZipFileException     LastUpdateStatusReasonCode = "InvalidZipFileException"
	LastUpdateStatusReasonCode_FunctionError               LastUpdateStatusReasonCode = "FunctionError"
)

type LogType string

const (
	LogType_None LogType = "None"
	LogType_Tail LogType = "Tail"
)

type PackageType string

const (
	PackageType_Zip   PackageType = "Zip"
	PackageType_Image PackageType = "Image"
)

type ProvisionedConcurrencyStatusEnum string

const (
	ProvisionedConcurrencyStatusEnum_IN_PROGRESS ProvisionedConcurrencyStatusEnum = "IN_PROGRESS"
	ProvisionedConcurrencyStatusEnum_READY       ProvisionedConcurrencyStatusEnum = "READY"
	ProvisionedConcurrencyStatusEnum_FAILED      ProvisionedConcurrencyStatusEnum = "FAILED"
)

type Runtime string

const (
	Runtime_nodejs         Runtime = "nodejs"
	Runtime_nodejs4_3      Runtime = "nodejs4.3"
	Runtime_nodejs6_10     Runtime = "nodejs6.10"
	Runtime_nodejs8_10     Runtime = "nodejs8.10"
	Runtime_nodejs10_x     Runtime = "nodejs10.x"
	Runtime_nodejs12_x     Runtime = "nodejs12.x"
	Runtime_nodejs14_x     Runtime = "nodejs14.x"
	Runtime_nodejs16_x     Runtime = "nodejs16.x"
	Runtime_java8          Runtime = "java8"
	Runtime_java8_al2      Runtime = "java8.al2"
	Runtime_java11         Runtime = "java11"
	Runtime_python2_7      Runtime = "python2.7"
	Runtime_python3_6      Runtime = "python3.6"
	Runtime_python3_7      Runtime = "python3.7"
	Runtime_python3_8      Runtime = "python3.8"
	Runtime_python3_9      Runtime = "python3.9"
	Runtime_dotnetcore1_0  Runtime = "dotnetcore1.0"
	Runtime_dotnetcore2_0  Runtime = "dotnetcore2.0"
	Runtime_dotnetcore2_1  Runtime = "dotnetcore2.1"
	Runtime_dotnetcore3_1  Runtime = "dotnetcore3.1"
	Runtime_dotnet6        Runtime = "dotnet6"
	Runtime_nodejs4_3_edge Runtime = "nodejs4.3-edge"
	Runtime_go1_x          Runtime = "go1.x"
	Runtime_ruby2_5        Runtime = "ruby2.5"
	Runtime_ruby2_7        Runtime = "ruby2.7"
	Runtime_provided       Runtime = "provided"
	Runtime_provided_al2   Runtime = "provided.al2"
	Runtime_nodejs18_x     Runtime = "nodejs18.x"
)

type SnapStartApplyOn string

const (
	SnapStartApplyOn_PublishedVersions SnapStartApplyOn = "PublishedVersions"
	SnapStartApplyOn_None              SnapStartApplyOn = "None"
)

type SnapStartOptimizationStatus string

const (
	SnapStartOptimizationStatus_On  SnapStartOptimizationStatus = "On"
	SnapStartOptimizationStatus_Off SnapStartOptimizationStatus = "Off"
)

type SourceAccessType string

const (
	SourceAccessType_BASIC_AUTH                  SourceAccessType = "BASIC_AUTH"
	SourceAccessType_VPC_SUBNET                  SourceAccessType = "VPC_SUBNET"
	SourceAccessType_VPC_SECURITY_GROUP          SourceAccessType = "VPC_SECURITY_GROUP"
	SourceAccessType_SASL_SCRAM_512_AUTH         SourceAccessType = "SASL_SCRAM_512_AUTH"
	SourceAccessType_SASL_SCRAM_256_AUTH         SourceAccessType = "SASL_SCRAM_256_AUTH"
	SourceAccessType_VIRTUAL_HOST                SourceAccessType = "VIRTUAL_HOST"
	SourceAccessType_CLIENT_CERTIFICATE_TLS_AUTH SourceAccessType = "CLIENT_CERTIFICATE_TLS_AUTH"
	SourceAccessType_SERVER_ROOT_CA_CERTIFICATE  SourceAccessType = "SERVER_ROOT_CA_CERTIFICATE"
)

type State string

const (
	State_Pending  State = "Pending"
	State_Active   State = "Active"
	State_Inactive State = "Inactive"
	State_Failed   State = "Failed"
)

type StateReasonCode string

const (
	StateReasonCode_Idle                        StateReasonCode = "Idle"
	StateReasonCode_Creating                    StateReasonCode = "Creating"
	StateReasonCode_Restoring                   StateReasonCode = "Restoring"
	StateReasonCode_EniLimitExceeded            StateReasonCode = "EniLimitExceeded"
	StateReasonCode_InsufficientRolePermissions StateReasonCode = "InsufficientRolePermissions"
	StateReasonCode_InvalidConfiguration        StateReasonCode = "InvalidConfiguration"
	StateReasonCode_InternalError               StateReasonCode = "InternalError"
	StateReasonCode_SubnetOutOfIPAddresses      StateReasonCode = "SubnetOutOfIPAddresses"
	StateReasonCode_InvalidSubnet               StateReasonCode = "InvalidSubnet"
	StateReasonCode_InvalidSecurityGroup        StateReasonCode = "InvalidSecurityGroup"
	StateReasonCode_ImageDeleted                StateReasonCode = "ImageDeleted"
	StateReasonCode_ImageAccessDenied           StateReasonCode = "ImageAccessDenied"
	StateReasonCode_InvalidImage                StateReasonCode = "InvalidImage"
	StateReasonCode_KMSKeyAccessDenied          StateReasonCode = "KMSKeyAccessDenied"
	StateReasonCode_KMSKeyNotFound              StateReasonCode = "KMSKeyNotFound"
	StateReasonCode_InvalidStateKMSKey          StateReasonCode = "InvalidStateKMSKey"
	StateReasonCode_DisabledKMSKey              StateReasonCode = "DisabledKMSKey"
	StateReasonCode_EFSIOError                  StateReasonCode = "EFSIOError"
	StateReasonCode_EFSMountConnectivityError   StateReasonCode = "EFSMountConnectivityError"
	StateReasonCode_EFSMountFailure             StateReasonCode = "EFSMountFailure"
	StateReasonCode_EFSMountTimeout             StateReasonCode = "EFSMountTimeout"
	StateReasonCode_InvalidRuntime              StateReasonCode = "InvalidRuntime"
	StateReasonCode_InvalidZipFileException     StateReasonCode = "InvalidZipFileException"
	StateReasonCode_FunctionError               StateReasonCode = "FunctionError"
)

type ThrottleReason string

const (
	ThrottleReason_ConcurrentInvocationLimitExceeded                 ThrottleReason = "ConcurrentInvocationLimitExceeded"
	ThrottleReason_FunctionInvocationRateLimitExceeded               ThrottleReason = "FunctionInvocationRateLimitExceeded"
	ThrottleReason_ReservedFunctionConcurrentInvocationLimitExceeded ThrottleReason = "ReservedFunctionConcurrentInvocationLimitExceeded"
	ThrottleReason_ReservedFunctionInvocationRateLimitExceeded       ThrottleReason = "ReservedFunctionInvocationRateLimitExceeded"
	ThrottleReason_CallerRateLimitExceeded                           ThrottleReason = "CallerRateLimitExceeded"
	ThrottleReason_ConcurrentSnapshotCreateLimitExceeded             ThrottleReason = "ConcurrentSnapshotCreateLimitExceeded"
)

type TracingMode string

const (
	TracingMode_Active      TracingMode = "Active"
	TracingMode_PassThrough TracingMode = "PassThrough"
)
