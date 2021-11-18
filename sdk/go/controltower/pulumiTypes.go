// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AwsAccountSso struct {
	Email     string `pulumi:"email"`
	FirstName string `pulumi:"firstName"`
	LastName  string `pulumi:"lastName"`
}

// AwsAccountSsoInput is an input type that accepts AwsAccountSsoArgs and AwsAccountSsoOutput values.
// You can construct a concrete instance of `AwsAccountSsoInput` via:
//
//          AwsAccountSsoArgs{...}
type AwsAccountSsoInput interface {
	pulumi.Input

	ToAwsAccountSsoOutput() AwsAccountSsoOutput
	ToAwsAccountSsoOutputWithContext(context.Context) AwsAccountSsoOutput
}

type AwsAccountSsoArgs struct {
	Email     pulumi.StringInput `pulumi:"email"`
	FirstName pulumi.StringInput `pulumi:"firstName"`
	LastName  pulumi.StringInput `pulumi:"lastName"`
}

func (AwsAccountSsoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsAccountSso)(nil)).Elem()
}

func (i AwsAccountSsoArgs) ToAwsAccountSsoOutput() AwsAccountSsoOutput {
	return i.ToAwsAccountSsoOutputWithContext(context.Background())
}

func (i AwsAccountSsoArgs) ToAwsAccountSsoOutputWithContext(ctx context.Context) AwsAccountSsoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountSsoOutput)
}

func (i AwsAccountSsoArgs) ToAwsAccountSsoPtrOutput() AwsAccountSsoPtrOutput {
	return i.ToAwsAccountSsoPtrOutputWithContext(context.Background())
}

func (i AwsAccountSsoArgs) ToAwsAccountSsoPtrOutputWithContext(ctx context.Context) AwsAccountSsoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountSsoOutput).ToAwsAccountSsoPtrOutputWithContext(ctx)
}

// AwsAccountSsoPtrInput is an input type that accepts AwsAccountSsoArgs, AwsAccountSsoPtr and AwsAccountSsoPtrOutput values.
// You can construct a concrete instance of `AwsAccountSsoPtrInput` via:
//
//          AwsAccountSsoArgs{...}
//
//  or:
//
//          nil
type AwsAccountSsoPtrInput interface {
	pulumi.Input

	ToAwsAccountSsoPtrOutput() AwsAccountSsoPtrOutput
	ToAwsAccountSsoPtrOutputWithContext(context.Context) AwsAccountSsoPtrOutput
}

type awsAccountSsoPtrType AwsAccountSsoArgs

func AwsAccountSsoPtr(v *AwsAccountSsoArgs) AwsAccountSsoPtrInput {
	return (*awsAccountSsoPtrType)(v)
}

func (*awsAccountSsoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAccountSso)(nil)).Elem()
}

func (i *awsAccountSsoPtrType) ToAwsAccountSsoPtrOutput() AwsAccountSsoPtrOutput {
	return i.ToAwsAccountSsoPtrOutputWithContext(context.Background())
}

func (i *awsAccountSsoPtrType) ToAwsAccountSsoPtrOutputWithContext(ctx context.Context) AwsAccountSsoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAccountSsoPtrOutput)
}

type AwsAccountSsoOutput struct{ *pulumi.OutputState }

func (AwsAccountSsoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsAccountSso)(nil)).Elem()
}

func (o AwsAccountSsoOutput) ToAwsAccountSsoOutput() AwsAccountSsoOutput {
	return o
}

func (o AwsAccountSsoOutput) ToAwsAccountSsoOutputWithContext(ctx context.Context) AwsAccountSsoOutput {
	return o
}

func (o AwsAccountSsoOutput) ToAwsAccountSsoPtrOutput() AwsAccountSsoPtrOutput {
	return o.ToAwsAccountSsoPtrOutputWithContext(context.Background())
}

func (o AwsAccountSsoOutput) ToAwsAccountSsoPtrOutputWithContext(ctx context.Context) AwsAccountSsoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AwsAccountSso) *AwsAccountSso {
		return &v
	}).(AwsAccountSsoPtrOutput)
}

func (o AwsAccountSsoOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v AwsAccountSso) string { return v.Email }).(pulumi.StringOutput)
}

func (o AwsAccountSsoOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v AwsAccountSso) string { return v.FirstName }).(pulumi.StringOutput)
}

func (o AwsAccountSsoOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v AwsAccountSso) string { return v.LastName }).(pulumi.StringOutput)
}

type AwsAccountSsoPtrOutput struct{ *pulumi.OutputState }

func (AwsAccountSsoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAccountSso)(nil)).Elem()
}

