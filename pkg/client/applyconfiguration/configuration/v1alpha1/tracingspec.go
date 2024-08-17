// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// TracingSpecApplyConfiguration represents an declarative configuration of the TracingSpec type for use
// with apply.
type TracingSpecApplyConfiguration struct {
	Enabled          *bool   `json:"enabled,omitempty"`
	ExporterType     *string `json:"exporterType,omitempty"`
	ExporterAddress  *string `json:"exporterAddress,omitempty"`
	IncludeEvent     *bool   `json:"includeEvent,omitempty"`
	IncludeEventBody *bool   `json:"includeEventBody,omitempty"`
}

// TracingSpecApplyConfiguration constructs an declarative configuration of the TracingSpec type for use with
// apply.
func TracingSpec() *TracingSpecApplyConfiguration {
	return &TracingSpecApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *TracingSpecApplyConfiguration) WithEnabled(value bool) *TracingSpecApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithExporterType sets the ExporterType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExporterType field is set to the value of the last call.
func (b *TracingSpecApplyConfiguration) WithExporterType(value string) *TracingSpecApplyConfiguration {
	b.ExporterType = &value
	return b
}

// WithExporterAddress sets the ExporterAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExporterAddress field is set to the value of the last call.
func (b *TracingSpecApplyConfiguration) WithExporterAddress(value string) *TracingSpecApplyConfiguration {
	b.ExporterAddress = &value
	return b
}

// WithIncludeEvent sets the IncludeEvent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IncludeEvent field is set to the value of the last call.
func (b *TracingSpecApplyConfiguration) WithIncludeEvent(value bool) *TracingSpecApplyConfiguration {
	b.IncludeEvent = &value
	return b
}

// WithIncludeEventBody sets the IncludeEventBody field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IncludeEventBody field is set to the value of the last call.
func (b *TracingSpecApplyConfiguration) WithIncludeEventBody(value bool) *TracingSpecApplyConfiguration {
	b.IncludeEventBody = &value
	return b
}
