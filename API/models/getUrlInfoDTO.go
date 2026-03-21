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
		Type string
		ID   string
	}
}

type RequestUrlDTO struct {
	ID string
}

type GetUrlsByIdResponse struct {
	Urls []string
}

// VirusTotalReportDTO represent the data analized
type VirusTotalReportDTO struct {
	Data struct {
		Attributes struct {
			Status string // "queued", "in_progress", "completed"

			// Statistics
			Stats struct {
				Malicious  int
				Suspicious int
				Harmless   int
				Undetected int
				Timeout    int
			}

			// Timestamp (Unix Format)
			Date    int64
			Results map[string]EngineResult
		}

		ID   string
		Type string
	}
}

// EngineResult antivirus/engine details
type EngineResult struct {
	Method     string
	EngineName string
	Category   string // "harmless", "malicious", etc.
	Result     string // Text description
}
