package main

type Config struct {
	General      GeneralConfig
	Instance     InstanceConfig
	Network      NetworkConfig
	Notification NotificationConfig
}

type GeneralConfig struct {
	AvailabilityDomain string
	Compartment        string
	SshKey             string
	Schedule           string
	InstanceName       string
}

type InstanceConfig struct {
	Shape  string
	Memory float32
	Cpus   float32
	Image  string
}

type NetworkConfig struct {
	VcnName    string
	SubnetName string
}

type NotificationConfig struct {
	From     string
	To       string
	Server   string
	Port     int
	Username string
	Password string
}
