// resource_server.go
package main

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &TextFileResource{}
)

// TextFileResource is the resource implementation.
type TextFileResource struct{}

// TextFileResourceModel maps the resource schema data.
type TextFileResourceModel struct {
	Path    types.String `tfsdk:"path"`
	Content types.String `tfsdk:"content"`
}

// NewTextFileResource is a helper function to simplify the provider server wiring.
func NewTextFileResource() resource.Resource {
	return &TextFileResource{}
}

// Metadata returns the resource type name.
func (r *TextFileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

// Schema defines the schema for the resource.
func (r *TextFileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a local text file.",
		Attributes: map[string]schema.Attribute{
			"path": schema.StringAttribute{
				Description: "The absolute path to the file.",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content to write in the file.",
				Required:    true,
			},
		},
	}
}

// Create creates the resource and sets the initial state.
func (r *TextFileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// 1. Read configuration data into the model
	var plan TextFileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Get values from the plan
	filePath := plan.Path.ValueString()
	fileContent := plan.Content.ValueString()

	// 3. Create the file on the local filesystem
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating File",
			"Could not create file at "+filePath+": "+err.Error(),
		)
		return
	}

	// 4. Set the final state in OpenTofu
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

}

// Read refreshes the resource state.
func (r *TextFileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// 1. Get the current state
	var state TextFileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Read the file from the filesystem
	filePath := state.Path.ValueString()
	content, err := os.ReadFile(filePath)

	// 3. Handle errors, especially if the file is gone
	if err != nil {
		// If the file no longer exists, we'll treat it as deleted.
		if os.IsNotExist(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading File",
			"Could not read file at "+filePath+": "+err.Error(),
		)
		return
	}

	// 4. Update the model with the actual content from the file
	state.Content = types.StringValue(string(content))

	// 5. Set the refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource.
func (r *TextFileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update resource logic
}

// Delete deletes the resource.
func (r *TextFileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// 1. Get the state
	var state TextFileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2. Get the file path from the state
	filePath := state.Path.ValueString()

	// 3. Delete the file
	err := os.Remove(filePath)
	if err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError(
			"Error Deleting File",
			"Could not delete file at "+filePath+": "+err.Error(),
		)
		return
	}
}
