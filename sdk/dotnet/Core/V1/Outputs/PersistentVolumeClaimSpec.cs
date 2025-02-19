// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
    /// </summary>
    [OutputType]
    public sealed class PersistentVolumeClaimSpec
    {
        /// <summary>
        /// AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
        /// </summary>
        public readonly ImmutableArray<string> AccessModes;
        /// <summary>
        /// This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.TypedLocalObjectReference DataSource;
        /// <summary>
        /// Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceRequirements Resources;
        /// <summary>
        /// A label query over volumes to consider for binding.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;
        /// <summary>
        /// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
        /// </summary>
        public readonly string StorageClassName;
        /// <summary>
        /// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
        /// </summary>
        public readonly string VolumeMode;
        /// <summary>
        /// VolumeName is the binding reference to the PersistentVolume backing this claim.
        /// </summary>
        public readonly string VolumeName;

        [OutputConstructor]
        private PersistentVolumeClaimSpec(
            ImmutableArray<string> accessModes,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.TypedLocalObjectReference dataSource,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceRequirements resources,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector,

            string storageClassName,

            string volumeMode,

            string volumeName)
        {
            AccessModes = accessModes;
            DataSource = dataSource;
            Resources = resources;
            Selector = selector;
            StorageClassName = storageClassName;
            VolumeMode = volumeMode;
            VolumeName = volumeName;
        }
    }
}
