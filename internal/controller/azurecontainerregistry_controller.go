/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/containers/azcontainerregistry"
	deploymentsv1beta1 "github.com/crodriguezde/acrops.git/api/v1beta1"
)

// AzureContainerRegistryReconciler reconciles a AzureContainerRegistry object
type AzureContainerRegistryReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=deployments.acrops.io,resources=azurecontainerregistries,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=deployments.acrops.io,resources=azurecontainerregistries/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=deployments.acrops.io,resources=azurecontainerregistries/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AzureContainerRegistry object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *AzureContainerRegistryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var acr deploymentsv1beta1.AzureContainerRegistry

	if err := r.Get(ctx, req.NamespacedName, &acr); err != nil {
		log.FromContext(ctx).Error(err, "Unable to fetch Ping")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Create a DefaultAzureCredential using managed identity.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.FromContext(ctx).Error(err, "Error creating Azure credential")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	registryURL := fmt.Sprintf("https://%s.azurecr.io", acr.Spec.Name)

	registryClient, err := azcontainerregistry.NewClient(registryURL, cred, nil)
	if err != nil {
		log.FromContext(ctx).Error(err, "Error creating Azure Container Registry client")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	pager := registryClient.NewListRepositoriesPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.FromContext(ctx).Error(err, "failed to advance page")
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		for _, v := range page.Repositories.Names {
			log.FromContext(ctx).Info(fmt.Sprintf("repository: %s\n", *v))
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AzureContainerRegistryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&deploymentsv1beta1.AzureContainerRegistry{}).
		Complete(r)
}
