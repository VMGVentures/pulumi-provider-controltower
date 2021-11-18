// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Controltower
{
    /// <summary>
    /// The provider type for the controltower package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [ControltowerResourceType("pulumi:providers:controltower")]
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment
        /// variable, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Output("accessKey")]
        public Output<string?> AccessKey { get; private set; } = null!;

        /// <summary>
        /// This is the AWS profile name as set in the shared credentials file.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment
        /// variables, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment
        /// variable, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Output("secretKey")]
        public Output<string?> SecretKey { get; private set; } = null!;

        /// <summary>
        /// This is the path to the shared credentials file. If this is not set and a profile is specified, `~/.aws/credentials`
        /// will be used.
        /// </summary>
        [Output("sharedCredentialsFile")]
        public Output<string?> SharedCredentialsFile { get; private set; } = null!;

        /// <summary>
        /// Session token for validating temporary credentials. Typically provided after successful identity federation or
        /// Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit
        /// MFA code used to get temporary credentials. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("controltower", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is the AWS access key. It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment
        /// variable, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Settings for making use of the AWS Assume Role functionality.
        /// </summary>
        [Input("assumeRole", json: true)]
        public Input<Inputs.ProviderAssumeRoleArgs>? AssumeRole { get; set; }

        /// <summary>
        /// This is the maximum number of times an API call is retried, in the case where requests are being throttled or
        /// experiencing transient failures. The delay between the subsequent API calls increases exponentially. If omitted, the
        /// default value is `25`.
        /// </summary>
        [Input("maxRetries", json: true)]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// This is the AWS profile name as set in the shared credentials file.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// This is the AWS region. It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment
        /// variables, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// This is the AWS secret key. It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment
        /// variable, or via a shared credentials file if `profile` is specified.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        /// <summary>
        /// This is the path to the shared credentials file. If this is not set and a profile is specified, `~/.aws/credentials`
        /// will be used.
        /// </summary>
        [Input("sharedCredentialsFile")]
        public Input<string>? SharedCredentialsFile { get; set; }

        /// <summary>
        /// Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or
        /// implemented.
        /// </summary>
        [Input("skipCredentialsValidation", json: true)]
        public Input<bool>? SkipCredentialsValidation { get; set; }

        /// <summary>
        /// Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to
        /// `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods
        /// like static credentials, configuration variables, or environment variables.
        /// </summary>
        [Input("skipMetadataApiCheck", json: true)]
        public Input<bool>? SkipMetadataApiCheck { get; set; }

        /// <summary>
        /// Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.
        /// </summary>
        [Input("skipRequestingAccountId", json: true)]
        public Input<bool>? SkipRequestingAccountId { get; set; }

        /// <summary>
        /// Session token for validating temporary credentials. Typically provided after successful identity federation or
        /// Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit
        /// MFA code used to get temporary credentials. It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ProviderArgs()
        {
        }
    }
}
