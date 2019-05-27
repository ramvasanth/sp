package config

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Config", t, func() {
		Convey("it should initialize the global config", func() {
			cfg := &Config{}
			cfg.WorkerRunMode = "testing"
			Initialize(cfg)

			cfg1 := &Config{}
			cfg1.WorkerRunMode = "testing 1"
			Initialize(cfg)

			So(appConfig.WorkerRunMode, ShouldEqual, cfg.WorkerRunMode)
		})

		Convey("when all the ENVS are set ", func() {
			Convey("it should load all the ENVS", func() {
				_, err := Load("../.envtest")
				So(err, ShouldBeNil)
				os.Setenv("MYSQL_URL", "")
			})
		})
	})
}