func (o AwsAccountSsoPtrOutput) ToAwsAccountSsoPtrOutput() AwsAccountSsoPtrOutput {
	return o
}

func (o AwsAccountSsoPtrOutput) ToAwsAccountSsoPtrOutputWithContext(ctx context.Context) AwsAccountSsoPtrOutput {
	return o
}

func (o AwsAccountSsoPtrOutput) Elem() AwsAccountSsoOutput {
	return o.ApplyT(func(v *AwsAccountSso) AwsAccountSso {
		if v != nil {
			return *v
		}
		var ret AwsAccountSso
		return ret
	}).(AwsAccountSsoOutput)
}

func (o AwsAccountSsoPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsAccountSso) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o AwsAccountSsoPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsAccountSso) *string {
		if v == nil {
			return nil
		}
		return &v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o AwsAccountSsoPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsAccountSso) *string {
		if v == nil {
			return nil
		}
		return &v.LastName
	}).(pulumi.StringPtrOutput)
}

type ProviderAssumeRole struct {
	DurationSeconds   *int              `pulumi:"durationSeconds"`
	ExternalId        *string           `pulumi:"externalId"`
	Policy            *string           `pulumi:"policy"`
	PolicyArns        []string          `pulumi:"policyArns"`
	RoleArn           *string           `pulumi:"roleArn"`
	SessionName       *string           `pulumi:"sessionName"`
	Tags              map[string]string `pulumi:"tags"`
	TransitiveTagKeys []string          `pulumi:"transitiveTagKeys"`
}

// ProviderAssumeRoleInput is an input type that accepts ProviderAssumeRoleArgs and ProviderAssumeRoleOutput values.
// You can construct a concrete instance of `ProviderAssumeRoleInput` via:
//
//          ProviderAssumeRoleArgs{...}
type ProviderAssumeRoleInput interface {
	pulumi.Input

	ToProviderAssumeRoleOutput() ProviderAssumeRoleOutput
	ToProviderAssumeRoleOutputWithContext(context.Context) ProviderAssumeRoleOutput
}

type ProviderAssumeRoleArgs struct {
	DurationSeconds   pulumi.IntPtrInput      `pulumi:"durationSeconds"`
	ExternalId        pulumi.StringPtrInput   `pulumi:"externalId"`
	Policy            pulumi.StringPtrInput   `pulumi:"policy"`
	PolicyArns        pulumi.StringArrayInput `pulumi:"policyArns"`
	RoleArn           pulumi.StringPtrInput   `pulumi:"roleArn"`
	SessionName       pulumi.StringPtrInput   `pulumi:"sessionName"`
	Tags              pulumi.StringMapInput   `pulumi:"tags"`
	TransitiveTagKeys pulumi.StringArrayInput `pulumi:"transitiveTagKeys"`
}

func (ProviderAssumeRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAssumeRole)(nil)).Elem()
}

func (i ProviderAssumeRoleArgs) ToProviderAssumeRoleOutput() ProviderAssumeRoleOutput {
	return i.ToProviderAssumeRoleOutputWithContext(context.Background())
}

func (i ProviderAssumeRoleArgs) ToProviderAssumeRoleOutputWithContext(ctx context.Context) ProviderAssumeRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAssumeRoleOutput)
}

func (i ProviderAssumeRoleArgs) ToProviderAssumeRolePtrOutput() ProviderAssumeRolePtrOutput {
	return i.ToProviderAssumeRolePtrOutputWithContext(context.Background())
}

func (i ProviderAssumeRoleArgs) ToProviderAssumeRolePtrOutputWithContext(ctx context.Context) ProviderAssumeRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAssumeRoleOutput).ToProviderAssumeRolePtrOutputWithContext(ctx)
}

// ProviderAssumeRolePtrInput is an input type that accepts ProviderAssumeRoleArgs, ProviderAssumeRolePtr and ProviderAssumeRolePtrOutput values.
// You can construct a concrete instance of `ProviderAssumeRolePtrInput` via:
//
//          ProviderAssumeRoleArgs{...}
//
//  or:
//
//          nil
type ProviderAssumeRolePtrInput interface {
	pulumi.Input

	ToProviderAssumeRolePtrOutput() ProviderAssumeRolePtrOutput
	ToProviderAssumeRolePtrOutputWithContext(context.Context) ProviderAssumeRolePtrOutput
}

type providerAssumeRolePtrType ProviderAssumeRoleArgs

func ProviderAssumeRolePtr(v *ProviderAssumeRoleArgs) ProviderAssumeRolePtrInput {
	return (*providerAssumeRolePtrType)(v)
}

