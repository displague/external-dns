/*
Copyright 2017 The Kubernetes Authors.

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

package provider

import (
	"fmt"
	"os"
	"testing"
)

func TestNewLinodeProvider(t *testing.T) {
	_ = os.Setenv("LINODE_TOKEN", "xxxxxxxxxxxxxxxxx")
	_, err := NewLinodeProvider(NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	_ = os.Unsetenv("LINODE_TOKEN")
	_, err = NewLinodeProvider(NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestNewLinodeFetchDomain(t *testing.T) {
	_ = os.Setenv("LINODE_TOKEN", "xxxxxxx")
	p, err := NewLinodeProvider(NewDomainFilter([]string{"ordcapital.com"}), true)
	// provider, err := NewLinodeProvider(NewDomainFilter([]string{}), true)

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	records, err := p.fetchRecords(1028256)

	for _, record := range records {
		fmt.Printf("Record Retrieved: %+v\n", record)
	}

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	// zones, err := p.fetchZones()

	// for _, zone := range zones {
	// 	fmt.Printf("Zone Retrieved: %+v\n", zone)
	// }
}
