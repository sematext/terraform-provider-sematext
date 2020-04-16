module github.com/sematext/terraform-provider-sematext

go 1.13

require (
	github.com/bflad/tfproviderdocs v0.5.0 // indirect
	github.com/bflad/tfproviderlint v0.11.0 // indirect
	github.com/client9/misspell v0.3.4
	github.com/davecgh/go-spew v1.1.1
	github.com/golangci/golangci-lint v1.22.2
	github.com/google/uuid v1.1.1
	github.com/hashicorp/terraform v0.12.23
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/sematext/sematext-api-client/api v0.0.0
	github.com/thoas/go-funk v0.5.0
)

replace github.com/sematext/sematext-api-client/stcloud => /home/euan/go/src/github.com/sematext/sematext-api-client/
