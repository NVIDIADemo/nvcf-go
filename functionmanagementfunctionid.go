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
)

// FunctionManagementFunctionIDService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementFunctionIDService] method instead.
type FunctionManagementFunctionIDService struct {
	Options []option.RequestOption
}

// NewFunctionManagementFunctionIDService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionManagementFunctionIDService(opts ...option.RequestOption) (r *FunctionManagementFunctionIDService) {
	r = &FunctionManagementFunctionIDService{}
	r.Options = opts
	return
}

// Lists ids of all the functions in the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *FunctionManagementFunctionIDService) GetAll(ctx context.Context, query FunctionManagementFunctionIDGetAllParams, opts ...option.RequestOption) (res *FunctionManagementFunctionIDGetAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response body containing list of function ids in an account
type FunctionManagementFunctionIDGetAllResponse struct {
	// List of function ids
	FunctionIDs []string                                       `json:"functionIds,required" format:"uuid"`
	JSON        functionManagementFunctionIDGetAllResponseJSON `json:"-"`
}

// functionManagementFunctionIDGetAllResponseJSON contains the JSON metadata for
// the struct [FunctionManagementFunctionIDGetAllResponse]
type functionManagementFunctionIDGetAllResponseJSON struct {
	FunctionIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionManagementFunctionIDGetAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionManagementFunctionIDGetAllResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionManagementFunctionIDGetAllParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]FunctionManagementFunctionIDGetAllParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [FunctionManagementFunctionIDGetAllParams]'s query
// parameters as `url.Values`.
func (r FunctionManagementFunctionIDGetAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionManagementFunctionIDGetAllParamsVisibility string

const (
	FunctionManagementFunctionIDGetAllParamsVisibilityAuthorized FunctionManagementFunctionIDGetAllParamsVisibility = "authorized"
	FunctionManagementFunctionIDGetAllParamsVisibilityPrivate    FunctionManagementFunctionIDGetAllParamsVisibility = "private"
	FunctionManagementFunctionIDGetAllParamsVisibilityPublic     FunctionManagementFunctionIDGetAllParamsVisibility = "public"
)

func (r FunctionManagementFunctionIDGetAllParamsVisibility) IsKnown() bool {
	switch r {
	case FunctionManagementFunctionIDGetAllParamsVisibilityAuthorized, FunctionManagementFunctionIDGetAllParamsVisibilityPrivate, FunctionManagementFunctionIDGetAllParamsVisibilityPublic:
		return true
	}
	return false
}
