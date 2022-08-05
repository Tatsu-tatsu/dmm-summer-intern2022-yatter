package object

type (
	RelationID = int64
	// Status account
	Relation struct {
		// The internal ID of the account
		ID RelationID `json:"-"`

		// フォローしたID
		FollowerId int64 `json:"-" db:"follower_id"`

		// フォローされたID
		FolloweeId int64 `json:"-" db:"followee_id"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}

	Follow struct {
		Following bool `json:"following"`

		Followed_by bool `json:"followed_by"`
	}
)
