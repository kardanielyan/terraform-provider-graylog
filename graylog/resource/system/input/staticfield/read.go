package staticfield

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/client"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.Input.Get(ctx, d.Id())
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf("failed to get a input %s: %w", d.Id(), err))
	}
	return setDataToResourceData(d, map[string]interface{}{
		KeyInputID: d.Id(),
		KeyFields:  data[KeyStaticFields].(map[string]interface{}),
	})
}
