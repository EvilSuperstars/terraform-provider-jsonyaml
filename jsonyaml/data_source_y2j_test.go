package jsonyaml

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceY2JConfig_basic = `
provider "jsonyaml" {}

data "jsonyaml_y2j" "foo" {
  yaml =<<EOS
---
Key1: Value1
Key2:
  - 42
  - true
Key3:
  Key1:
    - Key1: Value1
EOS
}
`

func TestDataSourceY2J_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceY2JConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.jsonyaml_y2j.foo", "json"),
				),
			},
		},
	})
}
