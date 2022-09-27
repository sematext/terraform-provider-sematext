package sematext

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// CommonResourceOperationCreateAlertRule is a common creation handler used by most resources.
func ResourceOperationCreateAlertRule(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var err error
	var id int64
	var alertRuleResponse stcloud.AlertRuleResponse
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient) // TODO - get client from context instead of meta?

	createAlertRule := &stcloud.AlertRule{}

	name, namePresent := d.GetOk("name")
	if namePresent {
		createAlertRule.Name = name.(string)
	} else {
		return diag.FromErr(errors.New("error : missing name field"))
	}

	description, descriptionPresent := d.GetOk("description")
	if descriptionPresent {
		createAlertRule.Description = description.(string)
	} else {
		return diag.FromErr(errors.New("error : missing description field"))
	}

	appID, appIDPresent := d.GetOk("appID")
	if appIDPresent {
		createAlertRule.AppId = appID.(int64)
	} else {
		return diag.FromErr(errors.New("error : missing appID field"))
	}

	backToNormalNeeded, backToNormalNeededPresent := d.GetOk("backToNormalNeeded")
	if backToNormalNeededPresent {
		createAlertRule.BackToNormalNeeded = backToNormalNeeded.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing backToNormalNeeded field"))
	}

	color, colorPresent := d.GetOk("Color") // TODO - check resource case
	if colorPresent {
		createAlertRule.Color = color.(string)
	} else {
		return diag.FromErr(errors.New("error : missing Color field"))
	}

	enabled, enabledPresent := d.GetOk("Enabled")
	if enabledPresent {
		createAlertRule.Enabled = enabled.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing enabled field"))
	}

	measureTrigger, measureTriggerPresent := d.GetOk("measure.trigger")
	if measureTriggerPresent {
		createAlertRule.EstimateOperation = measureTrigger.(string)
	} else {
		return diag.FromErr(errors.New("error : missing measure.trigger field"))
	}

	measureThreshold, measureThresholdPresent := d.GetOk("measure.threshold")
	if measureThresholdPresent {
		createAlertRule.EstimateValue = measureThreshold.(float64)
	} else {
		return diag.FromErr(errors.New("error : missing measure.trigger field"))
	}

	// Unpack filter subschema
	var filterValue stcloud.FilterValue
	var filterAccumulator []stcloud.FilterValue
	var filterEntry map[string]interface{}

	filterEntries, filtersPresent := d.GetOk("filter")
	if filtersPresent {

		filterAccumulator = []stcloud.FilterValue{}

		for _, entry := range filterEntries.([]interface{}) {
			filterValue = stcloud.FilterValue{}
			filterEntry = entry.(map[string]interface{})
			filterValue.AggType = filterEntry["aggType"].(string)
			filterValue.FilterName = filterEntry["filterName"].(string)
			filterValue.Key = filterEntry["key"].(string)
			filterValue.Label = filterEntry["label"].(string)
			filterValue.Name = filterEntry["name"].(string)
			filterValue.Values = filterEntry["values"].([]string)
			filterAccumulator = append(filterAccumulator, filterValue)
		}

		createAlertRule.FilterValuesObj = filterAccumulator

	} else {
		return diag.FromErr(errors.New("error : missing filter field"))
	}

	ignoreRegularEventsEnabled, ignoreRegularEventsEnabledPresent := d.GetOk("ignoreRegularEvents")
	if ignoreRegularEventsEnabledPresent {
		createAlertRule.IgnoreRegularEventsEnabled = ignoreRegularEventsEnabled.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing ignoreRegularEvents field"))
	}

	// TODO - reinstate when information available.
	// metadata, metadataPresent := d.GetOk("metadata")
	// if metadataPresent {

	// 	createAlertRule.Metadata = metadata.(string) // TODO unpack subschema - will be dependant on integration type

	// } else {
	// 	return diag.FromErr(errors.New("error : missing metadata field"))
	// }

	metricLabel, metricLabelPresent := d.GetOk("metricLabel")
	if metricLabelPresent {
		createAlertRule.MetricLabel = metricLabel.(string)
	} else {
		return diag.FromErr(errors.New("error : missing metricLabel field"))
	}

	minDelayBetweenNotificationsInMinutes, minDelayBetweenNotificationsInMinutesPresent := d.GetOk("minDelayBetweenNotificationsInMinutes")
	if minDelayBetweenNotificationsInMinutesPresent {
		createAlertRule.MinDelayBetweenNotificationsInMinutes = minDelayBetweenNotificationsInMinutes.(string)
	} else {
		return diag.FromErr(errors.New("error : missing minDelayBetweenNotificationsInMinutes field"))
	}

	notificationEmails, notificationEmailsPresent := d.GetOk("notificationEmails")
	if notificationEmailsPresent {
		createAlertRule.NotificationEmails = notificationEmails.([]string)
	} else {
		return diag.FromErr(errors.New("error : missing notificationEmails field"))
	}

	notificationIntegrationEntries, notificationIntegrationPresent := d.GetOk("notificationIntegrations")
	if notificationIntegrationPresent {

		entries := []stcloud.NotificationIntegration{}

		for i, entry := range notificationIntegrationEntries.([]interface{}) {

			ni := stcloud.NotificationIntegration{}
			ni.Applicability = entry.(map[string]interface{})["applicability"].(string)
			ni.IntegrationType = entry.(map[string]interface{})["type"].(string)
			ni.Name = entry.(map[string]interface{})["name"].(string)

			for k, v := range entry.(map[string]interface{})["params"].(map[string]string) {
				ni.Params[k] = v
			}

			ni.State = entry.(map[string]interface{})["state"].(string)

			entries[i] = ni
		}

		createAlertRule.NotificationIntegrations = entries

	} else {
		return diag.FromErr(errors.New("error : missing filter substanza"))
	}

	notificationsEnabled, notificationsEnabledPresent := d.GetOk("notificationsEnabled")
	if notificationsEnabledPresent {
		createAlertRule.NotificationsEnabled = notificationsEnabled.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing notificationsEnabled field"))
	}

	query, queryPresent := d.GetOk("query")
	if queryPresent {
		createAlertRule.Query = query.(string)
	} else {
		return diag.FromErr(errors.New("error : missing query field"))
	}

	reportName, reportNamePresent := d.GetOk("reportName")
	if reportNamePresent {
		createAlertRule.ReportName = reportName.(string)
	} else {
		return diag.FromErr(errors.New("error : missing reportName field"))
	}

	ruleType, ruleTypePresent := d.GetOk("ruleType")
	if ruleTypePresent {
		createAlertRule.RuleType = ruleType.(string)
	} else {
		return diag.FromErr(errors.New("error : missing ruleType field"))
	}

	runBook, runBookPresent := d.GetOk("runBook")
	if runBookPresent {
		createAlertRule.Runbook = runBook.(string)
	} else {
		return diag.FromErr(errors.New("error : missing runBook field"))
	}

	// Unpack Subschema for Schedule
	var alertRuleScheduleWeekdayDto stcloud.AlertRuleScheduleWeekdayDto
	var scheduleAccumulator []stcloud.AlertRuleScheduleWeekdayDto
	var intervalRangeEntry stcloud.AlertRuleScheduleTimeRangeDto

	scheduleEntries, scheduleEntriesPresent := d.GetOk("schedule")
	if scheduleEntriesPresent {

		scheduleAccumulator = []stcloud.AlertRuleScheduleWeekdayDto{}

		for _, scheduleEntry := range scheduleEntries.([]interface{}) {

			row := scheduleEntry.(map[string]interface{})

			alertRuleScheduleWeekdayDto = stcloud.AlertRuleScheduleWeekdayDto{}
			alertRuleScheduleWeekdayDto.Day = row["day"].(string)
			alertRuleScheduleWeekdayDto.Index = row["index"].(int32)
			alertRuleScheduleWeekdayDto.Label = row["label"].(string)
			alertRuleScheduleWeekdayDto.Type_ = row["type"].(string)

			for _, interval := range row["Intervals"].([]interface{}) {
				intervalRangeEntry.End = interval.(map[string]interface{})["end"].(string)
				intervalRangeEntry.Start = interval.(map[string]interface{})["start"].(string)
				alertRuleScheduleWeekdayDto.Intervals = append(alertRuleScheduleWeekdayDto.Intervals, intervalRangeEntry)
			}

			scheduleAccumulator = append(scheduleAccumulator, alertRuleScheduleWeekdayDto)
		}

		createAlertRule.Schedule = scheduleAccumulator

	} else {
		return diag.FromErr(errors.New("error : missing filter field"))
	}

	sendToEmail, sendToEmailPresent := d.GetOk("sendToEmail")
	if sendToEmailPresent {
		createAlertRule.SendToEmail = sendToEmail.(string)
	} else {
		return diag.FromErr(errors.New("error : missing sendToEmail field"))
	}

	timezone, timezonePresent := d.GetOk("timezone")
	if timezonePresent {
		createAlertRule.Timezone = timezone.(string)
	} else {
		return diag.FromErr(errors.New("error : missing timezone field"))
	}

	useOnlyAlertRuleIntegrations, useOnlyAlertRuleIntegrationsPresent := d.GetOk("useOnlyAlertRuleIntegrations")
	if useOnlyAlertRuleIntegrationsPresent {
		createAlertRule.UseOnlyAlertRuleIntegrations = useOnlyAlertRuleIntegrations.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing useOnlyAlertRuleIntegrations field"))
	}

	userPermissionsCanDelete, userPermissionsPresentCanDelete := d.GetOk("userPermissions.canDelete")
	if userPermissionsPresentCanDelete {
		createAlertRule.UserPermissions.CanDelete = userPermissionsCanDelete.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing userPermissions can Delete field"))
	}

	userPermissionsCanEdit, userPermissionsPresentCanEdit := d.GetOk("userPermissions.canDelete")
	if userPermissionsPresentCanEdit {
		createAlertRule.UserPermissions.CanEdit = userPermissionsCanEdit.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing userPermissions can Delete field"))
	}

	userPermissionsCanView, userPermissionsPresentCanView := d.GetOk("userPermissions.canDelete")
	if userPermissionsPresentCanView {
		createAlertRule.UserPermissions.CanView = userPermissionsCanView.(bool)
	} else {
		return diag.FromErr(errors.New("error : missing userPermissions can Delete field"))
	}

	valueColumnName, valueColumnNamePresent := d.GetOk("valueColumnName")
	if valueColumnNamePresent {
		createAlertRule.ValueColumnName = valueColumnName.(string)
	} else {
		return diag.FromErr(errors.New("error : missing valueColumnName field"))
	}

	valueName, valueNamePresent := d.GetOk("valueName")
	if valueNamePresent {
		createAlertRule.ValueColumnName = valueName.(string)
	} else {
		return diag.FromErr(errors.New("error : missing valueName field"))
	}

	alertRuleResponse, httpResponse, err = client.AlertsApi.CreateAlertUsingPOST(ctx, *createAlertRule)

	if httpResponse.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error creating alert rule",
			Detail:   "error creating alert rule '" + name.(string) + "'",
		})

		return diags
	}

	if err != nil {
		return diag.FromErr(err)
	}

	var alertRule *stcloud.AlertRule
	alertRule, err = extractAlertRule(alertRuleResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(id, 10))
	d.Set("appId", alertRule.AppId)

	// TODO - check computed field values to come back into state

	return diags

}

