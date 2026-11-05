package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type RulesScoringOut_f2a3b4c2 struct {
    Score     int     `mapstructure:"score_inner901"     json:"score"`
    Pd        float64 `mapstructure:"pd_inner902"        json:"pd"`
    RiskGrade string  `mapstructure:"riskGrade_inner903" json:"riskGrade"`
}
type RulesContextOut_f2a3b4c2 struct {
    Scoring   RulesScoringOut_f2a3b4c2 `mapstructure:"scoring_inner911"   json:"scoring"`
    Bki       map[string]interface{}   `mapstructure:"bki_inner912"       json:"bki"`
    Client    map[string]interface{}   `mapstructure:"client_inner913"    json:"client"`
    Antifraud map[string]interface{}   `mapstructure:"antifraud_inner914" json:"antifraud"`
    Product   map[string]interface{}   `mapstructure:"product_inner915"   json:"product"`
}
type OutMessage_f2a3b4c2 struct {
    RequestID     string                 `mapstructure:"requestId_inner921"     json:"requestId"`
    ApplicationID string                 `mapstructure:"applicationId_inner922" json:"applicationId"`
    RuleSet       string                 `mapstructure:"ruleSet_inner923"       json:"ruleSet"`
    Context       RulesContextOut_f2a3b4c2 `mapstructure:"context_inner924"     json:"context"`
}

func Process_b1111111(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID f2a3b4c2)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_f2a3b4c2
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