package sematext

// TODO - Expand Resource test cases to full checks.

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/api"
)

// ResourceTestFixture a common test fixture representing a resource
type ResourceTestFixture struct {
	ResourceType string
	ResourceName string
	AppType      string
	Name         string
	StatePath    string
	PlanID       int
	DiscountCode string
}

// Hydrate populates a test resource with some default values
func (rtf *ResourceTestFixture) Hydrate(resourceType string, appType string) *ResourceTestFixture {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AppType = appType
	rtf.Name = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = api.AssignPlanID(rtf.AppType)

	rtf.DiscountCode = api.TestDiscountCodeMetrics

	fmt.Println("------------------------------------")
	fmt.Println(rtf.PlanID)
	fmt.Println("------------------------------------")

	return rtf
}

// FixtureToHCL Formats a common struct containing a resource to HCL.
func (rtf *ResourceTestFixture) FixtureToHCL() string {

	// TODO Discount Code temprarily removed
	result := fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		billing_plan_id = %d
		discount_code = "%s"
	}
	`,
		rtf.ResourceType,
		rtf.ResourceName,
		rtf.Name,
		rtf.PlanID,
		rtf.DiscountCode,
	)

	return result
}

// CommonMonitorBasicTest is a common test of resource creation.
func CommonMonitorBasicTest(t *testing.T, resourceType string, appType string) {

	rtf := (&ResourceTestFixture{}).Hydrate(resourceType, appType)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: ConfirmMonitorDestruction(rtf),
		Steps: []resource.TestStep{
			{
				Config: rtf.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf),
					resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
					resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
					resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
				),
			},
		},
	})
}

// CommonMonitorUpdateTest tests for resource updates.
func CommonMonitorUpdateTest(t *testing.T, resourceType string, appType string) {

	rtf1 := (&ResourceTestFixture{}).Hydrate(resourceType, appType)
	rtf2 := rtf1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: ConfirmMonitorDestruction(rtf2),
		Steps: []resource.TestStep{
			{
				Config: rtf1.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf1),
					resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
					resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
					resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
				),
			},
			{
				Config: rtf2.FixtureToHCL(),
				Check: resource.ComposeTestCheckFunc(
					ConfirmMonitorCreation(rtf2),
					resource.TestCheckResourceAttr(rtf2.StatePath, "name", rtf2.Name),
					resource.TestCheckResourceAttr(rtf2.StatePath, "billing_plan_id", strconv.Itoa(rtf2.PlanID)),
					resource.TestCheckResourceAttr(rtf2.StatePath, "discount_code", rtf2.DiscountCode),
				),
			},
		},
	})
}

// ConfirmMonitorCreation checks the App ID exists in both state and API.
func ConfirmMonitorCreation(rtf *ResourceTestFixture) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		var id int
		var found bool
		var err error
		var app *api.App
		var rs *terraform.ResourceState

		client := testAccProvider.Meta().(*api.Client)

		spew.Dump(s.RootModule().Resources)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.Atoi(rs.Primary.ID); err != nil {
			return err
		}

		if app, err = (&api.App{}).Load(id, client); err != nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s, %s", rtf.StatePath, err) //TODO Check return value fs function sig
		}
		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmMonitorDestruction checks the App ID exists in both state and API and is marked as DISABLED.
func ConfirmMonitorDestruction(rtf *ResourceTestFixture) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		var id int
		var found bool
		var err error
		var app *api.App
		var rs *terraform.ResourceState

		client := testAccProvider.Meta().(*api.Client)

		if rs, found = s.RootModule().Resources[rtf.ResourceType]; !found {
			return nil
		}
		if app, err = (&api.App{}).Load(id, client); err != nil {
			return fmt.Errorf("ConfirmMonitorDestruction : Error in checking monitor %s, %s", rtf.StatePath, err)
		}
		if !app.IsArchived() {
			return fmt.Errorf("ConfirmMonitorDestruction : Error monitor %s is not marked as archived after deletion", rs.Primary.ID)
		}

		return nil
	}
}
