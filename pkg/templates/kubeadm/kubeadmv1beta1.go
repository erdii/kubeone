/*
Copyright 2019 The KubeOne Authors.

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

package kubeadm

import (
	kubeoneapi "github.com/kubermatic/kubeone/pkg/apis/kubeone"
	"github.com/kubermatic/kubeone/pkg/templates"
	"github.com/kubermatic/kubeone/pkg/templates/kubeadm/v1beta1"
	"github.com/kubermatic/kubeone/pkg/util"
)

type kubeadmv1beta1 struct{}

func (*kubeadmv1beta1) Config(ctx *util.Context, instance kubeoneapi.HostConfig) (string, error) {
	config, err := v1beta1.NewConfig(ctx, instance)
	if err != nil {
		return "", err
	}

	return templates.KubernetesToYAML(config)
}

func (*kubeadmv1beta1) UpgradeLeaderCommand() string {
	return "kubeadm upgrade apply"
}

func (*kubeadmv1beta1) UpgradeFollowerCommand() string {
	return "kubeadm upgrade node experimental-control-plane"
}
