package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Day struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Tasks   []string `json:"tasks"`
}

var days = map[string]Day{
	"1": {
		ID:      "1",
		Name:    "Понедельник",
		Comment: "Начало рабочей недели. Будь продуктивным!",
		Tasks: []string{
			"Проверить почту",
			"Составить план на неделю",
			"Планерка с директором по результатам",
		},
	},
	"2": {
		ID:      "2",
		Name:    "Вторник",
		Comment: "Сфокусируйся на текущих проектах!",
		Tasks: []string{
			"Проверить почту",
			"Продолжить работу над проектами",
			"Проверить прогресс команды",
			"Приоритизировать зависшие задачи",
		},
	},
}

// getAllDays возвращает все дни в формате JSON при GET-запросе.
func getAllDays(w http.ResponseWriter, req *http.Request) {
	resp, err := json.Marshal(days)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// в заголовок записываю тип контента - это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываю сериализованные в JSON данные в тело ответа, явно игнорируя ошибку
	_, _ = w.Write(resp)
}

// createNewDay добавляет новый день по запросу POST.
func createNewDay(w http.ResponseWriter, req *http.Request) {
	var newDay Day
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &newDay); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// проверяю существование задачи и если она есть - явно возвращаю ошибку с собщением и статусом. При ошибке ничего не добавляю.
	if _, exists := days[newDay.ID]; exists {
		http.Error(w, "День с таким ID уже существует", http.StatusConflict)
		return
	}

	// добавляю новый день в мапу
	days[newDay.ID] = newDay

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// getDay возвращает определенный день по ID при GET-запросе.
func getDay(w http.ResponseWriter, req *http.Request) {
	// эта функция возвращает значение параметра из URL
	id := chi.URLParam(req, "id")

	// проверяю, что ID не пустой
	if id == "" {
		http.Error(w, "ID дня не указан", http.StatusBadRequest)
		return
	}

	// проверяю, существует ли задача с таким ID
	day, ok := days[id]
	if !ok {
		http.Error(w, "День с таким ID не найден", http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(day)
	if err != nil {
		http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
		return
	}

	// возвращаю успешный ответ с днем
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// deleteDay удаляет день по ID при DELETE-запросе.
func deleteDay(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		http.Error(w, "ID дня не указан", http.StatusBadRequest)
		return
	}

	if _, exists := days[id]; !exists {
		http.Error(w, "День с таким ID не найден", http.StatusNotFound)
		return
	}

	// удаляю день из мапы
	delete(days, id)

	// вовращаю ответ с сообщением об удалении, для удобства
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "День успешно удален"}`))
}

func main() {
	r := chi.NewRouter()

	// регистрирую обработчики
	r.Get("/days", getAllDays)
	r.Post("/days", createNewDay)
	r.Get("/days/{id}", getDay)
	r.Delete("/days/{id}", deleteDay)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
