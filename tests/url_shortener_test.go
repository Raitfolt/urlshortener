package tests

import (
	"net/url"
	"testing"

	"github.com/Raitfolt/urlshortener/internal/http-server/middleware/logger/handlers/save"
	"github.com/Raitfolt/urlshortener/internal/lib/random"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
)

const (
	host = "127.0.0.1:8082"
)

func TestURLShortener_HappyPath(t *testing.T) {
	// Универсальный способ создания URL
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	// Создаем клиент httpexpect
	e := httpexpect.Default(t, u.String())

	e.POST("/"). // Отправляем POST-запрос, путь - '/url'
			WithJSON(save.Request{ // Формируем тело запроса
			URL:   gofakeit.URL(),             // Генерируем случайный URL
			Alias: random.NewRandomString(10), // Генерируем случайную строку
		}).
		WithBasicAuth("user", "password"). // Добавляем к запросу креды авторизации
		Expect().                          // Далее перечисляем наши ожидания от ответа
		Status(200).                       // Код должен быть 200
		JSON().Object().                   // Получаем JSON-объект тела ответа
		ContainsKey("alias")               // Проверяем, что в нём есть ключ 'alias'
}
