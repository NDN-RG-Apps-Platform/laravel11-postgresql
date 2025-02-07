package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("failed to hash password: %v", err)
		panic(err)
	}

	// Administrator
	user := models.Users{
		UUID:        uuid.New(),
		RegNumber:   "1101223080",
		Name:        "Fayyad Asadi",
		Username:    "ayad",
		Password:    string(password),
		Email:       "fayyad@goTelkom.id",
		PhoneNumber: "08123456789",
		Photo:       "https://example.com/profile.jpg",
		RoleID:      constants.Administrator,
		LibraryID:   constants.Telkom_University_Bandung,
	}
	// cek db
	err = db.FirstOrCreate(&user, models.Users{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)

	// Co-Administrator
	user = models.Users{
		UUID:        uuid.New(),
		RegNumber:   "1101223083",
		Name:        "Aliza Nurfitrian M",
		Username:    "alizaaja",
		Password:    string(password),
		Email:       "alizaaja@goTelkom.com",
		PhoneNumber: "08123456789",
		Photo:       "https://example.com/profile.jpg",
		RoleID:      constants.Co_Administrator,
		LibraryID:   constants.Telkom_University_Bandung,
	}
	err = db.FirstOrCreate(&user, models.Users{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)

	// Staff
	user = models.Users{
		UUID:        uuid.New(),
		RegNumber:   "1101223081",
		Name:        "Made Satya Yudha",
		Username:    "madeSat",
		Password:    string(password),
		Email:       "madeSat@goTelkom.com",
		PhoneNumber: "08123456789",
		Photo:       "https://example.com/profile.jpg",
		RoleID:      constants.Staff,
		LibraryID:   constants.Telkom_University_Bandung,
	}
	err = db.FirstOrCreate(&user, models.Users{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)

	// Lecture
	user = models.Users{
		UUID:        uuid.New(),
		RegNumber:   "1101223077",
		Name:        "Pak tody",
		Username:    "TAW",
		Password:    string(password),
		Email:       "TAW@goTelkom.com",
		PhoneNumber: "08123456789",
		Photo:       "https://example.com/profile.jpg",
		RoleID:      constants.Lecture,
		LibraryID:   constants.Telkom_University_Bandung,
	}
	err = db.FirstOrCreate(&user, models.Users{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)

	// Student
	user = models.Users{
		UUID:        uuid.New(),
		RegNumber:   "1101223078",
		Name:        "Naruto Uzumaki",
		Username:    "narUzu",
		Password:    string(password),
		Email:       "narUzu@goTelkom.com",
		PhoneNumber: "08123456789",
		Photo:       "https://example.com/profile.jpg",
		RoleID:      constants.Student,
		LibraryID:   constants.Telkom_University_Bandung,
	}
	err = db.FirstOrCreate(&user, models.Users{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)
}
