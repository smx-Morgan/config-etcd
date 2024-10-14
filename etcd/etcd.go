// Copyright 2023 CloudWeGo Authors
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

package etcd

import (
	cwetcd "github.com/cloudwego-contrib/cwgo-pkg/config/etcd/etcd"
)

type Key = cwetcd.Key

type Client = cwetcd.Client

// Options etcd config options. All the fields have default value.
type Options = cwetcd.Options

// NewClient Create a default etcd client
// It can create a client with default config by env variable.
// See: env.go
func NewClient(opts Options) (Client, error) {
	return cwetcd.NewClient(opts)
}
