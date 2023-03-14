package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Identity      string `json:"_id" bson:"_id"` // 定义Identity
	Name          string `json:"name"`
	Password      string `json:"password"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
	Register_date string `json:"register_date"`
	Status        int    `json:"status"`
}

func (User) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "users" // 返回Users
}

func GetUsers(name, password string) (*User, error) { // 定义GetUsers方法，参数为string，返回值为*User和error
	user := new(User)                                 // 定义一个User类型的user
	err := Mongo.Collection(User{}.CollectionName()). // 在MHQ数据库的Users集合中
								FindOne(context.Background(), bson.D{{"name", name}, {"password", password}}). // 查找name为name，password为password的数据
								Decode(user)                                                                   // 将查找到的数据解码到user中
	return user, err // 返回user和err
}

func GetUserByIdentity(identity string) (*User, error) { // 定义GetUserByIdentity方法，参数为string，返回值为*User和error
	user := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(user)
	return user, err
}

func GetUserByName(name string) (*User, error) {
	user := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"name", name}}).
		Decode(user)
	return user, err
}

func GetUserCountByEmail(email string) (int64, error) { // 定义GetUserCountByEmail方法，参数为string，返回值为int64和error
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

func GetUserCountByAccount(name string) (int64, error) { // 定义GetUserCountByAccount方法，参数为string，返回值为int64和error
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"name", name}})
}

// 插入数据
func InsertOneUser(user *User) error {
	_, err := Mongo.Collection(User{}.CollectionName()).InsertOne(context.Background(), user)
	return err
}

// 依据用户id更改用户密码和头像
func UpdateUserByIdentity(user *User) error {
	_, err := Mongo.Collection(User{}.CollectionName()).UpdateOne(context.Background(), bson.D{{"_id", user.Identity}}, bson.D{{"$set", bson.D{{"password", user.Password}, {"avatar", user.Avatar}}}})
	return err
}

func GetSurveyInfoByCID(_id string) (int64, error) {
	return Mongo.Collection(User{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"_id", _id}})
}

// 根据用户id获取用户邮箱
func GetUserEmailByIdentity(identity string) (string, error) {
	user := new(User)
	err := Mongo.Collection(User{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", identity}}).
		Decode(user)
	return user.Email, err
}
