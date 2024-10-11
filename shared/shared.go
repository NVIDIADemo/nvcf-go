// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/NVIDIADemo/nvcf-go/internal/apijson"
	"github.com/NVIDIADemo/nvcf-go/internal/param"
)

// Parties authorized to invoke function
type AuthorizedParties struct {
	// Data Transfer Object(DTO) representing a function with authorized accounts
	Function AuthorizedPartiesFunction `json:"function,required"`
	JSON     authorizedPartiesJSON     `json:"-"`
}

// authorizedPartiesJSON contains the JSON metadata for the struct
// [AuthorizedParties]
type authorizedPartiesJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizedParties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartiesJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function with authorized accounts
type AuthorizedPartiesFunction struct {
	// Function id
	ID string `json:"id,required" format:"uuid"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Authorized parties allowed to invoke the function
	AuthorizedParties []AuthorizedPartyDTO `json:"authorizedParties"`
	// Function version id
	VersionID string                        `json:"versionId" format:"uuid"`
	JSON      authorizedPartiesFunctionJSON `json:"-"`
}

// authorizedPartiesFunctionJSON contains the JSON metadata for the struct
// [AuthorizedPartiesFunction]
type authorizedPartiesFunctionJSON struct {
	ID                apijson.Field
	NcaID             apijson.Field
	AuthorizedParties apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuthorizedPartiesFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartiesFunctionJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedPartyDTO struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID string `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID string                 `json:"clientId"`
	JSON     authorizedPartyDTOJSON `json:"-"`
}

