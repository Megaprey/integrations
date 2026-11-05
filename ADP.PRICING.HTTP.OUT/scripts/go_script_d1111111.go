package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type PricingClientOut_d6e7f8a6 struct {
    ClientID     string `mapstructure:"clientId_inner1101"     json:"clientId"`
    RiskGrade    string `mapstructure:"riskGrade_inner1102"    json:"riskGrade"`
    SalaryClient bool   `mapstructure:"salaryClient_inner1103" json:"salaryClient"`
}
type PricingTermsOut_d6e7f8a6 struct {
    Amount     int    `mapstructure:"amount_inner1111"     json:"amount"`
    Currency   string `mapstructure:"currency_inner1112"   json:"currency"`
    TermMonths int    `mapstructure:"termMonths_inner1113" json:"termMonths"`
}
type PricingScoringOut_d6e7f8a6 struct {
    Pd    float64 `mapstructure:"pd_inner1121"    json:"pd"`
    Score int     `mapstructure:"score_inner1122" json:"score"`
}
type OutMessage_d6e7f8a6 struct {
    RequestID     string                  `mapstructure:"requestId_inner1131"     json:"requestId"`
    ApplicationID string                  `mapstructure:"applicationId_inner1132" json:"applicationId"`
    Product       string                  `mapstructure:"product_inner1133"       json:"product"`
    Client        PricingClientOut_d6e7f8a6 `mapstructure:"client_inner1134"      json:"client"`
    Terms         PricingTermsOut_d6e7f8a6  `mapstructure:"terms_inner1135"       json:"terms"`
    Scoring       PricingScoringOut_d6e7f8a6 `mapstructure:"scoring_inner1136"    json:"scoring"`
}

func Process_d1111111(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID d6e7f8a6)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_d6e7f8a6
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