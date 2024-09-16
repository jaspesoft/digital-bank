package mongo

import (
	"digital-bank/pkg/criteria_db/criteria"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type MongoQuery struct {
	Filter bson.D
	Sort   bson.D
	Skip   int
	Limit  int
}

type MongoConverter struct {
	criteria *criteria.Criteria
}

func NewMongoConverter() *MongoConverter {
	return &MongoConverter{}
}

func (convert *MongoConverter) PrepareSearch(criteria *criteria.Criteria) MongoQuery {
	convert.criteria = criteria

	bsonFilter := bson.D{}

	if convert.criteria.HasFilters() {
		bsonFilter = convert.generateFilter(*convert.criteria.GetFilters())
	}

	return MongoQuery{
		Filter: bsonFilter,
		Sort:   convert.generateSort(convert.criteria.GetOrdering()),
		Skip:   convert.criteria.GetOffset(),
		Limit:  convert.criteria.GetLimit(),
	}
}

func (convert *MongoConverter) generateFilter(filters []criteria.Filter) bson.D {
	bsonFilter := bson.D{}
	for _, filter := range filters {
		filter.GetOP()
		switch filter.GetOP() {
		case criteria.EQUAL:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: filter.GetValue()})
		case criteria.NOT_EQUAL:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$ne": filter.GetValue()}})
		case criteria.GT:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$gt": filter.GetValue()}})
		case criteria.GTE:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$gte": filter.GetValue()}})
		case criteria.CONTAINS:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: primitive.Regex{Pattern: filter.GetValue().(string)}})
		case criteria.IN:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$in": filter.GetValue()}})
		case criteria.NOT_IN:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$nin": filter.GetValue()}})
		case criteria.LT:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$lt": filter.GetValue()}})
		case criteria.LTE:
			bsonFilter = append(bsonFilter, bson.E{Key: filter.GetField(), Value: bson.M{"$lte": filter.GetValue()}})
		}

	}

	return bsonFilter
}

func (convert *MongoConverter) generateSort(order criteria.Order) bson.D {
	if order.IsAsc() {
		return bson.D{{order.GetField(), 1}}
	}
	return bson.D{{order.GetField(), -1}}

}

func (convert *MongoConverter) PrepareUpsert(data interface{}) (bson.M, error) {
	result := bson.M{}
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Struct {
		return result, fmt.Errorf("data must be a struct")
	}

	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("bson")
		if fieldName == "" {
			continue
		}

		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result, nil
}
