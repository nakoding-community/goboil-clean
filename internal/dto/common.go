package dto

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/nakoding-community/goboil-clean/internal/abstraction"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type cUUID struct {
	uuid.UUID
}

func CUUIDInit(value uuid.UUID) cUUID {
	return cUUID{value}
}

func (c *cUUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		c.UUID = uuid.Nil
		return nil
	}
	uuid, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	c.UUID = uuid
	return nil
}

func (c *cUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.UUID.String())
}

func (c *cUUID) ToUUID() uuid.UUID {
	return c.UUID
}

type cTime struct {
	time.Time
}

func (c *cTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		c.Time = time.Time{}
		return nil
	}
	var format = "2006-01-02"
	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c *cTime) ToTime() time.Time {
	return c.Time
}

type SearchGetRequest struct {
	abstraction.Pagination
	Search   string   `query:"search"`
	AscField []string `query:"asc_field"`
	DscField []string `query:"dsc_field"`
}

type SearchGetResponse[T any] struct {
	Datas          []T `json:"data"`
	PaginationInfo abstraction.PaginationInfo
}

type ByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

func (r *ByIDRequest) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}

	var err error
	r.ID, err = uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id is not valid")
	}

	return nil
}

type Filter struct {
	Field string
	Value string
}

func BindFilter[T any](c echo.Context, model T, name string) []Filter {
	var filters []Filter

	req := c.Request()
	queries := req.URL.Query()
	modelVal := reflect.ValueOf(model)

	for key, _ := range queries {
		if strings.HasPrefix(key, "filter") {

			field := strings.SplitN(key, "_", 2)
			if len(field) < 2 {
				continue
			}

			for i := 0; i < modelVal.NumField(); i++ {
				if modelVal.Type().Field(i).Tag.Get("json") == field[1] {
					filters = append(filters, Filter{
						Field: name + "." + field[1],
						Value: queries[key][0],
					})
				}
			}
		}
	}

	return filters
}
