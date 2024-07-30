// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/cdb84fa39f1401846dab6e1c76781fb3090527ed

package putautoscalingpolicy

import (
	"github.com/corneliusdavid97/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putautoscalingpolicy
//
// https://github.com/elastic/elasticsearch-specification/blob/cdb84fa39f1401846dab6e1c76781fb3090527ed/specification/autoscaling/put_autoscaling_policy/PutAutoscalingPolicyRequest.ts#L24-L35
type Request = types.AutoscalingPolicy

// NewRequest returns a Request
func NewRequest() *Request {
	r := types.NewAutoscalingPolicy()

	return r
}
