// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/NVIDIADemo/nvcf-go/internal/apierror"
	"github.com/NVIDIADemo/nvcf-go/shared"
)

type Error = apierror.Error

// Parties authorized to invoke function
//
// This is an alias to an internal type.
type AuthorizedParties = shared.AuthorizedParties

// Data Transfer Object(DTO) representing a function with authorized accounts
//
// This is an alias to an internal type.
type AuthorizedPartiesFunction = shared.AuthorizedPartiesFunction

// Data Transfer Object(DTO) representing an authorized party.
//
// This is an alias to an internal type.
type AuthorizedPartyDTO = shared.AuthorizedPartyDTO

// Data Transfer Object(DTO) representing an authorized party.
//
// This is an alias to an internal type.
type AuthorizedPartyDTOParam = shared.AuthorizedPartyDTOParam

// Response body for create function request.
//
// This is an alias to an internal type.
type CreateFunctionResponse = shared.CreateFunctionResponse

// Response body with function details
//
// This is an alias to an internal type.
type Function = shared.Function

// Data Transfer Object (DTO) representing a function
//
// This is an alias to an internal type.
type FunctionDTO = shared.FunctionDTO

// Used to indicate a STREAMING function. Defaults to DEFAULT.
//
// This is an alias to an internal type.
type FunctionDTOFunctionType = shared.FunctionDTOFunctionType

// This is an alias to an internal value.
const FunctionDTOFunctionTypeDefault = shared.FunctionDTOFunctionTypeDefault

// This is an alias to an internal value.
const FunctionDTOFunctionTypeStreaming = shared.FunctionDTOFunctionTypeStreaming

// Function status
//
// This is an alias to an internal type.
type FunctionDTOStatus = shared.FunctionDTOStatus

// This is an alias to an internal value.
const FunctionDTOStatusActive = shared.FunctionDTOStatusActive

// This is an alias to an internal value.
const FunctionDTOStatusDeploying = shared.FunctionDTOStatusDeploying

// This is an alias to an internal value.
const FunctionDTOStatusError = shared.FunctionDTOStatusError

// This is an alias to an internal value.
const FunctionDTOStatusInactive = shared.FunctionDTOStatusInactive

// This is an alias to an internal value.
const FunctionDTOStatusDeleted = shared.FunctionDTOStatusDeleted

// Data Transfer Object(DTO) representing a spot instance
//
// This is an alias to an internal type.
type FunctionDTOActiveInstance = shared.FunctionDTOActiveInstance

// Instance status
//
// This is an alias to an internal type.
type FunctionDTOActiveInstancesInstanceStatus = shared.FunctionDTOActiveInstancesInstanceStatus

// This is an alias to an internal value.
const FunctionDTOActiveInstancesInstanceStatusActive = shared.FunctionDTOActiveInstancesInstanceStatusActive

// This is an alias to an internal value.
const FunctionDTOActiveInstancesInstanceStatusErrored = shared.FunctionDTOActiveInstancesInstanceStatusErrored

// This is an alias to an internal value.
const FunctionDTOActiveInstancesInstanceStatusPreempted = shared.FunctionDTOActiveInstancesInstanceStatusPreempted

// This is an alias to an internal value.
const FunctionDTOActiveInstancesInstanceStatusDeleted = shared.FunctionDTOActiveInstancesInstanceStatusDeleted

// Invocation request body format
//
// This is an alias to an internal type.
type FunctionDTOAPIBodyFormat = shared.FunctionDTOAPIBodyFormat

// This is an alias to an internal value.
const FunctionDTOAPIBodyFormatPredictV2 = shared.FunctionDTOAPIBodyFormatPredictV2

// This is an alias to an internal value.
const FunctionDTOAPIBodyFormatCustom = shared.FunctionDTOAPIBodyFormatCustom

// Data Transfer Object(DTO) representing a container environment entry
//
// This is an alias to an internal type.
type FunctionDTOContainerEnvironment = shared.FunctionDTOContainerEnvironment

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionDTOModel = shared.FunctionDTOModel

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionDTOResource = shared.FunctionDTOResource

// Response body containing list of functions
//
// This is an alias to an internal type.
type FunctionsResponse = shared.FunctionsResponse

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type HealthDTO = shared.HealthDTO

// HTTP/gPRC protocol type for health endpoint
//
// This is an alias to an internal type.
type HealthDTOProtocol = shared.HealthDTOProtocol

// This is an alias to an internal value.
const HealthDTOProtocolHTTP = shared.HealthDTOProtocolHTTP

// This is an alias to an internal value.
const HealthDTOProtocolGRpc = shared.HealthDTOProtocolGRpc

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type HealthDTOParam = shared.HealthDTOParam

// Response body with result from a request for executing a job/task as a cloud
// function using a GPU powered spot/on-demand instance.
//
// This is an alias to an internal type.
type InvokeFunctionResponse = shared.InvokeFunctionResponse

// Status of the task/job executing cloud function.
//
// This is an alias to an internal type.
type InvokeFunctionResponseStatus = shared.InvokeFunctionResponseStatus

// This is an alias to an internal value.
const InvokeFunctionResponseStatusErrored = shared.InvokeFunctionResponseStatusErrored

// This is an alias to an internal value.
const InvokeFunctionResponseStatusInProgress = shared.InvokeFunctionResponseStatusInProgress

// This is an alias to an internal value.
const InvokeFunctionResponseStatusFulfilled = shared.InvokeFunctionResponseStatusFulfilled

// This is an alias to an internal value.
const InvokeFunctionResponseStatusPendingEvaluation = shared.InvokeFunctionResponseStatusPendingEvaluation

// This is an alias to an internal value.
const InvokeFunctionResponseStatusRejected = shared.InvokeFunctionResponseStatusRejected

// Request queue details of all the functions with same id in an account
//
// This is an alias to an internal type.
type Queues = shared.Queues

// Data Transfer Object(DTO) representing a request queue for function version
//
// This is an alias to an internal type.
type QueuesQueue = shared.QueuesQueue

// Function status
//
// This is an alias to an internal type.
type QueuesQueuesFunctionStatus = shared.QueuesQueuesFunctionStatus

// This is an alias to an internal value.
const QueuesQueuesFunctionStatusActive = shared.QueuesQueuesFunctionStatusActive

// This is an alias to an internal value.
const QueuesQueuesFunctionStatusDeploying = shared.QueuesQueuesFunctionStatusDeploying

// This is an alias to an internal value.
const QueuesQueuesFunctionStatusError = shared.QueuesQueuesFunctionStatusError

// This is an alias to an internal value.
const QueuesQueuesFunctionStatusInactive = shared.QueuesQueuesFunctionStatusInactive

// This is an alias to an internal value.
const QueuesQueuesFunctionStatusDeleted = shared.QueuesQueuesFunctionStatusDeleted
