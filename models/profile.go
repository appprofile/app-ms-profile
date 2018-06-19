package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// ProfileCollectionName Profile collection name constant.
	ProfileCollectionName = "profile"
)

// Profile struct.
type Profile struct {
	ID          string        `bson:"_id" json:"id" validate:"required"`
	Name        string        `bson:"name" json:"name" validate:"required"`
	Address     string        `bson:"address" json:"address"`
	Phone       uint64        `bson:"phone" json:"phone" validate:"required"`
	Email       string        `bson:"email" json:"email" validate:"required"`
	Description string        `bson:"description" json:"description" validate:"required"`
	Experiences []*Experience `bson:"experiences" json:"experiences"`
	Educations  []*Education  `bson:"educations" json:"educations"`
	Abilities   []string      `bson:"abilities" json:"abilities"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}

// CollectionName Returns the collection name.
func (p *Profile) CollectionName() string {
	return ProfileCollectionName
}

// ProfileDao Profile dao.
type ProfileDao struct {
	Dao
}

// NewProfileDao Builds a ProfileDao.
func NewProfileDao() *ProfileDao {
	d := new(ProfileDao)
	return d
}

// AddExperience Add experience to profile.
func (dao *ProfileDao) AddExperience(id string, experience *Experience) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$push": bson.M{"experiences": experience}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id}, query)
	return err
}

// AddEducation Add education to profile.
func (dao *ProfileDao) AddEducation(id string, education *Education) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$push": bson.M{"educations": education}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id}, query)
	return err
}

// GetExperience Returns an experience.
func (dao *ProfileDao) GetExperience(id, experienceID string, experience *Experience) error {
	// Get the session.
	d := GetDB()
	// Execute query.
	wrapper := new(ExperienceWrapper)
	err := d.C(ProfileCollectionName).Find(bson.M{"_id": id}).
		Select(bson.M{"experiences": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(experienceID)}}}).One(wrapper)
	// Validate.
	if len(wrapper.Experiences) == 0 {
		err = fmt.Errorf("not found")
		return err
	}
	// Update reference data.
	*experience = *wrapper.Experiences[0]
	return err
}

// GetEducation Returns an education.
func (dao *ProfileDao) GetEducation(id, educationID string, education *Education) error {
	// Get the session.
	d := GetDB()
	// Execute query.
	wrapper := new(EducationWrapper)
	err := d.C(ProfileCollectionName).Find(bson.M{"_id": id}).
		Select(bson.M{"educations": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(educationID)}}}).One(wrapper)
	// Validate.
	if len(wrapper.Educations) == 0 {
		err = fmt.Errorf("not found")
		return err
	}
	// Update reference data.
	*education = *wrapper.Educations[0]
	return err
}

// RemoveExperience Remove experience.
func (dao *ProfileDao) RemoveExperience(id, experienceID string) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$pull": bson.M{"experiences": bson.M{"_id": bson.ObjectIdHex(experienceID)}}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id}, query)
	return err
}

// RemoveEducation Remove education.
func (dao *ProfileDao) RemoveEducation(id, educationID string) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$pull": bson.M{"educations": bson.M{"_id": bson.ObjectIdHex(educationID)}}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id}, query)
	return err
}

// UpdateExperience Update experience.
func (dao *ProfileDao) UpdateExperience(id string, experience *Experience) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$set": bson.M{"experiences.$": experience}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id, "experiences._id": experience.ID}, query)
	return err
}

// UpdateEducation Update education.
func (dao *ProfileDao) UpdateEducation(id string, education *Education) error {
	// Get the session.
	d := GetDB()
	// Build query.
	query := bson.M{"$set": bson.M{"educations.$": education}}
	// Execute query.
	err := d.C(ProfileCollectionName).Update(bson.M{"_id": id, "educations._id": education.ID}, query)
	return err
}