// authorizedPartyDTOJSON contains the JSON metadata for the struct
// [AuthorizedPartyDTO]
type authorizedPartyDTOJSON struct {
	NcaID       apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizedPartyDTO) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartyDTOJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedPartyDTOParam struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID param.Field[string] `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID param.Field[string] `json:"clientId"`
}

func (r AuthorizedPartyDTOParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for create function request.
type CreateFunctionResponse struct {
	// Data Transfer Object (DTO) representing a function
	Function FunctionDTO                `json:"function,required"`
	JSON     createFunctionResponseJSON `json:"-"`
}

// createFunctionResponseJSON contains the JSON metadata for the struct
// [CreateFunctionResponse]
type createFunctionResponseJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseJSON) RawJSON() string {
	return r.raw
}

// Response body with function details
type Function struct {
	// Data Transfer Object (DTO) representing a function
	Function FunctionDTO  `json:"function,required"`
	JSON     functionJSON `json:"-"`
}

// functionJSON contains the JSON metadata for the struct [Function]
type functionJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object (DTO) representing a function
type FunctionDTO struct {
	// Unique function id
	ID string `json:"id,required" format:"uuid"`
	// Function creation timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Used to indicate a STREAMING function. Defaults to DEFAULT.
	FunctionType FunctionDTOFunctionType `json:"functionType,required"`
	// Health endpoint for the container or helmChart
	HealthUri string `json:"healthUri,required" format:"uri"`
	// Function name
	Name string `json:"name,required"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Function status
	Status FunctionDTOStatus `json:"status,required"`
	// Unique function version id
	VersionID string `json:"versionId,required" format:"uuid"`
	// List of active instances for this function.
	ActiveInstances []FunctionDTOActiveInstance `json:"activeInstances"`
	// Invocation request body format
	APIBodyFormat FunctionDTOAPIBodyFormat `json:"apiBodyFormat"`
	// Args used to launch the container
	ContainerArgs string `json:"containerArgs"`
	// Environment settings used to launch the container
	ContainerEnvironment []FunctionDTOContainerEnvironment `json:"containerEnvironment"`
	// Optional custom container
	ContainerImage string `json:"containerImage" format:"uri"`
	// Function/version description
	Description string `json:"description"`
	// Data Transfer Object(DTO) representing a function ne
	Health HealthDTO `json:"health"`
	// Optional Helm Chart
	HelmChart string `json:"helmChart" format:"uri"`
	// Helm Chart Service Name specified only when helmChart property is specified
	HelmChartServiceName string `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running - defaults to 8000
	// for Triton
	InferencePort int64 `json:"inferencePort"`
	// Entrypoint for invoking the container to process requests
	InferenceURL string `json:"inferenceUrl" format:"uri"`
	// Optional set of models
	Models []FunctionDTOModel `json:"models"`
	// Indicates whether the function is owned by another account. If the account that
	// is being used to lookup functions happens to be authorized to invoke/list this
	// function which is owned by a different account, then this field is set to true
	// and ncaId will contain the id of the account that owns the function. Otherwise,
	// this field is not set as it defaults to false.
	OwnedByDifferentAccount bool `json:"ownedByDifferentAccount"`
	// Optional set of resources.
	Resources []FunctionDTOResource `json:"resources"`
	// Optional secret names
	Secrets []string `json:"secrets"`
	// Optional set of tags. Maximum allowed number of tags per function is 64. Maximum
	// length of each tag is 128 chars.
	Tags []string        `json:"tags"`
	JSON functionDTOJSON `json:"-"`
}

// functionDTOJSON contains the JSON metadata for the struct [FunctionDTO]
type functionDTOJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	FunctionType            apijson.Field
	HealthUri               apijson.Field
	Name                    apijson.Field
	NcaID                   apijson.Field
	Status                  apijson.Field
	VersionID               apijson.Field
	ActiveInstances         apijson.Field
	APIBodyFormat           apijson.Field
	ContainerArgs           apijson.Field
	ContainerEnvironment    apijson.Field
	ContainerImage          apijson.Field
	Description             apijson.Field
	Health                  apijson.Field
	HelmChart               apijson.Field
	HelmChartServiceName    apijson.Field
	InferencePort           apijson.Field
	InferenceURL            apijson.Field
	Models                  apijson.Field
	OwnedByDifferentAccount apijson.Field
	Resources               apijson.Field
	Secrets                 apijson.Field
	Tags                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *FunctionDTO) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionDTOJSON) RawJSON() string {
	return r.raw
}

// Used to indicate a STREAMING function. Defaults to DEFAULT.
type FunctionDTOFunctionType string

const (
	FunctionDTOFunctionTypeDefault   FunctionDTOFunctionType = "DEFAULT"
	FunctionDTOFunctionTypeStreaming FunctionDTOFunctionType = "STREAMING"
)

func (r FunctionDTOFunctionType) IsKnown() bool {
	switch r {
	case FunctionDTOFunctionTypeDefault, FunctionDTOFunctionTypeStreaming:
		return true
	}
	return false
}

// Function status
type FunctionDTOStatus string

const (
	FunctionDTOStatusActive    FunctionDTOStatus = "ACTIVE"
	FunctionDTOStatusDeploying FunctionDTOStatus = "DEPLOYING"
	FunctionDTOStatusError     FunctionDTOStatus = "ERROR"
	FunctionDTOStatusInactive  FunctionDTOStatus = "INACTIVE"
	FunctionDTOStatusDeleted   FunctionDTOStatus = "DELETED"
)

func (r FunctionDTOStatus) IsKnown() bool {
	switch r {
	case FunctionDTOStatusActive, FunctionDTOStatusDeploying, FunctionDTOStatusError, FunctionDTOStatusInactive, FunctionDTOStatusDeleted:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a spot instance
type FunctionDTOActiveInstance struct {
	// Backend where the instance is running
	Backend string `json:"backend"`
	// Function executing on the instance
	FunctionID string `json:"functionId" format:"uuid"`
	// Function version executing on the instance
	FunctionVersionID string `json:"functionVersionId" format:"uuid"`
	// GPU name powering the instance
	GPU string `json:"gpu"`
	// Instance creation timestamp
	InstanceCreatedAt time.Time `json:"instanceCreatedAt" format:"date-time"`
	// Unique id of the instance
	InstanceID string `json:"instanceId"`
	// Instance status
	InstanceStatus FunctionDTOActiveInstancesInstanceStatus `json:"instanceStatus"`
	// GPU instance-type powering the instance
	InstanceType string `json:"instanceType"`
	// Instance's last updated timestamp
	InstanceUpdatedAt time.Time `json:"instanceUpdatedAt" format:"date-time"`
	// Location such as zone name or region where the instance is running
	Location string `json:"location"`
	// NVIDIA Cloud Account Id that owns the function running on the instance
	NcaID string `json:"ncaId"`
	// SIS request-id used to launch this instance
	SisRequestID string                        `json:"sisRequestId" format:"uuid"`
	JSON         functionDTOActiveInstanceJSON `json:"-"`
}

// functionDTOActiveInstanceJSON contains the JSON metadata for the struct
// [FunctionDTOActiveInstance]
type functionDTOActiveInstanceJSON struct {
	Backend           apijson.Field
	FunctionID        apijson.Field
	FunctionVersionID apijson.Field
	GPU               apijson.Field
	InstanceCreatedAt apijson.Field
	InstanceID        apijson.Field
	InstanceStatus    apijson.Field
	InstanceType      apijson.Field
	InstanceUpdatedAt apijson.Field
	Location          apijson.Field
	NcaID             apijson.Field
	SisRequestID      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FunctionDTOActiveInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionDTOActiveInstanceJSON) RawJSON() string {
	return r.raw
}

// Instance status
type FunctionDTOActiveInstancesInstanceStatus string

const (
	FunctionDTOActiveInstancesInstanceStatusActive    FunctionDTOActiveInstancesInstanceStatus = "ACTIVE"
	FunctionDTOActiveInstancesInstanceStatusErrored   FunctionDTOActiveInstancesInstanceStatus = "ERRORED"
	FunctionDTOActiveInstancesInstanceStatusPreempted FunctionDTOActiveInstancesInstanceStatus = "PREEMPTED"
	FunctionDTOActiveInstancesInstanceStatusDeleted   FunctionDTOActiveInstancesInstanceStatus = "DELETED"
)

func (r FunctionDTOActiveInstancesInstanceStatus) IsKnown() bool {
	switch r {
	case FunctionDTOActiveInstancesInstanceStatusActive, FunctionDTOActiveInstancesInstanceStatusErrored, FunctionDTOActiveInstancesInstanceStatusPreempted, FunctionDTOActiveInstancesInstanceStatusDeleted:
		return true
	}
	return false
}

// Invocation request body format
type FunctionDTOAPIBodyFormat string

const (
	FunctionDTOAPIBodyFormatPredictV2 FunctionDTOAPIBodyFormat = "PREDICT_V2"
	FunctionDTOAPIBodyFormatCustom    FunctionDTOAPIBodyFormat = "CUSTOM"
)

func (r FunctionDTOAPIBodyFormat) IsKnown() bool {
	switch r {
	case FunctionDTOAPIBodyFormatPredictV2, FunctionDTOAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type FunctionDTOContainerEnvironment struct {
	// Container environment key
	Key string `json:"key,required"`
	// Container environment value
	Value string                              `json:"value,required"`
	JSON  functionDTOContainerEnvironmentJSON `json:"-"`
}

// functionDTOContainerEnvironmentJSON contains the JSON metadata for the struct
// [FunctionDTOContainerEnvironment]
type functionDTOContainerEnvironmentJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionDTOContainerEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionDTOContainerEnvironmentJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an artifact
type FunctionDTOModel struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string               `json:"version,required"`
	JSON    functionDTOModelJSON `json:"-"`
}

// functionDTOModelJSON contains the JSON metadata for the struct
// [FunctionDTOModel]
type functionDTOModelJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionDTOModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionDTOModelJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an artifact
type FunctionDTOResource struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                  `json:"version,required"`
	JSON    functionDTOResourceJSON `json:"-"`
}

// functionDTOResourceJSON contains the JSON metadata for the struct
// [FunctionDTOResource]
type functionDTOResourceJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionDTOResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionDTOResourceJSON) RawJSON() string {
	return r.raw
}

// Response body containing list of functions
type FunctionsResponse struct {
	// List of functions
	Functions []FunctionDTO         `json:"functions,required"`
	JSON      functionsResponseJSON `json:"-"`
}

// functionsResponseJSON contains the JSON metadata for the struct
// [FunctionsResponse]
type functionsResponseJSON struct {
	Functions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionsResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function ne
type HealthDTO struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode int64 `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port int64 `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol HealthDTOProtocol `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout string `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri  string        `json:"uri,required" format:"uri"`
	JSON healthDTOJSON `json:"-"`
}

// healthDTOJSON contains the JSON metadata for the struct [HealthDTO]
type healthDTOJSON struct {
	ExpectedStatusCode apijson.Field
	Port               apijson.Field
	Protocol           apijson.Field
	Timeout            apijson.Field
	Uri                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HealthDTO) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDTOJSON) RawJSON() string {
	return r.raw
}

// HTTP/gPRC protocol type for health endpoint
type HealthDTOProtocol string

const (
	HealthDTOProtocolHTTP HealthDTOProtocol = "HTTP"
	HealthDTOProtocolGRpc HealthDTOProtocol = "gRPC"
)

func (r HealthDTOProtocol) IsKnown() bool {
	switch r {
	case HealthDTOProtocolHTTP, HealthDTOProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a function ne
type HealthDTOParam struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode param.Field[int64] `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port param.Field[int64] `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol param.Field[HealthDTOProtocol] `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout param.Field[string] `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri param.Field[string] `json:"uri,required" format:"uri"`
}

func (r HealthDTOParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body with result from a request for executing a job/task as a cloud
// function using a GPU powered spot/on-demand instance.
type InvokeFunctionResponse struct {
	// Error code from the container while executing cloud function.
	ErrorCode int64 `json:"errorCode"`
	// Progress indicator for the task/job executing cloud function.
	PercentComplete int64 `json:"percentComplete"`
	// Request id
	ReqID string `json:"reqId" format:"uuid"`
	// Response/result of size < 5MB size for the task/job executing cloud function.
	Response string `json:"response"`
	// For large results, responseReference will be a pre-signeddownload URL.
	ResponseReference string `json:"responseReference" format:"url"`
	// Status of the task/job executing cloud function.
	Status InvokeFunctionResponseStatus `json:"status"`
	JSON   invokeFunctionResponseJSON   `json:"-"`
}

// invokeFunctionResponseJSON contains the JSON metadata for the struct
// [InvokeFunctionResponse]
type invokeFunctionResponseJSON struct {
	ErrorCode         apijson.Field
	PercentComplete   apijson.Field
	ReqID             apijson.Field
	Response          apijson.Field
	ResponseReference apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvokeFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invokeFunctionResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the task/job executing cloud function.
type InvokeFunctionResponseStatus string

const (
	InvokeFunctionResponseStatusErrored           InvokeFunctionResponseStatus = "errored"
	InvokeFunctionResponseStatusInProgress        InvokeFunctionResponseStatus = "in-progress"
	InvokeFunctionResponseStatusFulfilled         InvokeFunctionResponseStatus = "fulfilled"
	InvokeFunctionResponseStatusPendingEvaluation InvokeFunctionResponseStatus = "pending-evaluation"
	InvokeFunctionResponseStatusRejected          InvokeFunctionResponseStatus = "rejected"
)

func (r InvokeFunctionResponseStatus) IsKnown() bool {
	switch r {
	case InvokeFunctionResponseStatusErrored, InvokeFunctionResponseStatusInProgress, InvokeFunctionResponseStatusFulfilled, InvokeFunctionResponseStatusPendingEvaluation, InvokeFunctionResponseStatusRejected:
		return true
	}
	return false
}

// Request queue details of all the functions with same id in an account
type QueuesResponse struct {
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Details of all the queues associated with same named functions
	Queues []QueuesResponseQueue `json:"queues,required"`
	JSON   queuesResponseJSON    `json:"-"`
}

// queuesResponseJSON contains the JSON metadata for the struct [QueuesResponse]
type queuesResponseJSON struct {
	FunctionID  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueuesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queuesResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a request queue for function version
type QueuesResponseQueue struct {
	// Function name
	FunctionName string `json:"functionName,required"`
	// Function status
	FunctionStatus QueuesResponseQueuesFunctionStatus `json:"functionStatus,required"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// Approximate number of messages in the request queue
	QueueDepth int64                   `json:"queueDepth"`
	JSON       queuesResponseQueueJSON `json:"-"`
}

