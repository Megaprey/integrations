package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerReasonCode_e1f2a3b1 struct {
    Code        string `mapstructure:"code"        json:"code_inner821"`
    Description string `mapstructure:"description" json:"description_inner822"`
}
type InnerMessage_e1f2a3b1 struct {
    RequestID            string                   `mapstructure:"requestId"            json:"requestId_inner831"`
    ApplicationID        string                   `mapstructure:"applicationId"        json:"applicationId_inner832"`
    Status               string                   `mapstructure:"status"               json:"status_inner833"`
    ModelID              string                   `mapstructure:"modelId"              json:"modelId_inner834"`
    ModelVersion         string                   `mapstructure:"modelVersion"         json:"modelVersion_inner835"`
    Score                int                      `mapstructure:"score"                json:"score_inner836"`
    ProbabilityOfDefault float64                  `mapstructure:"probabilityOfDefault" json:"probabilityOfDefault_inner837"`
    RiskGrade            string                   `mapstructure:"riskGrade"            json:"riskGrade_inner838"`
    ReasonCodes          []InnerReasonCode_e1f2a3b1 `mapstructure:"reasonCodes"          json:"reasonCodes_inner839"`
}

func Process_a3333333(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID e1f2a3b1)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_e1f2a3b1
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