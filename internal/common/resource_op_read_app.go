package sematext

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceOperationReadApp is a common read handler used by most resources.
func ResourceOpReadApp(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse, appType string, resourceModel *ResourceModel) {

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &ResourceModel)...)

	var appResponse stcloud.AppResponse
	var id int64
	var err error
	var app *stcloud.App
	var appTokenNames []string
	var tokensResponse stcloud.TokensResponse
	var tokenEntries *[]stcloud.TokenDto
	var tokenEntry stcloud.TokenDto
	var tokenAccumulator map[string]string
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)

	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	appResponse, httpResponse, err = client.AppsApi.GetUsingGET(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	//existance check
	if httpResponse.StatusCode == 404 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "resource not found during Read",
			Detail:   "resource '" + d.Get("name").(string) + "' is not present on Sematext Cloud during Read",
		})
		d.SetId("")
		return diags
	}

	app, err = extractApp(&appResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.Id)

	default:
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.Id)
	}

	// get the list of apptoken names that are supposed to be here
	appTokenNames = extractAppTokenNames(d.Get("apptoken"))

	// pull tokens for this app from SC.
	if tokensResponse, _, err = client.TokensApiControllerApi.GetAppTokens(ctx, id); err != nil {
		return diag.FromErr(err)
	}
	tokenEntries, err = extractAppTokens(tokensResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	// filter list of tokenEntries, ignore any not present in the Terraform HCL script
	//   if any are not writable output a warning
	//   if any are missing output a warning
	tokenAccumulator = map[string]string{}
	for _, tokenEntry = range *tokenEntries {
		if contains(appTokenNames, tokenEntry.Name) {
			tokenAccumulator[tokenEntry.Name] = tokenEntry.Token
			if !tokenEntry.Writeable {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "apptoken not writable",
					Detail:   "an apptoken named '" + tokenEntry.Name + "' exists but is not writable - update will fail unless made writable",
				})
			}
		}
	}

	for _, tokenName := range appTokenNames {
		if _, found := tokenAccumulator[tokenName]; !found {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "apptoken not present at sematext cloud",
				Detail:   "an apptoken named '" + tokenEntry.Name + "' is not yet present in your cloud account - terraform update will create this",
			})
		}
	}

	d.Set("sc_apptoken_entries", tokenAccumulator)

	return diags
}
