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
	_ datasource.DataSource              = &NetworksCellularGatewaySubnetPoolDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCellularGatewaySubnetPoolDataSource{}
)

func NewNetworksCellularGatewaySubnetPoolDataSource() datasource.DataSource {
	return &NetworksCellularGatewaySubnetPoolDataSource{}
}

type NetworksCellularGatewaySubnetPoolDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCellularGatewaySubnetPoolDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_subnet_pool"
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"cidr": schema.StringAttribute{
						Computed: true,
					},
					"deployment_mode": schema.StringAttribute{
						Computed: true,
					},
					"mask": schema.Int64Attribute{
						Computed: true,
					},
					"subnets": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"appliance_ip": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"serial": schema.StringAttribute{
									Computed: true,
								},
								"subnet": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCellularGatewaySubnetPool NetworksCellularGatewaySubnetPool
	diags := req.Config.Get(ctx, &networksCellularGatewaySubnetPool)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCellularGatewaySubnetPool")
		vvNetworkID := networksCellularGatewaySubnetPool.NetworkID.ValueString()

		response1, restyResp1, err := d.client.CellularGateway.GetNetworkCellularGatewaySubnetPool(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}

		networksCellularGatewaySubnetPool = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBody(networksCellularGatewaySubnetPool, response1)
		diags = resp.State.Set(ctx, &networksCellularGatewaySubnetPool)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCellularGatewaySubnetPool struct {
	NetworkID types.String                                                `tfsdk:"network_id"`
	Item      *ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool `tfsdk:"item"`
}

type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool struct {
	Cidr           types.String                                                         `tfsdk:"cidr"`
	DeploymentMode types.String                                                         `tfsdk:"deployment_mode"`
	Mask           types.Int64                                                          `tfsdk:"mask"`
	Subnets        *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets `tfsdk:"subnets"`
}

type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets struct {
	ApplianceIP types.String `tfsdk:"appliance_ip"`
	Name        types.String `tfsdk:"name"`
	Serial      types.String `tfsdk:"serial"`
	Subnet      types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBody(state NetworksCellularGatewaySubnetPool, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool) NetworksCellularGatewaySubnetPool {
	itemState := ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool{
		Cidr:           types.StringValue(response.Cidr),
		DeploymentMode: types.StringValue(response.DeploymentMode),
		Mask: func() types.Int64 {
			if response.Mask != nil {
				return types.Int64Value(int64(*response.Mask))
			}
			return types.Int64{}
		}(),
		Subnets: func() *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets {
			if response.Subnets != nil {
				result := make([]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets{
						ApplianceIP: types.StringValue(subnets.ApplianceIP),
						Name:        types.StringValue(subnets.Name),
						Serial:      types.StringValue(subnets.Serial),
						Subnet:      types.StringValue(subnets.Subnet),
					}
				}
				return &result
			}
			return &[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets{}
		}(),
	}
	state.Item = &itemState
	return state
}
