package config

type Bot struct {
	Token   string
	Webhook struct {
		Path    string
		BaseURL string
	}
}
