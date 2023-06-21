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

package credential

import (
	"os"
	"reflect"
	"testing"
)

func TestNewStoreError(t *testing.T) {
	filename := "testfile.txt"
	_, err := os.Create(filename)
	if err != nil {
		t.Errorf("error while creating test file %v", err)
	}
	defer func() {
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("error while removing test file %v", err)
			t.Fatal(err)
		}
	}()

	err = os.Chmod(filename, 000)
	if err != nil {
		t.Errorf("error: cannot change file permissions: %v", err)
	}

	var config string = filename
	credStore, err := NewStore(config)

	if !reflect.DeepEqual(credStore, nil) {
		t.Errorf("Expected NewStore to return nil but actually returned %v ", credStore)
	}
	if reflect.DeepEqual(err, nil) {
		t.Error("Expected Error to be not nil but actually returned nil ")
	}

}
