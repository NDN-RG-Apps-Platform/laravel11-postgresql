package seeders

import (
	"user-service/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunRoleSeeder(db *gorm.DB) {
	roles := []models.Roles{
		{Code: "ADMINISTRATOR", RoleName: "Administrator"},
		{Code: "CO_ADMINISTRATOR", RoleName: "Co-Administrator"},
		{Code: "STAFF", RoleName: "Staff"},
		{Code: "LECTURE", RoleName: "Lecture"},
		{Code: "STUDENT", RoleName: "User"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Roles{Code: role.Code}).Error; err != nil {
			logrus.Errorf("failed to seed role %v", err)
			panic(err)
		}
		logrus.Infof("Role %s successfully seeded", role.Code)
	}
}
