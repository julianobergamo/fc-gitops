# fc-gitops
Curso Full Cycle 3.0 GitOps com algumas alterações

## Na v0.2.0 foi incluído actions do github para o build da imagem docker em alternativa ao passo de build abaixo

# Build Image
```bash
docker build -t julianobergamo/webserver-go:gitops-v0.2.0 .
```

# Run Container
```bash
docker run --name webserver-go -d -p 8080:8080 julianobergamo/webserver-go:gitops-v0.2.0
```

# Sent Image to Docker Hub
```bash
docker push julianobergamo/webserver-go:gitops-v0.2.0
```