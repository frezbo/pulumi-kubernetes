// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Certificates.V1Beta1
{

    /// <summary>
    /// This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
    /// </summary>
    [OutputType]
    public sealed class CertificateSigningRequestSpec
    {
        /// <summary>
        /// Extra information about the requesting user. See user.Info interface for details.
        /// </summary>
        public readonly ImmutableDictionary<string, ImmutableArray<string>> Extra;
        /// <summary>
        /// Group information about the requesting user. See user.Info interface for details.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// Base64-encoded PKCS#10 CSR data
        /// </summary>
        public readonly string Request;
        /// <summary>
        /// Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
        ///  1. If it's a kubelet client certificate, it is assigned
        ///     "kubernetes.io/kube-apiserver-client-kubelet".
        ///  2. If it's a kubelet serving certificate, it is assigned
        ///     "kubernetes.io/kubelet-serving".
        ///  3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
        /// Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        /// </summary>
        public readonly string SignerName;
        /// <summary>
        /// UID information about the requesting user. See user.Info interface for details.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
        ///      https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        /// Valid values are:
        ///  "signing",
        ///  "digital signature",
        ///  "content commitment",
        ///  "key encipherment",
        ///  "key agreement",
        ///  "data encipherment",
        ///  "cert sign",
        ///  "crl sign",
        ///  "encipher only",
        ///  "decipher only",
        ///  "any",
        ///  "server auth",
        ///  "client auth",
        ///  "code signing",
        ///  "email protection",
        ///  "s/mime",
        ///  "ipsec end system",
        ///  "ipsec tunnel",
        ///  "ipsec user",
        ///  "timestamping",
        ///  "ocsp signing",
        ///  "microsoft sgc",
        ///  "netscape sgc"
        /// </summary>
        public readonly ImmutableArray<string> Usages;
        /// <summary>
        /// Information about the requesting user. See user.Info interface for details.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private CertificateSigningRequestSpec(
            ImmutableDictionary<string, ImmutableArray<string>> extra,

            ImmutableArray<string> groups,

            string request,

            string signerName,

            string uid,

            ImmutableArray<string> usages,

            string username)
        {
            Extra = extra;
            Groups = groups;
            Request = request;
            SignerName = signerName;
            Uid = uid;
            Usages = usages;
            Username = username;
        }
    }
}
