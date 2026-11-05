package main

import (
	"encoding/json"
	"log"

	"github.com/go-viper/mapstructure/v2"
)

type InnerResult_c3d4e5f3 struct { Result string `mapstructure:"result" json:"result_inner831"` }
type InnerFnsDetails_c3d4e5f3 struct { InnActive bool `mapstructure:"innActive" json:"innActive_inner667"` }
type InnerFnsCheck_c3d4e5f3 struct {
    Result  string                  `mapstructure:"result"  json:"result_inner904"`
    Details InnerFnsDetails_c3d4e5f3 `mapstructure:"details" json:"details_inner22"`
}
type InnerFsspCheck_c3d4e5f3 struct {
    Result      string        `mapstructure:"result"      json:"result_inner391"`
    Proceedings []interface{} `mapstructure:"proceedings" json:"proceedings_inner176"`
}
type InnerChecks_c3d4e5f3 struct {
    PassportValidity InnerResult_c3d4e5f3    `mapstructure:"passportValidity" json:"passportValidity_inner56"`
    InnSnilsMatch    InnerResult_c3d4e5f3    `mapstructure:"innSnilsMatch"    json:"innSnilsMatch_inner294"`
    FnsCheck         InnerFnsCheck_c3d4e5f3  `mapstructure:"fnsCheck"         json:"fnsCheck_inner375"`
    MvdCheck         InnerResult_c3d4e5f3    `mapstructure:"mvdCheck"         json:"mvdCheck_inner148"`
    FsspCheck        InnerFsspCheck_c3d4e5f3 `mapstructure:"fsspCheck"        json:"fsspCheck_inner820"`
    BankruptcyCheck  InnerResult_c3d4e5f3    `mapstructure:"bankruptcyCheck"  json:"bankruptcyCheck_inner459"`
}
type InnerBiometryResult_c3d4e5f3 struct {
    Liveness  string  `mapstructure:"liveness"  json:"liveness_inner888"`
    FaceMatch float64 `mapstructure:"faceMatch" json:"faceMatch_inner241"`
}
type InnerMessage_c3d4e5f3 struct {
    Status         string                      `mapstructure:"status"         json:"status_inner482"`
    VerifiedAt     string                      `mapstructure:"verifiedAt"     json:"verifiedAt_inner117"`
    Checks         InnerChecks_c3d4e5f3         `mapstructure:"checks"         json:"checks_inner739"`
    BiometryResult InnerBiometryResult_c3d4e5f3 `mapstructure:"biometryResult" json:"biometryResult_inner63"`
    RiskFlags      []interface{}               `mapstructure:"riskFlags"      json:"riskFlags_inner507"`
}

func Process_4b405a14(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script (out→inner, ID c3d4e5f3)")
    if msg.Headers == nil { msg.Headers = make(map[string]string) }
    for _, key := range []string{"Sender", "Receiver", "ProcessName"} {
        if _, ok := msg.Headers[key]; ok { continue }
        if msg.Data != nil { if v, ok := msg.Data[key].(string); ok && v != "" { msg.Headers[key] = v } }
    }
    var raw map[string]interface{}
    if err := json.Unmarshal(msg.Body, &raw); err != nil { log.Printf("❌ parse: %v", err); sendOrWarn(msg, output); return msg }
    var inner InnerMessage_c3d4e5f3
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