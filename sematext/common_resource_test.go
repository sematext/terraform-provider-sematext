package sematext

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

type AppTokenType struct {
	Name          string
	CreateMissing string
}

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
	AppToken     AppTokenType
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
	AppToken          AppTokenType
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
	rtf.PlanID = stcloud.AssignPlanID(rtf.AppType)

	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.AppToken = AppTokenType{
		Name:          "Test Token",
		CreateMissing: "true",
	}

}

func (rtf *ResourceTestFixtureAWS) hydrate(resourceType string, appType string) {

	rndID := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rtf.ResourceType = resourceType
	rtf.ResourceName = strings.ToLower(fmt.Sprintf("test_%s", rndID))
	rtf.AppType = appType
	rtf.Name = rtf.ResourceName
	rtf.StatePath = rtf.ResourceType + "." + rtf.ResourceName
	rtf.PlanID = stcloud.AssignPlanID(rtf.AppType)
	switch appType {
	case "Logsene":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	case "mobile-logs":
		rtf.DiscountCode = stcloud.TestDiscountCodeLogs
	default:
		rtf.DiscountCode = stcloud.TestDiscountCodeMetrics
	}

	rtf.AppToken = AppTokenType{
		Name:          "Test Token",
		CreateMissing: "true",
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
		apptoken_name = "%s"
		apptoken_create_missing = %s
	}
	`,
		rtf.ResourceType,
		rtf.ResourceName,
		rtf.Name,
		rtf.PlanID,
		rtf.DiscountCode,
		rtf.AppToken.Name,
		rtf.AppToken.CreateMissing,
	)

}

func (rtf *ResourceTestFixtureAWS) toHCL() string {

	return fmt.Sprintf(`
	resource "%s" "%s" {
		name = "%s"
		billing_plan_id = %d
		discount_code = "%s"
		apptoken_name = "%s"
		apptoken_create_missing = %s
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
		rtf.AppToken.Name,
		rtf.AppToken.CreateMissing,
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
						resource.TestCheckResourceAttr(rtf.StatePath, "name", rtf.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "billing_plan_id", strconv.Itoa(rtf.PlanID)),
						resource.TestCheckResourceAttr(rtf.StatePath, "discount_code", rtf.DiscountCode),
						resource.TestCheckResourceAttr(rtf.StatePath, "apptoken_name", rtf.AppToken.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "apptoken_create_missing", rtf.AppToken.CreateMissing),
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

		fmt.Println(fixture)

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
						resource.TestCheckResourceAttr(rtf.StatePath, "apptoken_name", rtf.AppToken.Name),
						resource.TestCheckResourceAttr(rtf.StatePath, "apptoken_create_missing", rtf.AppToken.CreateMissing),
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
						resource.TestCheckResourceAttr(rtf0.StatePath, "name", rtf0.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "billing_plan_id", strconv.Itoa(rtf0.PlanID)),
						resource.TestCheckResourceAttr(rtf0.StatePath, "discount_code", rtf0.DiscountCode),
						resource.TestCheckResourceAttr(rtf0.StatePath, "apptoken_name", rtf0.AppToken.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "apptoken_create_missing", rtf0.AppToken.CreateMissing),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationAWS(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
						resource.TestCheckResourceAttr(rtf1.StatePath, "apptoken_name", rtf1.AppToken.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "apptoken_create_missing", rtf1.AppToken.CreateMissing),
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
						resource.TestCheckResourceAttr(rtf0.StatePath, "apptoken_name", rtf0.AppToken.Name),
						resource.TestCheckResourceAttr(rtf0.StatePath, "apptoken_create_missing", rtf0.AppToken.CreateMissing),
					),
				},
				{
					Config: fixture1,
					Check: resource.ComposeTestCheckFunc(
						ConfirmMonitorCreationDefault(rtf1),
						resource.TestCheckResourceAttr(rtf1.StatePath, "name", rtf1.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "billing_plan_id", strconv.Itoa(rtf1.PlanID)),
						resource.TestCheckResourceAttr(rtf1.StatePath, "discount_code", rtf1.DiscountCode),
						resource.TestCheckResourceAttr(rtf1.StatePath, "apptoken_name", rtf1.AppToken.Name),
						resource.TestCheckResourceAttr(rtf1.StatePath, "apptoken_create_missing", rtf1.AppToken.CreateMissing),
					),
				},
			},
		})
	}

}

