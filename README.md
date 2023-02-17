# fc-gitops
Curso Full Cycle 3.0 GitOps com algumas alterações

# Build Image
```bash
docker build -t julianobergamo/webserver-go:gitops-v0.1.0 .
```

# Run Container
```bash
docker run -d -p 8080:8080 julianobergamo/webserver-go:gitops-v0.1.0 webserver-go
```

# Sent Image to Docker Hub
```bash
docker push julianobergamo/webserver-go:gitops-v0.1.0
```