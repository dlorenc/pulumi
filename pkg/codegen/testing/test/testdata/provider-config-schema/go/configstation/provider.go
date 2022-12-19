// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configstation

import (
	"context"
	"reflect"

	"config"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.FavoriteColor) {
		args.FavoriteColor = pulumi.StringPtr(getEnvOrDefault("", nil, "FAVE_COLOR").(string))
	}
	if args.SecretSandwiches != nil {
		args.SecretSandwiches = pulumi.ToSecret(args.SecretSandwiches).(config.SandwichArrayInput)
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:configstation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// this is a relaxed string enum which can also be set via env var
	FavoriteColor *string `pulumi:"favoriteColor"`
	// Super duper secret sandwiches.
	SecretSandwiches []config.Sandwich `pulumi:"secretSandwiches"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// this is a relaxed string enum which can also be set via env var
	FavoriteColor pulumi.StringPtrInput
	// Super duper secret sandwiches.
	SecretSandwiches config.SandwichArrayInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
