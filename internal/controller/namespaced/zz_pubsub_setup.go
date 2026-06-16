// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	litereservation "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/litereservation"
	litesubscription "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/litesubscription"
	litetopic "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/litetopic"
	schema "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/schema"
	subscription "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/subscription"
	subscriptioniammember "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/subscriptioniammember"
	topic "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/topic"
	topiciammember "github.com/aayushrangwala/provider-upjet-gcp/v2/internal/controller/namespaced/pubsub/topiciammember"
)

// Setup_pubsub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_pubsub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		litereservation.Setup,
		litesubscription.Setup,
		litetopic.Setup,
		schema.Setup,
		subscription.Setup,
		subscriptioniammember.Setup,
		topic.Setup,
		topiciammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_pubsub creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_pubsub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		litereservation.SetupGated,
		litesubscription.SetupGated,
		litetopic.SetupGated,
		schema.SetupGated,
		subscription.SetupGated,
		subscriptioniammember.SetupGated,
		topic.SetupGated,
		topiciammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
