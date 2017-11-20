package jsonyaml

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"gopkg.in/yaml.v2"
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
	var v interface{}

	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceJ2YConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckYamlExists("data.jsonyaml_j2y.foo", "yaml", &v),
					testAccCheckYamlStructure(&v),
				),
			},
		},
	})
}

func testAccCheckYamlExists(name, key string, v *interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}
		if is.ID == "" {
			return fmt.Errorf("No ID is set: %s", name)
		}

		y, ok := is.Attributes[key]
		if !ok {
			return fmt.Errorf("%s: Attribute '%s' not found", name, key)
		}

		return yaml.Unmarshal([]byte(y), v)
	}
}

func testAccCheckYamlStructure(y *interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		switch v := (*y).(type) {
		case map[interface{}]interface{}:
			k1 := v["Key1"]
			if k1.(string) != "Value1" {
				return fmt.Errorf("Incorrect value for Key1: %q", k1)
			}
			k2 := v["Key2"]
			if len(k2.([]interface{})) != 2 {
				return fmt.Errorf("Incorrect length for Key2's value: %q", k2)
			}
			break
		default:
			return fmt.Errorf("Incorrect top-level type: %T", v)
		}

		return nil
	}
}
