package main

import (
    "encoding/json"
    "log"
    "github.com/go-viper/mapstructure/v2"
)

type InnerDecApprovedTerms_c5d6e7f5 struct {
    Amount     int    `mapstructure:"amount"     json:"amount_inner1031"`
    Currency   string `mapstructure:"currency"   json:"currency_inner1032"`
    TermMonths int    `mapstructure:"termMonths" json:"termMonths_inner1033"`
}
type InnerDecNextAction_c5d6e7f5 struct {
    Action   string `mapstructure:"action"   json:"action_inner1041"`
    Required bool   `mapstructure:"required" json:"required_inner1042"`
}
type InnerMessage_c5d6e7f5 struct {
    RequestID            string                     `mapstructure:"requestId"            json:"requestId_inner1051"`
    ApplicationID        string                     `mapstructure:"applicationId"        json:"applicationId_inner1052"`
    Status               string                     `mapstructure:"status"               json:"status_inner1053"`
    Decision             string                     `mapstructure:"decision"             json:"decision_inner1054"`
    DecidedAt            string                     `mapstructure:"decidedAt"            json:"decidedAt_inner1055"`
    ApprovedTerms        InnerDecApprovedTerms_c5d6e7f5 `mapstructure:"approvedTerms"        json:"approvedTerms_inner1056"`
    ReasonCodes          []string                   `mapstructure:"reasonCodes"          json:"reasonCodes_inner1057"`
    RequiresManualReview bool                       `mapstructure:"requiresManualReview" json:"requiresManualReview_inner1058"`
    NextActions          []InnerDecNextAction_c5d6e7f5  `mapstructure:"nextActions"          json:"nextActions_inner1059"`
}

func Process_c3333333(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID c5d6e7f5)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_c5d6e7f5
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