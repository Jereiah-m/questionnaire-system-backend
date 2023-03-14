package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnswerInfo struct {
	Id         string `json:"_id" bson:"_id"`
	UserIp     string `json:"user_ip"`
	Survey_Id  string `json:"survey_id"`
	CreateDate string `json:"create_date"`
}

func (AnswerInfo) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "answer_info"
}

func AddAnswerSurvey(answerSurvey *AnswerInfo) error { // 新增答卷
	_, err := Mongo.Collection(AnswerInfo{}.CollectionName()).InsertOne(context.Background(), answerSurvey)
	return err
}

func GetAnswerInfoBySIdentity(Sidentity string) ([]*AnswerInfo, error) { // 依据问卷id查询答卷
	data := make([]*AnswerInfo, 0)
	cursor, err := Mongo.Collection(AnswerInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"survey_id": Sidentity},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		answerInfo := new(AnswerInfo)
		err = cursor.Decode(answerInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, answerInfo)
	}
	return data, nil
}
