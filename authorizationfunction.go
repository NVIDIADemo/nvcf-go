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

// AuthorizationFunctionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizationFunctionService] method instead.
type AuthorizationFunctionService struct {
	Options  []option.RequestOption
	Versions *AuthorizationFunctionVersionService
}

// NewAuthorizationFunctionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthorizationFunctionService(opts ...option.RequestOption) (r *AuthorizationFunctionService) {
	r = &AuthorizationFunctionService{}
	r.Options = opts
	r.Versions = NewAuthorizationFunctionVersionService(opts...)
	return
}

// Deletes all the extra NVIDIA Cloud Accounts that were authorized to invoke the
// function and all its versions. If a function version has its own set of
// authorized accounts, those are not deleted. If the specified function is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds the specified NVIDIA Cloud Account to the set of authorized accounts that
// are can invoke all the versions of the specified function. If the specified
// function does not have any existing inheritable authorized accounts, it results
// in a response with status 404. If the specified account is already in the set of
// existing inheritable authorized accounts, it results in a response with status
// code 409. If a function is public, then Account Admin cannot perform this
// operation. Access to this functionality mandates the inclusion of a bearer token
// with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) Add(ctx context.Context, functionID string, body AuthorizationFunctionAddParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/add", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Authorizes additional NVIDIA Cloud Accounts to invoke any version of the
// specified function. By default, a function belongs to the NVIDIA Cloud Account
// that created it, and the credentials used for function invocation must reference
// the same NVIDIA Cloud Account. Upon invocation of this endpoint, any existing
// authorized accounts will be overwritten by the newly specified authorized
// accounts. Access to this functionality mandates the inclusion of a bearer token
// with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) Authorize(ctx context.Context, functionID string, body AuthorizationFunctionAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes the specified NVIDIA Cloud Account from the set of authorized accounts
// that can invoke all the versions of the specified function. If the specified
// function does not have any existing inheritable authorized parties, it results
// in a response with status 404. Also, if the specified account is not in the
// existing set of inheritable authorized accounts, it results in a response with
// status 400. If the specified function is public, then Account Admin cannot
// perform this operation. Access to this functionality mandates the inclusion of a
// bearer token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) Remove(ctx context.Context, functionID string, body AuthorizationFunctionRemoveParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/remove", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists NVIDIA Cloud Account IDs that are authorized to invoke any version of the
// specified function. The response includes an array showing authorized accounts
// for each version. Individual versions of a function can have their own
// authorized accounts. So, each object in the array can have different authorized
// accounts listed. Access to this functionality mandates the inclusion of a bearer
// token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionService) GetAll(ctx context.Context, functionID string, opts ...option.RequestOption) (res *ListAuthorizedPartiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Parties authorized to invoke function
type ListAuthorizedPartiesResponse struct {
	// Functions with authorized parties and other details
	Functions []ListAuthorizedPartiesResponseFunction `json:"functions,required"`
	JSON      listAuthorizedPartiesResponseJSON       `json:"-"`
}

// listAuthorizedPartiesResponseJSON contains the JSON metadata for the struct
// [ListAuthorizedPartiesResponse]
type listAuthorizedPartiesResponseJSON struct {
	Functions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListAuthorizedPartiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAuthorizedPartiesResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function with authorized accounts
type ListAuthorizedPartiesResponseFunction struct {
	// Function id
	ID string `json:"id,required" format:"uuid"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Authorized parties allowed to invoke the function
	AuthorizedParties []shared.AuthorizedPartyDTO `json:"authorizedParties"`
	// Function version id
	VersionID string                                    `json:"versionId" format:"uuid"`
	JSON      listAuthorizedPartiesResponseFunctionJSON `json:"-"`
}

// listAuthorizedPartiesResponseFunctionJSON contains the JSON metadata for the
// struct [ListAuthorizedPartiesResponseFunction]
type listAuthorizedPartiesResponseFunctionJSON struct {
	ID                apijson.Field
	NcaID             apijson.Field
	AuthorizedParties apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ListAuthorizedPartiesResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listAuthorizedPartiesResponseFunctionJSON) RawJSON() string {
	return r.raw
}

type AuthorizationFunctionAddParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[shared.AuthorizedPartyDTOParam] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionAuthorizeParams struct {
	// Parties authorized to invoke function
	AuthorizedParties param.Field[[]shared.AuthorizedPartyDTOParam] `json:"authorizedParties,required"`
}

func (r AuthorizationFunctionAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionRemoveParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[shared.AuthorizedPartyDTOParam] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
