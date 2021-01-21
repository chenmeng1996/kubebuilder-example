/*


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
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "tutorial.kubebuilder.io/kubebuilder-example/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.tutorial.kubebuilder.io,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.tutorial.kubebuilder.io,resources=guestbooks/status,verbs=get;update;patch

func (r *GuestbookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("guestbook", req.NamespacedName)

	// your logic here
	// 新增逻辑
	ctx := context.Background()
	_ = r.Log.WithValues("apiexample", req.NamespacedName)

	// 获取当前的CR，并打印
	obj := &webappv1.Guestbook{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		log.Println(err, "Unable to fetch object")
	} else {
		log.Println("Geting from Kubebuilder to", obj.Spec.FirstName, obj.Spec.LastName)
	}

	// 初始化 CR 的Status 为 Running
	obj.Status.Status = "Running"
	if err := r.Status().Update(ctx, obj); err != nil {
		log.Println(err, "unable to update status")
	}

	return ctrl.Result{}, nil
}

func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}
