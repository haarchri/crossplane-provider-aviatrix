/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GatewayObservation struct {

	// List of available BGP LAN interface IPs for spoke external device connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.
	// List of available BGP LAN interface IPs for spoke external device connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.
	BGPLanIPList []*string `json:"bgpLanIpList,omitempty" tf:"bgp_lan_ip_list,omitempty"`

	// Cloud instance ID of the spoke gateway.
	// Cloud instance ID.
	CloudInstanceID *string `json:"cloudInstanceId,omitempty" tf:"cloud_instance_id,omitempty"`

	// List of available BGP LAN interface IPs for spoke external device HA connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.
	// List of available BGP LAN interface IPs for spoke external device HA connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.
	HaBGPLanIPList []*string `json:"haBgpLanIpList,omitempty" tf:"ha_bgp_lan_ip_list,omitempty"`

	// Cloud instance ID of the HA spoke gateway.
	// Cloud instance ID of HA spoke gateway.
	HaCloudInstanceID *string `json:"haCloudInstanceId,omitempty" tf:"ha_cloud_instance_id,omitempty"`

	// Aviatrix spoke gateway unique name of HA spoke gateway.
	// Aviatrix spoke gateway unique name of HA spoke gateway.
	HaGwName *string `json:"haGwName,omitempty" tf:"ha_gw_name,omitempty"`

	// Private IP address of HA spoke gateway.
	// Private IP address of the spoke gateway created.
	HaPrivateIP *string `json:"haPrivateIp,omitempty" tf:"ha_private_ip,omitempty"`

	// Public IP address of the HA Spoke Gateway.
	// Public IP address of the HA Spoke Gateway.
	HaPublicIP *string `json:"haPublicIp,omitempty" tf:"ha_public_ip,omitempty"`

	// HA security group used for the spoke gateway.
	// HA security group used for the spoke gateway.
	HaSecurityGroupID *string `json:"haSecurityGroupId,omitempty" tf:"ha_security_group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image version of the gateway. Use aviatrix_gateway_image data source to programmatically retrieve this value for the desired software_version. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the aviatrix_controller_config resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.
	// image_version can be used to set the desired image version of the gateway. If set, we will attempt to update the gateway to the specified version.
	ImageVersion *string `json:"imageVersion,omitempty" tf:"image_version,omitempty"`

	// Private IP address of the spoke gateway created.
	// Private IP address of the spoke gateway created.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Public IP address of the Spoke Gateway created.
	// Public IP address of the Spoke Gateway created.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Security group used for the spoke gateway.
	// Security group used for the spoke gateway.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// The software version of the gateway. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrade can be managed with the aviatrix_controller_config resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.
	// software_version can be used to set the desired software version of the gateway. If set, we will attempt to update the gateway to the specified version. If left blank, the gateway software version will continue to be managed through the aviatrix_controller_config resource.
	SoftwareVersion *string `json:"softwareVersion,omitempty" tf:"software_version,omitempty"`
}

type GatewayParameters struct {

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 4.7+. Valid values: true, false. Default: true.
	// If false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway.
	// +kubebuilder:validation:Optional
	AllocateNewEIP *bool `json:"allocateNewEip,omitempty" tf:"allocate_new_eip,omitempty"`

	// A set of approved learned CIDRs. Only valid when enable_learned_cidrs_approval is set to true. Example: ["10.250.0.0/16", "10.251.0.0/16"]. Available as of provider version R2.21+.
	// Approved learned CIDRs for BGP Spoke Gateway. Available as of provider version R2.21+.
	// +kubebuilder:validation:Optional
	ApprovedLearnedCidrs []*string `json:"approvedLearnedCidrs,omitempty" tf:"approved_learned_cidrs,omitempty"`

	// Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Availability domain for OCI.
	// +kubebuilder:validation:Optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// Name of public IP Address resource and its resource group in Azure to be assigned to the Spoke Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if allocate_new_eip is false and cloud_type is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.
	// The name of the public IP address and its resource group in Azure to assign to this Spoke Gateway.
	// +kubebuilder:validation:Optional
	AzureEIPNameResourceGroup *string `json:"azureEipNameResourceGroup,omitempty" tf:"azure_eip_name_resource_group,omitempty"`

	// Enable Equal Cost Multi Path (ECMP) routing for the next hop. Default value: false.
	// Enable Equal Cost Multi Path (ECMP) routing for the next hop for BGP Spoke Gateway.
	// +kubebuilder:validation:Optional
	BGPEcmp *bool `json:"bgpEcmp,omitempty" tf:"bgp_ecmp,omitempty"`

	// BGP hold time. Unit is in seconds. Valid values are between 12 and 360. Default value: 180.
	// BGP Hold Time for BGP Spoke Gateway. Unit is in seconds. Valid values are between 12 and 360.
	// +kubebuilder:validation:Optional
	BGPHoldTime *float64 `json:"bgpHoldTime,omitempty" tf:"bgp_hold_time,omitempty"`

	// Number of interfaces that will be created for BGP over LAN enabled Azure spoke. Applies on HA Transit as well if enabled. Available as of provider version R3.0.2+.
	// Number of interfaces that will be created for BGP over LAN enabled Azure spoke. Only valid for 8 (Azure), 32 (AzureGov) or AzureChina (2048). Default value: 1. Available as of provider version R3.0.2+.
	// +kubebuilder:validation:Optional
	BGPLanInterfacesCount *float64 `json:"bgpLanInterfacesCount,omitempty" tf:"bgp_lan_interfaces_count,omitempty"`

	// BGP route polling time. Unit is in seconds. Valid values are between 10 and 50. Default value: "50".
	// BGP route polling time for BGP Spoke Gateway. Unit is in seconds. Valid values are between 10 and 50.
	// +kubebuilder:validation:Optional
	BGPPollingTime *float64 `json:"bgpPollingTime,omitempty" tf:"bgp_polling_time,omitempty"`

	// Type of cloud service provider, requires an integer value. Currently, only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192), AWS Top Secret(16384) and AWS Secret (32768) are supported.
	// Type of cloud service provider.
	// +kubebuilder:validation:Required
	CloudType *float64 `json:"cloudType" tf:"cloud_type,omitempty"`

	// Customer managed key ID.
	// Customer managed key ID.
	// +kubebuilder:validation:Optional
	CustomerManagedKeysSecretRef *v1.SecretKeySelector `json:"customerManagedKeysSecretRef,omitempty" tf:"-"`

	// A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only. Example: "10.0.0.0/116,10.2.0.0/16".
	// A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only.
	// +kubebuilder:validation:Optional
	CustomizedSpokeVPCRoutes *string `json:"customizedSpokeVpcRoutes,omitempty" tf:"customized_spoke_vpc_routes,omitempty"`

	// Disables route propagation on BGP Spoke to attached Transit Gateway. Default value: false.
	// Disables route propagation on BGP Spoke to attached Transit Gateway. Default: false.
	// +kubebuilder:validation:Optional
	DisableRoutePropagation *bool `json:"disableRoutePropagation,omitempty" tf:"disable_route_propagation,omitempty"`

	// Required when allocate_new_eip is false. It uses the specified EIP for this gateway. Available in Controller 4.7+. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.
	// Required when allocate_new_eip is false. It uses specified EIP for this gateway.
	// +kubebuilder:validation:Optional
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// Enables Active-Standby Mode. Available only with HA enabled. Valid values: true, false. Default value: false.
	// Enables Active-Standby Mode, available only with HA enabled for BGP Spoke Gateway.
	// +kubebuilder:validation:Optional
	EnableActiveStandby *bool `json:"enableActiveStandby,omitempty" tf:"enable_active_standby,omitempty"`

	// Enables Preemptive Mode for Active-Standby. Available only with BGP enabled, HA enabled and Active-Standby enabled. Valid values: true, false. Default value: false.
	// Enables Preemptive Mode for Active-Standby, available only with Active-Standby enabled.
	// +kubebuilder:validation:Optional
	EnableActiveStandbyPreemptive *bool `json:"enableActiveStandbyPreemptive,omitempty" tf:"enable_active_standby_preemptive,omitempty"`

	// Auto Advertise Spoke Site2Cloud CIDRs. Default: false. Valid values: true or false. Available as of provider version R2.19+.
	// Automatically advertise remote CIDR to Aviatrix Transit Gateway when route based Site2Cloud Tunnel is created.
	// +kubebuilder:validation:Optional
	EnableAutoAdvertiseS2CCidrs *bool `json:"enableAutoAdvertiseS2CCidrs,omitempty" tf:"enable_auto_advertise_s2c_cidrs,omitempty"`

	// Enable BGP for this spoke gateway. Only available for AWS and Azure. Valid values: true, false. Default value: false. Available in provider R2.21.0+.
	// Enable BGP. Default: false.
	// +kubebuilder:validation:Optional
	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// Pre-allocate a network interface(eth4) for "BGP over LAN" functionality. Must be enabled to create a BGP over LAN aviatrix_spoke_external_device_conn resource with this Spoke Gateway. Only valid for 8 (Azure), 32 (AzureGov) or AzureChina (2048). Valid values: true or false. Default value: false. Available as of provider version R3.0.2+.
	// Pre-allocate a network interface(eth4) for "BGP over LAN" functionality. Only valid for 8 (Azure), 32 (AzureGov) or AzureChina (2048). Valid values: true or false. Default value: false. Available as of provider version R3.0.2+.
	// +kubebuilder:validation:Optional
	EnableBGPOverLan *bool `json:"enableBgpOverLan,omitempty" tf:"enable_bgp_over_lan,omitempty"`

	// Enable EBS volume encryption for Gateway. Only supports AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret providers. Valid values: true, false. Default value: false.
	// Enable encrypt gateway EBS volume. Only supported for AWS provider. Valid values: true, false. Default value: false.
	// +kubebuilder:validation:Optional
	EnableEncryptVolume *bool `json:"enableEncryptVolume,omitempty" tf:"enable_encrypt_volume,omitempty"`

	// Set to true to enable global VPC. Only supported for GCP.
	// +kubebuilder:validation:Optional
	EnableGlobalVPC *bool `json:"enableGlobalVpc,omitempty" tf:"enable_global_vpc,omitempty"`

	// Enable GRO/GSO for this transit gateway. Default value is true. Available in provider R3.1.0+.
	// Specify whether to disable GRO/GSO or not.
	// +kubebuilder:validation:Optional
	EnableGroGso *bool `json:"enableGroGso,omitempty" tf:"enable_gro_gso,omitempty"`

	// Enable jumbo frames for this spoke gateway. Default value is true.
	// Enable jumbo frame support for spoke gateway. Valid values: true or false. Default value: true.
	// +kubebuilder:validation:Optional
	EnableJumboFrame *bool `json:"enableJumboFrame,omitempty" tf:"enable_jumbo_frame,omitempty"`

	// Switch to enable/disable learned CIDR approval for BGP Spoke Gateway. Valid values: true, false. Default value: false.
	// Switch to enable/disable learned CIDR approval for BGP Spoke Gateway. Valid values: true, false.
	// +kubebuilder:validation:Optional
	EnableLearnedCidrsApproval *bool `json:"enableLearnedCidrsApproval,omitempty" tf:"enable_learned_cidrs_approval,omitempty"`

	// If set to true, the Monitor Gateway Subnets feature is enabled. Default value is false. Available in provider version R2.18+.
	// Enable [monitor gateway subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet). Only valid for cloud_type = 1 (AWS) or 256 (AWSGov). Valid values: true, false. Default value: false.
	// +kubebuilder:validation:Optional
	EnableMonitorGatewaySubnets *bool `json:"enableMonitorGatewaySubnets,omitempty" tf:"enable_monitor_gateway_subnets,omitempty"`

	// Enable preserve as_path when advertising manual summary cidrs on BGP spoke gateway. Valid values: true, false. Default value: false. Available as of provider version R.2.22.1+
	// Enable preserve as_path when advertising manual summary cidrs on BGP spoke gateway.
	// +kubebuilder:validation:Optional
	EnablePreserveAsPath *bool `json:"enablePreserveAsPath,omitempty" tf:"enable_preserve_as_path,omitempty"`

	// Enable Private OOB feature. Only available for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.
	// Enable private OOB.
	// +kubebuilder:validation:Optional
	EnablePrivateOob *bool `json:"enablePrivateOob,omitempty" tf:"enable_private_oob,omitempty"`

	// Program default route in VPC private route table. Default: false. Valid values: true or false. Available as of provider version R2.19+.
	// Config Private VPC Default Route.
	// +kubebuilder:validation:Optional
	EnablePrivateVPCDefaultRoute *bool `json:"enablePrivateVpcDefaultRoute,omitempty" tf:"enable_private_vpc_default_route,omitempty"`

	// Skip programming VPC public route table. Default: false. Valid values: true or false. Available as of provider version R2.19+.
	// Skip Public Route Table Update.
	// +kubebuilder:validation:Optional
	EnableSkipPublicRouteTableUpdate *bool `json:"enableSkipPublicRouteTableUpdate,omitempty" tf:"enable_skip_public_route_table_update,omitempty"`

	// Enable spot instance. NOT supported for production deployment.
	// Enable spot instance. NOT supported for production deployment.
	// +kubebuilder:validation:Optional
	EnableSpotInstance *bool `json:"enableSpotInstance,omitempty" tf:"enable_spot_instance,omitempty"`

	// Enable VPC DNS Server for Gateway. Currently only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.
	// Enable vpc_dns_server for Gateway. Valid values: true, false.
	// +kubebuilder:validation:Optional
	EnableVPCDNSServer *bool `json:"enableVpcDnsServer,omitempty" tf:"enable_vpc_dns_server,omitempty"`

	// Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Fault domain for OCI.
	// +kubebuilder:validation:Optional
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain,omitempty"`

	// A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only. Example: "10.2.0.0/116,10.3.0.0/16".
	// A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only.
	// +kubebuilder:validation:Optional
	FilteredSpokeVPCRoutes *string `json:"filteredSpokeVpcRoutes,omitempty" tf:"filtered_spoke_vpc_routes,omitempty"`

	// Size of the gateway instance. Example: AWS/AWSGov/AWSChina: "t2.large", Azure/AzureGov/AzureChina: "Standard_B1s", OCI: "VM.Standard2.2", GCP: "n1-standard-1".
	// Size of the gateway instance.
	// +kubebuilder:validation:Required
	GwSize *string `json:"gwSize" tf:"gw_size,omitempty"`

	// HA gateway availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// HA availability domain for OCI.
	// +kubebuilder:validation:Optional
	HaAvailabilityDomain *string `json:"haAvailabilityDomain,omitempty" tf:"ha_availability_domain,omitempty"`

	// Name of public IP Address resource and its resource group in Azure to be assigned to the HA Spoke Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ha_eip is set and cloud_type is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.
	// The name of the public IP address and its resource group in Azure to assign to the HA Spoke Gateway.
	// +kubebuilder:validation:Optional
	HaAzureEIPNameResourceGroup *string `json:"haAzureEipNameResourceGroup,omitempty" tf:"ha_azure_eip_name_resource_group,omitempty"`

	// Public IP address that you want to assign to the HA peering instance. If no value is given, a new EIP will automatically be allocated. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.
	// Public IP address that you want assigned to the HA Spoke Gateway.
	// +kubebuilder:validation:Optional
	HaEIP *string `json:"haEip,omitempty" tf:"ha_eip,omitempty"`

	// HA gateway fault domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// HA fault domain for OCI.
	// +kubebuilder:validation:Optional
	HaFaultDomain *string `json:"haFaultDomain,omitempty" tf:"ha_fault_domain,omitempty"`

	// HA Gateway Size. Mandatory if enabling HA.
	// HA Gateway Size.
	// +kubebuilder:validation:Optional
	HaGwSize *string `json:"haGwSize,omitempty" tf:"ha_gw_size,omitempty"`

	// The image version of the HA gateway. Use aviatrix_gateway_image data source to programmatically retrieve this value for the desired ha_software_version. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the aviatrix_controller_config resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.
	// ha_image_version can be used to set the desired image version of the HA gateway. If set, we will attempt to update the gateway to the specified version.
	// +kubebuilder:validation:Optional
	HaImageVersion *string `json:"haImageVersion,omitempty" tf:"ha_image_version,omitempty"`

	// west-1a".
	// AZ of subnet being created for Insane Mode Spoke HA Gateway. Required for AWS if insane_mode is true and ha_subnet is set.
	// +kubebuilder:validation:Optional
	HaInsaneModeAz *string `json:"haInsaneModeAz,omitempty" tf:"ha_insane_mode_az,omitempty"`

	// HA OOB availability zone. Required if enabling Private OOB and HA. Example: "us-west-1b".
	// OOB HA availability zone.
	// +kubebuilder:validation:Optional
	HaOobAvailabilityZone *string `json:"haOobAvailabilityZone,omitempty" tf:"ha_oob_availability_zone,omitempty"`

	// HA OOB management subnet. Required if enabling Private OOB and HA. Example: "11.0.0.48/28".
	// OOB HA management subnet.
	// +kubebuilder:validation:Optional
	HaOobManagementSubnet *string `json:"haOobManagementSubnet,omitempty" tf:"ha_oob_management_subnet,omitempty"`

	// Availability Zone of the HA subnet. Required when Private Mode is enabled on the Controller and cloud_type is AWS or AWSGov with HA. Available in Provider version R2.23+.
	// Private Mode HA subnet availability zone.
	// +kubebuilder:validation:Optional
	HaPrivateModeSubnetZone *string `json:"haPrivateModeSubnetZone,omitempty" tf:"ha_private_mode_subnet_zone,omitempty"`

	// The software version of the HA gateway. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the HA gateway upgrade can be managed with the aviatrix_controller_config resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.
	// ha_software_version can be used to set the desired software version of the HA gateway. If set, we will attempt to update the gateway to the specified version. If left blank, the gateway software version will continue to be managed through the aviatrix_controller_config resource.
	// +kubebuilder:validation:Optional
	HaSoftwareVersion *string `json:"haSoftwareVersion,omitempty" tf:"ha_software_version,omitempty"`

	// HA Subnet. Required if enabling HA for AWS, AWSGov, AWSChina, Azure, AzureGov, AzureChina, OCI, Alibaba Cloud, AWS Top Secret or AWS Secret gateways. Optional for GCP. Setting to empty/unsetting will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet. Example: "10.12.0.0/24"
	// HA Subnet. Required if enabling HA for AWS/AWSGov/AWSChina/Azure/AzureChina/OCI/Alibaba Cloud. Optional if enabling HA for GCP.
	// +kubebuilder:validation:Optional
	HaSubnet *string `json:"haSubnet,omitempty" tf:"ha_subnet,omitempty"`

	// HA Zone. Required if enabling HA for GCP gateway. Optional for Azure. For GCP, setting to empty/unsetting will disable HA and setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c". For Azure, this is an optional parameter to place the HA gateway in a specific availability zone. Valid values for Azure gateways are in the form "az-n". Example: "az-2". Available for Azure as of provider version R2.17+.
	// HA Zone. Required if enabling HA for GCP. Optional for Azure.
	// +kubebuilder:validation:Optional
	HaZone *string `json:"haZone,omitempty" tf:"ha_zone,omitempty"`

	// A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC. Example: "10.4.0.0/116,10.5.0.0/16".
	// A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC.
	// +kubebuilder:validation:Optional
	IncludedAdvertisedSpokeRoutes *string `json:"includedAdvertisedSpokeRoutes,omitempty" tf:"included_advertised_spoke_routes,omitempty"`

	// , please see notes here.
	// Enable Insane Mode for Spoke Gateway. Valid values: true, false. Supported for AWS/AWSGov, GCP, Azure and OCI. If insane mode is enabled, gateway size has to at least be c5 size for AWS and Standard_D3_v2 size for Azure.
	// +kubebuilder:validation:Optional
	InsaneMode *bool `json:"insaneMode,omitempty" tf:"insane_mode,omitempty"`

	// AZ of subnet being created for Insane Mode Spoke Gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if insane_mode is enabled. Example: AWS: "us-west-1a".
	// AZ of subnet being created for Insane Mode Spoke Gateway. Required if insane_mode is enabled for AWS cloud.
	// +kubebuilder:validation:Optional
	InsaneModeAz *string `json:"insaneModeAz,omitempty" tf:"insane_mode_az,omitempty"`

	// Learned CIDRs approval mode. Either "gateway" (approval on a per-gateway basis) or "connection" (approval on a per-connection basis). Only "gateway" is supported for BGP SPOKE Gateway. Default value: "gateway". Available as of provider version R2.21+.
	// Set the learned CIDRs approval mode for BGP Spoke Gateway. Only valid when 'enable_learned_cidrs_approval' is set to true. Currently, only 'gateway' is supported: learned CIDR approval applies to ALL connections. Default value: 'gateway'.
	// +kubebuilder:validation:Optional
	LearnedCidrsApprovalMode *string `json:"learnedCidrsApprovalMode,omitempty" tf:"learned_cidrs_approval_mode,omitempty"`

	// Changes the Aviatrix Spoke Gateway ASN number before you setup Aviatrix Spoke Gateway connection configurations.
	// Changes the Aviatrix BGP Spoke Gateway ASN number before you setup Aviatrix BGP Spoke Gateway connection configurations.
	// +kubebuilder:validation:Optional
	LocalAsNumber *string `json:"localAsNumber,omitempty" tf:"local_as_number,omitempty"`

	// Enable to manage Aviatrix spoke HA gateway using the aviatrix_spoke_gateway resource. If this is set to false, spoke HA gateways must be managed using the aviatrix_spoke_ha_gateway resource. Valid values: true, false. Default value: true. Available in provider R3.0+.
	// This parameter is a switch used to determine whether or not to manage spoke ha gateway using the aviatrix_spoke_gateway resource. If this is set to false, managing spoke ha gateway must be done using the aviatrix_spoke_ha_gateway resource. Valid values: true, false. Default value: true.
	// +kubebuilder:validation:Optional
	ManageHaGateway *bool `json:"manageHaGateway,omitempty" tf:"manage_ha_gateway,omitempty"`

	// Set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true. Available in provider version R2.18+.
	// A set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true.
	// +kubebuilder:validation:Optional
	MonitorExcludeList []*string `json:"monitorExcludeList,omitempty" tf:"monitor_exclude_list,omitempty"`

	// OOB availability zone. Required if enabling Private OOB. Example: "us-west-1a".
	// OOB subnet availability zone.
	// +kubebuilder:validation:Optional
	OobAvailabilityZone *string `json:"oobAvailabilityZone,omitempty" tf:"oob_availability_zone,omitempty"`

	// OOB management subnet. Required if enabling Private OOB. Example: "11.0.2.0/24".
	// OOB management subnet.
	// +kubebuilder:validation:Optional
	OobManagementSubnet *string `json:"oobManagementSubnet,omitempty" tf:"oob_management_subnet,omitempty"`

	// List of AS numbers to populate BGP AS_PATH field when it advertises to VGW or peer devices.
	// List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices. Only valid for BGP Spoke Gateway
	// +kubebuilder:validation:Optional
	PrependAsPath []*string `json:"prependAsPath,omitempty" tf:"prepend_as_path,omitempty"`

	// VPC ID of Private Mode load balancer. Required when Private Mode is enabled on the Controller. Available in provider version R2.23+.
	// Private Mode controller load balancer vpc_id.  Required when private mode is enabled for the Controller.
	// +kubebuilder:validation:Optional
	PrivateModeLBVPCID *string `json:"privateModeLbVpcId,omitempty" tf:"private_mode_lb_vpc_id,omitempty"`

	// Availability Zone of the subnet. Required when Private Mode is enabled on the Controller and cloud_type is AWS or AWSGov. Available in Provider version R2.23+.
	// Subnet availability zone. Required when Private Mode is enabled on the Controller and cloud_type is AWS.
	// +kubebuilder:validation:Optional
	PrivateModeSubnetZone *string `json:"privateModeSubnetZone,omitempty" tf:"private_mode_subnet_zone,omitempty"`

	// Gateway ethernet interface RX queue size. Applies on HA as well if enabled. Once set, can't be deleted or disabled. Available for AWS as of provider version R2.22+.
	// Gateway ethernet interface RX queue size. Supported for AWS related clouds only. Applies on HA as well if enabled.
	// +kubebuilder:validation:Optional
	RxQueueSize *string `json:"rxQueueSize,omitempty" tf:"rx_queue_size,omitempty"`

	// Set to true if this feature is desired. Valid values: true, false.
	// Set to 'enabled' if this feature is desired.
	// +kubebuilder:validation:Optional
	SingleAzHa *bool `json:"singleAzHa,omitempty" tf:"single_az_ha,omitempty"`

	// Specify whether to enable Source NAT feature in "single_ip" mode on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently, only supports AWS(1) and Azure(8). Valid values: true, false.
	// Specify whether to enable Source NAT feature in 'single_ip' mode on the gateway or not.
	// +kubebuilder:validation:Optional
	SingleIPSnat *bool `json:"singleIpSnat,omitempty" tf:"single_ip_snat,omitempty"`

	// Intended CIDR list to be advertised to external BGP router. Empty list is not valid. Example: ["10.2.0.0/16", "10.4.0.0/16"].
	// Intended CIDR list to be advertised to external BGP router.
	// +kubebuilder:validation:Optional
	SpokeBGPManualAdvertiseCidrs []*string `json:"spokeBgpManualAdvertiseCidrs,omitempty" tf:"spoke_bgp_manual_advertise_cidrs,omitempty"`

	// Price for spot instance. NOT supported for production deployment.
	// Price for spot instance. NOT supported for production deployment.
	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20". NOTE: If using
	// Public Subnet Info.
	// +kubebuilder:validation:Required
	Subnet *string `json:"subnet" tf:"subnet,omitempty"`

	// Map of tags to assign to the gateway. Only available for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret gateways. Allowed characters vary by cloud type but always include: letters, spaces, and numbers. AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret allow the use of any character. Azure, AzureGov and AzureChina allows the following special characters: + - = . _ : @. Example: {"key1" = "value1", "key2" = "value2"}.
	// A map of tags to assign to the spoke gateway.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The IPsec tunnel down detection time for the Spoke Gateway in seconds. Must be a number in the range [20-600]. The default value is set by the controller (60 seconds if nothing has been changed). NOTE: The controller UI has an option to set the tunnel detection time for all gateways. Available in provider R2.19+.
	// The IPSec tunnel down detection time for the Spoke Gateway.
	// +kubebuilder:validation:Optional
	TunnelDetectionTime *float64 `json:"tunnelDetectionTime,omitempty" tf:"tunnel_detection_time,omitempty"`

	// VPC-ID/VNet-Name of cloud provider. Example: AWS/AWSGov/AWSChina: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project-id", Azure/AzureGov/AzureChina: "vnet_name:rg_name:resource_guid", OCI: "ocid1.vcn.oc1.iad.aaaaaaaaba3pv6wkcr4jqae5f44n2b2m2yt2j6rx32uzr4h25vqstifsfdsq".
	// VPC-ID/VNet-Name of cloud provider.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-west-1, AWSChina: "cn-north-1", AzureChina: "China North", AWS Top Secret: "us-iso-east-1", AWS Secret: "us-isob-east-1".
	// Region of cloud provider.
	// +kubebuilder:validation:Required
	VPCReg *string `json:"vpcReg" tf:"vpc_reg,omitempty"`

	// Availability Zone. Only available for Azure (8), Azure GOV (32) and Azure CHINA (2048). Must be in the form 'az-n', for example, 'az-2'. Available in provider version R2.17+.
	// Availability Zone. Only available for Azure (8), Azure GOV (32) and Azure CHINA (2048). Must be in the form 'az-n', for example, 'az-2'.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API. Creates and manages Aviatrix spoke gateways
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aviatrix}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
