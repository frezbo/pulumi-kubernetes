// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class
type RuntimeClass struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec RuntimeClassSpecOutput `pulumi:"spec"`
}

// NewRuntimeClass registers a new resource with the given unique name, arguments, and options.
func NewRuntimeClass(ctx *pulumi.Context,
	name string, args *RuntimeClassArgs, opts ...pulumi.ResourceOption) (*RuntimeClass, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("node.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("RuntimeClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:node.k8s.io/v1:RuntimeClass"),
		},
		{
			Type: pulumi.String("kubernetes:node.k8s.io/v1beta1:RuntimeClass"),
		},
	})
	opts = append(opts, aliases)
	var resource RuntimeClass
	err := ctx.RegisterResource("kubernetes:node.k8s.io/v1alpha1:RuntimeClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeClass gets an existing RuntimeClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeClassState, opts ...pulumi.ResourceOption) (*RuntimeClass, error) {
	var resource RuntimeClass
	err := ctx.ReadResource("kubernetes:node.k8s.io/v1alpha1:RuntimeClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeClass resources.
type runtimeClassState struct {
}

type RuntimeClassState struct {
}

func (RuntimeClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassState)(nil)).Elem()
}

type runtimeClassArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec RuntimeClassSpec `pulumi:"spec"`
}

// The set of arguments for constructing a RuntimeClass resource.
type RuntimeClassArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec RuntimeClassSpecInput
}

func (RuntimeClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassArgs)(nil)).Elem()
}

type RuntimeClassInput interface {
	pulumi.Input

	ToRuntimeClassOutput() RuntimeClassOutput
	ToRuntimeClassOutputWithContext(ctx context.Context) RuntimeClassOutput
}

func (*RuntimeClass) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClass)(nil))
}

func (i *RuntimeClass) ToRuntimeClassOutput() RuntimeClassOutput {
	return i.ToRuntimeClassOutputWithContext(context.Background())
}

func (i *RuntimeClass) ToRuntimeClassOutputWithContext(ctx context.Context) RuntimeClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassOutput)
}

func (i *RuntimeClass) ToRuntimeClassPtrOutput() RuntimeClassPtrOutput {
	return i.ToRuntimeClassPtrOutputWithContext(context.Background())
}

func (i *RuntimeClass) ToRuntimeClassPtrOutputWithContext(ctx context.Context) RuntimeClassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassPtrOutput)
}

type RuntimeClassPtrInput interface {
	pulumi.Input

	ToRuntimeClassPtrOutput() RuntimeClassPtrOutput
	ToRuntimeClassPtrOutputWithContext(ctx context.Context) RuntimeClassPtrOutput
}

type runtimeClassPtrType RuntimeClassArgs

func (*runtimeClassPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClass)(nil))
}

func (i *runtimeClassPtrType) ToRuntimeClassPtrOutput() RuntimeClassPtrOutput {
	return i.ToRuntimeClassPtrOutputWithContext(context.Background())
}

func (i *runtimeClassPtrType) ToRuntimeClassPtrOutputWithContext(ctx context.Context) RuntimeClassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassPtrOutput)
}

// RuntimeClassArrayInput is an input type that accepts RuntimeClassArray and RuntimeClassArrayOutput values.
// You can construct a concrete instance of `RuntimeClassArrayInput` via:
//
//          RuntimeClassArray{ RuntimeClassArgs{...} }
type RuntimeClassArrayInput interface {
	pulumi.Input

	ToRuntimeClassArrayOutput() RuntimeClassArrayOutput
	ToRuntimeClassArrayOutputWithContext(context.Context) RuntimeClassArrayOutput
}

type RuntimeClassArray []RuntimeClassInput

func (RuntimeClassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeClass)(nil)).Elem()
}

func (i RuntimeClassArray) ToRuntimeClassArrayOutput() RuntimeClassArrayOutput {
	return i.ToRuntimeClassArrayOutputWithContext(context.Background())
}

