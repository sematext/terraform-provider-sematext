package sematext

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceOperationDeleteApp is a common retire handler used by most resources.
func ResourceOpDeleteApp(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var id int64
	var err error
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	updateAppInfo.Status = "DISABLED"
	_, _, err = client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)
	if err != nil {
		return diag.FromErr(err)
	}

	time.Sleep(2 * time.Second)

	_, httpResponse, err = client.AppsApi.DeleteUsingDELETE(context.Background(), id)
	if err != nil {
		return diag.FromErr(err)
	}

	if httpResponse.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error deleting app",
			Detail:   "error deleting an app named '" + d.Get("name").(string) + "'",
		})

		return diags
	}

	return diags
}
