package Routes

import (
	Controllers "Telegraf/Controllers/Reminder"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("user/:id/reminders", Controllers.GetAllReminders)
		api.GET("reminders", Controllers.GetWholeReminders)
		api.POST("reminders", Controllers.CreateAReminder)
		api.DELETE("reminder/:id", Controllers.DeleteAReminder)
		api.PUT("reminder/:id", Controllers.UpdateAReminder)
	}
	return r
}
