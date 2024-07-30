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
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/corneliusdavid97/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/multi-match-query.asciidoc#L275>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "dis_max": {
//       "queries": [
//         { "match_phrase_prefix": { "subject": "quick brown f" }},
//         { "match_phrase_prefix": { "message": "quick brown f" }}
//       ]
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_multi_match_query_33f148e3d8676de6cc52f58749898a13(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:33f148e3d8676de6cc52f58749898a13[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "dis_max": {
		      "queries": [
		        {
		          "match_phrase_prefix": {
		            "subject": "quick brown f"
		          }
		        },
		        {
		          "match_phrase_prefix": {
		            "message": "quick brown f"
		          }
		        }
		      ]
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:33f148e3d8676de6cc52f58749898a13[]
}
