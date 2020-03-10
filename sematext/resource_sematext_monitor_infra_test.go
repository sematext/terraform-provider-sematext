package sematext

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/golang/api"
)

// testAccSematextMonitorInfra_Create tests resource creation.
func testAccSematextMonitorInfraBasic(t *testing.T) {
	name := "test-infra-" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	description := "TESTING : SematextMonitorInfra_Create : Create Step"
	plan := "basic"
	discountCode := ""
	fixture := testAccSematextMonitorInfra_Fixture(name, description, plan, discountCode)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSematextMonitorInfra_ConfirmDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fixture,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorInfra_CheckConsistency(name),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "description", description),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "plan", plan),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discountCode),
				),
			},
		},
	})
}

// testAccSematextMonitorInfra_Update tests for resource updates.
func testAccSematextMonitorInfraUpdate(t *testing.T) {
	name := "test-infra-" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	description1 := "TESTING : SematextMonitorInfra_Update : Create Step"
	description2 := "TESTING : SematextMonitorInfra_Update : Update Step"
	plan := "basic"
	discountCode := ""
	fixture1 := testAccSematextMonitorInfra_Fixture(name, description_1, plan, discountCode)
	fixture2 := testAccSematextMonitorInfra_Fixture(name, description_2, plan, discountCode)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSematextMonitorInfra_ConfirmDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fixture_1,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorInfraCheckConsistency(name),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "description", description1),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "plan", plan),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discountCode),
				),
			},
			{
				Config: fixture_2,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorInfra_CheckConsistency(name),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "description", description2),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "plan", plan),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discountCode),
				),
			},
		},
	})
}

// testAccExampleResource returns a fixture resource.
func testAccSematextMonitorInfraFixture(name string, description string, billingPlan string, discountCode string) string {

	randomIdentifier := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	if name == "" {
		name = "infra-" + randomIdentifier
	}
	if description == "" {
		description = "TESTING : SematextMonitorInfra_Create : Default Step"
	}
	if billingPlan == "" {
		billingPlan = "basic"
	}
	if discountCode == "" {
		discountCode = ""
	}

	// TODO - check rest of CreateApp to see if any fields left out

	result := fmt.Sprintf(`
	resource "sematext_monitor_infra" "%s" {
		"name": "%s",
		"description": "%s",
		"billing_plan": "%s"
		"discount_code": "%s"
	}
	`, name, name, description, billingPlan, discountCode)

	return result
}

// testAccSematextMonitorInfra_CheckConsistency checks the App ID exists in both state and API.
func testAccSematextMonitorInfraCheckConsistency(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Resource not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return errors.New("No ID is set")
		}

		app := &api.App{}
		var err error

		client := testAccProvider.Meta().(*api.Client)
		app, err = app.Retrieve(rs.Primary.ID, client)
		if err != nil {
			return fmt.Errorf("Error in checking monitor %s, %s", name, err)
		}
		if app == nil {
			return fmt.Errorf("Error in checking monitor %s", name)
		}

		// TODO - test rest of values are consistent with both state and resource

		return nil
	}
}

// testAccSematextMonitorInfra_ConfirmDestroyed -  check is destroyed in API
func testAccSematextMonitorInfraConfirmDestroyed(s *terraform.State) error {

	app := new(api.App)
	err := new(error)
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "sematext_monitor_infra" {
			continue
		}

		app, err = app.Retrieve(rs.Primary.ID, client)

		if err != nil {
			return fmt.Errorf("Error in checking monitor deletion %s", err)
		}
		if !app.IsArchived() {
			return fmt.Errorf("Error monitor %s is not marked as archived after deletion", rs.Primary.ID)
		}

	}

	return nil

}
