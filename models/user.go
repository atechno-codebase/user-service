package models

const USER_COLLECTION = "user"

type Verification struct {
	IsVerified bool
	Code       string
}

type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`

	Institute   string `json:"institute" bson:"institute"`
	Designation string `json:"designation" bson:"designation"`
	Privilege   int    `json:"privilege" bson:"privilege"`

	MailSent int64 `json:"mailSentAt" bson:"mailSentAt"`

	Verification Verification `json:"verification" bson:"verification"`
	Faulty       []string     `json:"faulty" bson:"faulty"`

	CreatedBy string `json:"createdBy" bson:"createdBy"`
	CreatedOn int64  `json:"createdOn" bson:"createdOn"`
}
