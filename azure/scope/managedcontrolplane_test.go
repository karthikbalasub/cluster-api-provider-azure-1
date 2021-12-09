/*
Copyright 2021 The Kubernetes Authors.

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

package scope

import (
	"context"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1beta1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	capiv1exp "sigs.k8s.io/cluster-api/exp/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestManagedControlPlaneScope_Autoscaling(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = capiv1exp.AddToScheme(scheme)
	_ = infrav1.AddToScheme(scheme)

	cases := []struct {
		Name     string
		Input    ManagedControlPlaneScopeParams
		Expected azure.AgentPoolSpec
	}{
		{
			Name: "Without Autoscaling",
			Input: ManagedControlPlaneScopeParams{
				AzureClients: AzureClients{
					Authorizer: autorest.NullAuthorizer{},
				},
				Cluster: &clusterv1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cluster1",
						Namespace: "default",
					},
				},
				ControlPlane: &infrav1.AzureManagedControlPlane{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cluster1",
						Namespace: "default",
					},
					Spec: infrav1.AzureManagedControlPlaneSpec{
						SubscriptionID: "00000000-0000-0000-0000-000000000000",
					},
				},
				MachinePool:      getMachinePool("pool0"),
				InfraMachinePool: getAzureMachinePool("pool0", infrav1.NodePoolModeSystem),
				PatchTarget:      getAzureMachinePool("pool0", infrav1.NodePoolModeSystem),
			},
			Expected: azure.AgentPoolSpec{

				Name:         "pool0",
				SKU:          "Standard_D2s_v3",
				Replicas:     1,
				Mode:         "System",
				Cluster:      "cluster1",
				VnetSubnetID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/",
			},
		},
		{
			Name: "With Autoscaling",
			Input: ManagedControlPlaneScopeParams{
				AzureClients: AzureClients{
					Authorizer: autorest.NullAuthorizer{},
				},
				Cluster: &clusterv1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cluster1",
						Namespace: "default",
					},
				},
				ControlPlane: &infrav1.AzureManagedControlPlane{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cluster1",
						Namespace: "default",
					},
					Spec: infrav1.AzureManagedControlPlaneSpec{
						SubscriptionID: "00000000-0000-0000-0000-000000000000",
					},
				},
				MachinePool:      getMachinePool("pool1"),
				InfraMachinePool: getAzureMachinePoolWithScaling("pool1", 2, 10),
				PatchTarget:      getAzureMachinePoolWithScaling("pool1", 2, 10),
			},
			Expected: azure.AgentPoolSpec{
				Name:              "pool1",
				SKU:               "Standard_D2s_v3",
				Mode:              "User",
				Cluster:           "cluster1",
				Replicas:          1,
				EnableAutoScaling: to.BoolPtr(true),
				MinCount:          to.Int32Ptr(2),
				MaxCount:          to.Int32Ptr(10),
				VnetSubnetID:      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			g := NewWithT(t)
			fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(c.Input.MachinePool, c.Input.InfraMachinePool, c.Input.ControlPlane).Build()
			c.Input.Client = fakeClient
			s, err := NewManagedControlPlaneScope(context.TODO(), c.Input)
			g.Expect(err).To(Succeed())
			agentPool := s.AgentPoolSpec()
			g.Expect(agentPool).To(Equal(c.Expected))
		})
	}
}

func getAzureMachinePool(name string, mode infrav1.NodePoolMode) *infrav1.AzureManagedMachinePool {
	return &infrav1.AzureManagedMachinePool{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels: map[string]string{
				clusterv1.ClusterLabelName: "cluster1",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "cluster.x-k8s.io/v1beta1",
					Kind:       "MachinePool",
					Name:       name,
				},
			},
		},
		Spec: infrav1.AzureManagedMachinePoolSpec{
			Mode: string(mode),
			SKU:  "Standard_D2s_v3",
			Name: to.StringPtr(name),
		},
	}
}

func getAzureMachinePoolWithScaling(name string, min, max int32) *infrav1.AzureManagedMachinePool {
	managedPool := getAzureMachinePool(name, infrav1.NodePoolModeUser)
	managedPool.Spec.Scaling = &infrav1.ManagedMachinePoolScaling{
		MinSize: to.Int32Ptr(min),
		MaxSize: to.Int32Ptr(max),
	}
	return managedPool
}

func getMachinePool(name string) *capiv1exp.MachinePool {
	return &capiv1exp.MachinePool{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels: map[string]string{
				clusterv1.ClusterLabelName: "cluster1",
			},
		},
		Spec: capiv1exp.MachinePoolSpec{
			ClusterName: "cluster1",
		},
	}
}