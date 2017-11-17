# `jsonyaml` Provider

The Terraform [jsonyaml](https://github.com/EvilSuperstars/terraform-provider-jsonyaml) provider converts JSON to YAML and vice-versa.

This provider requires no configuration.

### Example Usage

```hcl
provider "jsonyaml" {}

data "jsonyaml_y2j" "foo" {
  yaml =<<EOS
---
name: "Seattle",
state: "WA"
EOS
}
```

## Data Sources

### `jsonyaml_j2y`

Converts JSON to YAML.

#### Argument Reference

The following arguments are supported:

* `json` - (Required, string) The JSON string that is to be converted to YAML.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `yaml` - (string) The YAML string

### `jsonyaml_y2j`

Converts YAML to JSON.

#### Argument Reference

The following arguments are supported:

* `yaml` - (Required, string) The YAML string that is to be converted to JSON.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `json` - (string) The JSON string
