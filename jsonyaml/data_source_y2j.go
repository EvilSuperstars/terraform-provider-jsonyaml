package jsonyaml

import (
	"time"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceYamlToJson() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceYamlToJsonRead,

		Schema: map[string]*schema.Schema{
			"yaml": {
				Type:     schema.TypeString,
				Required: true,
			},

			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceYamlToJsonRead(d *schema.ResourceData, meta interface{}) error {
	json, err := yaml.YAMLToJSON([]byte(d.Get("yaml").(string)))
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("json", string(json))

	return nil
}
