/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/warjiang/provider-volcengine-terraform/config/vpc"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	// "github.com/warjiang/provider-volcengine-terraform/config/null"
)

const (
	resourcePrefix = "volcengine-terraform"
	modulePath     = "github.com/warjiang/provider-volcengine-terraform"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("volcengine-terraform.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// null.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