// queuesResponseQueueJSON contains the JSON metadata for the struct
// [QueuesResponseQueue]
type queuesResponseQueueJSON struct {
	FunctionName      apijson.Field
	FunctionStatus    apijson.Field
	FunctionVersionID apijson.Field
	QueueDepth        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QueuesResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queuesResponseQueueJSON) RawJSON() string {
	return r.raw
}

// Function status
type QueuesResponseQueuesFunctionStatus string

const (
	QueuesResponseQueuesFunctionStatusActive    QueuesResponseQueuesFunctionStatus = "ACTIVE"
	QueuesResponseQueuesFunctionStatusDeploying QueuesResponseQueuesFunctionStatus = "DEPLOYING"
	QueuesResponseQueuesFunctionStatusError     QueuesResponseQueuesFunctionStatus = "ERROR"
	QueuesResponseQueuesFunctionStatusInactive  QueuesResponseQueuesFunctionStatus = "INACTIVE"
	QueuesResponseQueuesFunctionStatusDeleted   QueuesResponseQueuesFunctionStatus = "DELETED"
)

func (r QueuesResponseQueuesFunctionStatus) IsKnown() bool {
	switch r {
	case QueuesResponseQueuesFunctionStatusActive, QueuesResponseQueuesFunctionStatusDeploying, QueuesResponseQueuesFunctionStatusError, QueuesResponseQueuesFunctionStatusInactive, QueuesResponseQueuesFunctionStatusDeleted:
		return true
	}
	return false
}
