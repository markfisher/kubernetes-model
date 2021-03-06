/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5/typed/certificates/v1alpha1"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCertificatesV1alpha1 struct {
	*core.Fake
}

func (c *FakeCertificatesV1alpha1) CertificateSigningRequests() v1alpha1.CertificateSigningRequestInterface {
	return &FakeCertificateSigningRequests{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCertificatesV1alpha1) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}
