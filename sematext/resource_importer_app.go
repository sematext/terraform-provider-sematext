package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceImporterApp(appType string) *schema.ResourceImporter {

	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
