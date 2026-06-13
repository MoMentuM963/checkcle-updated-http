package pocketbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io"
	"log"
)

func (c *PocketBaseClient) GetServices() ([]Service, error) {
	req, err := http.NewRequest("GET", 
		fmt.Sprintf("%s/api/collections/services/records", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	// No authentication header needed for public access
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch services, status: %d", resp.StatusCode)
	}

	var servicesResponse ServicesResponse
	if err := json.NewDecoder(resp.Body).Decode(&servicesResponse); err != nil {
		return nil, err
	}

	return servicesResponse.Items, nil
}

func (c *PocketBaseClient) GetService(serviceID string) (*Service, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/api/collections/services/records/%s", c.baseURL, serviceID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to fetch service %s, status: %d",
			serviceID,
			resp.StatusCode,
		)
	}

	// Read raw response for debugging
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var service Service
	if err := json.Unmarshal(bodyBytes, &service); err != nil {
		return nil, fmt.Errorf("failed to decode service: %v", err)
	}

	return &service, nil
}

func (c *PocketBaseClient) GetActiveServices() ([]Service, error) {
	var allServices []Service
	page := 1
	perPage := 30 // Use default pagination size

	for {
		// Fetch services page by page with filter for non-paused services
		req, err := http.NewRequest("GET", 
			fmt.Sprintf("%s/api/collections/services/records?page=%d&perPage=%d&filter=(status!='paused')", 
				c.baseURL, page, perPage), nil)
		if err != nil {
			return nil, err
		}

		// No authentication header needed for public access
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch active services, status: %d", resp.StatusCode)
		}

		var servicesResponse ServicesResponse
		if err := json.NewDecoder(resp.Body).Decode(&servicesResponse); err != nil {
			return nil, err
		}

		// Add current page items to the result
		allServices = append(allServices, servicesResponse.Items...)

		// Check if we've fetched all pages
		if page >= servicesResponse.TotalPages || len(servicesResponse.Items) == 0 {
			break
		}

		page++
	}

	return allServices, nil
}

func (c *PocketBaseClient) UpdateServiceStatus(serviceID string, status string, responseTime int64, errorMessage string) error {
	// First check if the service is paused before updating
	service, err := c.GetService(serviceID)
	if err != nil {
		return fmt.Errorf("failed to check service status before update: %v", err)
	}
	
	if service.Status == "paused" {
		return fmt.Errorf("service %s is paused, skipping status update", serviceID)
	}

	updateData := map[string]interface{}{
		"status":        status,
		"response_time": responseTime,
		"last_checked":  time.Now().Format(time.RFC3339),
	}

	if errorMessage != "" {
		updateData["error_message"] = errorMessage
	}

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", 
		fmt.Sprintf("%s/api/collections/services/records/%s", c.baseURL, serviceID),
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// No authentication header needed
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to update service status, status: %d", resp.StatusCode)
	}

	return nil
}