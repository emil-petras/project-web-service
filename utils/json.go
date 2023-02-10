package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UnmarshalJSON(req *http.Request, into interface{}) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("cannot read body: %w", err)
	}

	err = json.Unmarshal(body, &into)
	if err != nil {
		return fmt.Errorf("cannot unmarshal json: %w", err)
	}

	return nil
}
