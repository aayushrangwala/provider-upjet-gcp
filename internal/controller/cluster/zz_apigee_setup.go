// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addonsconfig "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/addonsconfig"
	endpointattachment "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/endpointattachment"
	envgroup "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/envgroup"
	envgroupattachment "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/envgroupattachment"
	environment "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/environment"
	environmentiammember "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/environmentiammember"
	envkeystore "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/envkeystore"
	envreferences "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/envreferences"
	instance "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/instance"
	instanceattachment "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/instanceattachment"
	keystoresaliaseskeycertfile "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/keystoresaliaseskeycertfile"
	nataddress "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/nataddress"
	organization "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/organization"
	syncauthorization "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/syncauthorization"
	targetserver "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/cluster/apigee/targetserver"
)

// Setup_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.Setup,
		endpointattachment.Setup,
		envgroup.Setup,
		envgroupattachment.Setup,
		environment.Setup,
		environmentiammember.Setup,
		envkeystore.Setup,
		envreferences.Setup,
		instance.Setup,
		instanceattachment.Setup,
		keystoresaliaseskeycertfile.Setup,
		nataddress.Setup,
		organization.Setup,
		syncauthorization.Setup,
		targetserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.SetupGated,
		endpointattachment.SetupGated,
		envgroup.SetupGated,
		envgroupattachment.SetupGated,
		environment.SetupGated,
		environmentiammember.SetupGated,
		envkeystore.SetupGated,
		envreferences.SetupGated,
		instance.SetupGated,
		instanceattachment.SetupGated,
		keystoresaliaseskeycertfile.SetupGated,
		nataddress.SetupGated,
		organization.SetupGated,
		syncauthorization.SetupGated,
		targetserver.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
