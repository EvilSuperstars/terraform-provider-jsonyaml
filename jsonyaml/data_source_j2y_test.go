package jsonyaml

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceJ2YConfig_basic = `
provider "jsonyaml" {}

data "jsonyaml_j2y" "foo" {
  json =<<EOS
{
  "Key1": "Value1",
  "Key2": [42, true],
  "Key3": {
    "Key1": [
	  {
	    "Key1": "Value1"
	  }
	]
  }
}
EOS
}
`

func TestDataSourceJ2Y_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceJ2YConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.jsonyaml_j2y.foo", "yaml"),
				),
			},
		},
	})
}