func (*providerAssumeRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAssumeRole)(nil)).Elem()
}

func (i *providerAssumeRolePtrType) ToProviderAssumeRolePtrOutput() ProviderAssumeRolePtrOutput {
	return i.ToProviderAssumeRolePtrOutputWithContext(context.Background())
}

func (i *providerAssumeRolePtrType) ToProviderAssumeRolePtrOutputWithContext(ctx context.Context) ProviderAssumeRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAssumeRolePtrOutput)
}

type ProviderAssumeRoleOutput struct{ *pulumi.OutputState }

func (ProviderAssumeRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAssumeRole)(nil)).Elem()
}

func (o ProviderAssumeRoleOutput) ToProviderAssumeRoleOutput() ProviderAssumeRoleOutput {
	return o
}

func (o ProviderAssumeRoleOutput) ToProviderAssumeRoleOutputWithContext(ctx context.Context) ProviderAssumeRoleOutput {
	return o
}

func (o ProviderAssumeRoleOutput) ToProviderAssumeRolePtrOutput() ProviderAssumeRolePtrOutput {
	return o.ToProviderAssumeRolePtrOutputWithContext(context.Background())
}

func (o ProviderAssumeRoleOutput) ToProviderAssumeRolePtrOutputWithContext(ctx context.Context) ProviderAssumeRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderAssumeRole) *ProviderAssumeRole {
		return &v
	}).(ProviderAssumeRolePtrOutput)
}

func (o ProviderAssumeRoleOutput) DurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProviderAssumeRole) *int { return v.DurationSeconds }).(pulumi.IntPtrOutput)
}

func (o ProviderAssumeRoleOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAssumeRole) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRoleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAssumeRole) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRoleOutput) PolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderAssumeRole) []string { return v.PolicyArns }).(pulumi.StringArrayOutput)
}

func (o ProviderAssumeRoleOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAssumeRole) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRoleOutput) SessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAssumeRole) *string { return v.SessionName }).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRoleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProviderAssumeRole) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProviderAssumeRoleOutput) TransitiveTagKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderAssumeRole) []string { return v.TransitiveTagKeys }).(pulumi.StringArrayOutput)
}

type ProviderAssumeRolePtrOutput struct{ *pulumi.OutputState }

func (ProviderAssumeRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderAssumeRole)(nil)).Elem()
}

func (o ProviderAssumeRolePtrOutput) ToProviderAssumeRolePtrOutput() ProviderAssumeRolePtrOutput {
	return o
}

func (o ProviderAssumeRolePtrOutput) ToProviderAssumeRolePtrOutputWithContext(ctx context.Context) ProviderAssumeRolePtrOutput {
	return o
}

func (o ProviderAssumeRolePtrOutput) Elem() ProviderAssumeRoleOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) ProviderAssumeRole {
		if v != nil {
			return *v
		}
		var ret ProviderAssumeRole
		return ret
	}).(ProviderAssumeRoleOutput)
}

func (o ProviderAssumeRolePtrOutput) DurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) *int {
		if v == nil {
			return nil
		}
		return v.DurationSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ProviderAssumeRolePtrOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) *string {
		if v == nil {
			return nil
		}
		return v.ExternalId
	}).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRolePtrOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) *string {
		if v == nil {
			return nil
		}
		return v.Policy
	}).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRolePtrOutput) PolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) []string {
		if v == nil {
			return nil
		}
		return v.PolicyArns
	}).(pulumi.StringArrayOutput)
}

func (o ProviderAssumeRolePtrOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) *string {
		if v == nil {
			return nil
		}
		return v.RoleArn
	}).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRolePtrOutput) SessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) *string {
		if v == nil {
			return nil
		}
		return v.SessionName
	}).(pulumi.StringPtrOutput)
}

func (o ProviderAssumeRolePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ProviderAssumeRolePtrOutput) TransitiveTagKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderAssumeRole) []string {
		if v == nil {
			return nil
		}
		return v.TransitiveTagKeys
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAccountSsoInput)(nil)).Elem(), AwsAccountSsoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAccountSsoPtrInput)(nil)).Elem(), AwsAccountSsoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAssumeRoleInput)(nil)).Elem(), ProviderAssumeRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAssumeRolePtrInput)(nil)).Elem(), ProviderAssumeRoleArgs{})
	pulumi.RegisterOutputType(AwsAccountSsoOutput{})
	pulumi.RegisterOutputType(AwsAccountSsoPtrOutput{})
	pulumi.RegisterOutputType(ProviderAssumeRoleOutput{})
	pulumi.RegisterOutputType(ProviderAssumeRolePtrOutput{})
}
