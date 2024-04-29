package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsDevicesProvisioningStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesProvisioningStatusesDataSource{}
)

func NewOrganizationsDevicesProvisioningStatusesDataSource() datasource.DataSource {
	return &OrganizationsDevicesProvisioningStatusesDataSource{}
}

type OrganizationsDevicesProvisioningStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesProvisioningStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesProvisioningStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_provisioning_statuses"
}

func (d *OrganizationsDevicesProvisioningStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter device by network ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device by device product types. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `status query parameter. An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `The device MAC address.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The device name.`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network info.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID for the network containing the device.`,
									Computed:            true,
								},
							},
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Device product type.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The device serial number.`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `The device provisioning status. Possible statuses: unprovisioned, incomplete, complete.`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `List of custom tags for the device.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesProvisioningStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesProvisioningStatuses OrganizationsDevicesProvisioningStatuses
	diags := req.Config.Get(ctx, &organizationsDevicesProvisioningStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesProvisioningStatuses")
		vvOrganizationID := organizationsDevicesProvisioningStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesProvisioningStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesProvisioningStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesProvisioningStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesProvisioningStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesProvisioningStatuses.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesProvisioningStatuses.ProductTypes)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesProvisioningStatuses.Serials)
		queryParams1.Status = organizationsDevicesProvisioningStatuses.Status.ValueString()
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevicesProvisioningStatuses.Tags)
		queryParams1.TagsFilterType = organizationsDevicesProvisioningStatuses.TagsFilterType.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesProvisioningStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesProvisioningStatuses",
				err.Error(),
			)
			return
		}

		organizationsDevicesProvisioningStatuses = ResponseOrganizationsGetOrganizationDevicesProvisioningStatusesItemsToBody(organizationsDevicesProvisioningStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesProvisioningStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesProvisioningStatuses struct {
	OrganizationID types.String                                                           `tfsdk:"organization_id"`
	PerPage        types.Int64                                                            `tfsdk:"per_page"`
	StartingAfter  types.String                                                           `tfsdk:"starting_after"`
	EndingBefore   types.String                                                           `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                             `tfsdk:"network_ids"`
	ProductTypes   types.List                                                             `tfsdk:"product_types"`
	Serials        types.List                                                             `tfsdk:"serials"`
	Status         types.String                                                           `tfsdk:"status"`
	Tags           types.List                                                             `tfsdk:"tags"`
	TagsFilterType types.String                                                           `tfsdk:"tags_filter_type"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses struct {
	Mac         types.String                                                                `tfsdk:"mac"`
	Name        types.String                                                                `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork `tfsdk:"network"`
	ProductType types.String                                                                `tfsdk:"product_type"`
	Serial      types.String                                                                `tfsdk:"serial"`
	Status      types.String                                                                `tfsdk:"status"`
	Tags        types.List                                                                  `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesProvisioningStatusesItemsToBody(state OrganizationsDevicesProvisioningStatuses, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesProvisioningStatuses) OrganizationsDevicesProvisioningStatuses {
	var items []ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatuses{
			Mac:  types.StringValue(item.Mac),
			Name: types.StringValue(item.Name),
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork{
						ID: types.StringValue(item.Network.ID),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationDevicesProvisioningStatusesNetwork{}
			}(),
			ProductType: types.StringValue(item.ProductType),
			Serial:      types.StringValue(item.Serial),
			Status:      types.StringValue(item.Status),
			Tags:        StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
