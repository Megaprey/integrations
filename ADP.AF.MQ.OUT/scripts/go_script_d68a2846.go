package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerAfCheckResult_a7b8c9d7 struct { Result string `mapstructure:"result" json:"result_inner531"` }
type InnerAfChecks_a7b8c9d7 struct {
    BlacklistPhone       InnerAfCheckResult_a7b8c9d7 `mapstructure:"blacklistPhone"       json:"blacklistPhone_inner541"`
    BlacklistPassport    InnerAfCheckResult_a7b8c9d7 `mapstructure:"blacklistPassport"    json:"blacklistPassport_inner542"`
    DuplicateApplication InnerAfCheckResult_a7b8c9d7 `mapstructure:"duplicateApplication" json:"duplicateApplication_inner543"`
    EmployeeFraudPattern InnerAfCheckResult_a7b8c9d7 `mapstructure:"employeeFraudPattern" json:"employeeFraudPattern_inner544"`
}
type InnerMessage_a7b8c9d7 struct {
    RequestID     string              `mapstructure:"requestId"     json:"requestId_inner551"`
    ApplicationID string              `mapstructure:"applicationId" json:"applicationId_inner552"`
    Status        string              `mapstructure:"status"        json:"status_inner553"`
    RiskScore     int                 `mapstructure:"riskScore"     json:"riskScore_inner554"`
    RiskLevel     string              `mapstructure:"riskLevel"     json:"riskLevel_inner555"`
    Checks        InnerAfChecks_a7b8c9d7 `mapstructure:"checks"        json:"checks_inner556"`
    Flags         []interface{}       `mapstructure:"flags"         json:"flags_inner557"`
}

func Process_d68a2846(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID a7b8c9d7)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_a7b8c9d7
    if err := mapstructure.Decode(raw, &inner); err != nil { log.Printf("❌ decode: %v", err); sendOrWarn(msg, output); return msg }
    newBody, err := json.Marshal(inner)
    if err != nil { log.Printf("❌ marshal: %v", err); sendOrWarn(msg, output); return msg }
    msg.Body = newBody; msg.IsJSON = true; msg.Data = inner
    log.Printf("   • Body: %s", string(msg.Body))
    sendOrWarn(msg, output); return msg
}

func sendOrWarn_d68a2846(msg Message, output MessageSender) {
	if output != nil {
		if err := output.Send(msg); err != nil {
			log.Printf("❌ Ошибка отправки: %v", err)
		}
	} else {
		log.Printf("⚠️ output sender не установлен")
	}
}