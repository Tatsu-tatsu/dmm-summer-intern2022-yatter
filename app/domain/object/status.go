package object

type (
	StatusID = int64
	// Status account
	Status struct {
		// The internal ID of the account
		ID StatusID `json:"-"`

		AccountId int64 `json:"-" db:"account_id"`

		// URL to the avatar image
		Content string `json:"content,omitempty"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
