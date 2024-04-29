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
	_ datasource.DataSource              = &NetworksFirmwareUpgradesStagedStagesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFirmwareUpgradesStagedStagesDataSource{}
)

func NewNetworksFirmwareUpgradesStagedStagesDataSource() datasource.DataSource {
	return &NetworksFirmwareUpgradesStagedStagesDataSource{}
}

type NetworksFirmwareUpgradesStagedStagesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFirmwareUpgradesStagedStagesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFirmwareUpgradesStagedStagesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_stages"
}

func (d *NetworksFirmwareUpgradesStagedStagesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkFirmwareUpgradesStagedStages`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"group": schema.SingleNestedAttribute{
							MarkdownDescription: `The Staged Upgrade Group`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `Description of the Staged Upgrade Group`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Id of the Staged Upgrade Group`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the Staged Upgrade Group`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksFirmwareUpgradesStagedStagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFirmwareUpgradesStagedStages NetworksFirmwareUpgradesStagedStages
	diags := req.Config.Get(ctx, &networksFirmwareUpgradesStagedStages)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFirmwareUpgradesStagedStages")
		vvNetworkID := networksFirmwareUpgradesStagedStages.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkFirmwareUpgradesStagedStages(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedStages",
				err.Error(),
			)
			return
		}

		networksFirmwareUpgradesStagedStages = ResponseNetworksGetNetworkFirmwareUpgradesStagedStagesItemsToBody(networksFirmwareUpgradesStagedStages, response1)
		diags = resp.State.Set(ctx, &networksFirmwareUpgradesStagedStages)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFirmwareUpgradesStagedStages struct {
	NetworkID types.String                                                  `tfsdk:"network_id"`
	Items     *[]ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages struct {
	Group *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup `tfsdk:"group"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkFirmwareUpgradesStagedStagesItemsToBody(state NetworksFirmwareUpgradesStagedStages, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedStages) NetworksFirmwareUpgradesStagedStages {
	var items []ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStages{
			Group: func() *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup {
				if item.Group != nil {
					return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup{
						Description: types.StringValue(item.Group.Description),
						ID:          types.StringValue(item.Group.ID),
						Name:        types.StringValue(item.Group.Name),
					}
				}
				return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroup{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
