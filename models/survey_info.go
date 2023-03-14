package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SurveyInfo struct {
	Id                 string `json:"_id" bson:"_id"`
	Name               string `json:"name"`
	Creator_Id         string `json:"creator_id"`
	Start_Time         string `json:"start_time"`
	End_Time           string `json:"end_time"`
	Survey_Description string `json:"survey_description"`
	Status             int    `json:"status"`
	Create_Date        string `json:"create_date"`
}

func (SurveyInfo) CollectionName() string { // 为User定义CollectionName方法，返回值为string
	return "survey_info" // 返回Users
}

// 通过用户id得到用户所创建的所有问卷
func GetSurveyInfoByCIdentity(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "status": bson.M{"$ne": 0}},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

func GetSurveyInfoByCIdentityS(Cidentity string) ([]*SurveyInfo, error) { // 通过用户id得到用户所创建的所有问卷
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "status": 0},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

// 通过用户id得到start_time小于当前时间的问卷并且status不为0
func GetSurveyInfoByCIdentityW(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "start_time": bson.M{"$gt": time.Now().Format("2006-01-02 15:04:05")}, "status": bson.M{"$lt": 10, "$ne": 0}},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

// 通过用户id得到start_time小于当前时间的问卷并且status小于10且不为0
func GetSurveyInfoByCIdentityW1(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "start_time": bson.M{"$gt": time.Now().Format("2006-01-02 15:04:05")}, "status": bson.M{"$lt": 10, "$ne": 0}},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil

}

// 通过用户id得到end_time小于当前时间的问卷并且status不为0
func GetSurveyInfoByCIdentityE(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "end_time": bson.M{"$lt": time.Now().Format("2006-01-02 15:04:05")}, "status": bson.M{"$lt": 10, "$ne": 0}},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

// 通过用户id得到start_time大于当前时间并且end_time小于当前时间的问卷并且status不为0
func GetSurveyInfoByCIdentitySWE(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "start_time": bson.M{"$lt": time.Now().Format("2006-01-02 15:04:05")}, "end_time": bson.M{"$gt": time.Now().Format("2006-01-02 15:04:05")}, "status": bson.M{"$lt": 10, "$ne": 0}},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

func GetSurveyInfoByCIdentityG(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "status": 1},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

func GetSurveyInfoByCIdentityE1(Cidentity string) ([]*SurveyInfo, error) {
	data := make([]*SurveyInfo, 0)
	cursor, err := Mongo.Collection(SurveyInfo{}.CollectionName()).
		Find(context.Background(), bson.M{"creator_id": Cidentity, "status": 2},
			&options.FindOptions{
				Sort: bson.D{{"create_date", -1}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		surveyInfo := new(SurveyInfo)
		err = cursor.Decode(surveyInfo)
		if err != nil {
			return nil, err
		}
		data = append(data, surveyInfo)
	}
	return data, nil
}

// 用户新增问卷
func AddSurveyInfo(surveyInfo *SurveyInfo) error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).InsertOne(context.Background(), surveyInfo)
	return err
}

// 通过问卷id修改问卷部分信息
func UpdateSurveyInfoById(surveyInfo *SurveyInfo) error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateOne(context.Background(), bson.M{"_id": surveyInfo.Id}, bson.M{"$set": bson.M{"name": surveyInfo.Name, "start_time": surveyInfo.Start_Time, "end_time": surveyInfo.End_Time, "survey_description": surveyInfo.Survey_Description}})
	return err
}

// 通过问卷id修改问卷状态status为0
func UpdateSurveyInfoStatusById0(survey_id string) error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateOne(context.Background(), bson.M{"_id": survey_id}, bson.M{"$set": bson.M{"status": 0}})
	return err
}

// 通过问卷id为参数修改问卷状态status为1
func UpdateSurveyInfoStatusById1(survey_id string) error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateOne(context.Background(), bson.M{"_id": survey_id}, bson.M{"$set": bson.M{"status": 1}})
	return err
}

// 通过问卷id修改问卷状态使status为11
func UpdateSurveyInfoStatusByIdD(surveyInfo *SurveyInfo) error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateOne(context.Background(), bson.M{"_id": surveyInfo.Id}, bson.M{"$set": bson.M{"status": 11}})
	return err
}

// 通过问卷id查询问卷
func GetSurveyInfoById(id string) (*SurveyInfo, error) {
	surveyInfo := new(SurveyInfo)
	err := Mongo.Collection(SurveyInfo{}.CollectionName()).FindOne(context.Background(), bson.M{"_id": id}).Decode(surveyInfo)
	return surveyInfo, err
}

func GetSurveyInfoBySID(_id string) (int64, error) {
	return Mongo.Collection(SurveyInfo{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"_id", _id}})
}

// 当问卷的end_time小于当前时间时，将问卷的状态改为2,表示问卷已结束
func UpdateSurveyInfoStatus() error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateMany(context.Background(), bson.M{"end_time": bson.M{"$lt": time.Now().Format("2006-01-02 15:04:05")}}, bson.M{"$set": bson.M{"status": 2}})
	return err
}

// 当问卷的start_time大于当前时间时，将问卷的状态改为1,表示问卷进行中
func UpdateSurveyInfoStatus1() error {
	_, err := Mongo.Collection(SurveyInfo{}.CollectionName()).UpdateMany(context.Background(), bson.M{"start_time": bson.M{"$gt": time.Now().Format("2006-01-02 15:04:05")}}, bson.M{"$set": bson.M{"status": 1}})
	return err
}

func GetSurveyInfoByCCID(creator_id string) (int64, error) {
	return Mongo.Collection(SurveyInfo{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"creator_id", creator_id}})
}

// 根据问卷id获取问卷的creator_id
func GetSurveyInfoByCid(id string) (string, error) {
	surveyInfo := new(SurveyInfo)
	err := Mongo.Collection(SurveyInfo{}.CollectionName()).FindOne(context.Background(), bson.M{"_id": id}).Decode(surveyInfo)
	return surveyInfo.Creator_Id, err
}

// 依据问卷id获取问卷状态status
func GetSurveyInfoByStatus(id string) (int, error) {
	surveyInfo := new(SurveyInfo)
	err := Mongo.Collection(SurveyInfo{}.CollectionName()).FindOne(context.Background(), bson.M{"_id": id}).Decode(surveyInfo)
	return surveyInfo.Status, err
}
