package oficonnectbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Bot struct {
	OfiConnectID string
}

type EventsResponse struct {
	Status string   `json:"status"`
	Events []*Event `json:"eventos"`
}

type RegistrationRequest struct {
	ID       string `json:"id"`
	EventID  string `json:"id_evento"`
	UserID   string `json:"id_usuario"`
	Confimed string `json:"confirmado"`
}

type RegistrationResponse struct {
	Status         string `json:"status"`
	Limit          int    `json:"limite"`
	TotalConfirmed int    `json:"total_confirmados"`
}

type Event struct {
	ID           string `json:"id"`
	Active       string `json:"activo"`
	Open         string `json:"abierto"`
	EventID      string `json:"id_evento"`
	EventName    string `json:"nombre_evento"`
	CategoryName string `json:"nombre_categoria"`
	DivisionName string `json:"nombre_division"`
	Confimed     string `json:"confirmado"`
	Quota        string `json:"cupo"`
	FileURL      string `json:"url_file"`
	Role         string `json:"role"`
	PositionID   string `json:"id_puesto"`
	PositionName string `json:"nombre_puesto"`
	UserID       string `json:"id_usuario"`
	UpdatedAt    string `json:"fecha_actualizado"`
	CreatedAt    string `json:"fecha_creado"`
}

func (b *Bot) RetriveEvents() ([]*Event, error) {
	url := fmt.Sprintf("https://api.oficonnect.omdai.org/public/auth/eventos-usuario/obtener/%s", b.OfiConnectID)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("[RetriveEvents] Unable to create request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OFICONNECT_TOKEN")))

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("[RetriveEvents] Problems with request: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("[RetriveEvents] Unable to read body: %s", err.Error())
	}

	var response EventsResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, fmt.Errorf("[RetriveEvents] Unable to parse response body: %s", err.Error())
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("[RetriveEvents] Failed response: %+v", response)
	}

	return response.Events, nil
}

func (b *Bot) RegisterForEvent(evt *Event) (*RegistrationResponse, error) {
	url := "https://api.oficonnect.omdai.org/public/auth/eventos-usuario/evento/confirmar"

	client := &http.Client{}

	payload, err := json.Marshal(RegistrationRequest{
		ID:       evt.ID,
		EventID:  evt.EventID,
		UserID:   evt.UserID,
		Confimed: "1",
	})

	if err != nil {
		return nil, fmt.Errorf("[RegisterForEvent] Unable parse request body: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return nil, fmt.Errorf("[RegisterForEvent] Unable to create request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OFICONNECT_TOKEN")))

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("[RegisterForEvent] Problems with request: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("[RegisterForEvent] Unable to read body: %s", err.Error())
	}

	var response RegistrationResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, fmt.Errorf("[RegisterForEvent] Unable to parse response body: %s", err.Error())
	}

	return &response, nil
}

// # documentos
// 'https://api.oficonnect.omdai.org/public/auth/datos-personales/avisos/obtener/#{oficonnect_id}'
//
// # datos personales
// 'https://api.oficonnect.omdai.org/public/auth/datos-personales/obtener/#{oficonnect_id}'
//
// # confirmados por evento?
// 'https://api.oficonnect.omdai.org/public/auth/eventos-usuario/confirmados/obtener/#{event_id}'
//
// # cursos
// 'https://api.oficonnect.omdai.org/public/auth/cursos-usuario/obtener-todos/#{oficonnect_id}/normal'
