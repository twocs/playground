package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type Custom struct {
	ID uint
}

func (c *Custom) TableName() string {
	return "custom_table_name"
}

func TestGetTableName(t *testing.T) {
	tests := []struct {
		name        string
		tableStruct any
		expected    string
	}{
		{
			name:        "users",
			tableStruct: User{Name: "jinzhu"},
			expected:    "users",
		},
		{
			name:        "users with pointer",
			tableStruct: &User{Name: "jinzhu"},
			expected:    "users",
		},
		{
			name:        "custom",
			tableStruct: Custom{ID: 12345},
			expected:    "custom_table_name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tableSchema := GetSchema(tt.tableStruct)
			tableName := tableSchema.Table
			if tableName != tt.expected {
				t.Errorf("Expected: '%s', got: '%s'", tt.expected, tableName)
			}
		})
	}

}
