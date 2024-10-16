// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/nvcf-go/internal/apijson"
	"github.com/NVIDIADemo/nvcf-go/internal/apiquery"
	"github.com/NVIDIADemo/nvcf-go/internal/param"
	"github.com/NVIDIADemo/nvcf-go/internal/requestconfig"
	"github.com/NVIDIADemo/nvcf-go/option"
	"github.com/NVIDIADemo/nvcf-go/shared"
)

// FunctionService contains methods and other services that help with interacting
// with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	Options  []option.RequestOption
	Versions *FunctionVersionService
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r *FunctionService) {
	r = &FunctionService{}
	r.Options = opts
	r.Versions = NewFunctionVersionService(opts...)
	return
}

// Creates a new function within the authenticated NVIDIA Cloud Account. Requires a
// bearer token with 'register_function' scope in the HTTP Authorization header.
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *shared.FunctionDTO, err error) {
	var env shared.CreateFunctionResponse
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Function
	return
}

// Lists all the functions associated with the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *FunctionService) GetAll(ctx context.Context, query FunctionGetAllParams, opts ...option.RequestOption) (res *shared.FunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FunctionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[FunctionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]FunctionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[FunctionNewParamsFunctionType] `json:"functionType"`
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
	Models param.Field[[]FunctionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]FunctionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]FunctionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type FunctionNewParamsAPIBodyFormat string

const (
	FunctionNewParamsAPIBodyFormatPredictV2 FunctionNewParamsAPIBodyFormat = "PREDICT_V2"
	FunctionNewParamsAPIBodyFormatCustom    FunctionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r FunctionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case FunctionNewParamsAPIBodyFormatPredictV2, FunctionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type FunctionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type FunctionNewParamsFunctionType string

const (
	FunctionNewParamsFunctionTypeDefault   FunctionNewParamsFunctionType = "DEFAULT"
	FunctionNewParamsFunctionTypeStreaming FunctionNewParamsFunctionType = "STREAMING"
)

func (r FunctionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case FunctionNewParamsFunctionTypeDefault, FunctionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type FunctionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type FunctionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r FunctionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type FunctionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r FunctionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response body for create function request.
type CreateFunctionResponse struct {
	// Data Transfer Object (DTO) representing a function
	Function shared.FunctionDTO         `json:"function,required"`
	JSON     createFunctionResponseJSON `json:"-"`
}

// createFunctionResponseJSON contains the JSON metadata for the struct
// [shared.CreateFunctionResponse]
type createFunctionResponseJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *shared.CreateFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionGetAllParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]FunctionGetAllParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [FunctionGetAllParams]'s query parameters as `url.Values`.
func (r FunctionGetAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionGetAllParamsVisibility string

const (
	FunctionGetAllParamsVisibilityAuthorized FunctionGetAllParamsVisibility = "authorized"
	FunctionGetAllParamsVisibilityPrivate    FunctionGetAllParamsVisibility = "private"
	FunctionGetAllParamsVisibilityPublic     FunctionGetAllParamsVisibility = "public"
)

func (r FunctionGetAllParamsVisibility) IsKnown() bool {
	switch r {
	case FunctionGetAllParamsVisibilityAuthorized, FunctionGetAllParamsVisibilityPrivate, FunctionGetAllParamsVisibilityPublic:
		return true
	}
	return false
}
