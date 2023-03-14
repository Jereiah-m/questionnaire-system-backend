package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type AnswerOptionRelations struct {
	Id          string   `json:"_id" bson:"_id"` // 答案id
	Question_Id string   `json:"question_id"`    // 问题id
	Answer_Id   string   `json:"answer_id"`      //答卷标示
	Content     []string `json:"content"`        // 答案内容
}

func (AnswerOptionRelations) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "answer_option_relations"
}

func AddAnswerOptionRelations(answerOptionRelations *AnswerOptionRelations) error { // 添加答案
	_, err := Mongo.Collection(AnswerOptionRelations{}.CollectionName()).InsertOne(context.Background(), answerOptionRelations)
	return err
}

func GetAnswerOptionRelationsByAIdentity(Aidentity string) ([]*AnswerOptionRelations, error) { // 依据答卷id查询答案
	data := make([]*AnswerOptionRelations, 0)
	cursor, err := Mongo.Collection(AnswerOptionRelations{}.CollectionName()).
		Find(context.Background(), bson.M{"answer_id": Aidentity})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		answerOptionRelations := new(AnswerOptionRelations)
		err = cursor.Decode(answerOptionRelations)
		if err != nil {
			return nil, err
		}
		data = append(data, answerOptionRelations)
	}
	return data, nil
}

// 通过问题id统计答案内容的数量
func GetAnswerOptionRelationsByQIdentityAndContent(Qidentity string, content string) (int64, error) {
	count, err := Mongo.Collection(AnswerOptionRelations{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"question_id": Qidentity, "content": content}) // 统计数量
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 通过多个问题id统计多个答案内容的数量
func GetAnswerOptionRelationsByQIdentitysAndContent(Qidentitys []string, content string) (int64, error) {
	count, err := Mongo.Collection(AnswerOptionRelations{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"question_id": bson.M{"$in": Qidentitys}, "content": content}) // 统计数量
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 循环遍历数组，统计每个选项的数量
func GetAnswerOptionRelationsByQIdentitysAndContents(Qidentitys []string, contents []string) ([]int64, error) {
	counts := make([]int64, 0)
	for _, content := range contents {
		count, err := GetAnswerOptionRelationsByQIdentitysAndContent(Qidentitys, content)
		if err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}

// 通过问题id统计答案内容的数量
func GetAnswerOptionRelationsByQIdentityAndContents(Qidentity string, contents []string) ([]int64, error) {
	counts := make([]int64, 0)
	for _, content := range contents {
		count, err := GetAnswerOptionRelationsByQIdentityAndContent(Qidentity, content)
		if err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}
