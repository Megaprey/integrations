package main

import (
    "encoding/json"
    "fmt"
    "log"
)

// Process_ce39f6a8 - основная функция обработки сообщения.
// msg    - входящее сообщение
// output - MessageSender, куда отправлять результат (может быть nil).
//
// Вся отправка сообщений дальше делается через output.Send(...).
// Если функция ничего не отправляет, сообщение дальше не уйдёт.
// Возвращаемое значение игнорируется — можно вернуть msg или любое другое.
func Process_ce39f6a8(msg Message, output MessageSender) Message {
    log.Printf("🐹 Go Script: обработка сообщения")
    log.Printf("   • ID: %s", msg.ID)
    log.Printf("   • AggregateName: %s", msg.AggregateName)
    log.Printf("   • Body: %s", string(msg.Body))

    // TODO: Добавьте вашу логику обработки здесь
    // Пример: преобразование данных
    // var data map[string]interface{}
    // if err := json.Unmarshal(msg.Body, &data); err == nil {
    //     data["processed"] = true
    //     newBody, _ := json.Marshal(data)
    //     msg.Body = newBody
    //     msg.IsJSON = true
    //     msg.Data = data
    // }

    // Отправляем результат дальше. Если нужно отправить одно сообщение — оставьте
    // один вызов. Если нужно разбить на N копий — вызывайте output.Send для каждой.
    if output != nil {
        if err := output.Send(msg); err != nil {
            log.Printf("❌ Ошибка отправки: %v", err)
        }
    } else {
        log.Printf("⚠️ output sender не установлен")
    }

    return msg
}
