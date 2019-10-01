// Code generated from specification version 7.3.0 (de777fa3ccb): DO NOT EDIT

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
