# Kirana Club Backend Intern Assignment

This Project implements a backend service to process images collected from stores. It includes job creation, image processing and status tracking, all developed in Go with performance in mind by Utilizing the multi threaded nature of Go

---

## 📋 Description

This project is a robust and efficient application designed to handle high-volume image processing tasks. Core functionalities include:

- **Job Submission**: Users can submit jobs with image URLs and store IDs.
- **Image Processing**: The service calculates the perimeter of each image, introduces GPU-like delays, and stores results .
- **Job Status Tracking**: Users can query job statuses (ongoing, completed, or failed) and view error details for failed jobs.

---

## ✨ Assumptions

1. **In-memory Database**: The service uses a non-persistent in-memory store (`JobStore`) for simplicity and fast lookups. In a production environment, this would be replaced by a persistent database like PostgreSQL or Redis.
2. **Image Metadata**: Perimeters are calculated using placeholders for height and width; in a real-world scenario, metadata extraction from images would be required.
3. **Error Handling**: Store validation is based on a static `Store Master` file provided during the assignment. Failed downloads are simulated.
4. **Concurrency**: The system assumes that jobs can be processed concurrently without race conditions.

---

## 🚀 Installation and Setup

### Prerequisites

- **Go** (v1.23 or later)
- **Docker** (optional, for containerized setup)

### Running the service

- **With Docker** 
```bash
docker build -t image-service .
docker run -p 8080:8080 image-service
```

- **Without Docker**
```bash
go run ./api/main.go
```

### Testing in Postman

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/18156360-ae55516c-de87-4763-81c8-bcca4ba338a2?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D18156360-ae55516c-de87-4763-81c8-bcca4ba338a2%26entityType%3Dcollection%26workspaceId%3D14f44b01-f4f5-4cb2-808b-5eda8893429f)

### Work Environment

- **Operating System**: Windows 11 10.0.22631 Build 22631
- **Text Editor**: Visual Studio Code
- **External Libaries**: 
    1. gorilla/mux - For Implementing Router
    2. image - For providing the support of webp images to calculate perimeter

### 🔧 Improvements for the Future

1. **Persistent Storage**: Integrate a database (e.g., PostgreSQL ) to persist job data.
2. **Dynamic Image Processing**: Use libraries like image or GoCV to calculate image dimensions directly from URLs.
3. **Error Resilience**: Implement retries for failed downloads and handle edge cases like network outages.
4. **Authentication**: Add an authentication layer to secure endpoints.
5. **API Enhancements**: Include pagination, filtering, and more detailed error responses.