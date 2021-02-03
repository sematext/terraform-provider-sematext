# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**

# Sematext Nginx Resource

Creates a monitoring application within [Sematext Cloud](https://sematext.com/cloud/).
Refer to [Refer to Sematext Provider for authentication detail](../index.md)

## Example Usage

```hcl
terraform {
  required_providers {
    sematext = {
      source = "sematext/sematext"
      version = ">=0.1.10"
    }
  }
}

provider "sematext" {
  sematext_region = "US"
}

resource "sematext_monitor_nginx" "mymonitor" {
  name = "my monitor name"
  billing_plan_id = <[plan id](../guides/plans.md)>
  apptoken {
    name "my apptoken name"
    create_missing true
  }
}
```

## Argument Reference

* `name` - List attributes that this resource exports;
* `billing_plan_id` - List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);
* `apptoken.name` - Refer note below;


## App-tokens

Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation.
This is so that downstream resources can be grouped, with the instances within each group using the same app-token.

Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud.
When provisioning, downstream resources should set the app-token id in their configuration.

To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource)
to handle the provisioning.

## Outputs

* `id` - The app id. Provided on creation and used in terraform destroy operations.
* `apptoken.id` - On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.