func (i RuntimeClassArray) ToRuntimeClassArrayOutputWithContext(ctx context.Context) RuntimeClassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassArrayOutput)
}

// RuntimeClassMapInput is an input type that accepts RuntimeClassMap and RuntimeClassMapOutput values.
// You can construct a concrete instance of `RuntimeClassMapInput` via:
//
//          RuntimeClassMap{ "key": RuntimeClassArgs{...} }
type RuntimeClassMapInput interface {
	pulumi.Input

	ToRuntimeClassMapOutput() RuntimeClassMapOutput
	ToRuntimeClassMapOutputWithContext(context.Context) RuntimeClassMapOutput
}

type RuntimeClassMap map[string]RuntimeClassInput

func (RuntimeClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeClass)(nil)).Elem()
}

func (i RuntimeClassMap) ToRuntimeClassMapOutput() RuntimeClassMapOutput {
	return i.ToRuntimeClassMapOutputWithContext(context.Background())
}

func (i RuntimeClassMap) ToRuntimeClassMapOutputWithContext(ctx context.Context) RuntimeClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassMapOutput)
}

type RuntimeClassOutput struct {
	*pulumi.OutputState
}

func (RuntimeClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClass)(nil))
}

func (o RuntimeClassOutput) ToRuntimeClassOutput() RuntimeClassOutput {
	return o
}

func (o RuntimeClassOutput) ToRuntimeClassOutputWithContext(ctx context.Context) RuntimeClassOutput {
	return o
}

func (o RuntimeClassOutput) ToRuntimeClassPtrOutput() RuntimeClassPtrOutput {
	return o.ToRuntimeClassPtrOutputWithContext(context.Background())
}

func (o RuntimeClassOutput) ToRuntimeClassPtrOutputWithContext(ctx context.Context) RuntimeClassPtrOutput {
	return o.ApplyT(func(v RuntimeClass) *RuntimeClass {
		return &v
	}).(RuntimeClassPtrOutput)
}

type RuntimeClassPtrOutput struct {
	*pulumi.OutputState
}

func (RuntimeClassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClass)(nil))
}

func (o RuntimeClassPtrOutput) ToRuntimeClassPtrOutput() RuntimeClassPtrOutput {
	return o
}

func (o RuntimeClassPtrOutput) ToRuntimeClassPtrOutputWithContext(ctx context.Context) RuntimeClassPtrOutput {
	return o
}

type RuntimeClassArrayOutput struct{ *pulumi.OutputState }

func (RuntimeClassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeClass)(nil))
}

func (o RuntimeClassArrayOutput) ToRuntimeClassArrayOutput() RuntimeClassArrayOutput {
	return o
}

func (o RuntimeClassArrayOutput) ToRuntimeClassArrayOutputWithContext(ctx context.Context) RuntimeClassArrayOutput {
	return o
}

func (o RuntimeClassArrayOutput) Index(i pulumi.IntInput) RuntimeClassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuntimeClass {
		return vs[0].([]RuntimeClass)[vs[1].(int)]
	}).(RuntimeClassOutput)
}

type RuntimeClassMapOutput struct{ *pulumi.OutputState }

func (RuntimeClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RuntimeClass)(nil))
}

func (o RuntimeClassMapOutput) ToRuntimeClassMapOutput() RuntimeClassMapOutput {
	return o
}

func (o RuntimeClassMapOutput) ToRuntimeClassMapOutputWithContext(ctx context.Context) RuntimeClassMapOutput {
	return o
}

func (o RuntimeClassMapOutput) MapIndex(k pulumi.StringInput) RuntimeClassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RuntimeClass {
		return vs[0].(map[string]RuntimeClass)[vs[1].(string)]
	}).(RuntimeClassOutput)
}

func init() {
	pulumi.RegisterOutputType(RuntimeClassOutput{})
	pulumi.RegisterOutputType(RuntimeClassPtrOutput{})
	pulumi.RegisterOutputType(RuntimeClassArrayOutput{})
	pulumi.RegisterOutputType(RuntimeClassMapOutput{})
}
