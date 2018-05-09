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

package grpc

import "google.golang.org/grpc"

// Server defines the API for getting grpc.Server instance that
// is useful for registering new GRPC services
type Server interface {
	// Server is a getter for accessing grpc.Server (of a GRPC plugin)
	//
	// Example usage:
	//
	//   protocgenerated.RegisterServiceXY(plugin.Deps.GRPC.Server(), &ServiceXYImplP{})
	//
	//   type Deps struct {
	//       GRPC grps.Server // inject plugin implementing RegisterHandler
	//       // other dependencies ...
	//   }
	GetServer() *grpc.Server

	// GetClientFromServer allows to get client instance from the server
	GetClientFromServer() Client

	// Disabled informs other plugins about availability
	IsDisabled() bool
}

// Client defines API for connectiong to external servers and receiving information
// about their IP and port
type Client interface {
	// Connect dials provided address and creates new client connection
	Connect(address string) (*grpc.ClientConn, error)

	// GetNotificationEndpoints returns a list of addresses defined in grpc.conf. Those
	// addresses are expected as GRPC listeners for statistics
	GetNotificationEndpoints() []string

	// Disabled informs other plugins about availability
	IsDisabled() bool
}
