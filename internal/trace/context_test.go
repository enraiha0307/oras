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
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_New_Logger(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		ctx     context.Context
		debug   bool
		verbose bool
	}{
		{
			name:    "test NewLogger when both debug and verbose are false",
			ctx:     ctx,
			debug:   false,
			verbose: false,
		},
		{
			name:    "test NewLogger when  debug is true and verbose is false",
			ctx:     ctx,
			debug:   true,
			verbose: false,
		},
		{
			name:    "test NewLogger when  debug is false and verbose is true",
			ctx:     ctx,
			debug:   false,
			verbose: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnedCtx, value := NewLogger(tt.ctx, tt.debug, tt.verbose)
			assert.NotNil(t, value)
			assert.NotNil(t, returnedCtx)
			value = Logger(returnedCtx)
			assert.NotNil(t, value)
		})
	}
}

func Test_Logger_ok_Not_True(t *testing.T) {
	var loggerKey contextKey = 1
	ctx := context.Background()
	logger := logrus.New()
	entry := logger.WithContext(ctx)
	newctx := context.WithValue(ctx, loggerKey, entry)
	value := Logger(newctx)
	assert.NotNil(t, value)
}
