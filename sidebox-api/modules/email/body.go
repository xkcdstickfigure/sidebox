package email

import (
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
)

func parseBody(contentType string, encoding string, body io.Reader) (string, string, error) {
	// no content type
	if contentType == "" {
		b, err := io.ReadAll(quotedprintable.NewReader(body))
		if err != nil {
			return "", "", err
		}
		return decodeBody(b, encoding), "", nil
	}

	// parse content type
	ct1, params1, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", "", err
	}

	// plain
	if ct1 == "text/plain" {
		b, err := io.ReadAll(quotedprintable.NewReader(body))
		if err != nil {
			return "", "", err
		}
		return decodeBody(b, encoding), "", nil
	}

	// html
	if ct1 == "text/html" {
		b, err := io.ReadAll(quotedprintable.NewReader(body))
		if err != nil {
			return "", "", err
		}
		return "", decodeBody(b, encoding), nil
	}

	// alternative[plain,html]
	if ct1 == "multipart/alternative" {
		plain, html, err := parseAlternative(body, params1["boundary"])
		return plain, html, err
	}

	// mixed[plain|html|alternative[plain,html],attachment...]
	if ct1 == "multipart/mixed" {
		var plain, html string
		mr := multipart.NewReader(body, params1["boundary"])

		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				return plain, html, nil
			}
			if err != nil {
				return "", "", err
			}

			encoding2 := part.Header.Get("content-transfer-encoding")
			ct2, params2, err := mime.ParseMediaType(part.Header.Get("content-type"))
			if err != nil {
				return "", "", err
			}

			// plain
			if ct2 == "text/plain" {
				b, err := io.ReadAll(quotedprintable.NewReader(body))
				if err != nil {
					return "", "", err
				}
				plain = decodeBody(b, encoding2)
			}

			// html
			if ct2 == "text/html" {
				b, err := io.ReadAll(quotedprintable.NewReader(body))
				if err != nil {
					return "", "", err
				}
				html = decodeBody(b, encoding2)
			}

			// alternative[plain,html]
			if ct2 == "multipart/alternative" {
				plain, html, err = parseAlternative(part, params2["boundary"])
				if err != nil {
					return "", "", err
				}
			}
		}
	}

	return "", "", nil
}

func parseAlternative(r io.Reader, boundary string) (string, string, error) {
	var plain, html string
	mr := multipart.NewReader(r, boundary)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			return plain, html, nil
		}
		if err != nil {
			return plain, html, err
		}

		encoding := part.Header.Get("content-transfer-encoding")
		ct, _, err := mime.ParseMediaType(part.Header.Get("content-type"))
		if err != nil {
			return plain, html, err
		}

		// plain
		if ct == "text/plain" {
			b, err := io.ReadAll(part)
			if err != nil {
				return plain, html, err
			}
			plain = decodeBody(b, encoding)
		}

		// html
		if ct == "text/html" {
			b, err := io.ReadAll(part)
			if err != nil {
				return plain, html, err
			}
			html = decodeBody(b, encoding)
		}
	}
}

func decodeBody(b []byte, encoding string) string {
	if encoding == "base64" {
		decoded, err := base64.StdEncoding.DecodeString(string(b))
		if err != nil {
			return string(b)
		}
		return string(decoded)
	}

	return string(b)
}
