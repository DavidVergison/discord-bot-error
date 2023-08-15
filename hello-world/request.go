package main

type DiscordRequestDto struct {
	Type   int                   `json:"type"`
	Data   DiscordDataRequestDto `json:"data"`
	Member DiscordMemberDto      `json:"member"`
}

type DiscordDataRequestDto struct {
	Name    string      `json:"name"`
	Options DataOptions `json:"options"`
}

type DataOptions []DataOption
type DataOption struct {
	Name  string
	Value interface{}
}

type DiscordMemberDto struct {
	Nick string `json:"nick"`
}

type DiscordResponseDto struct {
	Type int                    `json:"type"` // 4 CHANNEL_MESSAGE_WITH_SOURCE
	Data DiscordDataResponseDto `json:"data"`
}
type DiscordDataResponseDto struct {
	Content string `json:"content"`
}
