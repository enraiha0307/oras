/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trace

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logheader_Empty(t *testing.T) {
	header := http.Header{}
	result := logHeader(header)
	assert.Equal(t, "   Empty header", result)

}

func Test_logheader(t *testing.T) {
	header := http.Header{}
	header.Add("Content", "application/json")
	header.Add("Authorization", "Basic")
	result := logHeader(header)
	assert.Contains(t, result, "****")

}
