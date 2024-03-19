# Create OCI VM

This is a small utlity created in [Go](https://go.dev) to use the Oracle Cloud Infrastructure (OCI) API and go SDK to request a Virtual machine on a schedule.

The purpose of this is on the OCI free tier you are eligible for 4 CPU cores and 24GB of memory, but there's only a certain amount of these available in each region.

This can be used to request them on a schedule (i.e. every 10 minutes).  

For this to work, the OCI connection details need to be placed in the .oci directory.  This will consist of a file named config and a private key used to log on to the Oracle cloud.  The easiest way to set this up is with the [OCI Command Line Interface](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/cliinstall.htm) and then copying the contents of ~/.oci into the .oci directory within this repository.  After this the key_file variable should be updated to /.oci/path-to-key.pem

Once configured, the config.toml.sample file should be copied to config.toml and updated, replacing the sections noted.  