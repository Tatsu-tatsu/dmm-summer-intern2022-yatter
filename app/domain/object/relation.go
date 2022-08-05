package object

type (
	RelationID = int64
	// relation information
	Relation struct {
		// The internal ID of the relation
		ID RelationID `json:"-"`

		// フォローしたID
		FollowerId int64 `json:"-" db:"follower_id"`

		// フォローされたID
		FolloweeId int64 `json:"-" db:"followee_id"`

		// The time the relation was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}

	Follow struct {
		Following bool `json:"following"`

		Followed_by bool `json:"followed_by"`
	}
)
