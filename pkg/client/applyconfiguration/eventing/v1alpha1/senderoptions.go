// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// SenderOptionsApplyConfiguration represents an declarative configuration of the SenderOptions type for use
// with apply.
type SenderOptionsApplyConfiguration struct {
	BatchOptions   *BatchOptionsApplyConfiguration `json:"batchOptions,omitempty"`
	DedupOptions   *DedupOptionsApplyConfiguration `json:"dedupOptions,omitempty"`
	NumWorkers     *int                            `json:"numWorkers,omitempty"`
	SendBufferSize *int                            `json:"sendBufferSize,omitempty"`
	QueueName      *string                         `json:"queueName,omitempty"`
}

// SenderOptionsApplyConfiguration constructs an declarative configuration of the SenderOptions type for use with
// apply.
func SenderOptions() *SenderOptionsApplyConfiguration {
	return &SenderOptionsApplyConfiguration{}
}

// WithBatchOptions sets the BatchOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BatchOptions field is set to the value of the last call.
func (b *SenderOptionsApplyConfiguration) WithBatchOptions(value *BatchOptionsApplyConfiguration) *SenderOptionsApplyConfiguration {
	b.BatchOptions = value
	return b
}

// WithDedupOptions sets the DedupOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DedupOptions field is set to the value of the last call.
func (b *SenderOptionsApplyConfiguration) WithDedupOptions(value *DedupOptionsApplyConfiguration) *SenderOptionsApplyConfiguration {
	b.DedupOptions = value
	return b
}

// WithNumWorkers sets the NumWorkers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NumWorkers field is set to the value of the last call.
func (b *SenderOptionsApplyConfiguration) WithNumWorkers(value int) *SenderOptionsApplyConfiguration {
	b.NumWorkers = &value
	return b
}

// WithSendBufferSize sets the SendBufferSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendBufferSize field is set to the value of the last call.
func (b *SenderOptionsApplyConfiguration) WithSendBufferSize(value int) *SenderOptionsApplyConfiguration {
	b.SendBufferSize = &value
	return b
}

// WithQueueName sets the QueueName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QueueName field is set to the value of the last call.
func (b *SenderOptionsApplyConfiguration) WithQueueName(value string) *SenderOptionsApplyConfiguration {
	b.QueueName = &value
	return b
}
