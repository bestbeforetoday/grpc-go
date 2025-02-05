/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package xdsclient

import (
	"google.golang.org/grpc/xds/internal/xdsclient/xdsresource"
)

// DumpLDS returns the status and contents of LDS.
func (c *clientImpl) DumpLDS() (string, map[string]xdsresource.UpdateWithMD) {
	return c.pubsub.Dump(xdsresource.ListenerResource)
}

// DumpRDS returns the status and contents of RDS.
func (c *clientImpl) DumpRDS() (string, map[string]xdsresource.UpdateWithMD) {
	return c.pubsub.Dump(xdsresource.RouteConfigResource)
}

// DumpCDS returns the status and contents of CDS.
func (c *clientImpl) DumpCDS() (string, map[string]xdsresource.UpdateWithMD) {
	return c.pubsub.Dump(xdsresource.ClusterResource)
}

// DumpEDS returns the status and contents of EDS.
func (c *clientImpl) DumpEDS() (string, map[string]xdsresource.UpdateWithMD) {
	return c.pubsub.Dump(xdsresource.EndpointsResource)
}
