package jsonyaml

import (
	"time"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceJsonToYaml() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJsonToYamlRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Required: true,
			},

			"yaml": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJsonToYamlRead(d *schema.ResourceData, meta interface{}) error {
	yaml, err := yaml.JSONToYAML([]byte(d.Get("json").(string)))
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("yaml", string(yaml))

	return nil
}
