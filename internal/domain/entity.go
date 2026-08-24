package domain

// EntityUid identifies a Cedar entity: Type::"id", e.g. User::"alice".
type EntityUid struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// String renders Cedar's Type::"id" textual syntax.
func (u EntityUid) String() string {
	return u.Type + `::"` + u.ID + `"`
}

// Entity is a principal/resource/context entity with attributes and parent
// entities (used for hierarchy and group-membership checks in policies).
type Entity struct {
	UID        EntityUid              `json:"uid"`
	Attributes map[string]any         `json:"attrs,omitempty"`
	Parents    []EntityUid            `json:"parents,omitempty"`
}
