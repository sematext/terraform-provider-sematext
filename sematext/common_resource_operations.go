package sematext

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/sematext-api-client/api"
)

// TODO - shift this back into terraform-provider-sematext

// CommonMonitorCreate TODO Doc Comment
func CommonMonitorCreate(d *schema.ResourceData, meta interface{}, apptype string) error {

	client := meta.(*api.Client)

	// TODO - Consider if pre-existence safeguard check should be done (if app is deleted it only changes status)
	// TODO Check for consistency between schema and update

	createAppInfo := &api.CreateAppInfo{}
	createAppInfo.AppType = apptype

	discountCode, discountCodePresent := d.GetOkExists("discount_code")
	if discountCodePresent {
		createAppInfo.DiscountCode = discountCode.(string)
	}

	initialPlanID, initialPlanIDPresent := d.GetOkExists("billing_plan_id")
	if initialPlanIDPresent {
		if _, found := api.LookupAppType2PlanID[initialPlanID.(int)]; !found {

			return fmt.Errorf("%i is invalid billing_plan_id for %s", initialPlanID, apptype)

		}
		createAppInfo.InitialPlanID = initialPlanID.(int64)
	} else {
		createAppInfo.InitialPlanID = int64(api.LookupPlanID2Apptypes[apptype][0])
	}

	name, namePresent := d.GetOkExists("name")
	if namePresent {
		createAppInfo.Name = name.(string)
	} else {
		return errors.New("Sematext monitor is missing a name field")
	}

	app, err := createAppInfo.Persist(client)
	if err != nil {
		return err
	}

	// Add fields that go in via app update packet

	updateAppInfo := &api.UpdateAppInfo{}

	description, descriptionExists := d.GetOkExists("description")
	if descriptionExists {
		updateAppInfo.Description = description.(string)
	} else {
		return fmt.Errorf("Sematext monitor is missing a description clause")
	}
	ignorePercentage, ignorePercentageExists := d.GetOkExists("ignore_percentage")
	if ignorePercentageExists {
		updateAppInfo.IgnorePercentage = ignorePercentage.(int)
	}
	maxEvents, maxEventsExists := d.GetOkExists("max_events")
	if maxEventsExists {
		updateAppInfo.MaxEvents = maxEvents.(int)
	}
	maxLimitMB, maxLimitMBExists := d.GetOkExists("max_limit_mb")
	if maxLimitMBExists {
		updateAppInfo.MaxLimitMB = maxLimitMB.(int)
	}
	sampling, samplingExists := d.GetOkExists("sampling")
	if samplingExists {
		updateAppInfo.Sampling = sampling.(bool)
	}
	samplingPercentage, samplingPercentageExists := d.GetOkExists("sampling_percentage")
	if samplingPercentageExists {
		updateAppInfo.SamplingPercentage = samplingPercentage.(int)
	}
	staggering, staggeringExists := d.GetOkExists("staggering")
	if staggeringExists {
		updateAppInfo.Staggering = staggering.(bool)
	}

	updateAppInfo.Status = "ACTIVE"

	app, err = updateAppInfo.Persist(app.ID, client)
	if err != nil {
		return err
	}

	d.SetId(app.ID)
	d.Set("token", app.Token) //TODO check this is available to other resources and decide if it should be a resource paramater in .tf script

	return nil

}

// CommonMonitorRead TODO Doc Comment
func CommonMonitorRead(d *schema.ResourceData, meta interface{}, apptype string) error {

	// TODO Check for consistency between schema and update

	client := meta.(*api.Client)
	id := d.Id()
	var app *api.App
	var err error
	app, err = app.Load(id, client)
	if err != nil {
		return err
	}

	d.Set("name", app.Name)
	d.Set("description", app.Description)
	d.Set("billing_plan_id", app.Plan.ID)
	//d.Set("discount_code", "???")     // TODO Where to read?
	//d.Set("ignore_percentage", "???") // TODO Where to read field?
	d.Set("max_events", app.Plan.MaxDailyEvents)
	//d.Set("max_limit_mb", "???")        // TODO Where to read field?
	//d.Set("sampling", "???")            // TODO Where to read field?
	//d.Set("sampling_percentage", "???") // TODO Where to read field?
	//d.Set("staggering", "???")          // TODO Where to read field?

	return nil
}

