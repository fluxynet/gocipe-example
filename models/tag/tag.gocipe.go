// generated by gocipe 2cdecd62f8cc1e3e1b51a96e61f1519e5d31bc3eb806037239eeb7707acf58b5; DO NOT EDIT

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
