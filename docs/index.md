# Sematext Provider

The Sematext provider is used to interact with Sematext Cloud related resources.



## Example Usage

```hcl
# Configure the Sematext Provider
provider "sematext" {
  sematext_region = "US"
}

# Create a monitoring application
resource "sematext_monitor_nodejs" "mymonitor" {
  name = "mymonitor"
  billing_plan_id = 6
  discount_code = "<discount code>"
}
```

## Argument Reference

The following arguments are supported:

* `sematext_region` - (Required) closest Sematext Cloud Region  "US" or "EU";


## Authentication

When the Sematext Terraform Provider connects to Sematext Cloud it authenticates with an access token.
You'll need to gain this via your login to your Sematext Cloud account.



## Enviropnment Variables

The following environment variables are required:

* export SEMATEXT_REGION="US"
* export SEMATEXT_API_TOKEN="&lt;Sematext-Cloud-Token&gt;"

If working with AWS Cloudwatch the the following environment vars should be set:

* export AWS_ACCESS_KEY_ID="&lt;access-key&gt;"
* export AWS_SECRET_ACCESS_KEY="&lt;secret-key&gt;"
* export AWS_REGION="&lt;aws-region&gt;"
