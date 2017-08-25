// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cassandrarpc

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/db/sql/cassandra"
	"github.com/ligato/cn-infra/flavors/rpc"
	"github.com/namsral/flag"
)

//init defines cassandra flags // TODO switch to viper to avoid global configuration
func init() {
	flag.String("cassandra-config", "cassandra.conf",
		"Location of the Cassandra Client configuration file; also set via 'CASSANDRA_CONFIG' env variable.")
}

// FlavorCassandraRPC glues together FlavorRPC plugins with:
// - CASSANDRA (for using with API to interact with Cassandra database)
type FlavorCassandraRPC struct {
	rpc.FlavorRPC
	Cassandra cassandra.Plugin
}

// Inject sets object references
func (f *FlavorCassandraRPC) Inject() (allReadyInjected bool) {
	if !f.FlavorRPC.Inject() {
		return false
	}

	f.Cassandra.Deps.PluginInfraDeps = *f.InfraDeps("cassandra")

	return true
}

// Plugins combines all Plugins in flavor to the list
func (f *FlavorCassandraRPC) Plugins() []*core.NamedPlugin {
	f.Inject()
	return core.ListPluginsInFlavor(f)
}
