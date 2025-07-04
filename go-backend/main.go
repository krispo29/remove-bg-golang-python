package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/fetch-from-url", fetchFromURLHandler)

	log.Println("Go server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

// uploadHandler จัดการการอัปโหลดไฟล์
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(30 << 20)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 1. จัดการไฟล์ภาพหลัก (image_fg)
	fgFile, fgHeader, err := r.FormFile("image_fg")
	if err != nil {
		http.Error(w, "Could not get foreground image", http.StatusBadRequest)
		return
	}
	defer fgFile.Close()
	partFg, _ := writer.CreateFormFile("image_fg", fgHeader.Filename)
	io.Copy(partFg, fgFile)

	// 2. จัดการไฟล์ภาพพื้นหลัง (image_bg) ถ้ามี
	bgFile, bgHeader, err := r.FormFile("image_bg")
	if err == nil {
		defer bgFile.Close()
		partBg, _ := writer.CreateFormFile("image_bg", bgHeader.Filename)
		io.Copy(partBg, bgFile)
	}

	// 3. จัดการสีพื้นหลัง (bg_color) ถ้ามี
	bgColor := r.FormValue("bg_color")
	if bgColor != "" {
		writer.WriteField("bg_color", bgColor)
	}

	// 4. จัดการชื่อโมเดล (model_name) ถ้ามี
	modelName := r.FormValue("model_name")
	if modelName != "" {
		writer.WriteField("model_name", modelName)
	}

	writer.Close()

	forwardRequestToPython(w, body, writer.FormDataContentType())
}

// fetchFromURLHandler จัดการการดึงรูปจาก URL
func fetchFromURLHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if payload.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(payload.URL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to download image from URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read image data", http.StatusInternalServerError)
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image_fg", "image_from_url.jpg")
	part.Write(imgBytes)
	writer.Close()

	forwardRequestToPython(w, body, writer.FormDataContentType())
}

// ฟังก์ชันช่วยสำหรับส่ง request ไปยัง Python
func forwardRequestToPython(w http.ResponseWriter, body io.Reader, contentType string) {
	pythonServiceURL := "http://python-processor:5000/remove-bg"
	req, err := http.NewRequest("POST", pythonServiceURL, body)
	if err != nil {
		http.Error(w, "Could not create request to Python service", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Could not reach Python service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		http.Error(w, "Error from Python service: "+string(bodyBytes), resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, resp.Body)
}
