package models

type User struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`
	FName string `bson:"f_name" json:"f_name"`
	LName string `bson:"l_name" json:"l_name"`
}
