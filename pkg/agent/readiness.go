/*
copyright 2023 the kubernetes authors.

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package agent

// ReadinessManager supports checking if the agent is ready.
type ReadinessManager interface {
	// Ready returns true the proxy server is ready.
	Ready() bool
}

var _ ReadinessManager = &ClientSet{}

func (cs *ClientSet) Ready() bool {
	// Returns true if the agent is connected to at least one server.
	return cs.HealthyClientsCount() > 0
}
