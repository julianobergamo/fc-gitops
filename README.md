# fc-gitops
Curso Full Cycle 3.0 GitOps com algumas alterações

# Build Image
```bash
docker build -t julianobergamo/webserver-go:gitops-v.0.1 .
```

# Run Container
```bash
docker run -d -p 8080:8080 julianobergamo/webserver-go:gitops-v.0.1 webserver-go
```

# Sent Image to Docker Hub
```bash
docker push julianobergamo/webserver-go:gitops-v.0.1
```