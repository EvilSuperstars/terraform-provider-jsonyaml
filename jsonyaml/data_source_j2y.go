package jsonyaml

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/yaml.v2"
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
	var input interface{}
	if err := yaml.Unmarshal([]byte(d.Get("input").(string)), &input); err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())

	return nil
}
