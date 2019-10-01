// Code generated from specification version 7.3.1 (a92f36162e5): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
//
type API struct {
	Bulk                                Bulk
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Bulk:                                newBulkFunc(t),
	}
}
