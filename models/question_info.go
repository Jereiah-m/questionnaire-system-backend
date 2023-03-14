package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionInfo struct {
	Id        string   `json:"_id" bson:"_id"`
	Title     string   `json:"title"`
	Survey_Id string   `json:"survey_id"`
	Type      int      `json:"type"`
	Options   []string `json:"options"`
	Flag      bool     `json:"flag"`
}

type QuestionInfoT struct {
	Id        string   `json:"_id" bson:"_id"`
	Title     string   `json:"title"`
	Survey_Id string   `json:"survey_id"`
	Type      int      `json:"type"`
	Options   []string `json:"options"`
	Flag      bool     `json:"flag"`
	Num       int      `json:"num"`
}

func (QuestionInfo) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "question_info" // 返回Users
}

func (QuestionInfoT) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "question_info" // 返回Users
}

// 通过问卷id得到问卷的所有问题
func GetQuestionInfoBySIdentity(Sidentity string) ([]*QuestionInfoT, error) {
	data := make([]*QuestionInfoT, 0) // 用于存储查询结果
	cursor, err := Mongo.Collection(QuestionInfoT{}.CollectionName()).
		Find(context.Background(), bson.M{"survey_id": Sidentity},
			&options.FindOptions{
				Sort: bson.D{{"type", 1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		questionInfoT := new(QuestionInfoT)
		err = cursor.Decode(questionInfoT)
		if err != nil {
			return nil, err
		}
		questionInfoT.Num = (questionInfoT.Type - (questionInfoT.Type)%10) / 10
		questionInfoT.Type = (questionInfoT.Type) % 10
		data = append(data, questionInfoT)
	}
	return data, nil
}

// 通过问卷id新增问题
func AddQuestion(questionInfo *QuestionInfo) error {
	_, err := Mongo.Collection(QuestionInfo{}.CollectionName()).InsertOne(context.Background(), questionInfo)
	return err
}

// 通过问题id修改部分问题信息
func UpdateQuestionInfoById(questionInfo *QuestionInfo) error {
	_, err := Mongo.Collection(QuestionInfo{}.CollectionName()).UpdateOne(context.Background(), bson.M{"_id": questionInfo.Id}, bson.M{"$set": bson.M{"title": questionInfo.Title, "type": questionInfo.Type, "options": questionInfo.Options, "flag": questionInfo.Flag}})
	return err
}

// 通过问题id删除问题
func DeleteQuestionInfoById(Qidentity string) error {
	_, err := Mongo.Collection(QuestionInfo{}.CollectionName()).DeleteOne(context.Background(), bson.M{"_id": Qidentity})
	return err
}

// 通过问题id查询问题options
func GetQuestionOptionsById(Qidentity string) ([]string, error) {
	var questionInfo *QuestionInfo
	questionInfo = new(QuestionInfo)
	err := Mongo.Collection(QuestionInfo{}.CollectionName()).FindOne(context.Background(), bson.M{"_id": Qidentity}).Decode(&questionInfo)
	return questionInfo.Options, err
}

// 通过问卷id查询所有问题id
func GetQuestionIdBySIdentity(Sidentity string) ([]string, error) {
	data := make([]string, 0) // 用于存储查询结果
	cursor, err := Mongo.Collection(QuestionInfo{}.CollectionName()).Find(context.Background(), bson.M{"survey_id": Sidentity})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		questionInfo := new(QuestionInfo)
		err = cursor.Decode(questionInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, questionInfo.Id)
	}
	return data, nil
}

// 通过问卷id查询问题id
