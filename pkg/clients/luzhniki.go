package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	SlotsAPIURL           = "https://api.vivacrm.ru/end-user/api/v1/ajV1T2/products/master-services/77075a2c-873a-411f-8073-028a2051cf2d/timeslots"
	LuznikiStudioID       = "0d1604c4-829e-4108-a5eb-71339f0075a3"
	SlotsAPISubServicesID = "5770b319-5402-4677-87aa-76df1c2ef5f0"
)

type SlotsRequest struct {
	SubServiceIds []string             `json:"subServiceIds"`
	StudioID      string               `json:"studioId"`
	Date          string               `json:"date"`
	Trainers      SlotsTrainersRequest `json:"trainers"`
}

type SlotsTrainersRequest struct {
	Type string `json:"type"`
}

type SlotsResponse struct {
	ByTrainer struct {
		NoTrainer struct {
			Trainer interface{} `json:"trainer"`
			Slots   [][]struct {
				TimeFrom          time.Time `json:"timeFrom"`
				TimeTo            time.Time `json:"timeTo"`
				SubServiceID      string    `json:"subServiceId"`
				StudioID          string    `json:"studioId"`
				RoomID            string    `json:"roomId"`
				AvailableDuration string    `json:"availableDuration"`
				Price             struct {
					From float64
				}
			}
		} `json:"NO_TRAINER"`
	}
}

type LuznikiAPIClient struct {
	httpClient http.Client
}

func NewLuznikiAPIClient(client http.Client) *LuznikiAPIClient {
	return &LuznikiAPIClient{
		httpClient: client,
	}
}

func (c *LuznikiAPIClient) GetSlotsForDate(ctx context.Context, date time.Time) (SlotsResponse, error) {
	reqData := SlotsRequest{
		SubServiceIds: []string{SlotsAPISubServicesID},
		StudioID:      LuznikiStudioID,
		Date:          date.Format("2006-01-02"),
		Trainers: SlotsTrainersRequest{
			Type: "NO_TRAINER",
		},
	}

	data, err := c.getSlotsDataFromAPI(ctx, reqData)
	defer data.Close()

	respBytes, err := io.ReadAll(data)
	//fmt.Println(string(respBytes))

	var slotsResponse SlotsResponse
	err = json.Unmarshal(respBytes, &slotsResponse)
	if err != nil {
		return SlotsResponse{}, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return slotsResponse, nil
}

func (c *LuznikiAPIClient) getSlotsDataFromAPI(ctx context.Context, sr SlotsRequest) (io.ReadCloser, error) {
	reqBytes, err := json.Marshal(sr)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		SlotsAPIURL,
		bytes.NewReader(reqBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("client.Do: invalid status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}
