package sematext

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/terraform-provider-sematext/sematext.com/api"
)

// testAccSematextMonitorInfra_Create tests resource creation.
func testAccSematextMonitorInfra_Basic(t *testing.T) {
	name := "test-infra-" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	description := "TESTING : SematextMonitorInfra_Create : Create Step"
	plan := "basic"
	discount_code := ""
	fixture := testAccSematextMonitorInfra_Fixture(name, description, plan, discount_code)

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
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discount_code),
				),
			},
		},
	})
}

// testAccSematextMonitorInfra_Update tests for resource updates.
func testAccSematextMonitorInfra_Update(t *testing.T) {
	name := "test-infra-" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	description_1 := "TESTING : SematextMonitorInfra_Update : Create Step"
	description_2 := "TESTING : SematextMonitorInfra_Update : Update Step"
	plan := "basic"
	discount_code := ""
	fixture_1 := testAccSematextMonitorInfra_Fixture(name, description_1, plan, discount_code)
	fixture_2 := testAccSematextMonitorInfra_Fixture(name, description_2, plan, discount_code)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSematextMonitorInfra_ConfirmDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fixture_1,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorInfra_CheckConsistency(name),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "description", description_1),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "plan", plan),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discount_code),
				),
			},
			{
				Config: fixture_2,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorInfra_CheckConsistency(name),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "description", description_2),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "plan", plan),
					resource.TestCheckResourceAttr("sematext_monitor_infra."+name, "discount_code", discount_code),
				),
			},
		},
	})
}

// testAccExampleResource returns a fixture resource.
func testAccSematextMonitorInfra_Fixture(name string, description string, billing_plan string, discount_code string) string {

	random_identifier := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	if name == "" {
		name = "infra-" + random_identifier
	}
	if description == "" {
		description := "TESTING : SematextMonitorInfra_Create : Default Step"
	}
	if billing_plan == "" {
		billing_plan = "basic"
	}
	if discount_code == "" {
		discount_code = ""
	}

	// TODO - check rest of CreateApp to see if any fields left out

	result := fmt.Sprintf(`
	resource "sematext_monitor_infra" "%s" {
		"name": "%s",
		"description": "%s",
		"billing_plan": "%s"
		"discount_code": "%s"
	}
	`, name, name, description, billing_plan, discount_code)

	return result
}

// testAccSematextMonitorInfra_CheckConsistency checks the App ID exists in both state and API.
func testAccSematextMonitorInfra_CheckConsistency(name string) resource.TestCheckFunc {
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
func testAccSematextMonitorInfra_ConfirmDestroyed(s *terraform.State) error {

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
