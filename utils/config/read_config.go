package config

type config struct {
	PromptPayScbId   string `yaml:"prompt_pay_scb_id"`
	PromptPayKbankId string `yaml:"prompt_pay_kbank_id"`
	TelegramToken    string `yaml:"telegram_token"`
}
