package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerMatchedRule_a3b4c5d3 struct {
    RuleID string `mapstructure:"ruleId" json:"ruleId_inner931"`
    Result string `mapstructure:"result" json:"result_inner932"`
}
type InnerRulesLimits_a3b4c5d3 struct {
    MaxAmount     int `mapstructure:"maxAmount"     json:"maxAmount_inner941"`
    MaxTermMonths int `mapstructure:"maxTermMonths" json:"maxTermMonths_inner942"`
}
type InnerMessage_a3b4c5d3 struct {
    RequestID      string                    `mapstructure:"requestId"      json:"requestId_inner951"`
    Status         string                    `mapstructure:"status"         json:"status_inner952"`
    PolicyDecision string                    `mapstructure:"policyDecision" json:"policyDecision_inner953"`
    MatchedRules   []InnerMatchedRule_a3b4c5d3 `mapstructure:"matchedRules"   json:"matchedRules_inner954"`
    Limits         InnerRulesLimits_a3b4c5d3  `mapstructure:"limits"         json:"limits_inner955"`
}

func Process_a3b4c5d3(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID a3b4c5d3)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_a3b4c5d3
    if err := mapstructure.Decode(raw, &inner); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(inner)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = inner
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