package Models

import (
	"Telegraf/Config"
	"time"
)

type Reminder struct {
	ID             uint   `json:"id"`
	Description    string `json:"description"`
	RemindDate     uint64 `json:"remind_date"`
	RemindFullDate string `json:"remind_full_date"`
	ChatId         uint   `json:"chat_id"`
	UserID         uint   `json:"user_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
}

func GetWholeReminders(reminders *[]Reminder) (err error) {
	// currTime := time.Now().UnixNano() / int64(time.Millisecond)
	if err := Config.DB.Order("remind_date asc").Find(reminders).Error; err != nil {
		return err
	}
	return nil
}

func GetAllReminders(reminders *[]Reminder, user_id string) (err error) {
	currTime := time.Now().UnixNano() / int64(time.Millisecond)
	if err := Config.DB.Order("remind_date asc").Where("user_id = ? and remind_date > ?", user_id, currTime).Find(reminders).Error; err != nil {
		return err
	}
	return nil
}

func CreateAReminder(reminder *Reminder) (err error) {
	if err := Config.DB.Create(reminder).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAReminder(reminder *Reminder, id string) (err error) {
	Config.DB.Save(reminder)
	return nil
}

func DeleteAReminder(reminder *Reminder, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(reminder)
	return nil
}

func GetAReminder(reminder *Reminder, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(reminder).Error; err != nil {
		return err
	}
	return nil
}
