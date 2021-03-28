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

package v2

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	v1 "tutorial.kubebuilder.io/kubebuilder-example/apis/webapp/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GuestbookSpec defines the desired state of Guestbook
type GuestbookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Guestbook. Edit Guestbook_types.go to remove/update
	Status   string   `json:"status"`
	Hobby    []string `json:"hobby"`    // v2版本修改类型
	Fullname string   `json:"fullname"` // v2版本新增
}

// GuestbookStatus defines the observed state of Guestbook
type GuestbookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Guestbook is the Schema for the guestbooks API
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuestbookSpec   `json:"spec,omitempty"`
	Status GuestbookStatus `json:"status,omitempty"`
}

// 当前版本 -> Hub版本（v1）
func (in *Guestbook) ConvertTo(dstRaw conversion.Hub) error {
	fmt.Println("v2 转 v1")
	dst := dstRaw.(*v1.Guestbook)
	dst.Spec.Name = in.Spec.Fullname
	hobby := ""
	for _, v := range in.Spec.Hobby {
		hobby += v
	}
	dst.Spec.Hobby = hobby

	dst.ObjectMeta = in.ObjectMeta
	dst.Spec.Status = in.Spec.Status
	//dst.Status = in.Status
	return nil
}

// Hub版本（v1）-> 当前版本
func (in *Guestbook) ConvertFrom(srcRaw conversion.Hub) error {
	fmt.Println("v1 转 v2")
	src := srcRaw.(*v1.Guestbook)

	in.Spec.Fullname = src.Spec.Name
	if in.Spec.Hobby == nil {
		in.Spec.Hobby = make([]string, 1)
	}
	in.Spec.Hobby[0] = src.Spec.Hobby
	in.Spec.Status = src.Spec.Status

	in.ObjectMeta = src.ObjectMeta

	return nil
}

// +kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
