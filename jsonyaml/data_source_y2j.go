package jsonyaml

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/yaml.v2"
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
	var input interface{}
	if err := yaml.Unmarshal([]byte(d.Get("input").(string)), &input); err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())

	return nil
}