// CommonMonitorUpdate TODO Doc Comment
func CommonMonitorUpdate(d *schema.ResourceData, meta interface{}, apptype string) error {

	client := meta.(*api.Client)
	id := d.Id()

	updateAppInfo := &api.UpdateAppInfo{}
	appInfoChanged := false

	if d.HasChange("name") {
		_, newName := d.GetChange("name")
		updateAppInfo.Name = newName.(string) // TODO name vs resource name/label in terraform
		appInfoChanged = true
	}
	if d.HasChange("description") {
		_, newDescription := d.GetChange("description")
		updateAppInfo.Description = newDescription.(string)
		appInfoChanged = true
	}
	if d.HasChange("ignore_percentage") {
		_, newIgnorePercentage := d.GetChange("ignore_percentage")
		updateAppInfo.IgnorePercentage = newIgnorePercentage.(int)
		appInfoChanged = true
	}
	if d.HasChange("max_events") {
		_, newMaxEvents := d.GetChange("max_events")
		updateAppInfo.MaxEvents = newMaxEvents.(int)
		appInfoChanged = true
	}
	if d.HasChange("max_limit_mb") {
		_, newMaxLimitMB := d.GetChange("max_limit_mb")
		updateAppInfo.MaxLimitMB = newMaxLimitMB.(int)
		appInfoChanged = true
	}
	if d.HasChange("sampling") {
		_, newSampling := d.GetChange("sampling")
		updateAppInfo.Sampling = newSampling.(bool)
		appInfoChanged = true
	}
	if d.HasChange("sampling_percentage") {
		_, newSamplingPercentage := d.GetChange("sampling_percentage")
		updateAppInfo.SamplingPercentage = newSamplingPercentage.(int)
		appInfoChanged = true
	}
	if d.HasChange("staggering") {
		_, newStaggering := d.GetChange("staggering")
		updateAppInfo.Staggering = newStaggering.(bool)
		appInfoChanged = true
	}

	if appInfoChanged {
		updateAppInfo.Status = "ACTIVE" // TODO Consider if is this correct or not?
		updateAppInfo.Persist(id, client)
	}

	billingInfo := &api.BillingInfo{}
	billingInfoChanged := false

	if d.HasChange("billing_plan_id") {
		_, newBillingPlanID := d.GetChange("billing_plan_id")
		billingInfo.PlanID = newBillingPlanID.(int64)
		billingInfoChanged = true
	}

	// update only if both are valid

	var err error
	var valid bool

	if appInfoChanged {
		valid, err = updateAppInfo.IsValid()
		if !valid {
			return err
		}
	}
	if billingInfoChanged {
		valid, err = billingInfo.IsValid()
		if !valid {
			return err
		}
	}
	if appInfoChanged {
		_, err = updateAppInfo.Persist(id, client)
		if err != nil {
			return err
		}
	}
	if billingInfoChanged {
		err = billingInfo.Persist(id, client)
		if err != nil {
			return err
		}
	}

	return nil

}

// CommonMonitorDelete TODO Doc Comment
func CommonMonitorDelete(d *schema.ResourceData, meta interface{}, apptype string) error {

	client := meta.(*api.Client)
	id := d.Id()
	updateAppInfo := &api.UpdateAppInfo{}
	updateAppInfo.Status = "ARCHIVED"
	_, err := updateAppInfo.Persist(id, client)
	if err != nil {
		return err
	}

	return nil
}

// CommonMonitorExists TODO Doc Comment
func CommonMonitorExists(d *schema.ResourceData, meta interface{}, apptype string) (b bool, e error) {

	// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
	client := meta.(*api.Client)
	id := d.Id()
	var app api.App
	b, e = app.Exists(id, client)
	return b, e
}

// CommonMonitorImport TODO Doc Comment
func CommonMonitorImport(d *schema.ResourceData, meta interface{}, apptype string) ([]*schema.ResourceData, error) {

	// TODO Decide if Resource Import necessary for MVP
	return nil, nil

}
