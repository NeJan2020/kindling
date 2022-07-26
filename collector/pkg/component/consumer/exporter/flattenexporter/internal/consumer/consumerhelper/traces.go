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

package consumerhelper

import (
	"context"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer"
)

// ConsumeTracesFunc is a helper function that is similar to ConsumeTraces.
type ConsumeTracesFunc func(ctx context.Context, request []byte) error

// ConsumeTraces calls f(ctx, ld).
func (f ConsumeTracesFunc) ConsumeTraces(ctx context.Context, request []byte) error {
	return f(ctx, request)
}

type baseTraces struct {
	*baseConsumer
	ConsumeTracesFunc
}

// NewTraces returns a consumer.Traces configured with the provided options.
func NewTraces(consume ConsumeTracesFunc, options ...Option) (consumer.Traces, error) {
	if consume == nil {
		return nil, errNilFunc
	}
	return &baseTraces{
		baseConsumer:      newBaseConsumer(options...),
		ConsumeTracesFunc: consume,
	}, nil
}
