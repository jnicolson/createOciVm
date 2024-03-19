package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Create Oracle Cloud VM")

	config := ReadConfig("config.toml")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		<-sigch
		cancel()
	}()

	schedule := NewSchedule(ctx, config.General.Schedule)
	schedule.Start()

	// Scheduler stuff

	// Setup Stuff
	compartmentId := GetCompartmentId(config.General.Compartment)
	imageId := GetImageId(config.Instance.Image)
	networkId := GetNetworkId(compartmentId, config.Network.VcnName)
	subnetId := GetSubnetId(compartmentId, networkId, config.Network.SubnetName)

	test := func() {
		available := GetCapacity(config.General.AvailabilityDomain, compartmentId, config.Instance.Shape, config.Instance.Memory, config.Instance.Cpus)

		if available {
			log.Info().Msg("Instance Available")
			CreateInstance(config.Instance.Shape, config.Instance.Memory, config.Instance.Cpus, config.General.InstanceName, config.General.AvailabilityDomain, compartmentId, imageId, subnetId, &config.General.SshKey)

			smtpServer := SmtpServer{
				Address:  config.Notification.Server,
				Port:     config.Notification.Port,
				Username: config.Notification.Username,
				Password: config.Notification.Password,
			}

			SendNotification(config.Notification.From, config.Notification.To, smtpServer)
			schedule.CancelJob()

		} else {
			log.Info().Msg("No availability")
		}
	}

	schedule.SetJob(test)

	// Loop Stuff

	<-ctx.Done()

	schedule.Stop()
	log.Info().Msg("Shutting Down...")

}
