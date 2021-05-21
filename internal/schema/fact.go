package schema

// import "github.com/google/uuid"

// User provides for unmarshalling the users contained withing a Group.
type Fact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
