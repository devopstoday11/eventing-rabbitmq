// +build e2e

/*
Copyright 2020 The Knative Authors

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

package rabbit_test

import (
	"testing"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// For our e2e testing, we want this linked first so that our
	// system namespace environment variable is defaulted prior to
	// logstream initialization.
	_ "knative.dev/eventing-rabbitmq/test/defaultsystem"

	"knative.dev/pkg/test/helpers"
	"knative.dev/pkg/test/logstream"
)

// TestSmokeBroker makes sure a Broker goes ready as a RabbitMQ Broker Class.
func TestSmokeBroker(t *testing.T) {
	t.Cleanup(logstream.Start(t))
	SmokeTestBrokerImpl(t, helpers.ObjectNameForTest(t))
}
