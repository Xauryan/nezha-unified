package app

import (
	"github.com/pocketbase/pocketbase/tools/cron"
	"time"
	_ "time/tzdata"
)

// The scheduler belongs to the current app, avoiding a stale global scheduler
// after an embedded engine is shut down and reopened in the same process.
func GetScheduler() *cron.Cron {
	scheduler := GetApp().Cron()
	scheduler.SetTimezone(time.Local)
	return scheduler
}
