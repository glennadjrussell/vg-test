package v1

type AuthBackend struct {
	// Name is a human-readable descriptor for this backend
	Name string `json:"name"`

	// Path refers to the mount point for the underlying backend
	Path string `json:"path"`

	// Type refers to the auth backend type that will be mounted
	Type string `json:"type"`

	AppRoleAuth AppRole `json:"approle,omitempty"`
}

type SecretBackend struct {
	// Name is a human-readable descriptor for this backend
	Name string `json:"name"`

	// Path refers to the mount point for the underlying backend
	Path string `json:"path"`

	// Type refers to the auth backend type that will be mounted
	Type string `json:"type"`
}

type AppRole struct {
	RoleName        string `json:"role_name"`
	SecretIdTtl     string `json:"secret_id_ttl,omitempty"`
	TokenNumUses    *int32 `json:"token_num_uses,omitempty"`
	TokenTtl        string `json:"token_ttl,omitempty"`
	TokenMaxTtl     string `json:"token_max_ttl,omitempty"`
	SecretIdNumUses *int32 `json:"secret_id_num_uses,omitempty"`
}

type RoleType string

const (
	IAMRole RoleType = "iam"
	GCERole RoleType = "gce"
)

type GcpRole struct {
	Type                 RoleType `json:"type"`
	Policies             []string `json:"policies"`
	BoundProjects        []string `json:"bound_projects,omitempty"`
	BoundZones           []string `json:"bound_zones,omitempty"`
	BoundLabels          []string `json:"bound_labels,omitempty"`
	BoundServiceAccounts []string `json:"bound_service_accounts,omitempty"`
}
