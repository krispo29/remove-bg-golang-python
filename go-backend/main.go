package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	// Serve a simple frontend
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	// Handle image uploads
	http.HandleFunc("/upload", uploadHandler)

	log.Println("Go server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Could not get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Prepare a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", header.Filename)
	if err != nil {
		http.Error(w, "Could not create form file", http.StatusInternalServerError)
		return
	}
	io.Copy(part, file)
	writer.Close()

	// Forward the request to the Python service
	pythonServiceURL := "http://python-processor:5000/remove-bg"
	req, err := http.NewRequest("POST", pythonServiceURL, body)
	if err != nil {
		http.Error(w, "Could not create request to Python service", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Could not reach Python service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Return the processed image from the Python service to the client
	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, resp.Body)
}
