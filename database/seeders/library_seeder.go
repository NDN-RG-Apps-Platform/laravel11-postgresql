package seeders

import (
	"user-service/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunLibrarySeeder(db *gorm.DB) {
	libraries := []models.Libraries{
		{Code: "TELU_BANDUNG", LibraryName: "TelU Bandung"},
		{Code: "TELU_SURABAYA", LibraryName: "TelU Surabaya"},
		{Code: "TELU_JAKARTA", LibraryName: "TelU Jakarta"},
		{Code: "TELU_PURWOKERTO", LibraryName: "TelU Purwokerto"},
	}

	for _, library := range libraries {
		if err := db.FirstOrCreate(&library, models.Libraries{Code: library.Code}).Error; err != nil {
			logrus.Errorf("failed to seed library %v", err)
			panic(err)
		}
		logrus.Infof("Library %s successfully seeded", library.Code)
	}

}
