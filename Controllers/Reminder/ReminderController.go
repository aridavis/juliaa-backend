package Controllers

import (
	Models "Telegraf/Models/Reminder"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWholeReminders(c *gin.Context) {
	var reminders []Models.Reminder
	err := Models.GetWholeReminders(&reminders)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reminders)
	}
}

func GetAllReminders(c *gin.Context) {
	id := c.Params.ByName("id")
	var reminders []Models.Reminder
	err := Models.GetAllReminders(&reminders, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reminders)
	}
}

func CreateAReminder(c *gin.Context) {
	var reminder Models.Reminder
	c.BindJSON(&reminder)
	err := Models.CreateAReminder(&reminder)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reminder)
	}
}

func UpdateAReminder(c *gin.Context) {
	var reminder Models.Reminder
	id := c.Params.ByName("id")
	err := Models.GetAReminder(&reminder, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reminder)
	}
	c.BindJSON(&reminder)
	err = Models.UpdateAReminder(&reminder, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reminder)
	}
	c.BindJSON(&reminder)
}

func DeleteAReminder(c *gin.Context) {
	var reminder Models.Reminder
	id := c.Params.ByName("id")
	err := Models.DeleteAReminder(&reminder, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
