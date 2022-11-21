package config

import (
	"os"
)

var C = new(config)

func init() {
	// * Load YAML configuration
	// yml, err := ioutil.ReadFile("config.yml")
	// if err != nil {
	// 	panic("UNABLE TO READ YAML CONFIGURATION FILE")
	// }
	// err = yaml.Unmarshal(yml, C)
	// if err != nil {
	// 	panic("UNABLE TO PARSE YAML CONFIGURATION FILE")
	// }
	C.PromptPayKbankId = os.Getenv("prompt_pay_kbank_id")
	C.PromptPayScbId = os.Getenv("prompt_pay_scb_id")
	C.TelegramToken = os.Getenv("telegram_token")
}
