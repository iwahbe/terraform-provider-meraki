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
	_ datasource.DataSource              = &NetworksApplianceFirewallSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallSettingsDataSource{}
)

func NewNetworksApplianceFirewallSettingsDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallSettingsDataSource{}
}

type NetworksApplianceFirewallSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_settings"
}

func (d *NetworksApplianceFirewallSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"spoofing_protection": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"ip_source_guard": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"mode": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceFirewallSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallSettings NetworksApplianceFirewallSettings
	diags := req.Config.Get(ctx, &networksApplianceFirewallSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallSettings")
		vvNetworkID := networksApplianceFirewallSettings.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallSettings",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallSettings = ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBody(networksApplianceFirewallSettings, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallSettings struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallSettings `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettings struct {
	SpoofingProtection *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection `tfsdk:"spoofing_protection"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection struct {
	IPSourceGuard *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuard `tfsdk:"ip_source_guard"`
}

type ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuard struct {
	Mode types.String `tfsdk:"mode"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallSettingsItemToBody(state NetworksApplianceFirewallSettings, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallSettings) NetworksApplianceFirewallSettings {
	itemState := ResponseApplianceGetNetworkApplianceFirewallSettings{
		SpoofingProtection: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection {
			if response.SpoofingProtection != nil {
				return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection{
					IPSourceGuard: func() *ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
						if response.SpoofingProtection.IPSourceGuard != nil {
							return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuard{
								Mode: types.StringValue(response.SpoofingProtection.IPSourceGuard.Mode),
							}
						}
						return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtectionIpSourceGuard{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceFirewallSettingsSpoofingProtection{}
		}(),
	}
	state.Item = &itemState
	return state
}
