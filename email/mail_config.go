package email

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type MailConfigs struct {
	Address  string
	Port     int
	User     string
	Password string
	cmd      commandLineConfigs
}

//根据用户输入的一些参数，从配置文件中获取相应的值
func (mc *MailConfigs) setMailConfigs() *MailConfigs {
	mc.readConfigByCommand()
	viper.AddConfigPath("configs")
	viper.SetConfigName("mailconfig")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		} else {
			log.Fatal("Config file was found but some other errors happened", err)
		}
	}
	pathString := mc.cmd.provider + ".hosts." + mc.cmd.host
	mc.Address = viper.GetString(pathString + ".address")
	mc.Port = viper.GetInt(pathString + ".port")
	mc.User = viper.GetString(pathString + ".username")
	mc.Password = viper.GetString(pathString + ".password")
	return mc
}

type commandLineConfigs struct {
	provider string
	host     string
	from     string
	to       string
	body     string
	subject  string
}

// 从命令行读取用户输入的参数
func (mc *MailConfigs) readConfigByCommand() {
	var provierType = flag.String("provider", "QQ", "Please provide a mail provider which you want to use.")
	var host = flag.String("host", "SMTP", "Please provide a mail server type which you want to use.")
	var from = flag.String("from", "example@qq.com", "Please provide a mail sender type which you want to use.")
	var to = flag.String("to", "example@qq.com", "Please provide a mail subject which you want to use.")
	var subject = flag.String("subject", "xxxxxx", "Please provide a mail sender type which you want to use.")
	var body = flag.String("body", "xxxxxx.", "Please provide a mail contents which you want to send")

	flag.Parse()

	if *to == "" && *provierType == "" {
		log.Fatal("Must specify the from, to and providerType parameters.")
	}

	mc.cmd = commandLineConfigs{
		provider: *provierType,
		host:     *host,
		from:     *from,
		to:       *to,
		subject:  *subject,
		body:     *body,
	}

	return
}

