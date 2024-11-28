package parser

//type SlotStorage interface {
//	Add(ctx context.Context, slot domain.BookingSlot) (domain.BookingSlot, error)
//}
//
//type SlotParser struct {
//	storage SlotStorage
//	client  http.Client
//}
//
//func NewSlotService(storage SlotStorage, client http.Client) *SlotParser {
//	return &SlotParser{
//		storage: storage,
//		client:  client,
//	}
//}
//
//func (s *SlotParser) Add(ctx context.Context, slot BookingSlotDTO) (domain.BookingSlot, error) {
//	slotModel := domain.BookingSlot{
//		TimeFrom: slot.TimeFrom,
//		TimeTo:   slot.TimeTo,
//		Price:    slot.Price,
//	}
//	slotModel, err := s.storage.Add(ctx, slotModel)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("storage.Add: %w", err)
//	}
//
//	return slotModel, nil
//}
//
//func (s *SlotParser) Parse(ctx context.Context, date time.Time) (domain.BookingSlot, error) {
//	reqData := &SlotsRequest{
//		SubServiceIds: []string{SlotsAPISubServicesID},
//		StudioID:      LuznikiStudioID,
//		Date:          date.Format("2006-01-02"),
//		Trainers: SlotsTrainersRequest{
//			Type: "NO_TRAINER",
//		},
//	}
//
//	reqBytes, err := json.Marshal(reqData)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("json.Marshal: %w", err)
//	}
//
//	req, err := http.NewRequestWithContext(
//		ctx, http.MethodPost,
//		SlotsAPIURL,
//		bytes.NewReader(reqBytes),
//	)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("http.NewRequestWithContext: %w", err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := s.client.Do(req)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("client.Do: %w", err)
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return domain.BookingSlot{}, fmt.Errorf("client.Do: invalid status code: %d", resp.StatusCode)
//	}
//	defer resp.Body.Close()
//
//	data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("io.ReadAll: %w", err)
//	}
//	fmt.Println(string(data))
//
//	var slot domain.BookingSlot
//	err = json.Unmarshal(data, &slot)
//	if err != nil {
//		return domain.BookingSlot{}, fmt.Errorf("json.Unmarshal: %w", err)
//	}
//	return slot, nil
//}
