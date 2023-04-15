package lead

import (
	"github.com/jinzhu/gorm"
	"github.com/nwochaadim/go-crm-fiber/database"
)

var (
	db *gorm.DB
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(Lead{})
}

func GetLead(id int64) *Lead {
	var lead Lead

	db.Where("ID=?", id).Find(&lead)

	return &lead
}

func GetAllLeads() []Lead {
	var leads []Lead

	db.Find(&leads)

	return leads
}

func (l *Lead) Create() *Lead {
	db.NewRecord(l)
	db.Create(&l)
	return l
}

func (l *Lead) Update() *Lead {
	db.Save(&l)

	return l
}

func DeleteLead(id int64) Lead {
	var lead Lead
	db.Where("ID=?", id).Delete(lead)
	return lead
}
