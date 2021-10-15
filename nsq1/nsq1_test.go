package nsq1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

var nullLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func TestSendMessage(t *testing.T) {
	config := nsq.NewConfig() // 1. 创建生产者
	producer, err := nsq.NewProducer("10.100.64.251:14152", config)
	if err != nil {
		log.Fatalln("连接失败: (127.0.0.1:4150)", err)
	}

	errPing := producer.Ping() // 2. 生产者ping
	if errPing != nil {
		log.Fatalln("无法ping通: 127.0.0.1:4150", errPing)
	}

	producer.SetLogger(nullLogger, nsq.LogLevelInfo) // 3. 设置不输出info级别的日志

	for i := 0; i < 1; i++ { // 4. 生产者发布消息
		message := "消息发送测试 " + strconv.Itoa(i+10000)
		err2 := producer.Publish("one-test", []byte(message)) // 注意one-test　对应消费者consumer.go　保持一致
		if err2 != nil {
			log.Panic("生产者推送消息失败!")
		}
		fmt.Println("--message--", message)
	}

	producer.Stop() // 5. 生产者停止执行
}

// func main() {
// 	sendMessage()
// }
type analyzeHandler struct {
}

func TestConsumerTask(t *testing.T) {
	// 1. 创建消费者
	config := nsq.NewConfig()
	q, errNewCsmr := nsq.NewConsumer("one-test", "ch-one-test", config)
	if errNewCsmr != nil {
		fmt.Printf("fail to new consumer!, topic=%s, channel=%s", "one-test", "ch-one-test")
	}

	// 2. 添加处理消息的方法
	h := &analyzeHandler{}
	q.AddHandler(h)
	// q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
	// 	fmt.Println("message: ", string(message.Body))
	// 	// message.Finish()
	// 	return fmt.Errorf("---error--")
	// }))

	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	lookupAddr := []string{
		"http://127.0.0.1:14161",
	}
	err := q.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Panic("[ConnectToNSQLookupds] Could not find nsqd!")
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("-----")

	}

	// 4. 接收消费者停止通知
	// <-q.StopChan

	// 5. 获取统计结果
	// stats := q.Stats()
	// fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
	// 	stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

func (h *analyzeHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("message: ", string(msg.Body))
	// message.Finish()
	return fmt.Errorf("---error--")
}
