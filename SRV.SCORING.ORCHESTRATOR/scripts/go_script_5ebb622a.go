package main

import (
    "log"
)

func Process_5ebb622a(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script: обработка сообщения")
    log.Printf("   • ID: %s", msg.ID)
    log.Printf("   • AggregateName: %s", msg.AggregateName)
    log.Printf("   • Body: %s", string(msg.Body))

    if output != nil {
        if err := output.Send(msg); err != nil {
            log.Printf("❌ Ошибка отправки: %v", err)
        }
    } else {
        log.Printf("⚠️ output sender не установлен")
    }

    return msg
}