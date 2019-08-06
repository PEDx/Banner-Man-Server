package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	validator "gopkg.in/go-playground/validator.v9"
)

// The User holds
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	Username    string             `bson:"username" json:"username" validate:"min=1,max=32"`
	Password    string             `bson:"password" json:"password" validate:"min=5,max=128"`
	Email       string             `bson:"email" json:"email"`
	Role        string             `bson:"role" json:"role"`
	IsGroupUser bool               `bson:"is_group_user" json:"is_group_user"` // 是否是组用户
	MembersID   []*string          `bson:"members_id" json:"members_id"`
	Phone       string             `bson:"phone" json:"phone"`
	OwnWidgets  []*Widget          `bson:"own_widgets" json:"own_widgets"`
	Created     time.Time          `bson:"created" json:"created"`
	Updated     time.Time          `bson:"updated" json:"updated"`
}

type Widget struct {
	WidgetName string `bson:"widget_name" json:"widget_name"`
	Name       string `bson:"name" json:"name"`
	GroupName  string `bson:"group_name" json:"group_name"`
}

func (u *User) New() *User {
	return &User{
		ID:          primitive.NewObjectID(),
		Name:        u.Name,
		Username:    u.Username,
		Email:       u.Email,
		Avatar:      u.Avatar,
		Password:    u.Password,
		Phone:       u.Phone,
		Role:        u.Role,
		IsGroupUser: u.IsGroupUser,
		MembersID:   u.MembersID,
		OwnWidgets:  u.OwnWidgets,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}

func (u *User) CreateUser() error {

	if _, err := DB.Self.Collection("User").InsertOne(context.Background(), u); err != nil {
		return err
	}
	return nil
}
func (u *User) DeleteUserByID(id primitive.ObjectID) error {

	if _, err := DB.Self.Collection("User").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}}); err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByIDs(ids *[]primitive.ObjectID) error {

	if _, err := DB.Self.Collection("User").Find(context.Background(), bson.D{{
		Key: "_id",
		Value: bson.D{{
			Key:   "$in",
			Value: ids,
		}},
	}}); err != nil {
		return err
	}
	return nil
}
func (u *User) GetUserByUsername(username string) error {

	if _, err := DB.Self.Collection("User").Find(context.Background(), bson.D{{
		Key: "username",
		Value: bson.D{{
			Key:   "$in",
			Value: username,
		}},
	}}); err != nil {
		return err
	}
	return nil
}
func (u *User) GetUserList(limit int, page int) error {
	return nil
}

func (u *User) UpdateUser() *User {
	result := DB.Self.Collection("User").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: u.ID}},
			u,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return u
	}
	return nil
}

// Validate the fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
