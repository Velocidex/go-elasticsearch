# Modified Elastic Client for Use in Velociraptor.

Velociraptor only uses the Bulk upload API to implement the
[elastic_upload](https://docs.velociraptor.app/vql_reference/other/elastic_upload/)
plugin.

However, the elastic API is designed in such a way that by
instantiating an API object, every conceivable API function is added
to the binary as dead code.

See https://github.com/elastic/go-elasticsearch/blob/62f06fd1bf95ea28476f117632b04635396d8133/esapi/api._.go#L695

This results in about 16mb additional size for the Velociraptor binary
for just a single plugin which is unacceptable.

## Reducing binary size

This repository is a fork of the upstream
https://github.com/elastic/go-elasticsearch with some modifications:

1. All other APIs are removed but the Bulk upload API. This is the
   only one we need.

2. The go API client also checks that the server reports itself as
   pure elasticsearch. In practice, the opensearch server also works
   fine with the Bulk upload API so this un-necessary check is
   removed.

These changes result in a huge reduction in binary size. It now adds
about 200kb vs the 16mb of the upstream `go-elastic` client library.

## Maintaining

To make maintaining future versions of this library easy, the
magefile.go includes mutations to sync the codebase from the latest
upstream git repository.

Simply update the release branch information in the `magefile.go` and
run `make`
