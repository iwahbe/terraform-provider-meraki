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
	_ datasource.DataSource              = &NetworksSwitchAccessControlListsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchAccessControlListsDataSource{}
)

func NewNetworksSwitchAccessControlListsDataSource() datasource.DataSource {
	return &NetworksSwitchAccessControlListsDataSource{}
}

type NetworksSwitchAccessControlListsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchAccessControlListsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchAccessControlListsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_access_control_lists"
}

func (d *NetworksSwitchAccessControlListsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the access control list rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Computed:            true,
								},
								"dst_cidr": schema.StringAttribute{
									MarkdownDescription: `Destination IP address (in IP or CIDR notation)`,
									Computed:            true,
								},
								"dst_port": schema.StringAttribute{
									MarkdownDescription: `Destination port`,
									Computed:            true,
								},
								"ip_version": schema.StringAttribute{
									MarkdownDescription: `IP address version`,
									Computed:            true,
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol`,
									Computed:            true,
								},
								"src_cidr": schema.StringAttribute{
									MarkdownDescription: `Source IP address (in IP or CIDR notation)`,
									Computed:            true,
								},
								"src_port": schema.StringAttribute{
									MarkdownDescription: `Source port`,
									Computed:            true,
								},
								"vlan": schema.StringAttribute{
									MarkdownDescription: `ncoming traffic VLAN`,
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

func (d *NetworksSwitchAccessControlListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchAccessControlLists NetworksSwitchAccessControlLists
	diags := req.Config.Get(ctx, &networksSwitchAccessControlLists)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchAccessControlLists")
		vvNetworkID := networksSwitchAccessControlLists.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchAccessControlLists(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessControlLists",
				err.Error(),
			)
			return
		}

		networksSwitchAccessControlLists = ResponseSwitchGetNetworkSwitchAccessControlListsItemToBody(networksSwitchAccessControlLists, response1)
		diags = resp.State.Set(ctx, &networksSwitchAccessControlLists)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchAccessControlLists struct {
	NetworkID types.String                                      `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchAccessControlLists `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchAccessControlLists struct {
	Rules *[]ResponseSwitchGetNetworkSwitchAccessControlListsRules `tfsdk:"rules"`
}

type ResponseSwitchGetNetworkSwitchAccessControlListsRules struct {
	Comment   types.String `tfsdk:"comment"`
	DstCidr   types.String `tfsdk:"dst_cidr"`
	DstPort   types.String `tfsdk:"dst_port"`
	IPVersion types.String `tfsdk:"ip_version"`
	Policy    types.String `tfsdk:"policy"`
	Protocol  types.String `tfsdk:"protocol"`
	SrcCidr   types.String `tfsdk:"src_cidr"`
	SrcPort   types.String `tfsdk:"src_port"`
	VLAN      types.String `tfsdk:"vlan"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchAccessControlListsItemToBody(state NetworksSwitchAccessControlLists, response *merakigosdk.ResponseSwitchGetNetworkSwitchAccessControlLists) NetworksSwitchAccessControlLists {
	itemState := ResponseSwitchGetNetworkSwitchAccessControlLists{
		Rules: func() *[]ResponseSwitchGetNetworkSwitchAccessControlListsRules {
			if response.Rules != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessControlListsRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseSwitchGetNetworkSwitchAccessControlListsRules{
						Comment:   types.StringValue(rules.Comment),
						DstCidr:   types.StringValue(rules.DstCidr),
						DstPort:   types.StringValue(rules.DstPort),
						IPVersion: types.StringValue(rules.IPVersion),
						Policy:    types.StringValue(rules.Policy),
						Protocol:  types.StringValue(rules.Protocol),
						SrcCidr:   types.StringValue(rules.SrcCidr),
						SrcPort:   types.StringValue(rules.SrcPort),
						VLAN:      types.StringValue(rules.VLAN),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchAccessControlListsRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
