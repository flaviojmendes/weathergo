package config

import (
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	convey.Convey("get configuration", t, func() {
		dir,_ := os.Getwd()
		configuration := readConfigFile( dir+ "/../config_sample.yml")

		convey.So(configuration.DebugPort, convey.ShouldEqual, ":8000")
	})
}


func TestGetConfigWithoutFile(t *testing.T) {
	convey.Convey("get configuration", t, func() {
		dir,_ := os.Getwd()
		convey.ShouldPanicWith(func(){readConfigFile( dir+ "/../config_sampleeee.yml")}, "The file " + dir + "/../config_sampleeee.yml " +  "doesn't exist.")
	})
}
