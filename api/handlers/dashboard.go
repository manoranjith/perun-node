// Copyright (c) 2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/hyperledger-labs/perun-node
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import "fmt"

// Dashboard is a boolean to enable/disable dashboard messages for funding and
// watching handlers.
//
// For now, this is initialized directly at package level. To make this
// configurable, a flag can be added to the "perunnode run" command and the
// value can be passed to this package via the  grpc or tcp pacakge.
var Dashboard = true

func printf(format string, a ...any) {
	if Dashboard {
		fmt.Printf(format, a...)
	}
}
