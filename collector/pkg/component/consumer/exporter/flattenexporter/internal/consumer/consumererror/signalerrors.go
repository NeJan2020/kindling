// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consumererror

import (
	"errors"
)

// Traces is an error that may carry associated Trace data for a subset of received data
// that failed to be processed or sent.
type Traces struct {
	error
	failed []byte
}

// NewTraces creates a Traces that can encapsulate received data that failed to be processed or sent.
func NewTraces(err error, failed []byte) error {
	return Traces{
		error:  err,
		failed: failed,
	}
}

// AsTraces finds the first error in err's chain that can be assigned to target. If such an error is found
// it is assigned to target and true is returned, otherwise false is returned.
func AsTraces(err error, target *Traces) bool {
	if err == nil {
		return false
	}
	return errors.As(err, target)
}

// GetTraces returns failed traces from the associated error.
func (err Traces) GetTraces() []byte {
	return err.failed
}

// Unwrap returns the wrapped error for functions Is and As in standard package errors.
func (err Traces) Unwrap() error {
	return err.error
}

// Metrics is an error that may carry associated Metrics data for a subset of received data
// that failed to be processed or sent.
type Metrics struct {
	error
	failed []byte
}

// NewMetrics creates a Metrics that can encapsulate received data that failed to be processed or sent.
func NewMetrics(err error, failed []byte) error {
	return Metrics{
		error:  err,
		failed: failed,
	}
}

// AsMetrics finds the first error in err's chain that can be assigned to target. If such an error is found
// it is assigned to target and true is returned, otherwise false is returned.
func AsMetrics(err error, target *Metrics) bool {
	if err == nil {
		return false
	}
	return errors.As(err, target)
}

// GetMetrics returns failed metrics from the associated error.
func (err Metrics) GetMetrics() []byte {
	return err.failed
}

// Unwrap returns the wrapped error for functions Is and As in standard package errors.
func (err Metrics) Unwrap() error {
	return err.error
}