// ResourceOperationReadAlertRule is a common read handler used by most resources.
func ResourceOperationReadAlertRule(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	//var id int64 TODO reinstate when id available.
	var err error
	var httpResponse *http.Response
	var readAlertRuleResponse stcloud.AlertRuleResponse
	var readAlertRule *stcloud.AlertRule

	//client := meta.(*stcloud.APIClient) TODO reinstate when id available.

	// TODO - reinstate when available.
	// if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil { /// TODO - there is no id in the resource
	// 	return diag.FromErr(err)
	// }

	// TODO - reinstate when available.
	// readAlertRuleResponse, httpResponse, err = client.AlertsApi.GetAlertRulesForAppUsingGET1(ctx, id)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

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

	readAlertRule, err = extractAlertRule(readAlertRuleResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", readAlertRule.Name)
	d.Set("description", readAlertRule.Description)
	d.Set("appId", readAlertRule.AppId)
	d.Set("backToNormalNeeded", readAlertRule.BackToNormalNeeded)
	d.Set("color", readAlertRule.Color)
	d.Set("enabled", readAlertRule.Enabled)
	d.Set("measure.trigger", readAlertRule.EstimateOperation)
	d.Set("measure.threshold", readAlertRule.EstimateValue)
	d.Set("filter", readAlertRule.FilterValuesObj) // TODO - restate object into correct type
	d.Set("ignoreRegularEvents", readAlertRule.IgnoreRegularEventsEnabled)
	d.Set("metadata", readAlertRule.Metadata) // TODO - restate object into correct type
	d.Set("metricLabel", readAlertRule.MetricLabel)
	d.Set("minDelayBetweenNotificationsInMinutes", readAlertRule.MinDelayBetweenNotificationsInMinutes)
	d.Set("notificationEmails", readAlertRule.NotificationEmails)

	if readAlertRule.NotificationIntegrations != nil {
		entries := make([]interface{}, len(readAlertRule.NotificationIntegrations), len(readAlertRule.NotificationIntegrations))
		for i, notificationIntegration := range readAlertRule.NotificationIntegrations {
			entry := make(map[string]interface{})
			entry["id"] = notificationIntegration.Id
			entry["name"] = notificationIntegration.Name
			entry["state"] = notificationIntegration.State
			entry["applicability"] = notificationIntegration.Applicability
			entry["params"] = make(map[string]interface{}, len(notificationIntegration.Params))
			for k, v := range notificationIntegration.Params {
				entry["params"].(map[string]string)[k] = v
			}
			entries[i] = entry
		}
		d.Set("notificationIntegrations", entries)
	}

	d.Set("notificationsEnabled", readAlertRule.NotificationsEnabled)
	d.Set("query", readAlertRule.Query)
	d.Set("reportName", readAlertRule.ReportName)
	d.Set("ruleType", readAlertRule.RuleType)
	d.Set("runbook", readAlertRule.Runbook)

	if readAlertRule.Schedule != nil {
		entries := make([]interface{}, len(readAlertRule.Schedule), len(readAlertRule.Schedule))
		for i, schedule := range readAlertRule.Schedule {
			entry := make(map[string]interface{})
			entry["day"] = schedule.Day
			entry["index"] = schedule.Index
			entry["interval"] = make([]stcloud.AlertRuleScheduleTimeRangeDto, len(schedule.Intervals))
			for j, interval := range schedule.Intervals {
				entry["interval"].([]stcloud.AlertRuleScheduleTimeRangeDto)[j] = stcloud.AlertRuleScheduleTimeRangeDto{
					Start: interval.Start,
					End:   interval.End,
				}
			}
			entry["label"] = schedule.Label
			entries[i] = entry
		}
		d.Set("schedule", entries)
	}

	return diags
}

// ResourceOperationUpdateAlertRule is a common update handler used by most resources.
func ResourceOperationUpdateAlertRule(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var id int64
	var err error
	var httpResponse *http.Response
	var updateAlertRule *stcloud.AlertRule
	var alertRuleResponse stcloud.AlertRuleResponse
	var alertRuleChanged bool

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	updateAlertRule = &stcloud.AlertRule{}
	alertRuleChanged = false

	if d.HasChange("name") {
		_, newName := d.GetChange("name")
		updateAlertRule.Name = newName.(string)
		alertRuleChanged = true
	}

	if d.HasChange("description") {
		_, newDescription := d.GetChange("description")
		updateAlertRule.Description = newDescription.(string)
		alertRuleChanged = true
	}

	if d.HasChange("backToNormalNeeded") { // TODO check field name
		_, newBackToNormalNeeded := d.GetChange("BackToNormalNeeded")
		updateAlertRule.BackToNormalNeeded = newBackToNormalNeeded.(bool)
		alertRuleChanged = true
	}

	if d.HasChange("color") { // TODO check field name
		_, newColor := d.GetChange("color")
		updateAlertRule.Color = newColor.(string)
		alertRuleChanged = true
	}

	if d.HasChange("enabled") { // TODO check field name
		_, newEnabled := d.GetChange("enabled")
		updateAlertRule.Enabled = newEnabled.(bool)
		alertRuleChanged = true
	}

	if d.HasChange("measure.trigger") {
		_, newMeasureTrigger := d.GetChange("measure.trigger")
		updateAlertRule.EstimateOperation = newMeasureTrigger.(string)
		alertRuleChanged = true
	}

	if d.HasChange("measure.threshold") {
		_, newMeasureThreshold := d.GetChange("measure.threashold")
		updateAlertRule.EstimateOperation = newMeasureThreshold.(string)
		alertRuleChanged = true
	}

	if d.HasChange("filter") {

		var filterValue stcloud.FilterValue
		var filterAccumulator []stcloud.FilterValue
		var filterEntry map[string]interface{}

		_, filterEntries := d.GetChange("filter")
		filterAccumulator = []stcloud.FilterValue{}

		for _, entry := range filterEntries.([]interface{}) {
			filterValue = stcloud.FilterValue{}
			filterEntry = entry.(map[string]interface{})
			filterValue.AggType = filterEntry["aggType"].(string)
			filterValue.FilterName = filterEntry["filterName"].(string)
			filterValue.Key = filterEntry["key"].(string)
			filterValue.Label = filterEntry["label"].(string)
			filterValue.Name = filterEntry["name"].(string)
			filterValue.Values = filterEntry["values"].([]string)
			filterAccumulator = append(filterAccumulator, filterValue)
		}

		updateAlertRule.FilterValuesObj = filterAccumulator
		alertRuleChanged = true

	}

	// TODO - reinstate and adjust to unpack object once we have a proper schema
	// if d.HasChange("metadata") { // TODO unpack subobject
	// 	_, newMetadata := d.GetChange("metadata")
	// 	updateAlertRule.Metadata = newMetadata.(bool)
	// 	alertRuleChanged = true
	// }

	if d.HasChange("metricLabel") {
		_, newMetricLabel := d.GetChange("metricLabel")
		updateAlertRule.MetricLabel = newMetricLabel.(string)
		alertRuleChanged = true
	}

	if d.HasChange("minDelayBetweenNotificationsInMinutes") {
		_, newMinDelayBetweenNotificationsInMinutes := d.GetChange("minDelayBetweenNotificationsInMinutes")
		updateAlertRule.MinDelayBetweenNotificationsInMinutes = newMinDelayBetweenNotificationsInMinutes.(string)
		alertRuleChanged = true
	}

	if d.HasChange("notificationEmails") { // TODO check field name, unpack subobject
		_, newNotificationEmails := d.GetChange("notificationEmails")
		updateAlertRule.NotificationEmails = newNotificationEmails.([]string)
		alertRuleChanged = true
	}

	if d.HasChange("notificationIntegrations") {
		_, newNotificationIntegrations := d.GetChange("notify")
		entries := make([]stcloud.NotificationIntegration, len(newNotificationIntegrations.([]interface{})), len(newNotificationIntegrations.([]interface{})))
		for i, entry := range newNotificationIntegrations.([]stcloud.NotificationIntegration) {
			entries[i] = stcloud.NotificationIntegration{}
			entries[i].Name = entry.Name
			entries[i].Applicability = entry.Applicability
			entries[i].Id = entry.Id
			entries[i].IntegrationType = entry.IntegrationType
			entries[i].State = entry.State
			entries[i].Params = make(map[string]string)
			for k, v := range entry.Params {
				entries[i].Params[k] = v
			}
		}
		updateAlertRule.NotificationIntegrations = entries
		alertRuleChanged = true
	}

	if d.HasChange("notificationsEnabled") {
		_, newNotificationsEnabled := d.GetChange("notificationsEnabled")
		updateAlertRule.NotificationsEnabled = newNotificationsEnabled.(bool)
		alertRuleChanged = true
	}

	if d.HasChange("query") {
		_, newQuery := d.GetChange("query")
		updateAlertRule.Query = newQuery.(string)
		alertRuleChanged = true
	}

	if d.HasChange("reportName") {
		_, newReportName := d.GetChange("reportName")
		updateAlertRule.ReportName = newReportName.(string)
		alertRuleChanged = true
	}

	if d.HasChange("ruleType") {

		_, newRuleType := d.GetChange("ruleType")
		updateAlertRule.RuleType = newRuleType.(string)
		alertRuleChanged = true
	}

	if d.HasChange("runBook") {
		_, newRunbook := d.GetChange("runbook")
		updateAlertRule.Runbook = newRunbook.(string)
		alertRuleChanged = true
	}

	if d.HasChange("schedule") {
		_, newSchedule := d.GetChange("schedule")
		entries := make([]stcloud.AlertRuleScheduleWeekdayDto, len(newSchedule.([]interface{})), len(newSchedule.([]interface{})))
		for i, schedule := range newSchedule.([]interface{}) {
			entry := stcloud.AlertRuleScheduleWeekdayDto{}
			entry.Day = schedule.(map[string]interface{})["day"].(string)
			entry.Index = schedule.(map[string]interface{})["index"].(int32)
			entry.Intervals = make([]stcloud.AlertRuleScheduleTimeRangeDto, len(entry.Intervals))
			for j, interval := range entry.Intervals {
				entry.Intervals[j] = stcloud.AlertRuleScheduleTimeRangeDto{
					Start: interval.Start,
					End:   interval.End,
				}
			}
			entry.Label = schedule.(map[string]interface{})["label"].(string)
			entries[i] = entry
		}

		updateAlertRule.Schedule = entries
		alertRuleChanged = true
	}

	if d.HasChange("sendToEmail") { // TODO check field name, unpack array copy
		_, newSendToEmail := d.GetChange("sendToEmail")
		updateAlertRule.SendToEmail = newSendToEmail.(string)
		alertRuleChanged = true
	}

	if d.HasChange("timezone") { // TODO check field name, unpack subobject
		_, newTimezone := d.GetChange("timezone")
		updateAlertRule.Timezone = newTimezone.(string)
		alertRuleChanged = true
	}

	if d.HasChange("useOnlyAlertRuleIntegrations") { // TODO check field name, unpack subobject
		_, newUseOnlyAlertRuleIntegrations := d.GetChange("useOnlyAlertRuleIntegrations")
		updateAlertRule.UseOnlyAlertRuleIntegrations = newUseOnlyAlertRuleIntegrations.(bool)
		alertRuleChanged = true
	}

	if d.HasChange("userPermissions") { // TODO check field name, unpack subobject
		_, newUserPermissions := d.GetChange("userPermissions")
		updateAlertRule.UserPermissions = newUserPermissions.(*stcloud.UserPermissions) // TODO check this works
		alertRuleChanged = true
	}

	if d.HasChange("valueColumnName") { // TODO check field name, unpack subobject
		_, newValueColumnName := d.GetChange("valueColumnName")
		updateAlertRule.ValueColumnName = newValueColumnName.(string)
		alertRuleChanged = true
	}

	if d.HasChange("valueName") { // TODO check field name, unpack subobject
		_, newValueName := d.GetChange("valueName")
		updateAlertRule.ValueName = newValueName.(string)
		alertRuleChanged = true
	}

	//First disable and delete Id then create another one -> This will update with a new value Id
	if alertRuleChanged {

		_, _, err = client.AlertsApi.DisableAlertRuleUsingPUT(ctx, id)

		if err != nil {
			return diag.FromErr(err)
		}

		time.Sleep(2 * time.Second)

		_, httpResponse, err = client.AlertsApi.DeleteAlertRuleUsingDELETE1(ctx, id)
		if err != nil {
			return diag.FromErr(err)
		}

		if httpResponse.StatusCode != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "error deleting alert rule",
				Detail:   "error deleting an alert rule named '" + d.Get("name").(string) + "'",
			})

			return diags
		}

		alertRuleResponse, _, err = client.AlertsApi.CreateAlertUsingPOST(ctx, *updateAlertRule)

		if err != nil {
			return diag.FromErr(err)
		}

		if httpResponse.StatusCode != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "error replacing alert rule",
				Detail:   "error replacing an alert rule named '" + d.Get("name").(string) + "'",
			})

			return diags
		}

		var newAlertRule *stcloud.AlertRule

		newAlertRule, err = extractAlertRule(alertRuleResponse)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(strconv.FormatInt(id, 10)) // TODO - id not present in response

		// TODO - are there other values we need to extract from the resource into state?

	}

	return diags

}

// ResourceOperationDeleteAlertRule is a common retire handler used by most resources.
func ResourceOperationDeleteAlertRule(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var id int64
	var err error
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	_, _, err = client.AlertsApi.DisableAlertRuleUsingPUT(ctx, id)

	if err != nil {
		return diag.FromErr(err)
	}

	time.Sleep(2 * time.Second)

	_, httpResponse, err = client.AlertsApi.DeleteAlertRuleUsingDELETE1(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if httpResponse.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error deleting alert rule",
			Detail:   "error deleting an alert rule named '" + d.Get("name").(string) + "'",
		})

		return diags
	}

	return diags
}

// ResourceOperationImportApp  is a common import handler used by most resources.
func ResourceOperationImportAlertRule() *schema.ResourceImporter {

	return &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
}
