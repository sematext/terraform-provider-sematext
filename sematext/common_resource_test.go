package sematext

// TODO - Expand Resource test cases to full checks.

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/api"
)

type ResourceTestFixture interface {
	hydrate(resourceType string, appType string)
	toHCL() string
}

// ResourceTestFixture a common test fixture representing most resources
type ResourceTestFixtureDefault struct {
	ResourceType string
	ResourceName string
	AppType      string
	Name         string
	StatePath    string
	PlanID       int
	DiscountCode string
}

// ResourceTestFixtureAWS a test fixture representing a resource - AWS EBS, AWS EC2, AWS ELB
type ResourceTestFixtureAWS struct {
	ResourceType      string
	ResourceName      string
	AppType           string
	Name              string
	StatePath         string
	PlanID            int
	DiscountCode      string
	AwsRegion         string
	AwsSecretKey      string
	AwsAccessKey      string
	AwsFetchFrequency string
}

func (rtf *ResourceTestFixtureDefault) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AppType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = api.AssignPlanID(rtf.AppType)
	if appType != "Logsene" {
		rtf.DiscountCode = api.TestDiscountCodeMetrics
	}

}

func (rtf *ResourceTestFixtureAWS) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AppType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = api.AssignPlanID(rtf.AppType)
	if appType != "Logsene" {
		rtf.DiscountCode = api.TestDiscountCodeMetrics
	}
	rtf.AwsRegion = os.Getenv("AWS_REGION")
	rtf.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	rtf.AwsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	rtf.AwsFetchFrequency = "FIVE_MINUTES"
}

func (rtf *ResourceTestFixtureDefault) toHCL() string {

	return fmt.Sprintf(`
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

}

func (rtf *ResourceTestFixtureAWS) toHCL() string {

	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		billing_plan_id = %d
		discount_code = "%s"
		aws_region = "%s"
		aws_secret_key = "%s"
		aws_access_key = "%s"
		aws_fetch_frequency = "%s"
	}
	`,
		rtf.ResourceType,
		rtf.ResourceName,
		rtf.Name,
		rtf.PlanID,
		rtf.DiscountCode,
		rtf.AwsRegion,
		rtf.AwsSecretKey,
		rtf.AwsAccessKey,
		rtf.AwsFetchFrequency,
	)

}

// CommonMonitorBasicTest is a common test of resource creation.
func CommonMonitorBasicTest(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorBasicTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureAWS{}
		rtf.hydrate(resourceType, appType)
		//workaround := fmt.Sprintf("%s %s", appType, rtf.Name) // WORKAROUND

		fixture := rtf.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: ConfirmMonitorDestructionAWS(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationAWS(rtf),
						//resource.TestCheckResourceAttr(rtf.StatePath, "name", workaround),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})

	default:

		fmt.Println("---------------------------------------")
		fmt.Println("default identified")
		fmt.Println("---------------------------------------")

		rtf := ResourceTestFixtureDefault{}
		rtf.hydrate(resourceType, appType)
		fixture := rtf.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: ConfirmMonitorDestructionDefault(rtf),
			Steps: []resource.TestStep{
				{
					Config: fixture,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationDefault(rtf),
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
					),
				},
			},
		})
	}
}