// ConfirmMonitorCreationDefault checks the App ID exists in both state and stcloud.
func ConfirmMonitorCreationDefault(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorCreationDefault func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.App
		var rs *terraform.ResourceState
		var genericAPIResponse stcloud.GenericAPIResponse

		client := testAccProvider.Meta().(*stcloud.APIClient)

		/*
			fmt.Println("---------------------------------------")
			fmt.Println("rtf.StatePath")
			fmt.Println(rtf.StatePath)
			spew.Dump(s.RootModule().Resources)
			fmt.Println("---------------------------------------")
		*/
		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s, %s", rtf.StatePath, err)
		}

		if app, err = genericAPIResponse.ExtractApp(); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmMonitorCreationAWS checks the App ID exists in both state and stcloud.
func ConfirmMonitorCreationAWS(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorCreationAWS func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var found bool
		var err error
		var app *stcloud.App
		var rs *terraform.ResourceState
		var genericAPIResponse stcloud.GenericAPIResponse

		client := testAccProvider.Meta().(*stcloud.APIClient)

		if rs, found = s.RootModule().Resources[rtf.StatePath]; !found {
			return fmt.Errorf("ConfirmMonitorCreation : Resource not found in state: %s %s", rtf.ResourceType, rtf.ResourceName)
		}

		if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
			return err
		}

		if genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id); err != nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s, %s", rtf.StatePath, err)
		}

		if app, err = genericAPIResponse.ExtractApp(); err != nil {
			return err
		}

		if app == nil {
			return fmt.Errorf("ConfirmMonitorCreation : Error in checking monitor %s", rtf.StatePath)
		}

		return nil
	}
}

// ConfirmMonitorDestructionDefault checks the App ID has been removed from state and the API has marked the app as DELETED.
func ConfirmMonitorDestructionDefault(rtf ResourceTestFixtureDefault) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorDestructionDefault func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.App
		var genericAPIResponse stcloud.GenericAPIResponse

		client := testAccProvider.Meta().(*stcloud.APIClient)

		//spew.Dump(s.RootModule().Resources)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO Consider shift to template and make check tighter after MVP
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			time.Sleep(time.Second * 4) // TODO - workaround for cache latency on SC API, no longer required, confirm and obsolete after MVP.

			if genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Failed to pull app in checking monitor %s, %s", rtf.StatePath, err)
			}

			if app, err = genericAPIResponse.ExtractApp(); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Missing app in genericAPIResponse checking monitor %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmMonitorDestructionDefault : Unexpected status in checking monitor %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}

// ConfirmMonitorDestructionAWS checks the App ID exists in both state and API and is marked as DELETED.
func ConfirmMonitorDestructionAWS(rtf ResourceTestFixtureAWS) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		fmt.Println("---------------------------------------")
		fmt.Println("ConfirmMonitorDestructionDefaultAWS func Called")
		fmt.Println("---------------------------------------")

		var id int64
		var err error
		var rs *terraform.ResourceState
		var app *stcloud.App
		var genericAPIResponse stcloud.GenericAPIResponse

		client := testAccProvider.Meta().(*stcloud.APIClient)

		for _, rs = range s.RootModule().Resources {

			if !strings.HasPrefix(rs.Type, "sematext_") { // TODO Consider shift to template and make check tighter after MVP
				continue
			}
			if id, err = strconv.ParseInt(rs.Primary.ID, 10, 64); err != nil {
				return err
			}

			if genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id); err != nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Failed to pull app in checking monitor %s, %s", rtf.StatePath, err)
			}

			if app, err = genericAPIResponse.ExtractApp(); err != nil {
				return err
			}

			if app == nil {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Missing app in genericAPIResponse checking monitor %s", rtf.StatePath)
			}

			if app.Status != "DELETED" {
				return fmt.Errorf("ConfirmMonitorDestructionDefaultAWS : Unexpected status in checking monitor %s : Expected DELETED, got %s ", rtf.StatePath, app.Status)
			}

		}

		return nil
	}
}
