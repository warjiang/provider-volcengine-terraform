package vpc

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("volcengine_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "vpc"
		r.Kind = "VPC"
	})

	p.AddResourceConfigurator("volcengine_subnet", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "subnet"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "volcengine_vpc",
		}
	})
}
