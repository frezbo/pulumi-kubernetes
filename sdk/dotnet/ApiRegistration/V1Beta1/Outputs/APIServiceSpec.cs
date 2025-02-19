// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiRegistration.V1Beta1
{

    /// <summary>
    /// APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
    /// </summary>
    [OutputType]
    public sealed class APIServiceSpec
    {
        /// <summary>
        /// CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
        /// </summary>
        public readonly string CaBundle;
        /// <summary>
        /// Group is the API group name this server hosts
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
        /// </summary>
        public readonly int GroupPriorityMinimum;
        /// <summary>
        /// InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.
        /// </summary>
        public readonly bool InsecureSkipTLSVerify;
        /// <summary>
        /// Service is a reference to the service for this API server.  It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiRegistration.V1Beta1.ServiceReference Service;
        /// <summary>
        /// Version is the API version this server hosts.  For example, "v1"
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA &gt; beta &gt; alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
        /// </summary>
        public readonly int VersionPriority;

        [OutputConstructor]
        private APIServiceSpec(
            string caBundle,

            string group,

            int groupPriorityMinimum,

            bool insecureSkipTLSVerify,

            Pulumi.Kubernetes.Types.Outputs.ApiRegistration.V1Beta1.ServiceReference service,

            string version,

            int versionPriority)
        {
            CaBundle = caBundle;
            Group = group;
            GroupPriorityMinimum = groupPriorityMinimum;
            InsecureSkipTLSVerify = insecureSkipTLSVerify;
            Service = service;
            Version = version;
            VersionPriority = versionPriority;
        }
    }
}
