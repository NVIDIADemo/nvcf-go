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

// AuthorizationFunctionVersionService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizationFunctionVersionService] method instead.
type AuthorizationFunctionVersionService struct {
	Options []option.RequestOption
}

// NewAuthorizationFunctionVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAuthorizationFunctionVersionService(opts ...option.RequestOption) (r *AuthorizationFunctionVersionService) {
	r = &AuthorizationFunctionVersionService{}
	r.Options = opts
	return
}

// Gets NVIDIA Cloud Account IDs that are authorized to invoke specified function
// version. Response includes authorized accounts that were added specifically to
// the function version and the inherited authorized accounts that were added at
// the function level. Access to this functionality mandates the inclusion of a
// bearer token with the 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Get(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all the authorized accounts that are directly associated with the
// specified function version. Authorized parties that are inherited by the
// function version are not deleted. If the specified function version is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Delete(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds the specified NVIDIA Cloud Account to the set of authorized accounts that
// can invoke the specified function version. If the specified function version
// does not have any existing inheritable authorized accounts, it results in a
// response with status 404. If the specified account is already in the set of
// existing authorized accounts that are directly associated with the function
// version, it results in a response wit status code 409. If a function is public,
// then Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Add(ctx context.Context, functionID string, functionVersionID string, body AuthorizationFunctionVersionAddParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s/add", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Authorizes additional NVIDIA Cloud Accounts to invoke a specific function
// version. By default, a function belongs to the NVIDIA Cloud Account that created
// it, and the credentials used for function invocation must reference the same
// NVIDIA Cloud Account. Upon invocation of this endpoint, any existing authorized
// accounts will be overwritten by the newly specified authorized accounts. Access
// to this functionality mandates the inclusion of a bearer token with the
// 'authorize_clients' scope in the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Authorize(ctx context.Context, functionID string, functionVersionID string, body AuthorizationFunctionVersionAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes the specified NVIDIA Cloud Account from the set of authorized accounts
// that are directly associated with specified function version. If the specified
// function version does not have any of its own(not inherited) authorized
// accounts, it results in a response with status 404. Also, if the specified
// authorized account is not in the set of existing authorized parties that are
// directly associated with the specified function version, it results in a
// response with status code 400. If the specified function version is public, then
// Account Admin cannot perform this operation. Access to this functionality
// mandates the inclusion of a bearer token with the 'authorize_clients' scope in
// the HTTP Authorization header
func (r *AuthorizationFunctionVersionService) Remove(ctx context.Context, functionID string, functionVersionID string, body AuthorizationFunctionVersionRemoveParams, opts ...option.RequestOption) (res *shared.AuthorizedParties, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/authorizations/functions/%s/versions/%s/remove", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AuthorizationFunctionVersionAddParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[shared.AuthorizedPartyDTOParam] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionVersionAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionVersionAuthorizeParams struct {
	// Parties authorized to invoke function
	AuthorizedParties param.Field[[]shared.AuthorizedPartyDTOParam] `json:"authorizedParties,required"`
}

func (r AuthorizationFunctionVersionAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthorizationFunctionVersionRemoveParams struct {
	// Data Transfer Object(DTO) representing an authorized party.
	AuthorizedParty param.Field[shared.AuthorizedPartyDTOParam] `json:"authorizedParty,required"`
}

func (r AuthorizationFunctionVersionRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
