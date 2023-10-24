package repohelpers

import (
	"database/sql"
	"testing"
	"time"
)

type dtoStruct struct {
	FieldInt64  sql.NullInt64
	FieldString sql.NullString
	FieldTime   sql.NullTime
	FieldBool   sql.NullBool
}

type entityStruct struct {
	FieldInt64  int64
	FieldString string
	FieldTime   time.Time
	FieldBool   bool
}

func TestDtoToEntity(t *testing.T) {
	type args struct {
		dto    dtoStruct
		entity *entityStruct
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{
				dto: dtoStruct{
					FieldInt64: sql.NullInt64{
						Int64: 64,
						Valid: true,
					},
					FieldString: sql.NullString{
						String: "strinf",
						Valid:  true,
					},
					FieldTime: sql.NullTime{
						Time:  time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
						Valid: true,
					},
					FieldBool: sql.NullBool{
						Bool:  false,
						Valid: true,
					},
				},
				entity: &entityStruct{
					FieldInt64:  64,
					FieldString: "strinf",
					FieldTime:   time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
					FieldBool:   false,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DtoToEntity(tt.args.dto, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("DtoToEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkDtoToEntity(b *testing.B) {
	dto := dtoStruct{
		FieldInt64: sql.NullInt64{
			Int64: 64,
			Valid: true,
		},
		FieldString: sql.NullString{
			String: "strinf",
			Valid:  true,
		},
		FieldTime: sql.NullTime{
			Time:  time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
			Valid: true,
		},
		FieldBool: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
	}

	entity := entityStruct{
		FieldInt64:  64,
		FieldString: "strinf",
		FieldTime:   time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC),
		FieldBool:   false,
	}

	for i := 0; i < b.N; i++ {
		_ = DtoToEntity(dto, &entity)
	}
}
