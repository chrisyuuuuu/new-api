package message

const (
	TEXT        string = "text"
	POST        string = "post"
	IMAGE       string = "image"
	SHARE_CHAT  string = "share_chat"
	INTERACTIVE string = "interactive"
)

const (
	PostMessageContentPostZhCnContentTagText    = "text"
	PostMessageContentPostZhCnContentTagA       = "a"
	PostMessageContentPostZhCnContentTagAt      = "at"
	PostMessageContentPostZhCnContentTagImg     = "img"
	PostMessageContentPostZhCnContentTagMedia   = "media"
	PostMessageContentPostZhCnContentTagEmotion = "emotion"
)

// TextMessage 文本消息
type TextMessage struct {
	MsgType string             `json:"msg_type"`
	Content TextMessageContent `json:"content"`
}

type TextMessageContent struct {
	Text string `json:"text"`
}

// PostMessage 富文本消息
type PostMessage struct {
	MsgType string             `json:"msg_type"`
	Content PostMessageContent `json:"content"`
}

type PostMessageContent struct {
	Post PostMessageContentPost `json:"post"`
}

type PostMessageContentPost struct {
	ZhCn PostMessageContentPostZhCn `json:"zh_cn"`
}

type PostMessageContentPostZhCn struct {
	Title   string                                `json:"title"`
	Content [][]PostMessageContentPostZhCnContent `json:"content"`
}

type PostMessageContentPostZhCnContent struct {
	Tag       string `json:"tag"`
	Text      string `json:"text,omitempty"`
	Href      string `json:"href,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	ImageKey  string `json:"image_key,omitempty"`
	FileKey   string `json:"file_key,omitempty"`
	EmojiType string `json:"emoji_type,omitempty"`
}

// ShareChatMessage 群名片
type ShareChatMessage struct {
	MsgType string                  `json:"msg_type"`
	Content ShareChatMessageContent `json:"content"`
}

type ShareChatMessageContent struct {
	ShareChatId string `json:"share_chat_id"`
}

// ImageMessage 图片
type ImageMessage struct {
	MsgType string              `json:"msg_type"`
	Content ImageMessageContent `json:"content"`
}

type ImageMessageContent struct {
	ImageKey string `json:"image_key"`
}

// InteractiveMessage 消息卡片
type InteractiveMessage struct {
	MsgType string                 `json:"msg_type"`
	Card    InteractiveMessageCard `json:"card"`
}

type InteractiveMessageCard struct {
	Elements InteractiveMessageCardElements `json:"elements"`
	Header   InteractiveMessageCardHeader   `json:"header"`
}

type InteractiveMessageCardElements []struct {
	Tag     string                                `json:"tag"`
	Text    InteractiveMessageCardElementsText    `json:"text,omitempty"`
	Actions InteractiveMessageCardElementsActions `json:"actions,omitempty"`
}

type InteractiveMessageCardElementsText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type InteractiveMessageCardElementsActions []struct {
	Tag   string                                    `json:"tag"`
	Text  InteractiveMessageCardElementsActionsText `json:"text"`
	Url   string                                    `json:"url"`
	Type  string                                    `json:"type"`
	Value struct {
	} `json:"value"`
}

type InteractiveMessageCardElementsActionsText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type InteractiveMessageCardHeader struct {
	Title    InteractiveMessageCardHeaderTitle `json:"title"`
	Template string                            `json:"template"`
}
type InteractiveMessageCardHeaderTitle struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

func NewTextMessage(text string) *TextMessage {
	return &TextMessage{
		MsgType: TEXT,
		Content: TextMessageContent{
			Text: text,
		},
	}
}

func NewPostMessageContentPostZhCnContent(tag string, text string, href string, userId string, userName string, imageKey string, fileKey string, emojiType string) *PostMessageContentPostZhCnContent {
	return &PostMessageContentPostZhCnContent{
		Tag:       tag,
		Text:      text,
		Href:      href,
		UserId:    userId,
		UserName:  userName,
		ImageKey:  imageKey,
		FileKey:   fileKey,
		EmojiType: emojiType,
	}
}

func NewPostMessage(title string, content [][]PostMessageContentPostZhCnContent) *PostMessage {
	return &PostMessage{
		MsgType: POST,
		Content: PostMessageContent{
			Post: PostMessageContentPost{
				ZhCn: PostMessageContentPostZhCn{
					Title:   title,
					Content: content,
				},
			},
		},
	}
}

func NewShareChatMessage(shareChatId string) *ShareChatMessage {
	return &ShareChatMessage{
		MsgType: SHARE_CHAT,
		Content: ShareChatMessageContent{
			ShareChatId: shareChatId,
		},
	}
}

func NewImageMessage(imageKey string) *ImageMessage {
	return &ImageMessage{
		MsgType: IMAGE,
		Content: ImageMessageContent{
			ImageKey: imageKey,
		},
	}
}

func NewInteractiveMessage(elements InteractiveMessageCardElements, header InteractiveMessageCardHeader) *InteractiveMessage {
	return &InteractiveMessage{
		MsgType: INTERACTIVE,
		Card: InteractiveMessageCard{
			Elements: elements,
			Header:   header,
		},
	}
}

// AlertCard 定义了告警卡片的内容
type AlertCard struct {
	Config struct {
		WideScreenMode bool `json:"wide_screen_mode"`
	} `json:"config"`
	Header struct {
		Title struct {
			Content string `json:"tag"`
			Text    string `json:"content"`
		} `json:"title"`
	} `json:"header"`
	Elements []interface{} `json:"elements"`
}

// NewAlertCardMessage 创建一个新的告警卡片
func NewAlertCardMessage(title, errorText, requestParams string) *AlertCard {
	card := &AlertCard{
		Config: struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		}{
			WideScreenMode: true,
		},
		Header: struct {
			Title struct {
				Content string `json:"tag"`
				Text    string `json:"content"`
			} `json:"title"`
		}{
			Title: struct {
				Content string `json:"tag"`
				Text    string `json:"content"`
			}{
				Content: "markdown",
				Text:    title,
			},
		},
		Elements: []interface{}{
			map[string]interface{}{
				"tag": "div",
				"text": map[string]string{
					"content": "错误信息: " + errorText,
				},
			},
			map[string]interface{}{
				"tag": "div",
				"text": map[string]string{
					"content": "请求参数: " + requestParams,
				},
			},
		},
	}
	return card
}
