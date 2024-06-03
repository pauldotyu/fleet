/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package e2e

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	placementv1beta1 "go.goms.io/fleet/apis/placement/v1beta1"
	"go.goms.io/fleet/pkg/utils"
	"go.goms.io/fleet/test/e2e/framework"
	testutils "go.goms.io/fleet/test/e2e/v1alpha1/utils"
	"go.goms.io/fleet/test/utils/controller"
)

// Note that this container will run in parallel with other containers.
var _ = Describe("placing wrapped resources using a CRP", Ordered, func() {
	Context("Test a CRP place enveloped objects successfully", Ordered, func() {
		crpName := fmt.Sprintf(crpNameTemplate, GinkgoParallelProcess())
		workNamespaceName := appNamespace().Name
		var wantSelectedResources []placementv1beta1.ResourceIdentifier
		var testEnvelopDeployment corev1.ConfigMap
		var testDeployment appv1.Deployment

		BeforeAll(func() {
			readDeploymentTestManifest(&testDeployment)
			readEnvelopConfigMapTestManifest(&testEnvelopDeployment)
			wantSelectedResources = []placementv1beta1.ResourceIdentifier{
				{
					Kind:    "Namespace",
					Name:    workNamespaceName,
					Version: "v1",
				},
				{
					Kind:      "ConfigMap",
					Name:      testEnvelopDeployment.Name,
					Version:   "v1",
					Namespace: workNamespaceName,
				},
			}
		})

		It("Create the wrapped deployment resources in the namespace", func() {
			createWrappedResourcesForRollout(&testEnvelopDeployment, &testDeployment)
		})

		It("Create the CRP that select the name space", func() {
			crp := &placementv1beta1.ClusterResourcePlacement{
				ObjectMeta: metav1.ObjectMeta{
					Name: crpName,
					// Add a custom finalizer; this would allow us to better observe
					// the behavior of the controllers.
					Finalizers: []string{customDeletionBlockerFinalizer},
				},
				Spec: placementv1beta1.ClusterResourcePlacementSpec{
					ResourceSelectors: workResourceSelector(),
					Strategy: placementv1beta1.RolloutStrategy{
						Type: placementv1beta1.RollingUpdateRolloutStrategyType,
						RollingUpdate: &placementv1beta1.RollingUpdateConfig{
							UnavailablePeriodSeconds: ptr.To(2),
						},
					},
				},
			}
			Expect(hubClient.Create(ctx, crp)).To(Succeed(), "Failed to create CRP")
		})

		It("should update CRP status as expected", func() {
			crpStatusUpdatedActual := crpStatusUpdatedActual(wantSelectedResources, allMemberClusterNames, nil, "0")
			// For the development, at least it will take 4 minutes to be ready.
			Eventually(crpStatusUpdatedActual, 6*time.Minute, eventuallyInterval).Should(Succeed(), "Failed to update CRP status as expected")
		})

		It("should place the resources on all member clusters", func() {
			for idx := range allMemberClusters {
				memberCluster := allMemberClusters[idx]
				workResourcesPlacedActual := waitForDeploymentPlacementToReady(memberCluster, &testDeployment)
				Eventually(workResourcesPlacedActual, eventuallyDuration, eventuallyInterval).Should(Succeed(), "Failed to place work resources on member cluster %s", memberCluster.ClusterName)
			}
		})

		It("should mark the work as available", func() {
			for idx := range allMemberClusters {
				memberCluster := allMemberClusters[idx]
				var works placementv1beta1.WorkList
				listOpts := []client.ListOption{
					client.InNamespace(fmt.Sprintf(utils.NamespaceNameFormat, memberCluster.ClusterName)),
				}
				Eventually(func() string {
					if err := hubClient.List(ctx, &works, listOpts...); err != nil {
						return err.Error()
					}
					for i := range works.Items {
						work := works.Items[i]
						wantConditions := []metav1.Condition{
							{
								Type:               placementv1beta1.WorkConditionTypeApplied,
								Status:             metav1.ConditionTrue,
								Reason:             "WorkAppliedCompleted",
								ObservedGeneration: 1,
							},
							{
								Type:               placementv1beta1.WorkConditionTypeAvailable,
								Status:             metav1.ConditionTrue,
								Reason:             "WorkAvailable",
								ObservedGeneration: 1,
							},
						}
						diff := controller.CompareConditions(wantConditions, work.Status.Conditions)
						if len(diff) != 0 {
							return diff
						}
					}
					if len(works.Items) == 0 {
						return "no available work found"
					}
					return ""
				}, testutils.PollTimeout, testutils.PollInterval).Should(BeEmpty(),
					"work condition mismatch for work %s (-want, +got):", memberCluster.ClusterName)
			}
		})

		AfterAll(func() {
			// Remove the custom deletion blocker finalizer from the CRP.
			ensureCRPAndRelatedResourcesDeletion(crpName, allMemberClusters)
		})
	})

	Context("Test a CRP place workload objects successfully, block rollout based on availability", Ordered, func() {
		crpName := fmt.Sprintf(crpNameTemplate, GinkgoParallelProcess())
		workNamespaceName := appNamespace().Name
		var wantSelectedResources []placementv1beta1.ResourceIdentifier
		var testDeployment appv1.Deployment

		BeforeAll(func() {
			// Create the test resources.
			readDeploymentTestManifest(&testDeployment)
			wantSelectedResources = []placementv1beta1.ResourceIdentifier{
				{
					Kind:    "Namespace",
					Name:    workNamespaceName,
					Version: "v1",
				},
				{
					Group:     "apps",
					Version:   "v1",
					Kind:      "Deployment",
					Name:      testDeployment.Name,
					Namespace: workNamespaceName,
				},
			}
		})

		It("create the deployment resource in the namespace", func() {
			createDeploymentForRollout(&testDeployment)
		})

		It("create the CRP that select the name space", func() {
			crp := &placementv1beta1.ClusterResourcePlacement{
				ObjectMeta: metav1.ObjectMeta{
					Name: crpName,
					// Add a custom finalizer; this would allow us to better observe
					// the behavior of the controllers.
					Finalizers: []string{customDeletionBlockerFinalizer},
				},
				Spec: placementv1beta1.ClusterResourcePlacementSpec{
					ResourceSelectors: workResourceSelector(),
					Strategy: placementv1beta1.RolloutStrategy{
						Type: placementv1beta1.RollingUpdateRolloutStrategyType,
						RollingUpdate: &placementv1beta1.RollingUpdateConfig{
							MaxUnavailable: &intstr.IntOrString{
								Type:   intstr.Int,
								IntVal: 1,
							},
						},
					},
				},
			}
			Expect(hubClient.Create(ctx, crp)).To(Succeed(), "Failed to create CRP")
		})

		It("should update CRP status as expected", func() {
			crpStatusUpdatedActual := crpStatusUpdatedActual(wantSelectedResources, allMemberClusterNames, nil, "0")
			// For deployment, at the least it will take 4 minutes to be ready.
			Eventually(crpStatusUpdatedActual, 6*time.Minute, eventuallyInterval).Should(Succeed(), "Failed to update CRP status as expected")
		})

		It("should place the resources on all member clusters", func() {
			for idx := range allMemberClusters {
				memberCluster := allMemberClusters[idx]
				workResourcesPlacedActual := waitForDeploymentPlacementToReady(memberCluster, &testDeployment)
				Eventually(workResourcesPlacedActual, eventuallyDuration, eventuallyInterval).Should(Succeed(), "Failed to place work resources on member cluster %s", memberCluster.ClusterName)
			}
		})

		It("change the image name in deployment, to make it unavailable", func() {
			Eventually(func() error {
				var dep appv1.Deployment
				err := hubClient.Get(ctx, types.NamespacedName{Name: testDeployment.Name, Namespace: testDeployment.Namespace}, &dep)
				if err != nil {
					return err
				}
				dep.Spec.Template.Spec.Containers[0].Image = "random-image-name"
				return hubClient.Update(ctx, &dep)
			}, eventuallyDuration, eventuallyInterval).Should(Succeed(), "Failed to change the image name in deployment")
		})

		It("should update CRP status as expected", func() {
			failedDeploymentResourceIdentifier := placementv1beta1.ResourceIdentifier{
				Group:     appv1.SchemeGroupVersion.Group,
				Version:   appv1.SchemeGroupVersion.Version,
				Kind:      "Deployment",
				Name:      testDeployment.Name,
				Namespace: testDeployment.Namespace,
			}
			crpStatusActual := safeDeploymentCRPStatusUpdatedActual(wantSelectedResources, failedDeploymentResourceIdentifier, allMemberClusterNames, "1")
			// For deployment, at the least it will take 4 minutes to be ready.
			Eventually(crpStatusActual, 6*time.Minute, eventuallyInterval).Should(Succeed(), "Failed to update CRP status as expected")
		})

		AfterAll(func() {
			// Remove the custom deletion blocker finalizer from the CRP.
			ensureCRPAndRelatedResourcesDeletion(crpName, allMemberClusters)
		})
	})
})

