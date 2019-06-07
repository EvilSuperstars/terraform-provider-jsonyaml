Terraform `jsonyaml` Provider
==============================

## This Terraform provider has been obsoleted by the new Terraform 0.12.2 [`yamldecode`](https://github.com/hashicorp/terraform/pull/21459) interpolation function.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/EvilSuperstars/terraform-provider-jsonyaml`

```sh
$ mkdir -p $GOPATH/src/github.com/EvilSuperstars; cd $GOPATH/src/github.com/EvilSuperstars
$ git clone git@github.com:EvilSuperstars/terraform-provider-jsonyaml
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/EvilSuperstars/terraform-provider-jsonyaml
$ make build
```

Run acceptance tests

```sh
$ cd $GOPATH/src/github.com/EvilSuperstars/terraform-provider-jsonyaml
$ make testacc TEST=./jsonyaml/ TESTARGS='-run=TestDataSource'
```

Using The Provider
------------------

See the [documentation](using.md) to get started using the [jsonyaml](https://github.com/EvilSuperstars/terraform-provider-jsonyaml) provider.
