package kafka_adapter_publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

//
// func (a *KafkaAdapterPublisher) SendBook(ctx context.Context, b book.Entity1) error {
//	r := Request(b)
//
//	bookJSONBytes, err := json.Marshal(r)
//	if err != nil {
//		return err
//	}
//
//	message := kafka.Message{
//		Key:   []byte("Key"),
//		Value: bookJSONBytes,
//	}
//
//	err = a.writer.WriteMessages(ctx, message)
//	if err != nil {
//		return err
//	}
//
//	return nil
// }

type Request struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Date          time.Time `json:"date"`
	NumberOfPages int       `json:"number_of_pages"`
}

type Message struct {
	Payload []byte `json:"payload"`
}

func (a *KafkaAdapterPublisher) SendBooks(ctx context.Context, books ...entity1.Entity1) error {
	var records, mdmRecords []*kgo.Record

	for _, book := range books {
		var body []byte
		var err error

		r := Request(book)

		bookJSONBytes, err := json.Marshal(r)
		if err != nil {
			return err
		}

		message := Message{
			Payload: bookJSONBytes,
		}

		record := kgo.KeySliceRecord([]byte(event.Key), message)

		records = append(records, record)

	}

	ctx, cancel := context.WithTimeout(ctx, p.producingTimeout)
	defer cancel()

	span.LogFields(log.Object("events_cnt", len(records)))
	results := p.client.ProduceSync(ctx, records...)
	for _, result := range results {
		if result.Err != nil {
			return contexts.SpanError(span, result.Err)
		}
	}

	if len(mdmRecords) > 0 {
		span.LogFields(log.Object("mdm_events_cnt", len(mdmRecords)))
		results = p.mdmClient.ProduceSync(ctx, mdmRecords...)
		for _, result := range results {
			if result.Err != nil {
				return contexts.SpanError(span, result.Err)
			}
		}
	}

	return nil
}
