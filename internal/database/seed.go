package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/utils"
)

func SeedDatabase() {
	seedTable("roles")
	seedTable("users")
	seedTable("user_roles")
}

func seedTable(tableName string) {
	var recordCount int64
	DB.Table(tableName).Count(&recordCount)

	if recordCount != 0 {
		return
	}

	seedFile := fmt.Sprintf("./internal/database/seeds/%s.csv", tableName)
	file, err := os.Open(seedFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error when opening file for %s seed", tableName), err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error when reading file for %s seed", tableName), err)
	}

	createRecords(tableName, records[1:])
}

func createRecords(tableName string, records [][]string) {
	switch tableName {
	case "users":
		for _, record := range records {
			user := models.User{
				FirstName:    record[0],
				LastName:     record[1],
				Email:        record[2],
				PasswordHash: utils.HashSha256String(record[3]),
			}

			DB.Create(&user)
		}
	case "roles":
		for _, record := range records {
			role := models.Role{
				Name: record[0],
			}

			DB.Create(&role)
		}
	case "user_roles":
		for _, record := range records {
			roleName := record[0]
			userEmail := record[1]

			var role models.Role
			if err := DB.Find(&role).Where("Name = ?", roleName).Error; err != nil {
				fmt.Printf("Unable to find role with name %s", roleName)
				continue
			}

			var user models.User
			if err := DB.Find(&user).Where("Email = ?", userEmail).Error; err != nil {
				fmt.Printf("Unable to user with email %s", userEmail)
				continue
			}

			userRole := models.UserRole{
				UserID: user.ID,
				RoleID: role.ID,
			}

			DB.Create(&userRole)
		}
	}
}
