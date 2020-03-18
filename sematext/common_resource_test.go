package sematext

// TODO - Expand Resource test cases to full checks.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/api"
)

// ResourceTestFixture a common test fixture representing a resource
type ResourceTestFixture struct {
	ResourceType       string
	ResourceTag        string
	Name               string
	StatePath          string
	Description        string
	Plan               string
	DiscountCode       string
	IgnorePercentage   int
	MaxEvents          int
	MaxLimitMb         int
	Sampling           bool
	SamplingPercentage int
	Staggering         bool
}

// Hydrate populates a test resource with some default values
func (rtf *ResourceTestFixture) Hydrate(resourceType string) *ResourceTestFixture {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceTag = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.Name = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.StatePath = rtf.ResourceType + "." + rtf.Name
	rtf.Description = fmt.Sprintf("TESTING : SematextMonitorBasic : %s : Create", rtf.ResourceType)
	rtf.Plan = api.TestBillingPlan
	rtf.DiscountCode = api.TestDiscountCodeMetrics
	rtf.IgnorePercentage = 10
	rtf.MaxEvents = 10
	rtf.MaxLimitMb = 10
	rtf.Sampling = false
	rtf.SamplingPercentage = 0
	rtf.Staggering = false

	return rtf
}

// FixtureToHCL Formats a common struct containing a resource to HCL.
func (rtf *ResourceTestFixture) FixtureToHCL() string {

	// TODO Discount Code temprarily removed
	result := fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		description = "%s"
		billing_plan = "%s"
		discount_code = "%s"
		ignore_percentage = %d
		max_events = %d
		max_limit_mb = %d
		sampling = %t
		sampling_percentage = %d
		staggering = %t
	}
	`,
		rtf.ResourceType,
		rtf.ResourceTag,
		rtf.Name,
		rtf.Description,
		rtf.Plan,
		rtf.DiscountCode,
		rtf.IgnorePercentage,
		rtf.MaxEvents,
		rtf.MaxLimitMb,
		rtf.Sampling,
		rtf.SamplingPercentage,
		rtf.Staggering,
	)

	return result
}

// CommonMonitorBasicTest is a common test of resource creation.
func CommonMonitorBasicTest(t *testing.T, resourceType string) {

	rtf := (&ResourceTestFixture{}).Hydrate(resourceType)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: ConfirmMonitorDestruction,
		Steps: []resource.TestStep{
			{
				Config: rtf.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf),
					resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
					resource.TestCheckResourceAttr(rtf.StatePath, "description", rtf.Description),
					resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan", rtf.Plan),
					resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					resource.TestCheckResourceAttr(rtf.StatePath, "ignore_percentage", strconv.Itoa(rtf.IgnorePercentage)),
					resource.TestCheckResourceAttr(rtf.StatePath, "max_events", strconv.Itoa(rtf.MaxEvents)),
					resource.TestCheckResourceAttr(rtf.StatePath, "max_limit_mb", strconv.Itoa(rtf.MaxLimitMb)),
					resource.TestCheckResourceAttr(rtf.StatePath, "sampling", strconv.FormatBool(rtf.Sampling)),
					resource.TestCheckResourceAttr(rtf.StatePath, "sampling_percentage", strconv.Itoa(rtf.SamplingPercentage)),
					resource.TestCheckResourceAttr(rtf.StatePath, "staggering", strconv.FormatBool(rtf.Staggering)),
				),
			},
		},
	})
}

// CommonMonitorUpdateTest tests for resource updates.
func CommonMonitorUpdateTest(t *testing.T, resourceType string) {

	rtf1 := (&ResourceTestFixture{}).Hydrate(resourceType)
	rtf2 := rtf1
	rtf2.Description = "TESTING : SematextMonitorBasic : " + rtf2.ResourceType + " : Update"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: ConfirmMonitorDestruction,
		Steps: []resource.TestStep{
			{
				Config: rtf1.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf1),
					resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
					resource.TestCheckResourceAttr(rtf1.StatePath, "description", rtf1.Description),
					resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan", rtf1.Plan),
					resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					resource.TestCheckResourceAttr(rtf1.StatePath, "ignore_percentage", strconv.Itoa(rtf1.IgnorePercentage)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "max_events", strconv.Itoa(rtf1.MaxEvents)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "max_limit_mb", strconv.Itoa(rtf1.MaxLimitMb)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "sampling", strconv.FormatBool(rtf1.Sampling)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "sampling_percentage", strconv.Itoa(rtf1.SamplingPercentage)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "staggering", strconv.FormatBool(rtf1.Staggering)),
				),
			},
			{
				Config: rtf2.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf2),
					resource.TestCheckResourceAttr(rtf2.StatePath, "name", rtf2.Name),
					resource.TestCheckResourceAttr(rtf2.StatePath, "description", rtf2.Description),
					resource.TestCheckResourceAttr(rtf2.StatePath, "billing_plan", rtf2.Plan),
					resource.TestCheckResourceAttr(rtf2.StatePath, "discount_code", rtf2.DiscountCode),
					resource.TestCheckResourceAttr(rtf2.StatePath, "ignore_percentage", strconv.Itoa(rtf2.IgnorePercentage)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "max_events", strconv.Itoa(rtf2.MaxEvents)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "max_limit_mb", strconv.Itoa(rtf2.MaxLimitMb)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "sampling", strconv.FormatBool(rtf2.Sampling)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "sampling_percentage", strconv.Itoa(rtf2.SamplingPercentage)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "staggering", strconv.FormatBool(rtf2.Staggering)),
				),
			},
		},
	})
}

// ConfirmMonitorCreation checks the App ID exists in both state and API.
func ConfirmMonitorCreation(rtf *ResourceTestFixture) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(*api.Client)

		rs, ok := s.RootModule().Resources[rtf.Name]
		if !ok {
			return fmt.Errorf("Resource not found: %s", rtf.Name)
		}
		if rs.Primary.ID == "" {
			return errors.New("No ID is set")
		}

		app, err := (&api.App{}).Load(rs.Primary.ID, client)
		if err != nil {
			return fmt.Errorf("Error in checking monitor %s, %s", rtf.StatePath, err) //TODO Check return value fs function sig
		}
		if app == nil {
			return fmt.Errorf("Error in checking monitor %s", rtf.StatePath)
		}

		// TODO - test rest of values are consistent with both state and resource

		return nil
	}
}

// ConfirmMonitorDestruction -  check is destroyed in API
func ConfirmMonitorDestruction(s *terraform.State) error {

	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "sematext_monitor_<<MONITOR_TYPE>>" {
			continue
		}

		app, err := (&api.App{}).Load(rs.Primary.ID, client)

		if err != nil {
			return fmt.Errorf("Error in checking monitor deletion %s", err)
		}
		if !app.IsArchived() {
			return fmt.Errorf("Error monitor %s is not marked as archived after deletion", rs.Primary.ID)
		}

	}

	return nil

}
