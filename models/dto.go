package models

type (
	LinksRequest struct {
		ChatId string `json:"chatId"`
		Link   string `json:"link"`
	}
	LinksResponse struct {
		Links []string `json:"links"`
	}
)
