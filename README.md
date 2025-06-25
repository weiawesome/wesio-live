# Wesio Live - Live Streaming Platform

A scalable live streaming platform built with microservices architecture, designed to provide real-time video streaming, chat functionality, and comprehensive media management capabilities.

## 🚀 Overview

Wesio Live is a modern live streaming platform that supports:
- **Real-time video streaming** with WebRTC technology
- **Interactive chat system** for viewer engagement
- **User management** with role-based authentication
- **Media storage and CDN** for optimal content delivery
- **Room management** for organized streaming sessions
- **Scalable microservices architecture** for high availability

## 🏗️ Architecture

The platform follows a microservices architecture pattern with the following core services:

### Core Services
- **Signaling Server (WebRTC)** - Handles WebRTC peer connections and media negotiation
- **Media Server** - Manages media storage, upload, and CDN distribution
- **Chat Service (WebSocket)** - Real-time messaging and chat functionality
- **User Service** - User authentication, authorization, and profile management
- **Room Service** - Live streaming room creation and management
- **Identity Service** - JWT-based authentication and authorization

### Infrastructure Components
- **Object Storage** - MinIO for scalable media file storage
- **Message Queue** - NATS for inter-service communication
- **Load Balancer** - Traffic distribution and service discovery
- **CDN Integration** - Content delivery optimization

## 📁 Project Structure

```
wesio-live/
├── docs/                    # Documentation and architecture diagrams
├── frontend/               # Client applications
│   ├── web/               # Web application
│   └── mobile-app/        # Mobile applications (iOS/Android)
├── services/              # Microservices
│   ├── signaling/         # WebRTC signaling service
│   ├── media/             # Media storage service
│   ├── chat/              # Real-time chat service
│   ├── user/              # User management service
│   ├── room/              # Room management service
│   └── identity/          # Authentication service
├── storage/               # Data layer
│   ├── media/             # Media storage interface
│   ├── user/              # User data models
│   ├── chat/              # Chat message models
│   └── room/              # Room data models
├── libs/                  # Shared libraries
│   ├── auth/              # JWT authentication (gRPC)
│   ├── logger/            # Structured logging
│   └── config/            # Configuration management
├── infra/                 # Infrastructure as Code
│   ├── docker/            # Docker configurations
│   ├── k8s/              # Kubernetes manifests
│   ├── terraform/         # Infrastructure provisioning
│   ├── helm/              # Helm charts
│   └── ansible/           # Configuration management
└── main.go               # Application entry point
```

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.24+
- **Communication**: gRPC, WebSocket, REST API
- **Authentication**: JWT with protobuf definitions
- **Storage**: MinIO (S3-compatible object storage)
- **Message Queue**: NATS with JetStream
- **Logging**: Structured logging with trace IDs

### Frontend
- **Web**: Modern web technologies (React/Vue.js)
- **Mobile**: Cross-platform development (React Native/Flutter)

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Orchestration**: Kubernetes
- **IaC**: Terraform
- **Configuration**: Helm Charts
- **Automation**: Ansible

## 🚀 Quick Start

### Prerequisites
- Go 1.24 or later
- Docker and Docker Compose
- MinIO (for object storage)
- NATS (for messaging)

### 1. Clone the Repository
```bash
git clone https://github.com/your-org/wesio-live.git
cd wesio-live
```

### 2. Start Infrastructure Services
```bash
cd infra/container
docker-compose up -d
```

This will start:
- MinIO server (API: http://localhost:9000, Console: http://localhost:9001)
- NATS server (Client: localhost:4222, Dashboard: http://localhost:8222)

### 3. Install Dependencies
```bash
go mod download
```

### 4. Run the Application
```bash
go run main.go
```

### 5. Access Services
- **MinIO Console**: http://localhost:9001 (minioadmin/minioadmin)
- **NATS Dashboard**: http://localhost:8222
- **Application**: Check logs for service endpoints

## 🔧 Configuration

### Environment Variables
```bash
# MinIO Configuration
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_USE_SSL=false

# CDN Configuration
CDN_DOMAIN=https://cdn.example.com
CDN_SIGNING_KEY=your-cdn-signing-key

# NATS Configuration
NATS_URL=nats://localhost:4222

# Application Configuration
LOG_LEVEL=info
PORT=8080
```

## 🏃‍♂️ Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o wesio-live main.go
```

### Docker Build
```bash
docker build -t wesio-live:latest .
```

## 📦 Deployment

### Kubernetes Deployment
```bash
# Apply Kubernetes manifests
kubectl apply -f infra/k8s/

# Or use Helm
helm install wesio-live infra/helm/wesio-live
```

### Terraform Infrastructure
```bash
cd infra/terraform
terraform init
terraform plan
terraform apply
```

## 🔒 Security Features

- **JWT Authentication** with role-based access control
- **Secure media storage** with signed URLs
- **CDN integration** for content protection
- **Rate limiting** and request validation
- **Encrypted communication** between services

## 📊 Monitoring & Observability

- **Structured Logging** with trace ID correlation
- **Metrics Collection** for performance monitoring
- **Health Checks** for service availability
- **Distributed Tracing** for request flow analysis

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support, please:
1. Check the [documentation](docs/)
2. Search existing [issues](https://github.com/your-org/wesio-live/issues)
3. Create a new issue with detailed information

## 🚧 Roadmap

- [ ] Enhanced video quality controls
- [ ] Mobile SDK for third-party integration
- [ ] Advanced analytics and reporting
- [ ] Multi-language support
- [ ] Enhanced security features
- [ ] Performance optimizations

---

**Wesio Live** - Building the future of live streaming technology 🎥✨