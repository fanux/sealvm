/*
Copyright 2019 fanux.

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

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "github.com/fanux/sealvm/api/v1"
	v1 "github.com/fanux/sealvm/api/v1"
)

// VirtulMachineReconciler reconciles a VirtulMachine object
type VirtulMachineReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=infra.sealyun.com,resources=virtulmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infra.sealyun.com,resources=virtulmachines/status,verbs=get;update;patch

//Reconcile is
func (r *VirtulMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("virtulmachine", req.NamespacedName)

	vm := &v1.VirtulMachine{}
	if err := r.Get(ctx, req.NamespacedName, vm); err != nil {
		log.Error(err, "unable to fetch vm")
	} else {
		fmt.Println(vm.Spec.CPU, vm.Spec.Memory)
	}

	vm.Status.Status = "Running"
	if err := r.Status().Update(ctx, vm); err != nil {
		log.Error(err, "unable to update vm status")
	}

	time.Sleep(time.Second * 10)
	if err := r.Delete(ctx, vm); err != nil {
		log.Error(err, "unable to delete vm ", "vm", vm)
	}

	return ctrl.Result{}, nil
}

//SetupWithManager is
func (r *VirtulMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.VirtulMachine{}).
		Complete(r)
}
