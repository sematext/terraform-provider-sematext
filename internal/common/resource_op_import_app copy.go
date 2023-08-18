package sematext

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// ResourceOperationImportApp  is a common import handler used by most resources.
func ResourceOpImportApp(appType string) *schema.ResourceImporter {

	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
