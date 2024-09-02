package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"one-api/feishu/message"
	"os"
	"strings"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

type Client struct {
	token   string
	keyWork string // 自定义关键词
}

func NewFeiShuClient(token string, keyWork string) *Client {
	return &Client{
		token:   token,
		keyWork: keyWork,
	}
}

func (client *Client) Send(msg any) error {
	url := os.Getenv("WEBHOOK_URL")

	messageContent, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(messageContent))
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	request.Header.Add("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if response.StatusCode != http.StatusOK {
		errMessageResponse := &message.ErrMessageResponse{}
		err := json.Unmarshal(body, errMessageResponse)
		if err != nil {
			return err
		}
		return errors.New(errMessageResponse.Msg)
	}
	return nil
}

func (client *Client) SendTextMessage(text string) error {
	text = text + client.keyWork
	message := message.NewTextMessage(text)
	return client.Send(message)
}

func (client *Client) SendPostMessage(title string, content [][]message.PostMessageContentPostZhCnContent) error {
	title = title + client.keyWork
	message := message.NewPostMessage(title, content)
	return client.Send(message)
}

func (client *Client) SendImageMessage(imageKey string) error {
	message := message.NewImageMessage(imageKey)
	return client.Send(message)
}

func (client *Client) SendShareChatMessage(shareChatId string) error {
	message := message.NewShareChatMessage(shareChatId)
	return client.Send(message)
}

func (client *Client) SendInteractiveMessage(elements message.InteractiveMessageCardElements, header message.InteractiveMessageCardHeader) error {
	message := message.NewInteractiveMessage(elements, header)
	return client.Send(message)
}
