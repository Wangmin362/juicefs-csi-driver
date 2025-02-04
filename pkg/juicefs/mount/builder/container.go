/*
 Copyright 2022 Juicedata Inc

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

package builder

import (
	corev1 "k8s.io/api/core/v1"
)

func (r *Builder) NewMountSidecar() *corev1.Pod {
	pod := r.NewMountPod("")
	for i, v := range pod.Spec.Volumes {
		if v.Name == JfsRootDirName {
			v = corev1.Volume{
				Name: v.Name,
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			}
			pod.Spec.Volumes[i] = v
		}
	}
	return pod
}
