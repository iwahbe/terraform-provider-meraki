terraform {
  required_providers {
    meraki = {
      version = "0.2.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

variable "my_network_id" {
  type    = string
  default = "L_828099381482775375" # site 3
}
resource "meraki_networks_wireless_settings" "example" {

  ipv6_bridge_enabled        = false
  led_lights_on              = false
  location_analytics_enabled = false
  meshing_enabled            = true
  network_id                 = var.my_network_id
  upgrade_strategy           = "minimizeUpgradeTime"
}

