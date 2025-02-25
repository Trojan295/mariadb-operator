/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func (r *Grant) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//nolint
//+kubebuilder:webhook:path=/validate-mariadb-mmontes-io-v1alpha1-grant,mutating=false,failurePolicy=fail,sideEffects=None,groups=mariadb.mmontes.io,resources=grants,verbs=create;update,versions=v1alpha1,name=vgrant.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Grant{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Grant) ValidateCreate() (admission.Warnings, error) {
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Grant) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	if err := inmutableWebhook.ValidateUpdate(r, old.(*Grant)); err != nil {
		return nil, err
	}
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Grant) ValidateDelete() (admission.Warnings, error) {
	return nil, nil
}
