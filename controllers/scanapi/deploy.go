/*
Copyright 2022 Mondoo, Inc.

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

package scanapi

import (
	"context"

	"go.mondoo.com/mondoo-operator/pkg/utils/k8s"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	mondoov1alpha2 "go.mondoo.com/mondoo-operator/api/v1alpha2"
)

var logger = ctrl.Log.WithName("scan-api-deploy")

// Deploy deploys the scan API for a given MondooAuditConfig. The function checks if the scan API is already deployed.
// If that is the case, the existing resources are compared with the ones that are desired and the necessary updates are applied.
func Deploy(ctx context.Context, kubeClient client.Client, ns, image string, mondoo mondoov1alpha2.MondooAuditConfig) error {
	if err := createSecret(ctx, kubeClient, mondoo); err != nil {
		return err
	}
	if err := createDeployment(ctx, kubeClient, ns, image, mondoo); err != nil {
		return err
	}
	return createService(ctx, kubeClient, ns, mondoo)
}

// Cleanup cleans up the scan API for a given MondooAuditConfig. The function returns no errors if the scan API is already
// deleted.
func Cleanup(ctx context.Context, kubeClient client.Client, ns string, mondoo mondoov1alpha2.MondooAuditConfig) error {
	scanApiTokenSecret := ScanApiSecret(mondoo)
	if err := k8s.DeleteIfExists(ctx, kubeClient, scanApiTokenSecret); err != nil {
		logger.Error(err, "failed to clean up scan API token Secret resource")
		return err
	}
	scanApiDeployment := ScanApiDeployment(ns, "", mondoo) // Image is not relevant when deleting.
	if err := k8s.DeleteIfExists(ctx, kubeClient, scanApiDeployment); err != nil {
		logger.Error(err, "failed to clean up scan API Deployment resource")
		return err
	}

	scanApiService := ScanApiService(ns, mondoo)
	if err := k8s.DeleteIfExists(ctx, kubeClient, scanApiService); err != nil {
		logger.Error(err, "failed to clean up scan API Service resource")
		return err
	}
	return nil
}

func createSecret(ctx context.Context, kubeClient client.Client, mondoo mondoov1alpha2.MondooAuditConfig) error {
	scanApiTokenSecret := ScanApiSecret(mondoo)
	if err := ctrl.SetControllerReference(&mondoo, scanApiTokenSecret, kubeClient.Scheme()); err != nil {
		return err
	}

	// Doing a direct Create() so that we don't have to do the Get()->IfNotExists->Create() dance
	// which lets us avoid asking for Get/List on Secrets across all Namespaces.
	err := kubeClient.Create(ctx, scanApiTokenSecret)
	if err == nil {
		logger.Info("Created token Secret for scan API")
		return nil
	} else if errors.IsAlreadyExists(err) {
		logger.Info("Token Secret for scan API already exists")
		return nil
	} else {
		logger.Error(err, "Faled to create/check for existence of token Secret for scan API")
		return err
	}
}

func createDeployment(ctx context.Context, kubeClient client.Client, ns, image string, mondoo mondoov1alpha2.MondooAuditConfig) error {
	deployment := ScanApiDeployment(ns, image, mondoo)
	if err := ctrl.SetControllerReference(&mondoo, deployment, kubeClient.Scheme()); err != nil {
		return err
	}
	existingDeployment := appsv1.Deployment{}
	created, err := k8s.CreateIfNotExist(ctx, kubeClient, &existingDeployment, deployment)
	if err != nil {
		logger.Error(err, "Failed to create Deployment for scan API")
		return err
	}

	if created {
		logger.Info("Created Deployment for scan API")
	} else if !k8s.AreDeploymentsEqual(*deployment, existingDeployment) {
		logger.Info("Updated needed for scan API Deployment")
		// If the deployment exists but it is different from what we actually want it to be, then update.
		k8s.UpdateDeployment(&existingDeployment, *deployment)
		if err := kubeClient.Update(ctx, &existingDeployment); err != nil {
			return err
		}
	}
	return nil
}

func createService(ctx context.Context, kubeClient client.Client, ns string, mondoo mondoov1alpha2.MondooAuditConfig) error {
	service := ScanApiService(ns, mondoo)
	if err := ctrl.SetControllerReference(&mondoo, service, kubeClient.Scheme()); err != nil {
		return err
	}
	existingService := corev1.Service{}
	created, err := k8s.CreateIfNotExist(ctx, kubeClient, &existingService, service)
	if err != nil {
		logger.Error(err, "Failed to create Service for scan API")
		return err
	}

	if created {
		logger.Info("Created Service for scan API")
	} else if !k8s.AreServicesEqual(*service, existingService) {
		k8s.UpdateService(&existingService, *service)
		// If the service exists but it is different from what we actually want it to be, then update.
		if err := kubeClient.Update(ctx, &existingService); err != nil {
			return err
		}
	}
	return nil
}
