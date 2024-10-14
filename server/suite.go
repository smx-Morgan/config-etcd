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

package server

import (
	cwserver "github.com/cloudwego-contrib/cwgo-pkg/config/etcd/server"
	"github.com/kitex-contrib/config-etcd/etcd"
	"github.com/kitex-contrib/config-etcd/utils"
)

const (
	limiterConfigName = "limit"
)

// EtcdServerSuite etcd server config suite, configure limiter config dynamically from etcd.
type EtcdServerSuite = cwserver.EtcdServerSuite

// NewSuite service is the destination service.
func NewSuite(service string, cli etcd.Client,
	opts ...utils.Option,
) *EtcdServerSuite {
	return cwserver.NewSuite(service, cli, opts...)
}
