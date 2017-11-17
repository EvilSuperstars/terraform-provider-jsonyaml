package jsonyaml

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceConfig_boolean = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
---
	true
EOS
}
`

const testDataSourceConfig_string = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
---
	"SSS"
EOS
}
`

const testDataSourceConfig_number = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
	123
EOS
}
`

const testDataSourceConfig_string_array = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
	["ABC", "zyxwv"]
EOS
}
`

const testDataSourceConfig_object = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
{
	"name": "Washington",
	"state": "DC"
}
EOS
}
`

const testDataSourceConfig_empty_object = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
	{}
EOS
}
`

const testDataSourceConfig_empty_array = `
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
	[]
EOS
}
`

func TestDataSource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceConfig_boolean,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "boolean", "true"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
			{
				Config: testDataSourceConfig_string,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "string", "SSS"),
				),
			},
			{
				Config: testDataSourceConfig_number,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "number", "123"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
			{
				Config: testDataSourceConfig_string_array,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "string_array.#", "2"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "string_array.0", "ABC"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
			{
				Config: testDataSourceConfig_empty_array,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "string_array.#", "0"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
			{
				Config: testDataSourceConfig_object,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckResourceAttr("data.yamldecode_decode.foo", "object.state", "DC"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
			{
				Config: testDataSourceConfig_empty_object,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "boolean"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object_array"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "object"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "number"),
					resource.TestCheckNoResourceAttr("data.yamldecode_decode.foo", "string"),
				),
			},
		},
	})
}
