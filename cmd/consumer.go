package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wptest/configs"
	"github.com/wptest/internal/repositories"
	"github.com/wptest/internal/services"
	"github.com/wptest/pkg/kafka"
	"github.com/wptest/pkg/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	messagingConsumer     = "device-consumer"
	messagingConsumerInfo = `
Name     : %s
Topic    : %s
Group ID : %s
------------------------------------------------------------------------------
`
)

type msgConsumer struct {
	stop <-chan bool

	BaseCmd       *cobra.Command
	filename      string
	config        *configs.Config
	DeviceService services.IDeviceService
}

func NewConsumerCmd(
	configuration *configs.Config,
) *msgConsumer {
	return NewConsumerCmdSignaled(configuration, nil)
}

func NewConsumerCmdSignaled(
	configuration *configs.Config,
	stop <-chan bool,
) *msgConsumer {
	cc := &msgConsumer{stop: stop}
	cc.config = configuration
	cc.BaseCmd = &cobra.Command{
		Use:   "consumer",
		Short: "Used to run the http service",
		Run:   cc.RunConsumerMessaging,
	}
	fs := pflag.NewFlagSet("Root", pflag.ContinueOnError)
	fs.StringVarP(&cc.filename, "file", "f", "", "Custom configuration filename")
	cc.BaseCmd.Flags().AddFlagSet(fs)
	return cc
}

func NewMessagingConsumer() *msgConsumer {
	return &msgConsumer{}
}

func (mc *msgConsumer) initService() {
	if len(mc.filename) > 1 {
		mc.config = configs.New(mc.filename,
			"./configs",
			"../configs",
			"../../configs",
			"../../../configs")
	}

	conn := fmt.Sprintf(
		mysql.MysqlDataSourceFormat,
		mc.config.MariaDB.User,
		mc.config.MariaDB.Password,
		mc.config.MariaDB.Host,
		mc.config.MariaDB.Port,
		mc.config.MariaDB.DbName,
		mc.config.MariaDB.Charset,
	)

	db := mysql.NewMySQL()
	db.OpenConnection(conn, mc.config)
	db.SetConnMaxLifetime(mc.config.MariaDB.MaxLifeTime)
	db.SetMaxIdleConn(mc.config.MariaDB.MaxIdleConnection)
	db.SetMaxOpenConn(mc.config.MariaDB.MaxOpenConnection)

	// init kafka
	// Set kafka
	kafkaAddrs := strings.Split(mc.config.Kafka.BrokerList, ",")
	k := kafka.NewKafka(kafkaAddrs) // Set kafka

	// init service
	messageRepo := repositories.NewMessageRepository(*k, mc.config)
	deviceRepo := repositories.NewDeviceRepository(db, mc.config)
	deviceService := services.NewDeviceService(deviceRepo, messageRepo)

	mc.DeviceService = deviceService
}

func (mc *msgConsumer) kafkaMessageHandler(key []byte, msg []byte) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	err := mc.DeviceService.ReceiveDevice(msg)
	if err != nil {
		log.Println(err)
	}
}

func (buc *msgConsumer) kafkaErrorHandler(err error) {
	log.Println(err)
}

func (buc *msgConsumer) kafkaNotificationHandler(notification interface{}) {
	log.Println(notification)
}

func (buc *msgConsumer) printInfo(topics []string, groupID string) {
	fmt.Printf(messagingConsumerInfo, messagingConsumer, topics, groupID)
}

// Run Consumer Message
func (mc *msgConsumer) RunConsumerMessaging(cmd *cobra.Command, args []string) {

	mc.initService()

	var (
		topics  = []string{mc.config.Kafka.MessagingConsumer.Topic}
		groupID = mc.config.Kafka.MessagingConsumer.Group
		broker  = strings.Split(mc.config.Kafka.BrokerList, ",")
	)

	mc.printInfo(topics, groupID)
	done := make(chan bool)
	k := kafka.NewKafka(broker)

	err := k.Consume(groupID, topics, mc.kafkaMessageHandler, mc.kafkaErrorHandler, mc.kafkaNotificationHandler, done)
	if err != nil {
		log.Println(err)
	}

	log.Println("shutting down")
	os.Exit(0)
}