func readDeploymentTestManifest(testDeployment *appv1.Deployment) {
	By("Read the deployment resources")
	err := utils.GetObjectFromManifest("resources/test-deployment.yaml", testDeployment)
	Expect(err).Should(Succeed())
}

func readEnvelopConfigMapTestManifest(testEnvelopDeployment *corev1.ConfigMap) {
	By("Read testEnvelopConfigMap resource")
	err := utils.GetObjectFromManifest("resources/test-envelop-deployment.yaml", testEnvelopDeployment)
	Expect(err).Should(Succeed())
}

func createDeploymentForRollout(testDeployment *appv1.Deployment) {
	ns := appNamespace()
	Expect(hubClient.Create(ctx, &ns)).To(Succeed(), "Failed to create namespace %s", ns.Namespace)
	testDeployment.Namespace = ns.Name
	Expect(hubClient.Create(ctx, testDeployment)).To(Succeed(), "Failed to create test deployment %s", testDeployment.Name)
}

// createWrappedResourcesForRollout creates some enveloped resources on the hub cluster with a deployment for testing purposes.
func createWrappedResourcesForRollout(testEnvelopDeployment *corev1.ConfigMap, testDeployment *appv1.Deployment) {
	ns := appNamespace()
	Expect(hubClient.Create(ctx, &ns)).To(Succeed(), "Failed to create namespace %s", ns.Namespace)
	// modify the enveloped configMap according to the namespace
	testEnvelopDeployment.Namespace = ns.Name

	// modify the embedded namespaced resource according to the namespace
	testDeployment.Namespace = ns.Name
	resourceDeploymentByte, err := json.Marshal(testDeployment)
	Expect(err).Should(Succeed())
	testEnvelopDeployment.Data["deployment.yaml"] = string(resourceDeploymentByte)
	Expect(hubClient.Create(ctx, testEnvelopDeployment)).To(Succeed(), "Failed to create testEnvelop deployment %s", testEnvelopDeployment.Name)
}

func waitForDeploymentPlacementToReady(memberCluster *framework.Cluster, testDeployment *appv1.Deployment) func() error {
	workNamespaceName := appNamespace().Name
	return func() error {
		if err := validateWorkNamespaceOnCluster(memberCluster, types.NamespacedName{Name: workNamespaceName}); err != nil {
			return err
		}
		By("check the placedDeployment")
		placedDeployment := &appv1.Deployment{}
		if err := memberCluster.KubeClient.Get(ctx, types.NamespacedName{Namespace: workNamespaceName, Name: testDeployment.Name}, placedDeployment); err != nil {
			return err
		}
		By("check the placedDeployment is ready")
		var depCond *appv1.DeploymentCondition
		for i := range placedDeployment.Status.Conditions {
			if placedDeployment.Status.Conditions[i].Type == appv1.DeploymentAvailable {
				depCond = &placedDeployment.Status.Conditions[i]
				break
			}
		}
		if placedDeployment.Status.ObservedGeneration == placedDeployment.Generation && depCond != nil && depCond.Status == corev1.ConditionTrue {
			return nil
		}
		return nil
	}
}