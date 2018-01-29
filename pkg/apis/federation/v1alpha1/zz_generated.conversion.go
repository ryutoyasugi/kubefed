// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	federation "github.com/marun/fnord/pkg/apis/federation"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_FederatedSecret_To_federation_FederatedSecret,
		Convert_federation_FederatedSecret_To_v1alpha1_FederatedSecret,
		Convert_v1alpha1_FederatedSecretList_To_federation_FederatedSecretList,
		Convert_federation_FederatedSecretList_To_v1alpha1_FederatedSecretList,
		Convert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec,
		Convert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec,
		Convert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus,
		Convert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus,
		Convert_v1alpha1_FederatedSecretStatusStrategy_To_federation_FederatedSecretStatusStrategy,
		Convert_federation_FederatedSecretStatusStrategy_To_v1alpha1_FederatedSecretStatusStrategy,
		Convert_v1alpha1_FederatedSecretStrategy_To_federation_FederatedSecretStrategy,
		Convert_federation_FederatedSecretStrategy_To_v1alpha1_FederatedSecretStrategy,
	)
}

func autoConvert_v1alpha1_FederatedSecret_To_federation_FederatedSecret(in *FederatedSecret, out *federation.FederatedSecret, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_FederatedSecret_To_federation_FederatedSecret is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecret_To_federation_FederatedSecret(in *FederatedSecret, out *federation.FederatedSecret, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecret_To_federation_FederatedSecret(in, out, s)
}

func autoConvert_federation_FederatedSecret_To_v1alpha1_FederatedSecret(in *federation.FederatedSecret, out *FederatedSecret, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_federation_FederatedSecret_To_v1alpha1_FederatedSecret is an autogenerated conversion function.
func Convert_federation_FederatedSecret_To_v1alpha1_FederatedSecret(in *federation.FederatedSecret, out *FederatedSecret, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecret_To_v1alpha1_FederatedSecret(in, out, s)
}

func autoConvert_v1alpha1_FederatedSecretList_To_federation_FederatedSecretList(in *FederatedSecretList, out *federation.FederatedSecretList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]federation.FederatedSecret)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_FederatedSecretList_To_federation_FederatedSecretList is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecretList_To_federation_FederatedSecretList(in *FederatedSecretList, out *federation.FederatedSecretList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecretList_To_federation_FederatedSecretList(in, out, s)
}

func autoConvert_federation_FederatedSecretList_To_v1alpha1_FederatedSecretList(in *federation.FederatedSecretList, out *FederatedSecretList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]FederatedSecret)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_federation_FederatedSecretList_To_v1alpha1_FederatedSecretList is an autogenerated conversion function.
func Convert_federation_FederatedSecretList_To_v1alpha1_FederatedSecretList(in *federation.FederatedSecretList, out *FederatedSecretList, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecretList_To_v1alpha1_FederatedSecretList(in, out, s)
}

func autoConvert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec(in *FederatedSecretSpec, out *federation.FederatedSecretSpec, s conversion.Scope) error {
	out.Template = in.Template
	return nil
}

// Convert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec(in *FederatedSecretSpec, out *federation.FederatedSecretSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecretSpec_To_federation_FederatedSecretSpec(in, out, s)
}

func autoConvert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec(in *federation.FederatedSecretSpec, out *FederatedSecretSpec, s conversion.Scope) error {
	out.Template = in.Template
	return nil
}

// Convert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec is an autogenerated conversion function.
func Convert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec(in *federation.FederatedSecretSpec, out *FederatedSecretSpec, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecretSpec_To_v1alpha1_FederatedSecretSpec(in, out, s)
}

func autoConvert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus(in *FederatedSecretStatus, out *federation.FederatedSecretStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus(in *FederatedSecretStatus, out *federation.FederatedSecretStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecretStatus_To_federation_FederatedSecretStatus(in, out, s)
}

func autoConvert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus(in *federation.FederatedSecretStatus, out *FederatedSecretStatus, s conversion.Scope) error {
	return nil
}

// Convert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus is an autogenerated conversion function.
func Convert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus(in *federation.FederatedSecretStatus, out *FederatedSecretStatus, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecretStatus_To_v1alpha1_FederatedSecretStatus(in, out, s)
}

func autoConvert_v1alpha1_FederatedSecretStatusStrategy_To_federation_FederatedSecretStatusStrategy(in *FederatedSecretStatusStrategy, out *federation.FederatedSecretStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_FederatedSecretStatusStrategy_To_federation_FederatedSecretStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecretStatusStrategy_To_federation_FederatedSecretStatusStrategy(in *FederatedSecretStatusStrategy, out *federation.FederatedSecretStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecretStatusStrategy_To_federation_FederatedSecretStatusStrategy(in, out, s)
}

func autoConvert_federation_FederatedSecretStatusStrategy_To_v1alpha1_FederatedSecretStatusStrategy(in *federation.FederatedSecretStatusStrategy, out *FederatedSecretStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_federation_FederatedSecretStatusStrategy_To_v1alpha1_FederatedSecretStatusStrategy is an autogenerated conversion function.
func Convert_federation_FederatedSecretStatusStrategy_To_v1alpha1_FederatedSecretStatusStrategy(in *federation.FederatedSecretStatusStrategy, out *FederatedSecretStatusStrategy, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecretStatusStrategy_To_v1alpha1_FederatedSecretStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_FederatedSecretStrategy_To_federation_FederatedSecretStrategy(in *FederatedSecretStrategy, out *federation.FederatedSecretStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_FederatedSecretStrategy_To_federation_FederatedSecretStrategy is an autogenerated conversion function.
func Convert_v1alpha1_FederatedSecretStrategy_To_federation_FederatedSecretStrategy(in *FederatedSecretStrategy, out *federation.FederatedSecretStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_FederatedSecretStrategy_To_federation_FederatedSecretStrategy(in, out, s)
}

func autoConvert_federation_FederatedSecretStrategy_To_v1alpha1_FederatedSecretStrategy(in *federation.FederatedSecretStrategy, out *FederatedSecretStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_federation_FederatedSecretStrategy_To_v1alpha1_FederatedSecretStrategy is an autogenerated conversion function.
func Convert_federation_FederatedSecretStrategy_To_v1alpha1_FederatedSecretStrategy(in *federation.FederatedSecretStrategy, out *FederatedSecretStrategy, s conversion.Scope) error {
	return autoConvert_federation_FederatedSecretStrategy_To_v1alpha1_FederatedSecretStrategy(in, out, s)
}
