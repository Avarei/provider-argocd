// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-argocd/apis/cluster/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Application.
func (mg *Application) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination.Server),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.Destination.ServerRef,
		Selector:     mg.Spec.ForProvider.Destination.ServerSelector,
		To: reference.To{
			List:    &v1alpha1.ClusterList{},
			Managed: &v1alpha1.Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Destination.Server")
	}
	mg.Spec.ForProvider.Destination.Server = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.Destination.ServerRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Status.AtProvider.Sync.ComparedTo.Destination.Server),
		Extract:      reference.ExternalName(),
		Reference:    mg.Status.AtProvider.Sync.ComparedTo.Destination.ServerRef,
		Selector:     mg.Status.AtProvider.Sync.ComparedTo.Destination.ServerSelector,
		To: reference.To{
			List:    &v1alpha1.ClusterList{},
			Managed: &v1alpha1.Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Status.AtProvider.Sync.ComparedTo.Destination.Server")
	}
	mg.Status.AtProvider.Sync.ComparedTo.Destination.Server = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Status.AtProvider.Sync.ComparedTo.Destination.ServerRef = rsp.ResolvedReference

	return nil
}
