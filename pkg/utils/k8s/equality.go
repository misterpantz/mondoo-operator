package k8s

import (
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

// AreDeploymentsEqual returns a value indicating whether 2 deployments are equal. Note that it does not perform a full
// comparison but checks just some of the properties of a deployment (only the ones we are currently interested at).
func AreDeploymentsEqual(a, b appsv1.Deployment) bool {
	return len(a.Spec.Template.Spec.Containers) == len(b.Spec.Template.Spec.Containers) &&
		reflect.DeepEqual(a.Spec.Replicas, b.Spec.Replicas) &&
		reflect.DeepEqual(a.Spec.Selector, b.Spec.Selector) &&
		a.Spec.Template.Spec.ServiceAccountName == b.Spec.Template.Spec.ServiceAccountName &&
		reflect.DeepEqual(a.Spec.Template.Spec.Containers[0].Image, b.Spec.Template.Spec.Containers[0].Image) &&
		reflect.DeepEqual(a.Spec.Template.Spec.Containers[0].Command, b.Spec.Template.Spec.Containers[0].Command) &&
		reflect.DeepEqual(a.Spec.Template.Spec.Containers[0].Args, b.Spec.Template.Spec.Containers[0].Args) &&
		reflect.DeepEqual(a.Spec.Template.Spec.Containers[0].VolumeMounts, b.Spec.Template.Spec.Containers[0].VolumeMounts) &&
		reflect.DeepEqual(a.Spec.Template.Spec.Containers[0].Env, b.Spec.Template.Spec.Containers[0].Env) &&
		reflect.DeepEqual(a.Spec.Template.Spec.Volumes, b.Spec.Template.Spec.Volumes) &&
		reflect.DeepEqual(a.GetOwnerReferences(), b.GetOwnerReferences())
}

// AreServicesEqual return a value indicating whether 2 services are equal. Note that it
// does not perform a full comparison but checks just some of the properties of a deployment
// (only the ones we are currently interested at).
func AreServicesEqual(a, b corev1.Service) bool {
	return reflect.DeepEqual(a.Spec.Ports, b.Spec.Ports) &&
		reflect.DeepEqual(a.Spec.Selector, b.Spec.Selector) &&
		reflect.DeepEqual(a.GetOwnerReferences(), b.GetOwnerReferences()) &&
		a.Spec.Type == b.Spec.Type
}

// AreCronJobsEqual returns a value indicating whether 2 cron jobs are equal. Note that it does not perform a full
// comparison but checks just some of the properties of a deployment (only the ones we are currently interested at).
func AreCronJobsEqual(a, b batchv1.CronJob) bool {
	aPodSpec := a.Spec.JobTemplate.Spec.Template.Spec
	bPodSpec := b.Spec.JobTemplate.Spec.Template.Spec
	return len(aPodSpec.Containers) == len(bPodSpec.Containers) &&
		aPodSpec.ServiceAccountName == bPodSpec.ServiceAccountName &&
		reflect.DeepEqual(aPodSpec.Tolerations, bPodSpec.Tolerations) &&
		reflect.DeepEqual(aPodSpec.NodeName, bPodSpec.NodeName) &&
		reflect.DeepEqual(aPodSpec.Containers[0].Image, bPodSpec.Containers[0].Image) &&
		reflect.DeepEqual(aPodSpec.Containers[0].Command, bPodSpec.Containers[0].Command) &&
		reflect.DeepEqual(aPodSpec.Containers[0].Args, bPodSpec.Containers[0].Args) &&
		reflect.DeepEqual(aPodSpec.Containers[0].VolumeMounts, bPodSpec.Containers[0].VolumeMounts) &&
		reflect.DeepEqual(aPodSpec.Containers[0].Env, bPodSpec.Containers[0].Env) &&
		reflect.DeepEqual(aPodSpec.Volumes, bPodSpec.Volumes) &&
		reflect.DeepEqual(a.Spec.SuccessfulJobsHistoryLimit, b.Spec.SuccessfulJobsHistoryLimit) &&
		reflect.DeepEqual(a.Spec.FailedJobsHistoryLimit, b.Spec.FailedJobsHistoryLimit) &&
		reflect.DeepEqual(a.GetOwnerReferences(), b.GetOwnerReferences())
}

// AreResouceRequirementsEqual returns a value indicating whether 2 resource requirements are equal.
func AreResouceRequirementsEqual(x corev1.ResourceRequirements, y corev1.ResourceRequirements) bool {
	if x.Limits.Cpu().Equal(*y.Limits.Cpu()) &&
		x.Limits.Memory().Equal(*y.Limits.Memory()) &&
		x.Requests.Cpu().Equal(*y.Requests.Cpu()) &&
		x.Requests.Memory().Equal(*y.Requests.Memory()) {
		return true
	}
	return false
}
