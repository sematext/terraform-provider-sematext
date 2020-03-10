package sematext


// TODO - Expand Resource test cases to full checks.

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/golang/api"
)

// ResourceTestFixtureSolrcloud TODO Doc Comment
struct ResourceTestFixtureSolrcloud {
	Name string
	Description string
	Plan string
	Discount_code string
	Ignore_percentage int
	Max_events int
	Max_limit_mb int
	Sampling bool
	Sampling_percentage int
	Staggering bool
}

// HydrateBasic TODO Doc Comment
func (rtf *ResourceTestFixtureSolrcloud) HydrateBasic() (*ResourceTestFixtureSolrcloud){
	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.Name = strings.ToLowercase(fmt.Sprintf("Solrcloud_test_%s", rndID))
	rtf.Description = "TESTING : SematextMonitorSolrcloud_Basic : Create"
	rtf.Plan = "basic"
	rtf.Discount_code = "testing"
	rtf.Ignore_percentage = 10
	rtf.Max_events  = 10
	rtf.Max_limit_mb = 10
	rtf.Sampling = false
	rtf.Sampling_percentage = 0
	rtf.Staggering = false

	return rtf
}

// FormatToHCL TODO Doc Comment
func (rtf *ResourceTestFixtureSolrcloud) FormatToHCL() {

	monitortype = strings.ToLowerCase("Solrcloud")

	result := fmt.Sprintf(`
	resource "sematext_monitor_%s" "%s" {
		name : %s,
		description : %s,
		billing_plan : %s,
		discount_code : %s,
		ignore_percentage : %i,
		max_events  : %i,
		max_limit_mb : %i,
		sampling : %b,
		sampling_percentage : %i,
		staggering : %b
	}
	`,
	monitortype,
	rtf.Name,
	rtf.Description,
	rtf.Plan,
	rtf.Discount_code,
	rtf.Ignore_percentage,
	rtf.Max_events,
	rtf.Max_limit_mb,
	rtf.Sampling,
	rtf.Sampling_percentage,
	rtf.Staggering,
	)

	return result
}


// testAccSematextMonitorSolrcloud_Basic tests resource creation.
func testAccSematextMonitorSolrcloudBasic(t *testing.T) {

	monitortype = strings.ToLowerCase("Solrcloud")
	rtf := (&ResourceTestFixtureSolrcloud{}).HydrateBasic()
	fixture := rtf.FormatToHCL()
	statepath := fmt.Sprintf("sematext_monitor_%s.%s", monitortype, rtf.Name)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSematextMonitorSolrcloud_ConfirmDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fixture,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorSolrcloud_CheckConsistency(name),
					resource.TestCheckResourceAttr(statepath,"name",rtf.Name),
					resource.TestCheckResourceAttr(statepath,"description",rtf.Description),
					resource.TestCheckResourceAttr(statepath,"billing_plan", rtf.Plan),
					resource.TestCheckResourceAttr(statepath,"discount_code", rtf.Discount_code),
					resource.TestCheckResourceAttr(statepath,"ignore_percentage", rtf.Ignore_percentage,
					resource.TestCheckResourceAttr(statepath,"max_events", rtf.Max_events),
					resource.TestCheckResourceAttr(statepath,"max_limit_mb", rtf.Max_limit_mb),
					resource.TestCheckResourceAttr(statepath,"sampling", rtf.Sampling),
					resource.TestCheckResourceAttr(statepath,"sampling_percentage", rtf.Sampling_percentage),
					resource.TestCheckResourceAttr(statepath,"staggering", rtf.Staggering),
				),
			},
		},
	})
}

// testAccSematextMonitorSolrcloud_Update tests for resource updates.
func testAccSematextMonitorSolrcloudUpdate(t *testing.T) {

	monitortype = strings.ToLowerCase("Solrcloud")

	rtf1 := (&ResourceTestFixtureSolrcloud{}).HydrateBasic()
	statepath := fmt.Sprintf("sematext_monitor_%s.%s", monitortype, rtf1.Name)

	rtf2 := rtf1
	rtf2.Description = "TESTING : SematextMonitorSolrcloud_Basic : Update"
	fixture1 := rtf1.FormatToHCL()
	fixture2 := rtf2.FormatToHCL()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSematextMonitorSolrcloud_ConfirmDestroyed,
		Steps: []resource.TestStep{
			{
				Config: fixture1,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorSolrcloudCheckConsistency(rtf1.Name),
					resource.TestCheckResourceAttr(statepath,"name",rtf1.Name),
					resource.TestCheckResourceAttr(statepath,"description",rtf1.Description),
					resource.TestCheckResourceAttr(statepath,"billing_plan", rtf1.Plan),
					resource.TestCheckResourceAttr(statepath,"discount_code", rtf1.Discount_code),
					resource.TestCheckResourceAttr(statepath,"ignore_percentage", rtf1.Ignore_percentage,
					resource.TestCheckResourceAttr(statepath,"max_events", rtf1.Max_events),
					resource.TestCheckResourceAttr(statepath,"max_limit_mb", rtf1.Max_limit_mb),
					resource.TestCheckResourceAttr(statepath,"sampling", rtf1.Sampling),
					resource.TestCheckResourceAttr(statepath,"sampling_percentage", rtf1.Sampling_percentage),
					resource.TestCheckResourceAttr(statepath,"staggering", rtf1.Staggering),
				),
			},
			{
				Config: fixture2,
				Check: resource.ComposeTestCheckFunc(
					testAccSematextMonitorSolrcloud_CheckConsistency(rtf2.name),
					resource.TestCheckResourceAttr(statepath,"name",rtf2.Name),
					resource.TestCheckResourceAttr(statepath,"description",rtf2.Description),
					resource.TestCheckResourceAttr(statepath,"billing_plan", rtf2.Plan),
					resource.TestCheckResourceAttr(statepath,"discount_code", rtf2.Discount_code),
					resource.TestCheckResourceAttr(statepath,"ignore_percentage", rtf2.Ignore_percentage,
					resource.TestCheckResourceAttr(statepath,"max_events", rtf2.Max_events),
					resource.TestCheckResourceAttr(statepath,"max_limit_mb", rtf2.Max_limit_mb),
					resource.TestCheckResourceAttr(statepath,"sampling", rtf2.Sampling),
					resource.TestCheckResourceAttr(statepath,"sampling_percentage", rtf2.Sampling_percentage),
					resource.TestCheckResourceAttr(statepath,"staggering", rtf2.Staggering),
				),
			},
		},
	})
}


// testAccSematextMonitorSolrcloud_CheckConsistency checks the App ID exists in both state and API.
func testAccSematextMonitorSolrcloudCheckConsistency(name string) resource.TestCheckFunc {
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

// testAccSematextMonitorSolrcloud_ConfirmDestroyed -  check is destroyed in API
func testAccSematextMonitorSolrcloudConfirmDestroyed(s *terraform.State) error {

	app := new(api.App)
	err := new(error)
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "sematext_monitor_Solrcloud" {
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
