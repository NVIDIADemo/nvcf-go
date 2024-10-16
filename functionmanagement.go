// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/NVIDIADemo/nvcf-go/option"
)

// FunctionManagementService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementService] method instead.
type FunctionManagementService struct {
	Options   []option.RequestOption
	Functions *FunctionManagementFunctionService
}

// NewFunctionManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFunctionManagementService(opts ...option.RequestOption) (r *FunctionManagementService) {
	r = &FunctionManagementService{}
	r.Options = opts
	r.Functions = NewFunctionManagementFunctionService(opts...)
	return
}
