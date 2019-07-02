// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
	"time"
)

func newMLOpenJobFunc(t Transport) MLOpenJob {
	return func(job_id string, o ...func(*MLOpenJobRequest)) (*Response, error) {
		var r = MLOpenJobRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html.
//
type MLOpenJob func(job_id string, o ...func(*MLOpenJobRequest)) (*Response, error)

// MLOpenJobRequest configures the Ml  Open Job API request.
//
type MLOpenJobRequest struct {
	IgnoreDowntime *bool
	JobID          string
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLOpenJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("_open"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("_open")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f MLOpenJob) WithContext(v context.Context) func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.ctx = v
	}
}

// WithIgnoreDowntime - controls if gaps in data are treated as anomalous or as a maintenance window after a job re-start.
//
func (f MLOpenJob) WithIgnoreDowntime(v bool) func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.IgnoreDowntime = &v
	}
}

// WithTimeout - controls the time to wait until a job has opened. default to 30 minutes.
//
func (f MLOpenJob) WithTimeout(v time.Duration) func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLOpenJob) WithPretty() func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLOpenJob) WithHuman() func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLOpenJob) WithErrorTrace() func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLOpenJob) WithFilterPath(v ...string) func(*MLOpenJobRequest) {
	return func(r *MLOpenJobRequest) {
		r.FilterPath = v
	}
}
