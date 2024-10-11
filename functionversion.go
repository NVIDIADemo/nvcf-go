// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/NVIDIADemo/nvcf-go/internal/apijson"
	"github.com/NVIDIADemo/nvcf-go/internal/param"
	"github.com/NVIDIADemo/nvcf-go/internal/requestconfig"
	"github.com/NVIDIADemo/nvcf-go/option"
	"github.com/NVIDIADemo/nvcf-go/shared"
)

// FunctionVersionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionVersionService] method instead.
type FunctionVersionService struct {
	Options []option.RequestOption
}

// NewFunctionVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionVersionService(opts ...option.RequestOption) (r *FunctionVersionService) {
	r = &FunctionVersionService{}
	r.Options = opts
	return
}

// Creates a version of the specified function within the authenticated NVIDIA
// Cloud Account. Requires a bearer token with 'register_function' scope in the
// HTTP Authorization header.
func (r *FunctionVersionService) New(ctx context.Context, functionID string, body FunctionVersionNewParams, opts ...option.RequestOption) (res *shared.CreateFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists details of all the versions of the specified function in the authenticated
// NVIDIA Cloud Account. Requires either a bearer token or an api-key with
// 'list_functions' or 'list_functions_details' scopes in the HTTP Authorization
// header.
func (r *FunctionVersionService) GetAll(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.FunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FunctionVersionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[FunctionVersionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]FunctionVersionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[FunctionVersionNewParamsFunctionType] `json:"functionType"`
	// Data Transfer Object(DTO) representing a function ne
	Health param.Field[shared.HealthDTOParam] `json:"health"`
	// Health endpoint for the container or the helmChart
	HealthUri param.Field[string] `json:"healthUri" format:"uri"`
	// Optional Helm Chart
	HelmChart param.Field[string] `json:"helmChart" format:"uri"`
	// Helm Chart Service Name is required when helmChart property is specified
	HelmChartServiceName param.Field[string] `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running. Defaults to 8000
	// for Triton.
	InferencePort param.Field[int64] `json:"inferencePort"`
	// Optional set of models
	Models param.Field[[]FunctionVersionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]FunctionVersionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]FunctionVersionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type FunctionVersionNewParamsAPIBodyFormat string

const (
	FunctionVersionNewParamsAPIBodyFormatPredictV2 FunctionVersionNewParamsAPIBodyFormat = "PREDICT_V2"
	FunctionVersionNewParamsAPIBodyFormatCustom    FunctionVersionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r FunctionVersionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case FunctionVersionNewParamsAPIBodyFormatPredictV2, FunctionVersionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type FunctionVersionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionVersionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type FunctionVersionNewParamsFunctionType string

const (
	FunctionVersionNewParamsFunctionTypeDefault   FunctionVersionNewParamsFunctionType = "DEFAULT"
	FunctionVersionNewParamsFunctionTypeStreaming FunctionVersionNewParamsFunctionType = "STREAMING"
)

func (r FunctionVersionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionVersionNewParamsFunctionTypeDefault, FunctionVersionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type FunctionVersionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionVersionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type FunctionVersionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionVersionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type FunctionVersionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionVersionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
