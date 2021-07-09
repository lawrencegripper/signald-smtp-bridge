package v0

type LegacyResponse struct {
	ID   string             `json:",omitempty"`
	Data LegacyResponseData `json:",omitempty"`
	Type string             `json:",omitempty"`
}

// ResponseData is where most of the data in the response is stored.
type LegacyResponseData struct {
	Groups          []Group     `json:",omitempty"`
	Accounts        []Account   `json:",omitempty"`
	URI             string      `json:",omitempty"`
	DataMessage     DataMessage `json:",omitempty"`
	Message         string      `json:",omitempty"`
	Username        string      `json:",omitempty"`
	Source          JsonAddress `json:",omitempty"`
	SourceDevice    int         `json:",omitempty"`
	Type            string      `json:",omitempty"`
	IsReceipt       bool        `json:",omitempty"`
	Timestamp       int64       `json:",omitempty"`
	ServerTimestamp int64       `json:",omitempty"`
}

// Group represents a group in signal
type Group struct {
	GroupID  string   `json:",omitempty"`
	Members  []string `json:",omitempty"`
	Name     string   `json:",omitempty"`
	AvatarID int      `json:",omitempty"`
}

// Account represents a user account registered to signald
type Account struct {
	Username   string `json:",omitempty"`
	DeviceID   int    `json:",omitempty"`
	Filename   string `json:",omitempty"`
	Registered bool   `json:",omitempty"`
	HasKeys    bool   `json:"has_keys,omitempty"`
	Subscribed bool
	UUID       string `json:",omitempty"`
}

// DataMessage is the main component of incoming text messages
type DataMessage struct {
	Timestamp        int64               `json:",omitempty"`
	Body             string              `json:",omitempty"`
	ExpiresInSeconds int64               `json:",omitempty"`
	GroupInfo        IncomingGroupInfo   `json:"group,omitempty"`
	GroupV2          IncomingGroupV2Info `json:"groupV2,omitempty"`
}

// IncomingGroupInfo is information about a particular legacy group
type IncomingGroupInfo struct {
	GroupID string `json:",omitempty"`
	Type    string `json:",omitempty"`
}

// IncomingGroupV2Info is a stripped down copy of v1.JsonGroupV2Info because v0 can't depend on v1
type IncomingGroupV2Info struct {
	Avatar     string `json:"avatar,omitempty" yaml:"avatar,omitempty"` // path to the group's avatar on local disk, if available
	ID         string `json:"id,omitempty" yaml:"id,omitempty"`
	InviteLink string `json:"inviteLink,omitempty" yaml:"inviteLink,omitempty"` // the signal.group link, if applicable
	Revision   int32  `json:"revision,omitempty" yaml:"revision,omitempty"`
	Timer      int32  `json:"timer,omitempty" yaml:"timer,omitempty"`
	Title      string `json:"title,omitempty" yaml:"title,omitempty"`
}
