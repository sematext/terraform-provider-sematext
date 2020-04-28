Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.svg)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-sematext`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-sematext.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-sematext
$ make build
# or if you're on a mac:
$ gnumake build
```

Using the provider
----------------------

Detailed documentation for the Sematext provider can be found [here](https://www.terraform.io/docs/providers/sematext/index.html).


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


In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Acceptance test prerequisites
-----------------------------
In order to successfully run the full suite of acceptance tests, you will need to have the following:


### Sematext Cloud Access Token

You will need get your [Sematext Cloud API Access Token](https://apps.sematext.com/ui/account/api) for
testing. 

Once you have your token the following should be exported in your environment:

export TF_ACC=1
export SEMATEXT_REGION="US"
export SEMATEXT_API_TOKEN="&lt;your api roken&gt;"
export AWS_ACCESS_KEY_ID="&lt;your aws access key&gt;"
export AWS_SECRET_ACCESS_KEY="&lt;bAZztV1M+dekhod5AIQ/K5GCIOPeJnY2fckrsjV+&gt;"
export AWS_REGION="&lt;us-east-1&gt;"
