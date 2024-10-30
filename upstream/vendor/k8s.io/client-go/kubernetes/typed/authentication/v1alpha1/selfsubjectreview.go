/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "k8s.io/api/authentication/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// SelfSubjectReviewsGetter has a method to return a SelfSubjectReviewInterface.
// A group's client should implement this interface.
type SelfSubjectReviewsGetter interface {
	SelfSubjectReviews() SelfSubjectReviewInterface
}

// SelfSubjectReviewInterface has methods to work with SelfSubjectReview resources.
type SelfSubjectReviewInterface interface {
	Create(ctx context.Context, selfSubjectReview *v1alpha1.SelfSubjectReview, opts v1.CreateOptions) (*v1alpha1.SelfSubjectReview, error)
	SelfSubjectReviewExpansion
}

// selfSubjectReviews implements SelfSubjectReviewInterface
type selfSubjectReviews struct {
	*gentype.Client[*v1alpha1.SelfSubjectReview]
}

// newSelfSubjectReviews returns a SelfSubjectReviews
func newSelfSubjectReviews(c *AuthenticationV1alpha1Client) *selfSubjectReviews {
	return &selfSubjectReviews{
		gentype.NewClient[*v1alpha1.SelfSubjectReview](
			"selfsubjectreviews",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.SelfSubjectReview { return &v1alpha1.SelfSubjectReview{} }),
	}
}