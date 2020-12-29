package structs

type SearchAllResponse struct {
	Content          []Content `json:"content"`
	Pageable         string    `json:"pageable"`
	TotalPages       int       `json:"totalPages"`
	TotalElements    int       `json:"totalElements"`
	Last             bool      `json:"last"`
	Number           int       `json:"number"`
	Sort             Sort      `json:"sort"`
	Size             int       `json:"size"`
	NumberOfElements int       `json:"numberOfElements"`
	First            bool      `json:"first"`
	Empty            bool      `json:"empty"`
}

type Content struct {
	Tenant                  string     `json:"tenant"`
	LoginName               string     `json:"loginName"`
	LoginNameAlt            string     `json:"loginNameAlt"`
	ModeratorMode           bool       `json:"moderatorMode"`
	Status                  string     `json:"status"`
	Active                  bool       `json:"active"`
	Anonymized              bool       `json:"anonymized"`
	Firstname               string     `json:"firstname"`
	Lastname                string     `json:"lastname"`
	Email                   string     `json:"email"`
	Language                string     `json:"language"`
	Timezone                string     `json:"timezone"`
	Properties              Properties `json:"properties"`
	Manager                 string     `json:"manager"`
	UpdatedID               string     `json:"updatedId"`
	ID                      string     `json:"id"`
	EntityID                EntityID   `json:"entityId"`
	Slug                    string     `json:"slug"`
	TypeName                string     `json:"typeName"`
	DisplayName             string     `json:"displayName"`
	DisplayNameInitials     string     `json:"displayNameInitials"`
	Color                   string     `json:"color"`
	Target                  Target     `json:"target"`
	ExternalWorkspaceMember bool       `json:"externalWorkspaceMember"`
	ImageUrls               ImageUrls  `json:"imageUrls"`
	Public                  bool       `json:"public"`
}

type Properties struct {
	Unknown string
}

type Sort struct {
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
	Empty    bool `json:"empty"`
}

type CreateUsersRequest struct {
	Email                string      `json:"email"`
	LoginName            string      `json:"loginName"`
	LoginNameAlt         interface{} `json:"loginNameAlt"`
	Firstname            string      `json:"firstname"`
	Lastname             string      `json:"lastname"`
	Active               bool        `json:"active"`
	Superadmin           bool        `json:"superadmin"`
	GroupIds             []string    `json:"groupIds"`
	RoleIds              []string    `json:"roleIds"`
	RemoteLogonName      interface{} `json:"remoteLogonName"`
	PersistedDisplayName interface{} `json:"persistedDisplayName"`
	Password             string      `json:"password"`
	WelcomeMail          bool        `json:"welcomeMail"`
	GeneratePassword     bool        `json:"generatePassword"`
	Language             interface{} `json:"language"`
	TemporaryPassword    bool        `json:"temporaryPassword"`
	InitialUser          bool        `json:"initialUser"`
}

type CreateUsersResponse struct {
	Tenant                  interface{} `json:"tenant"`
	LoginName               interface{} `json:"loginName"`
	LoginNameAlt            interface{} `json:"loginNameAlt"`
	ModeratorMode           bool        `json:"moderatorMode"`
	Status                  string      `json:"status"`
	Active                  bool        `json:"active"`
	Anonymized              bool        `json:"anonymized"`
	Firstname               string      `json:"firstname"`
	Lastname                string      `json:"lastname"`
	Email                   string      `json:"email"`
	Language                interface{} `json:"language"`
	Timezone                interface{} `json:"timezone"`
	Properties              interface{} `json:"properties"`
	Manager                 interface{} `json:"manager"`
	UpdatedID               interface{} `json:"updatedId"`
	ID                      interface{} `json:"id"`
	EntityID                EntityID    `json:"entityId"`
	Slug                    interface{} `json:"slug"`
	TypeName                string      `json:"typeName"`
	DisplayName             string      `json:"displayName"`
	DisplayNameInitials     string      `json:"displayNameInitials"`
	Color                   string      `json:"color"`
	Target                  Target      `json:"target"`
	ExternalWorkspaceMember bool        `json:"externalWorkspaceMember"`
	ImageUrls               ImageUrls   `json:"imageUrls"`
	Public                  bool        `json:"public"`
}

type EntityID struct {
	ID       interface{} `json:"id"`
	TypeName string      `json:"typeName"`
}

type Target struct {
	Name   string `json:"name"`
	Params Params `json:"params"`
}

type Params struct {
	ID   interface{} `json:"id"`
	Slug interface{} `json:"slug"`
}

type ImageUrls struct {
	Cover  string `json:"cover"`
	Avatar string `json:"avatar"`
}
