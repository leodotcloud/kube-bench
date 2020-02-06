// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha3 "sigs.k8s.io/kind/pkg/apis/config/v1alpha3"
	cri "sigs.k8s.io/kind/pkg/container/cri"
	config "sigs.k8s.io/kind/pkg/internal/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha3.Cluster)(nil), (*config.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Cluster_To_config_Cluster(a.(*v1alpha3.Cluster), b.(*config.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Cluster)(nil), (*v1alpha3.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Cluster_To_v1alpha3_Cluster(a.(*config.Cluster), b.(*v1alpha3.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.Networking)(nil), (*config.Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Networking_To_config_Networking(a.(*v1alpha3.Networking), b.(*config.Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Networking)(nil), (*v1alpha3.Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Networking_To_v1alpha3_Networking(a.(*config.Networking), b.(*v1alpha3.Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.Node)(nil), (*config.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Node_To_config_Node(a.(*v1alpha3.Node), b.(*config.Node), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Node)(nil), (*v1alpha3.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Node_To_v1alpha3_Node(a.(*config.Node), b.(*v1alpha3.Node), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.PatchJSON6902)(nil), (*config.PatchJSON6902)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_PatchJSON6902_To_config_PatchJSON6902(a.(*v1alpha3.PatchJSON6902), b.(*config.PatchJSON6902), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PatchJSON6902)(nil), (*v1alpha3.PatchJSON6902)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PatchJSON6902_To_v1alpha3_PatchJSON6902(a.(*config.PatchJSON6902), b.(*v1alpha3.PatchJSON6902), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_Cluster_To_config_Cluster(in *v1alpha3.Cluster, out *config.Cluster, s conversion.Scope) error {
	out.Nodes = *(*[]config.Node)(unsafe.Pointer(&in.Nodes))
	if err := Convert_v1alpha3_Networking_To_config_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]config.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	return nil
}

// Convert_v1alpha3_Cluster_To_config_Cluster is an autogenerated conversion function.
func Convert_v1alpha3_Cluster_To_config_Cluster(in *v1alpha3.Cluster, out *config.Cluster, s conversion.Scope) error {
	return autoConvert_v1alpha3_Cluster_To_config_Cluster(in, out, s)
}

func autoConvert_config_Cluster_To_v1alpha3_Cluster(in *config.Cluster, out *v1alpha3.Cluster, s conversion.Scope) error {
	out.Nodes = *(*[]v1alpha3.Node)(unsafe.Pointer(&in.Nodes))
	if err := Convert_config_Networking_To_v1alpha3_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]v1alpha3.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	return nil
}

// Convert_config_Cluster_To_v1alpha3_Cluster is an autogenerated conversion function.
func Convert_config_Cluster_To_v1alpha3_Cluster(in *config.Cluster, out *v1alpha3.Cluster, s conversion.Scope) error {
	return autoConvert_config_Cluster_To_v1alpha3_Cluster(in, out, s)
}

func autoConvert_v1alpha3_Networking_To_config_Networking(in *v1alpha3.Networking, out *config.Networking, s conversion.Scope) error {
	out.IPFamily = config.ClusterIPFamily(in.IPFamily)
	out.APIServerPort = in.APIServerPort
	out.APIServerAddress = in.APIServerAddress
	out.PodSubnet = in.PodSubnet
	out.ServiceSubnet = in.ServiceSubnet
	out.DisableDefaultCNI = in.DisableDefaultCNI
	return nil
}

// Convert_v1alpha3_Networking_To_config_Networking is an autogenerated conversion function.
func Convert_v1alpha3_Networking_To_config_Networking(in *v1alpha3.Networking, out *config.Networking, s conversion.Scope) error {
	return autoConvert_v1alpha3_Networking_To_config_Networking(in, out, s)
}

func autoConvert_config_Networking_To_v1alpha3_Networking(in *config.Networking, out *v1alpha3.Networking, s conversion.Scope) error {
	out.IPFamily = v1alpha3.ClusterIPFamily(in.IPFamily)
	out.APIServerPort = in.APIServerPort
	out.APIServerAddress = in.APIServerAddress
	out.PodSubnet = in.PodSubnet
	out.ServiceSubnet = in.ServiceSubnet
	out.DisableDefaultCNI = in.DisableDefaultCNI
	return nil
}

// Convert_config_Networking_To_v1alpha3_Networking is an autogenerated conversion function.
func Convert_config_Networking_To_v1alpha3_Networking(in *config.Networking, out *v1alpha3.Networking, s conversion.Scope) error {
	return autoConvert_config_Networking_To_v1alpha3_Networking(in, out, s)
}

func autoConvert_v1alpha3_Node_To_config_Node(in *v1alpha3.Node, out *config.Node, s conversion.Scope) error {
	out.Role = config.NodeRole(in.Role)
	out.Image = in.Image
	out.ExtraMounts = *(*[]cri.Mount)(unsafe.Pointer(&in.ExtraMounts))
	out.ExtraPortMappings = *(*[]cri.PortMapping)(unsafe.Pointer(&in.ExtraPortMappings))
	return nil
}

// Convert_v1alpha3_Node_To_config_Node is an autogenerated conversion function.
func Convert_v1alpha3_Node_To_config_Node(in *v1alpha3.Node, out *config.Node, s conversion.Scope) error {
	return autoConvert_v1alpha3_Node_To_config_Node(in, out, s)
}

func autoConvert_config_Node_To_v1alpha3_Node(in *config.Node, out *v1alpha3.Node, s conversion.Scope) error {
	out.Role = v1alpha3.NodeRole(in.Role)
	out.Image = in.Image
	out.ExtraMounts = *(*[]cri.Mount)(unsafe.Pointer(&in.ExtraMounts))
	out.ExtraPortMappings = *(*[]cri.PortMapping)(unsafe.Pointer(&in.ExtraPortMappings))
	return nil
}

// Convert_config_Node_To_v1alpha3_Node is an autogenerated conversion function.
func Convert_config_Node_To_v1alpha3_Node(in *config.Node, out *v1alpha3.Node, s conversion.Scope) error {
	return autoConvert_config_Node_To_v1alpha3_Node(in, out, s)
}

func autoConvert_v1alpha3_PatchJSON6902_To_config_PatchJSON6902(in *v1alpha3.PatchJSON6902, out *config.PatchJSON6902, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Patch = in.Patch
	return nil
}

// Convert_v1alpha3_PatchJSON6902_To_config_PatchJSON6902 is an autogenerated conversion function.
func Convert_v1alpha3_PatchJSON6902_To_config_PatchJSON6902(in *v1alpha3.PatchJSON6902, out *config.PatchJSON6902, s conversion.Scope) error {
	return autoConvert_v1alpha3_PatchJSON6902_To_config_PatchJSON6902(in, out, s)
}

func autoConvert_config_PatchJSON6902_To_v1alpha3_PatchJSON6902(in *config.PatchJSON6902, out *v1alpha3.PatchJSON6902, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Patch = in.Patch
	return nil
}

// Convert_config_PatchJSON6902_To_v1alpha3_PatchJSON6902 is an autogenerated conversion function.
func Convert_config_PatchJSON6902_To_v1alpha3_PatchJSON6902(in *config.PatchJSON6902, out *v1alpha3.PatchJSON6902, s conversion.Scope) error {
	return autoConvert_config_PatchJSON6902_To_v1alpha3_PatchJSON6902(in, out, s)
}
