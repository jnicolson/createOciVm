package main

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/rs/zerolog/log"
)

func GetAvailabilityDomains(tenancyID string) {
	client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creating Identity Client")
	}

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := client.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	for i, ad := range r.Items {
		fmt.Printf("Availability Domain %d: Id: %s, Name: %s\n", i, *ad.Id, *ad.Name)
	}
}

func GetCompartmentId(compartmentName string) *string {
	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creating Identity Client")
	}

	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		log.Fatal().Err(err).Msg("Fetching Tenancy ID")
	}

	request := identity.ListCompartmentsRequest{
		CompartmentId: &tenancyID,
		Name:          &compartmentName,
	}

	compartments, err := identityClient.ListCompartments(context.Background(), request)
	if err != nil {
		log.Fatal().Err(err).Msg("Listing Compartments")
	}

	/*for i, c := range compartments.Items {
		fmt.Printf("Compartment %d: Id: %s, Name: %s\n", i, *c.Id, *c.Name)
	}*/

	return compartments.Items[0].Id

}
