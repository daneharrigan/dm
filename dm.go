package dm

import (
	"reflect"
	"strings"
	"log"
)

func Find(model interface{}, primaryId string) error {
	typeOf := reflect.TypeOf(model)
	value  := reflect.Indirect(reflect.ValueOf(model))

	field  := value.FieldByName("Id")
	field.Set(reflect.ValueOf(primaryId))

	// temporary
	field  = value.FieldByName("GenderId")
	field.Set(reflect.ValueOf("gender-id"))

	tableName := getTableName(typeOf)
	//From(tableName).Where(Eq("id", primaryId))
	sql := "SELECT * FROM "+tableName+" WHERE id = $1"
	// wrap with "
	// replace all " with ""

	log.Printf("[SQL]: %s | %s", sql, primaryId)
	return nil
}

func HasOne(model, relatedModel interface{}) error {
	typeOf := reflect.TypeOf(relatedModel)
	value  := reflect.Indirect(reflect.ValueOf(model))
	relatedFieldName := getRelatedFieldName(typeOf)
	field  := value.FieldByName(relatedFieldName)

	tableName := getTableName(typeOf)
	foreignKey := field.String()

	sql := "SELECT * FROM "+tableName+" WHERE id = $1"

	log.Printf("[SQL]: %s | %s", sql, foreignKey)
	return nil
}

func HasMany(model, manyModel interface{}) error {
	typeOf := reflect.TypeOf(manyModel)
	value  := reflect.Indirect(reflect.ValueOf(model))
	field  := value.FieldByName("Id")

	tableName := getTableName(typeOf)
	foreignFieldName := getForeignFieldName(typeOf)
	foreignKey := field.String()

	sql := "SELECT * FROM "+tableName+" WHERE "+foreignFieldName+" = $1"

	log.Printf("[SQL]: %s | %s", sql, foreignKey)
	return nil
}

func getTableName(typeOf reflect.Type) string {
	structName := getStructName(typeOf)
	return strings.ToLower(structName)
}

func getForeignFieldName(typeOf reflect.Type) string {
	structName := getStructName(typeOf)
	return strings.ToLower(structName)+"_id"
}

func getRelatedFieldName(typeOf reflect.Type) string {
	return getStructName(typeOf)+"Id"
}

func getStructName(typeOf reflect.Type) string {
	values := strings.Split(typeOf.String(), ".")

	switch len(values) {
	case 1: return values[0]
	case 2: return values[1]
	}

	panic("can not identify struct name")
}
