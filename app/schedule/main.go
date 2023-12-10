package schedule

import (
	"github.com/futurisen-solution/symphonic-skeleton/bootstrap"
	"github.com/fwidjaya20/symphonic/facades"
)

func RunServer() {
	bootstrap.Boot()

	facades.App().GetSchedule().Run()

	select {}
}
