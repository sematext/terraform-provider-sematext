# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**


- Website: https://www.terraform.io
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.svg)](https://gitter.im/hashicorp-terraform/Lobby)


Requirements
------------

- [Terraform v0.13+](https://www.terraform.io/downloads.html)
- [Go](https://golang.org/doc/install) v1.14 (to build the provider plugin)



Installation
------------

Build the provider and put it in Terraform's third-party providers directory in `~/.terraform.d/plugins`:

```bash
go get github.com/sematext/terraform-provider-sematext
mkdir -p ~/.terraform.d/plugins
go build -o ~/.terraform.d/plugins/terraform-provider-sematext github.com/sematext/terraform-provider-sematext
```

I recommend using [Go modules](https://github.com/golang/go/wiki/Modules) to ensure
using the same version in development and production.

Configuration
------------

In your Terraform configuration:

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

resource "sematext_monitor_elasticsearch" "mymonitor" {
    name = "Node.js Monitor Example"
    billing_plan_id = 12
    apptoken_name = "my apptoken name"
    apptoken_create_missing = true
}
```

Refer to [documentation](docs/index.md) for more detail.
Refer to [plan guide](docs/guides/plans.md) for more detail.


Using the Provider
----------------------

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

For either installation method, documentation about the provider specific configuration options can be found on the [provider's website](https://www.terraform.io/docs/providers/aws/index.html).


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-sematext
...
```

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-sematext`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-sematext.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-sematext
$ make build
```



Testing the Provider
---------------------------

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Acceptance test prerequisites
-----------------------------
In order to successfully run the full suite of acceptance tests, you will need to have the following:


### Sematext Cloud Access Token

You will need get your [Sematext Cloud API Key](https://apps.sematext.com/ui/account/api) for
testing.

Once you have your token the following should be exported in your environment:

````sh
export TF_ACC=1
export SEMATEXT_REGION="US"
export SEMATEXT_API_KEY="<your api token>"
export AWS_ACCESS_KEY_ID="<your aws access key>"
export AWS_SECRET_ACCESS_KEY="<aws secret key>"
export AWS_REGION="<us-east-1>"
````
