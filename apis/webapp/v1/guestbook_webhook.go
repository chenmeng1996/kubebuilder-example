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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var guestbooklog = logf.Log.WithName("guestbook-resource")

// +kubebuilder:webhook:verbs=create;update,path=/validate-webapp-tutorial-kubebuilder-io-v1-guestbook,mutating=false,failurePolicy=fail,groups=webapp.tutorial.kubebuilder.io,resources=guestbooks,versions=v1,name=vcronjob.kb.io
// +kubebuilder:webhook:path=/mutate-webapp-tutorial-kubebuilder-io-v1-guestbook,mutating=true,failurePolicy=fail,groups=webapp.tutorial.kubebuilder.io,resources=guestbooks,verbs=create;update,versions=v1,name=mcronjob.kb.io

func (r *Guestbook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

var _ webhook.Defaulter = &Guestbook{}

func (r *Guestbook) Default() {

}

var _ webhook.Validator = &Guestbook{}

func (r *Guestbook) ValidateCreate() error {
	return nil
}

func (r *Guestbook) ValidateUpdate(old runtime.Object) error {
	return nil
}

func (r *Guestbook) ValidateDelete() error {
	return nil
}
