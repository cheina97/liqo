// Copyright 2019-2022 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LiqoHelm", func() {
	var (
		lhc                         *LiqoHelmClient
		expectedClusterLabels       map[string]interface{}
		expectedClusterLabelsString map[string]string
		clusterLabels               map[string]string
		err                         error
	)
	BeforeEach(func() {
		expectedClusterLabels = map[string]interface{}{"label1": "value1", "label2": "value2", "label3": "value3"}
		expectedClusterLabelsString = map[string]string{"label1": "value1", "label2": "value2", "label3": "value3"}
	})
	Context("Creating a new LiqoHelmClient", func() {
		BeforeEach(func() {
			lhc = &LiqoHelmClient{
				getValues: func() (map[string]interface{}, error) {
					return map[string]interface{}{
						"discovery": map[string]interface{}{
							"config": map[string]interface{}{
								"clusterLabels": expectedClusterLabels,
							},
						},
					}, nil
				},
			}
		})
		JustBeforeEach(func() {
			clusterLabels, err = lhc.GetClusterLabels()
		})
		It("should return cluster labels", func() {
			Expect(expectedClusterLabelsString).To(Equal(clusterLabels))
			Expect(err).To(BeNil())
		})
	})
})
