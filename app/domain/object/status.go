package object

type (
	StatusID = int64
	// Status
	Status struct {
		// The internal ID of the status
		ID StatusID `json:"-"`

		AccountId int64 `json:"-" db:"account_id"`

		// content of status
		Content string `json:"content,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
