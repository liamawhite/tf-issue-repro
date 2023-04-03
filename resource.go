package repro

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

type ServiceAccountModel struct {
	Id   types.String `tfsdk:"id"`
	Keys []*KeyModel  `tfsdk:"keys"`
}

type KeyModel struct {
	Token types.String `tfsdk:"token"`
}

func NewResource() resource.Resource {
	return &ServiceAccountResource{}
}

type ServiceAccountResource struct{}

func (r *ServiceAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_account"
}

func (r *ServiceAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *ServiceAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (*ServiceAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"keys": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"token": schema.StringAttribute{
					Computed: true,
				},
			}},
		},
	}}
}

func (r *ServiceAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	model.Id = types.StringValue("some-id")
	model.Keys = []*KeyModel{{Token: types.StringValue("some-token")}}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *ServiceAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ServiceAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	model.Id = types.StringValue("some-id")
	model.Keys = []*KeyModel{{Token: types.StringValue("some-token")}}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *ServiceAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	model.Id = types.StringValue("some-id")
	model.Keys = []*KeyModel{{Token: types.StringValue("some-token")}}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
