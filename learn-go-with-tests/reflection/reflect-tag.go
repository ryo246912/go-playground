package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name" db:"user_name,omitempty" validate:"required"`
	Age   int    `json:"age" db:"user_age"`
	Email string
}

func main() {
	user := User{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// reflect.TypeOf を使って、User 型の型情報を取得
	t := reflect.TypeOf(user)

	// 型が構造体であることを確認
	if t.Kind() == reflect.Struct {

		for i := 0; i < t.NumField(); i++ {
			// 各フィールドの reflect.StructField オブジェクトを取得
			field := t.Field(i)

			fmt.Printf("Field: %s (Type: %s)\n", field.Name, field.Type)

			// field.Tag は reflect.StructTag 型
			// Get() メソッドで特定のキーのタグ値を取得できる
			jsonTag := field.Tag.Get("json")
			dbTag := field.Tag.Get("db")
			validateTag := field.Tag.Get("validate")

			if jsonTag != "" {
				fmt.Printf("  - JSON Tag: \"%s\"\n", jsonTag)
			}
			if dbTag != "" {
				fmt.Printf("  - DB Tag:   \"%s\"\n", dbTag)
			}
			if validateTag != "" {
				fmt.Printf("  - Validate Tag: \"%s\"\n", validateTag)
			}

			// タグ文字列全体を取得したい場合
			if field.Tag != "" {
				fmt.Printf("  - Raw Tag String: \"%s\"\n", field.Tag)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Input is not a struct.")
	}
}
