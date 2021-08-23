package tests_test

import (
	"testing"

	. "gorm.io/gorm/utils/tests"
)

func BenchmarkCreate(b *testing.B) {
	var user = *GetUser("bench", Config{})

	for x := 0; x < b.N; x++ {
		user.ID = 0
		DB.Insert(&user)
	}
}

func BenchmarkFind(b *testing.B) {
	var user = *GetUser("find", Config{})
	DB.Insert(&user)

	for x := 0; x < b.N; x++ {
		DB.Find(&User{}, "id = ?", user.ID)
	}
}

func BenchmarkUpdate(b *testing.B) {
	var user = *GetUser("find", Config{})
	DB.Insert(&user)

	for x := 0; x < b.N; x++ {
		DB.Model(&user).Updates(map[string]interface{}{"Age": x})
	}
}

func BenchmarkDelete(b *testing.B) {
	var user = *GetUser("find", Config{})

	for x := 0; x < b.N; x++ {
		user.ID = 0
		DB.Insert(&user)
		DB.Delete(&user)
	}
}
