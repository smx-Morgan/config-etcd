// Copyright 2024 CloudWeGo Authors
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

package degradation

import (
	"errors"

	cwdegr "github.com/cloudwego-contrib/cwgo-pkg/config/etcd/pkg/degradation"
)

var errRejected = errors.New("rejected by client degradation config")

var defaultConfig = &Config{
	Enable:     false,
	Percentage: 0,
}

type Config = cwdegr.Config

// Container is a wrapper for Config
type Container = cwdegr.Container

func NewContainer() *Container {

	return cwdegr.NewContainer()
}
