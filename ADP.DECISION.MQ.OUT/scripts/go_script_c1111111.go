package main

import (
    "encoding/json"
    "log"
    "github.com/go-viper/mapstructure/v2"
)

type DecProductOut_b4c5d6e4 struct {
    Type       string `mapstructure:"type_inner1001"       json:"type"`
    Amount     int    `mapstructure:"amount_inner1002"     json:"amount"`
    Currency   string `mapstructure:"currency_inner1003"   json:"currency"`
    TermMonths int    `mapstructure:"termMonths_inner1004" json:"termMonths"`
}
type DecAggregatedOut_b4c5d6e4 struct {
    Kyc       map[string]interface{} `mapstructure:"kyc_inner1011"       json:"kyc"`
    Bki       map[string]interface{} `mapstructure:"bki_inner1012"       json:"bki"`
    Antifraud map[string]interface{} `mapstructure:"antifraud_inner1013" json:"antifraud"`
    Scoring   map[string]interface{} `mapstructure:"scoring_inner1014"   json:"scoring"`
    Rules     map[string]interface{} `mapstructure:"rules_inner1015"     json:"rules"`
}
type OutMessage_b4c5d6e4 struct {
    RequestID         string                    `mapstructure:"requestId_inner1021"         json:"requestId"`
    ApplicationID     string                    `mapstructure:"applicationId_inner1022"     json:"applicationId"`
    Product           DecProductOut_b4c5d6e4    `mapstructure:"product_inner1023"           json:"product"`
    AggregatedResults DecAggregatedOut_b4c5d6e4 `mapstructure:"aggregatedResults_inner1024" json:"aggregatedResults"`
}

func Process_b4c5d6e4(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID b4c5d6e4)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_b4c5d6e4
    if err := mapstructure.Decode(raw, &out); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(out)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = out
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
}

func sendOrWarn(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}