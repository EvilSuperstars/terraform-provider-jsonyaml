# `yamldecode` Provider

The Terraform [yamldecode](https://github.com/EvilSuperstars/terraform-provider-yamldecode) provider decodes a YAML string.

This provider requires no configuration.

### Example Usage

```hcl
provider "yamldecode" {}

data "yamldecode_decode" "foo" {
  input =<<EOS
---
name: "Seattle",
state: "WA"
EOS
}
```

## Data Sources

### `yamldecode_decode`

#### Argument Reference

The following arguments are supported:

* `input` - (Required, string) The YAML string that is to be decoded. The subset of YAML that can be decoded is limited - boolean, number, string, object with string values and array of strings.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `boolean` - (boolean) Boolean
* `number` - (float) Number
* `string` - (string) String
* `object` - (map) Object with string values
* `string_array` - (list) Array of strings
