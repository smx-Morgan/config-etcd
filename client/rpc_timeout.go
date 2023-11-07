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

package client

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/rpctimeout"
	"github.com/kitex-contrib/config-etcd/etcd"
	"github.com/kitex-contrib/config-etcd/utils"
)

// WithRPCTimeout sets the RPC timeout policy from etcd configuration center.
func WithRPCTimeout(dest, src string, etcdClient etcd.Client, uniqueID int64, opts utils.Options) []client.Option {
	param, err := etcdClient.ClientConfigParam(&etcd.ConfigParamConfig{
		Category:          rpcTimeoutConfigName,
		ServerServiceName: dest,
		ClientServiceName: src,
	})
	if err != nil {
		panic(err)
	}

	for _, f := range opts.EtcdCustomFunctions {
		f(&param)
	}
	key := param.Prefix + "/" + param.Path
	return []client.Option{
		client.WithTimeoutProvider(initRPCTimeoutContainer(key, dest, etcdClient, uniqueID)),
		client.WithCloseCallbacks(func() error {
			// cancel the configuration listener when client is closed.
			etcdClient.DeregisterConfig(key, uniqueID)
			return nil
		}),
	}
}

func initRPCTimeoutContainer(key, dest string,
	etcdClient etcd.Client, uniqueID int64,
) rpcinfo.TimeoutProvider {
	rpcTimeoutContainer := rpctimeout.NewContainer()

	onChangeCallback := func(eventType mvccpb.Event_EventType, data string, parser etcd.ConfigParser) {
		if data == "" && eventType == mvccpb.PUT {
			klog.Debugf("[etcd] %s client etcd rpc timeout: get config failed: empty config, skip...", key)
			return
		}
		configs := map[string]*rpctimeout.RPCTimeout{}
		err := parser.Decode(data, &configs)
		if err != nil {
			klog.Warnf("[etcd] %s client etcd rpc timeout: unmarshal data %s failed: %s, skip...", key, data, err)
			return
		}
		rpcTimeoutContainer.NotifyPolicyChange(configs)
	}

	etcdClient.RegisterConfigCallback(context.Background(), key, uniqueID, onChangeCallback)

	return rpcTimeoutContainer
}
