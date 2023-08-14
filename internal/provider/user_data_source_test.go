// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserDataSource(t *testing.T) {
	dataSourceName := "data.petstore_user.example"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccUserDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO: implement attribute checks
					resource.TestCheckResourceAttr(dataSourceName, "id", "example-id"),
				),
			},
		},
	})
}

// TODO: implement a valid test config.
const testAccUserDataSourceConfig = `
data "scaffolding_example" "test" {
  configurable_attribute = "example"
}
`