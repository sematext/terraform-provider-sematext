package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

func ResourceOpImportAppAWS(client *stcloud.APIClient, ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse, appType string) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
