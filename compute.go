package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func ListImageByFamily(operatingSystem string) {
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creaing Compute Client")
	}

	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		log.Fatal().Err(err).Msg("Fetching Tenancy ID")
	}

	request := core.ListImagesRequest{
		CompartmentId:   &tenancyID,
		OperatingSystem: &operatingSystem,
	}

	r, err := client.ListImages(context.Background(), request)
	if err != nil {
		log.Fatal().Err(err).Msg("Listing Images")
	}

	for i, image := range r.Items {
		fmt.Printf("Image %d: Id: %s, Name: %s, Operating System: %s\n", i, *image.Id, *image.DisplayName, *image.OperatingSystem)
	}
}

func GetImageId(imageName string) *string {
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Creaing Compute Client")
	}

	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		log.Fatal().Err(err).Msg("Fetching Tenancy ID")
	}

	request := core.ListImagesRequest{
		CompartmentId: &tenancyID,
		DisplayName:   &imageName,
	}

	r, err := client.ListImages(context.Background(), request)
	if err != nil {
		log.Fatal().Err(err).Msg("Listing Images")
	}
	return r.Items[0].Id

}

func GetCapacity(availabilityDomain string, compartmentID *string, shape string, memory, cpus float32) bool {
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Fatal().Err(err).Msg("Create Compute Client")
	}

	req := core.CreateComputeCapacityReportRequest{
		CreateComputeCapacityReportDetails: core.CreateComputeCapacityReportDetails{
			AvailabilityDomain: common.String(availabilityDomain),
			CompartmentId:      compartmentID,
			ShapeAvailabilities: []core.CreateCapacityReportShapeAvailabilityDetails{{
				InstanceShape: common.String(shape),
				InstanceShapeConfig: &core.CapacityReportInstanceShapeConfig{
					MemoryInGBs: common.Float32(memory),
					Ocpus:       common.Float32(cpus),
				},
			},
			},
		},
	}

	resp, err := client.CreateComputeCapacityReport(context.Background(), req)
	if err != nil {
		log.Fatal().Err(err).Msg("Create Capacity Report")
	}

	if resp.ShapeAvailabilities[0].AvailabilityStatus == core.CapacityReportShapeAvailabilityAvailabilityStatusOutOfHostCapacity {
		return false
	} else {
		return true
	}

}

func CreateInstance(shape string, memory, cpus float32, availabilityDomain, instanceName string, compartmentID, imageID, subnetID, sshKey *string) {
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	request := core.LaunchInstanceRequest{}
	request.CompartmentId = compartmentID
	request.DisplayName = common.String(instanceName)
	request.AvailabilityDomain = common.String(availabilityDomain)
	request.Shape = &shape

	request.Metadata = map[string]string{"ssh_authorized_keys": *sshKey}

	request.SourceDetails = core.InstanceSourceViaImageDetails{ImageId: imageID}
	request.CreateVnicDetails = &core.CreateVnicDetails{SubnetId: subnetID}

	request.ShapeConfig = &core.LaunchInstanceShapeConfigDetails{
		Ocpus:       common.Float32(cpus),
		MemoryInGBs: common.Float32(memory),
	}

	createResp, err := client.LaunchInstance(context.Background(), request)

	if err != nil {
		log.Fatal().Err(err).Msg("Launching Instance")
	}

	fmt.Printf("Instance ID: %s", *createResp.Id)

}
