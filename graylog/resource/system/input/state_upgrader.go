package input

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/convert"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/util"
)

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    util.UpgraderType(),
	Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		if err := convert.ConvertOneSizeListToJSON(rawState, keyAttributes); err != nil {
			return nil, err
		}
		return rawState, nil
	},
}
