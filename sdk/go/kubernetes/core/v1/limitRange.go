// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LimitRange sets resource usage limits for each kind of resource in a Namespace.
type LimitRange struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the limits enforced. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LimitRangeSpecPtrOutput `pulumi:"spec"`
}

// NewLimitRange registers a new resource with the given unique name, arguments, and options.
func NewLimitRange(ctx *pulumi.Context,
	name string, args *LimitRangeArgs, opts ...pulumi.ResourceOption) (*LimitRange, error) {
	if args == nil {
		args = &LimitRangeArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("LimitRange")
	var resource LimitRange
	err := ctx.RegisterResource("kubernetes:core/v1:LimitRange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLimitRange gets an existing LimitRange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLimitRange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LimitRangeState, opts ...pulumi.ResourceOption) (*LimitRange, error) {
	var resource LimitRange
	err := ctx.ReadResource("kubernetes:core/v1:LimitRange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LimitRange resources.
type limitRangeState struct {
}

type LimitRangeState struct {
}

func (LimitRangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*limitRangeState)(nil)).Elem()
}

type limitRangeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the limits enforced. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *LimitRangeSpec `pulumi:"spec"`
}

// The set of arguments for constructing a LimitRange resource.
type LimitRangeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the limits enforced. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LimitRangeSpecPtrInput
}

func (LimitRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*limitRangeArgs)(nil)).Elem()
}

type LimitRangeInput interface {
	pulumi.Input

	ToLimitRangeOutput() LimitRangeOutput
	ToLimitRangeOutputWithContext(ctx context.Context) LimitRangeOutput
}

func (*LimitRange) ElementType() reflect.Type {
	return reflect.TypeOf((*LimitRange)(nil))
}

func (i *LimitRange) ToLimitRangeOutput() LimitRangeOutput {
	return i.ToLimitRangeOutputWithContext(context.Background())
}

func (i *LimitRange) ToLimitRangeOutputWithContext(ctx context.Context) LimitRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeOutput)
}

func (i *LimitRange) ToLimitRangePtrOutput() LimitRangePtrOutput {
	return i.ToLimitRangePtrOutputWithContext(context.Background())
}

func (i *LimitRange) ToLimitRangePtrOutputWithContext(ctx context.Context) LimitRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangePtrOutput)
}

type LimitRangePtrInput interface {
	pulumi.Input

	ToLimitRangePtrOutput() LimitRangePtrOutput
	ToLimitRangePtrOutputWithContext(ctx context.Context) LimitRangePtrOutput
}

type limitRangePtrType LimitRangeArgs

func (*limitRangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LimitRange)(nil))
}

func (i *limitRangePtrType) ToLimitRangePtrOutput() LimitRangePtrOutput {
	return i.ToLimitRangePtrOutputWithContext(context.Background())
}

func (i *limitRangePtrType) ToLimitRangePtrOutputWithContext(ctx context.Context) LimitRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangePtrOutput)
}

// LimitRangeArrayInput is an input type that accepts LimitRangeArray and LimitRangeArrayOutput values.
// You can construct a concrete instance of `LimitRangeArrayInput` via:
//
//          LimitRangeArray{ LimitRangeArgs{...} }
type LimitRangeArrayInput interface {
	pulumi.Input

	ToLimitRangeArrayOutput() LimitRangeArrayOutput
	ToLimitRangeArrayOutputWithContext(context.Context) LimitRangeArrayOutput
}

type LimitRangeArray []LimitRangeInput

func (LimitRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LimitRange)(nil)).Elem()
}

func (i LimitRangeArray) ToLimitRangeArrayOutput() LimitRangeArrayOutput {
	return i.ToLimitRangeArrayOutputWithContext(context.Background())
}

func (i LimitRangeArray) ToLimitRangeArrayOutputWithContext(ctx context.Context) LimitRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeArrayOutput)
}

// LimitRangeMapInput is an input type that accepts LimitRangeMap and LimitRangeMapOutput values.
// You can construct a concrete instance of `LimitRangeMapInput` via:
//
//          LimitRangeMap{ "key": LimitRangeArgs{...} }
type LimitRangeMapInput interface {
	pulumi.Input

	ToLimitRangeMapOutput() LimitRangeMapOutput
	ToLimitRangeMapOutputWithContext(context.Context) LimitRangeMapOutput
}

type LimitRangeMap map[string]LimitRangeInput

func (LimitRangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LimitRange)(nil)).Elem()
}

func (i LimitRangeMap) ToLimitRangeMapOutput() LimitRangeMapOutput {
	return i.ToLimitRangeMapOutputWithContext(context.Background())
}

func (i LimitRangeMap) ToLimitRangeMapOutputWithContext(ctx context.Context) LimitRangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeMapOutput)
}

type LimitRangeOutput struct {
	*pulumi.OutputState
}

func (LimitRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LimitRange)(nil))
}

func (o LimitRangeOutput) ToLimitRangeOutput() LimitRangeOutput {
	return o
}

func (o LimitRangeOutput) ToLimitRangeOutputWithContext(ctx context.Context) LimitRangeOutput {
	return o
}

func (o LimitRangeOutput) ToLimitRangePtrOutput() LimitRangePtrOutput {
	return o.ToLimitRangePtrOutputWithContext(context.Background())
}

func (o LimitRangeOutput) ToLimitRangePtrOutputWithContext(ctx context.Context) LimitRangePtrOutput {
	return o.ApplyT(func(v LimitRange) *LimitRange {
		return &v
	}).(LimitRangePtrOutput)
}

type LimitRangePtrOutput struct {
	*pulumi.OutputState
}

func (LimitRangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LimitRange)(nil))
}

func (o LimitRangePtrOutput) ToLimitRangePtrOutput() LimitRangePtrOutput {
	return o
}

func (o LimitRangePtrOutput) ToLimitRangePtrOutputWithContext(ctx context.Context) LimitRangePtrOutput {
	return o
}

type LimitRangeArrayOutput struct{ *pulumi.OutputState }

func (LimitRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LimitRange)(nil))
}

func (o LimitRangeArrayOutput) ToLimitRangeArrayOutput() LimitRangeArrayOutput {
	return o
}

func (o LimitRangeArrayOutput) ToLimitRangeArrayOutputWithContext(ctx context.Context) LimitRangeArrayOutput {
	return o
}

func (o LimitRangeArrayOutput) Index(i pulumi.IntInput) LimitRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LimitRange {
		return vs[0].([]LimitRange)[vs[1].(int)]
	}).(LimitRangeOutput)
}

type LimitRangeMapOutput struct{ *pulumi.OutputState }

func (LimitRangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LimitRange)(nil))
}

func (o LimitRangeMapOutput) ToLimitRangeMapOutput() LimitRangeMapOutput {
	return o
}

func (o LimitRangeMapOutput) ToLimitRangeMapOutputWithContext(ctx context.Context) LimitRangeMapOutput {
	return o
}

func (o LimitRangeMapOutput) MapIndex(k pulumi.StringInput) LimitRangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LimitRange {
		return vs[0].(map[string]LimitRange)[vs[1].(string)]
	}).(LimitRangeOutput)
}

func init() {
	pulumi.RegisterOutputType(LimitRangeOutput{})
	pulumi.RegisterOutputType(LimitRangePtrOutput{})
	pulumi.RegisterOutputType(LimitRangeArrayOutput{})
	pulumi.RegisterOutputType(LimitRangeMapOutput{})
}