// CommonMonitorUpdateTest tests for resource updates.
func CommonMonitorUpdateTest(t *testing.T, resourceType string, appType string) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorUpdateTest Called")
	fmt.Println("---------------------------------------")
	fmt.Println("appType")
	fmt.Println(appType)
	fmt.Println("---------------------------------------")

	switch appType {
	case "AWS EBS", "AWS EC2", "AWS ELB":

		fmt.Println("---------------------------------------")
		fmt.Println("AWS identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureAWS{}
		rtf0.hydrate(resourceType, appType)
		//workaround := fmt.Sprintf("%s %s", appType, rtf0.Name) // WORKAROUND
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0

		fixture1 := rtf1.toHCL()

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: ConfirmMonitorDestructionAWS(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationAWS(rtf0),
						//resource.TestCheckResourceAttr(rtf0.StatePath, "name", workaround), //TODO-workaround.
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationAWS(rtf1),
						//resource.TestCheckResourceAttr(rtf1.StatePath, "name", workaround), //TODO-workaround.
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})

	default:

		fmt.Println("---------------------------------------")
		fmt.Println("default identified")
		fmt.Println("---------------------------------------")

		rtf0 := ResourceTestFixtureDefault{}
		rtf0.hydrate(resourceType, appType)
		fixture0 := rtf0.toHCL()

		rtf1 := rtf0
		fixture1 := rtf1.toHCL()
		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: ConfirmMonitorDestructionDefault(rtf0),
			Steps: []resource.TestStep{
				{
					Config: fixture0,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationDefault(rtf0),
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationDefault(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
					),
				},
			},
		})
	}

}

// ConfirmMonitorCreationDefault checks the App ID exists in both state and API.
func ConfirmMonitorCreationDefault(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorCreationDefault func Called")
		fmt.Println("---------------------------------------")

		var id int
		var found bool
		var err error
		var app *api.App
		var rs *terraform.ResourceState

		client := testAccProvider.Meta().(*api.Client)

		fmt.Println("---------------------------------------")
		fmt.Println("rtf.StatePath")
		fmt.Println(rtf.StatePath)
		spew.Dump(s.RootModule().Resources)
		fmt.Println("---------------------------------------")

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			fmt.Println("HERE A")
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}
		fmt.Println("HERE B")

		if id, err = strconv.Atoi(rs.Primary.ID); err != nil {
			fmt.Println("HERE C")

			return err
		}

		fmt.Println("HERE D")

		fmt.Println("Loading App " + string(id))
		if app, err = (&api.App{}).Load(id, client); err != nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s, %s", rtf.StatePath, err) //TODO Check return value fs function sig
		}
		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmMonitorCreationAWS checks the App ID exists in both state and API.
func ConfirmMonitorCreationAWS(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorCreationAWS func Called")
		fmt.Println("---------------------------------------")

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

// TODO - shift this into the temaplate
// ConfirmMonitorDestructionDefault checks the App ID has been removed from state and the API has marked the app as DISABLED.
func ConfirmMonitorDestructionDefault(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorDestructionDefault func Called")
		fmt.Println("---------------------------------------")

		var id int
		var retired bool
		var err error
		var rs *terraform.ResourceState
		client := testAccProvider.Meta().(*api.Client)

		fmt.Println("HERE 1")

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO shift to template and make check more explicit after MVP
				continue
			}

			fmt.Println("HERE 2")

			if id, err = strconv.Atoi(rs.Primary.ID); err != nil {
				fmt.Println("HERE 2 " + string(id))
				retired, err = (&api.App{}).Retired(id, client)
				if !retired {
					fmt.Println("HERE 3")
					return fmt.Errorf("ConfirmMonitorDestructionDefault : Error in checking monitor %s : %s", rtf.StatePath, err)
				}
			}
			if err != nil {
				fmt.Println("HERE 4")
				break
			}
		}

		fmt.Println("EXIT ")

		return err
	}
}

// ConfirmMonitorDestructionAWS checks the App ID exists in both state and API and is marked as DISABLED.
func ConfirmMonitorDestructionAWS(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorDestructionAWS func Called")
		fmt.Println("---------------------------------------")

		var id int
		var retired bool
		var err error
		var rs *terraform.ResourceState
		client := testAccProvider.Meta().(*api.Client)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO shift to template and make check more explicit after MVP
				continue
			}

			if id, err = strconv.Atoi(rs.Primary.ID); err != nil {
				retired, err = (&api.App{}).Retired(id, client)
				if !retired {
					return fmt.Errorf("ConfirmMonitorDestructionDefault : Error in checking monitor %s : %s", rtf.StatePath, err)
				}
			}
			if err != nil {
				break
			}
		}

		return err
	}
}
