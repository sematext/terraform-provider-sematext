# Sematext AWS EBS Monitoring Resource

Creates a monitoring application within Sematext Cloud. 
Refer to [Refer to Sematext Provider for authentication detail](../index.md) 

## Example Usage

```hcl
# Configure the Sematext Provider
provider "sematext" {
  sematext_region = "US"
}

# Create a monitoring application
resource "sematext_monitor_awsebs" "mymonitor" {
  name = "my monitor name"
  billing_plan_id = 6
  discount_code = "<discount code>"
}
```

## Argument Reference

* `name` - List attributes that this resource exports.
* `billing_plan_id` - List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md) 
* `discount_code` - List attributes that this resource exports.
* `aws_access_key` - (optional) if not set then reads from env AWS_ACCESS_KEY_ID.
* `aws_secret_key` - (optionl) is not present set from env AWS_SECRET_ACCESS_KEY
* `aws_fetch_frequency` - (required) one of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES.
* `aws_region` - (optional) if not present withh set from env AWS_REGION.

