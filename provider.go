// provider.go
package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &OperatingSystemProvider{}
)

// OperatingSystemProvider is the provider implementation.
type OperatingSystemProvider struct{}

// New is a helper function to simplify provider server wiring.
func New() provider.Provider {
	return &OperatingSystemProvider{}
}

// Metadata returns the provider type name.
func (p *OperatingSystemProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "os"
}

// Schema defines the provider-level schema for configuration data.
func (p *OperatingSystemProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure prepares a provider for data sources and resources.
func (p *OperatingSystemProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *OperatingSystemProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *OperatingSystemProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTextFileResource,
	}
}
