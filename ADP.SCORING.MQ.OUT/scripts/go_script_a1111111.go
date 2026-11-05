package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type ScoringFeaturesOut_d0e1f2a0 struct {
    Age                int     `mapstructure:"age_inner801"                json:"age"`
    IncomeMonthly      int     `mapstructure:"incomeMonthly_inner802"      json:"incomeMonthly"`
    Dti                float64 `mapstructure:"dti_inner803"                json:"dti"`
    BkiScore           int     `mapstructure:"bkiScore_inner804"           json:"bkiScore"`
    ActiveLoans        int     `mapstructure:"activeLoans_inner805"        json:"activeLoans"`
    OverdueMaxDays     int     `mapstructure:"overdueMaxDays_inner806"     json:"overdueMaxDays"`
    AntifraudRiskScore int     `mapstructure:"antifraudRiskScore_inner807" json:"antifraudRiskScore"`
}
type OutMessage_d0e1f2a0 struct {
    RequestID       string                   `mapstructure:"requestId_inner811"       json:"requestId"`
    ApplicationID   string                   `mapstructure:"applicationId_inner812"   json:"applicationId"`
    ModelHint       string                   `mapstructure:"modelHint_inner813"       json:"modelHint"`
    FeatureVectorID string                   `mapstructure:"featureVectorId_inner814" json:"featureVectorId"`
    Features        ScoringFeaturesOut_d0e1f2a0 `mapstructure:"features_inner815"     json:"features"`
    Context         map[string]string        `mapstructure:"context_inner816"         json:"context"`
}

func Process_a1111111(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (inner→out, ID d0e1f2a0)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v; delete(msg.Data, key) } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var out OutMessage_d0e1f2a0
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