// generated by gocipe 98bee01dcbcfef2fa8be1efb59ffb2d095d1ca8b8d0200a0c99f31ad111a69db; DO NOT EDIT

package tag

// Tag Tags can be used to categories countries
type Tag struct {
	ID        *string  `json:"id"`
	Name      *string  `json:""`
	Countries []string `json:""`
}

// New returns an instance of Tag
func New() *Tag {
	return &Tag{
		ID:   new(string),
		Name: new(string),
	}
}