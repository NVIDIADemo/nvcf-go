// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/option"
)

// EnvelopeFunctionInvocationExecService contains methods and other services that
// help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvelopeFunctionInvocationExecService] method instead.
type EnvelopeFunctionInvocationExecService struct {
	Options []option.RequestOption
	Status  *EnvelopeFunctionInvocationExecStatusService
}

// NewEnvelopeFunctionInvocationExecService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnvelopeFunctionInvocationExecService(opts ...option.RequestOption) (r *EnvelopeFunctionInvocationExecService) {
	r = &EnvelopeFunctionInvocationExecService{}
	r.Options = opts
	r.Status = NewEnvelopeFunctionInvocationExecStatusService(opts...)
	return
}
