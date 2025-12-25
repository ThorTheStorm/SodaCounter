package request

import (
	"encoding/json"
	"io"
)

// JSON decodes a JSON payload from the provided io.Reader into the specified output structure.
func JSON(body io.Reader, out any) error {
	return json.NewDecoder(body).Decode(out)
}
