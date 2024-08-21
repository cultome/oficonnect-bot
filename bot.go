package oficonnectbot

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Bot struct {
	OfiConnectID string
	Client       *Client
}

type PersonalInformation struct {
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

func BuildBot(oficonnect_id string) *Bot {
	return &Bot{
		OfiConnectID: oficonnect_id,
		Client:       BuildClient(),
	}
}

func (b *Bot) RetriveEvents() ([]*Event, error) {
	url := fmt.Sprintf("https://api.oficonnect.omdai.org/public/auth/eventos-usuario/obtener/%s", b.OfiConnectID)

	body, err := b.Client.Get(url)

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

	payload, err := json.Marshal(RegistrationRequest{
		ID:       evt.ID,
		EventID:  evt.EventID,
		UserID:   evt.UserID,
		Confimed: "1",
	})

	body, err := b.Client.Post(url, bytes.NewBuffer(payload))

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

// func (b *Bot) RetrivePersonalInformation() *PersonalInformation {
//
// }

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
