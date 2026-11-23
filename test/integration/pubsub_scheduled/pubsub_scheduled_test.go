// Copyright 2022 Google LLC
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

package pubsub_scheduled

import (
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestPubsubScheduled(t *testing.T) {
	bpt := tft.NewTFBlueprintTest(t)

	bpt.DefineVerify(func(assert *assert.Assertions) {
		bpt.DefaultVerify(assert)

		projectId := bpt.GetStringOutput("project_id")
		name := bpt.GetStringOutput("name")

		op := gcloud.Runf(t, "beta scheduler jobs describe %s --location us-central1 --project=%s", name, projectId)
		assert.Contains(op.Get("state").String(), "ENABLED", "should be enabled")

		functionName := bpt.GetStringOutput("function_name")

		fn := gcloud.Runf(t, "functions describe %s --gen2 --region us-central1 --project=%s", functionName, projectId)
		assert.Equal("ACTIVE", fn.Get("state").String(), "function should be active")
		assert.Equal("GEN_2", fn.Get("environment").String(), "function should be 2nd gen")
		assert.Equal("google.cloud.pubsub.topic.v1.messagePublished", fn.Get("eventTrigger.eventType").String(), "function should be triggered by Pub/Sub")
	})

	bpt.Test()
}
