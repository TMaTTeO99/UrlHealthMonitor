package models

type GetUrlInfoDTO struct {
	Time  string
	Info  string
	Error UrlError
}

type UrlError struct {
	Code    int
	Message string
}

type UrlIdDTO struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

type RequestUrlDTO struct {
	URL string
}

// VirusTotalReportDTO represent the data analized
type VirusTotalReportDTO struct {
	Data struct {
		Attributes struct {
			Status string `json:"status"` // "queued", "in_progress", "completed"

			// Statistics
			Stats struct {
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Harmless   int `json:"harmless"`
				Undetected int `json:"undetected"`
				Timeout    int `json:"timeout"`
			} `json:"stats"`

			// Timestamp (Unix Format)
			Date    int64                   `json:"date"`
			Results map[string]EngineResult `json:"results"`
		} `json:"attributes"`

		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// EngineResult antivirus/engine details
type EngineResult struct {
	Method     string `json:"method"`
	EngineName string `json:"engine_name"`
	Category   string `json:"category"` // "harmless", "malicious", etc.
	Result     string `json:"result"`   // Text description
}
