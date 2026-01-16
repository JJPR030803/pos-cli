package reporter

import "io"

type Reporter struct {
	output io.Writer
}
