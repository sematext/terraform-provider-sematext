# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**

# Sematext SolrCloud Resource

Creates a monitoring application within [Sematext Cloud](https://sematext.com/cloud/).
Refer to [Refer to Sematext Provider for authentication detail](../index.md)

## Example Usage

```hcl
terraform {
  required_providers {
    sematext = {
      source = "sematext/sematext"
      version = ">=0.1.3"
    }
  }
}

provider "sematext" {
  sematext_region = "US"
}

resource "sematext_monitor_solrcloud" "mymonitor" {
  name = "my monitor name"
  billing_plan_id = <[plan id](../guides/plans.md)>
}
```

## Argument Reference

* `name` - List attributes that this resource exports.
* `billing_plan_id` - List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md)

