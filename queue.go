// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/NVIDIADemo/nvcf-go/option"
)

// QueueService contains methods and other services that help with interacting with
// the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options   []option.RequestOption
	Functions *QueueFunctionService
	Position  *QueuePositionService
}

// NewQueueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueueService(opts ...option.RequestOption) (r *QueueService) {
	r = &QueueService{}
	r.Options = opts
	r.Functions = NewQueueFunctionService(opts...)
	r.Position = NewQueuePositionService(opts...)
	return
}
