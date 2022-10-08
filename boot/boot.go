package boot

import (
	"aliyunoss/pkg/database"
	"aliyunoss/pkg/logger"
	"aliyunoss/pkg/oss"
	"aliyunoss/pkg/viperlib"
	"github.com/spf13/viper"
)

func Boot() {
	viperlib.InitViper(".", "config.yml")

	logger.InitLogger(viper.GetString("logger.logPath"),
		viper.GetInt("logger.maxSize"),
		viper.GetInt("logger.maxBackups"),
		viper.GetInt("logger.maxAge"),
		viper.GetBool("logger.compress"),
		viper.GetString("logger.level"),
	)

	oss.InitOss(viper.GetString("oss.endpoint"),
		viper.GetString("oss.accessKeyID"),
		viper.GetString("oss.accessKeySecret"),
		viper.GetString("oss.BucketName"))

	database.InitDatabase(viper.GetString("database.dbName"))

}
