package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerMessage_e7f8a9b7 struct {
    RequestID      string  `mapstructure:"requestId"      json:"requestId_inner1141"`
    Rate           float64 `mapstructure:"rate"           json:"rate_inner1142"`
    RateType       string  `mapstructure:"rateType"       json:"rateType_inner1143"`
    MonthlyPayment int     `mapstructure:"monthlyPayment" json:"monthlyPayment_inner1144"`
    TotalCost      int     `mapstructure:"totalCost"      json:"totalCost_inner1145"`
    EffectiveRate  float64 `mapstructure:"effectiveRate"  json:"effectiveRate_inner1146"`
}

func Process_d3333333(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID e7f8a9b7)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_e7f8a9b7
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