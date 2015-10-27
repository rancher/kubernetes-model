package model

const (
	VOLUME_TYPE = "v1.Volume"
)

type Volume struct {
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" yaml:"aws_elastic_block_store,omitempty"`

	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty" yaml:"empty_dir,omitempty"`

	GcePersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" yaml:"gce_persistent_disk,omitempty"`

	GitRepo *GitRepoVolumeSource `json:"gitRepo,omitempty" yaml:"git_repo,omitempty"`

	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" yaml:"glusterfs,omitempty"`

	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" yaml:"host_path,omitempty"`

	Iscsi *ISCSIVolumeSource `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Nfs *NFSVolumeSource `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" yaml:"persistent_volume_claim,omitempty"`

	Rbd *RBDVolumeSource `json:"rbd,omitempty" yaml:"rbd,omitempty"`

	Secret *SecretVolumeSource `json:"secret,omitempty" yaml:"secret,omitempty"`
}
