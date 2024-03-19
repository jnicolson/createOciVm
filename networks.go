package main

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/rs/zerolog/log"
)

func GetNetworkId(compartmentId *string, networkName string) *string {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creating network client")
	}

	vcnRequest := core.ListVcnsRequest{
		CompartmentId: compartmentId,
		DisplayName:   common.String(networkName),
	}

	vcns, err := client.ListVcns(context.Background(), vcnRequest)
	if err != nil {
		log.Fatal().Err(err).Msg("List VCNs")
	}

	/*for i, vcn := range vcns.Items {
		fmt.Printf("VCN %d: Id: %s, Name: %s\n", i, *vcn.Id, *vcn.DisplayName)
	}*/

	return vcns.Items[0].Id

}

func GetSubnetId(compartmentID, vcnID *string, subnetName string) *string {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creating network client")
	}

	subnetRequest := core.ListSubnetsRequest{
		CompartmentId: compartmentID,
		VcnId:         vcnID,
		DisplayName:   &subnetName,
	}

	sub, err := client.ListSubnets(context.Background(), subnetRequest)
	if err != nil {
		log.Fatal().Err(err).Msg("List Subnets")
	}

	/*for i, s := range sub.Items {
		fmt.Printf("Subnet %d: Id: %s, Name: %s\n", i, *s.Id, *s.DisplayName)
	}*/

	return sub.Items[0].Id

}
