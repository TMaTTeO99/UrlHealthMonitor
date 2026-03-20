package service

/*
	Functions to handle the business logic of the APIs requests
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TMaTTeO99/UrlHealthMonitor/API/external"
	"github.com/TMaTTeO99/UrlHealthMonitor/API/models"
	"github.com/TMaTTeO99/UrlHealthMonitor/Repository/connection"
	"github.com/TMaTTeO99/UrlHealthMonitor/config"

	"github.com/jackc/pgx/v5"
)

// Struct used to see the env variable retrieved from .env file
// to find the url informations
type UrlService struct {
	Config *config.ConfigData
	Client *http.Client
	DBConn *pgx.Conn
}

// Method to find the URL's ID that it is going to used to retrieve informations
func (s *UrlService) SearchHandling(w http.ResponseWriter, r *http.Request) {

	// URL value retrieved from request
	urlToAnalyze := FixUrlFormatting(r.PathValue("urlValue"))

	// save url within db
	go connection.InsertUrl(s.DBConn, urlToAnalyze)

	// Build external service req factory
	var reqFactory external.ReqFactoryInterface = &external.ReqFactoryImpl{
		Config: s.Config,
	}

	// Create the request
	req, err := reqFactory.BuildTotalVirusPostReq(urlToAnalyze)
	if err != nil {
		http.Error(w, "Error in build request", http.StatusBadGateway)
		return
	}

	// Do request
	resp, err := s.Client.Do(req)
	if err != nil {
		http.Error(w, "Error in external service http request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		fmt.Println(urlToAnalyze)
		http.Error(w, "VirusTotal doesn't work", resp.StatusCode)
		return
	}

	var vtResponse models.UrlIdDTO
	if err := json.NewDecoder(resp.Body).Decode(&vtResponse); err != nil {
		http.Error(w, "Error during json reading", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vtResponse)

}

// Method to handle the analize API request
func (s *UrlService) AnalizeHandling(w http.ResponseWriter, r *http.Request) {

	// Deserialize the body object
	var request models.RequestUrlDTO
	json.NewDecoder(r.Body).Decode(&request)

	var reqFactory external.ReqFactoryInterface = &external.ReqFactoryImpl{
		Config: s.Config,
	}
	req, err := reqFactory.BuildTotalVirusGetReq(request.ID)
	if err != nil {
		http.Error(w, "Error in build request", http.StatusInternalServerError)
	}

	// Do the request
	resp, err := s.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error in external service http request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "VirusTotal doesn't work", resp.StatusCode)
		return
	}

	var extResponse models.VirusTotalReportDTO
	if err := json.NewDecoder(resp.Body).Decode(&extResponse); err != nil {
		http.Error(w, "Error during json reading", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(extResponse)

}

// Function to fix the url adding http protocol if absent
func FixUrlFormatting(url string) string {

	var fixedUrl string
	if !strings.HasPrefix("http://", url) {
		fixedUrl = "http://" + url
	}

	return fixedUrl

}
