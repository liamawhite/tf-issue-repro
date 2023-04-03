// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repro

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var AccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"tsb": providerserver.NewProtocol6WithError(New("test")()),
}

func TestAccServiceAccountResource(t *testing.T) {

	initial := `
provider "provider" {}

resource provider_service_account "some-name" {}
	`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: AccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: initial,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("provider_service_account.some-name", "name", "some-name"),
					resource.TestCheckResourceAttr("provider_service_account.some-name", "keys.0.token", "some-token"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "provider_service_account.some-name",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: initial,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("provider_service_account.some-name", "name", "some-name"),
					resource.TestCheckResourceAttr("provider_service_account.some-name", "keys.0.token", "some-token"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
