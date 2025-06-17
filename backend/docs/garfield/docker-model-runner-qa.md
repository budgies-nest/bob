# 20 Frequently Asked Questions About Docker Model Runner

### 1. What is Docker Model Runner?
Docker Model Runner is a feature that allows you to run and interact with AI models directly from the command line or Docker Desktop GUI. It supports pulling, running, and managing AI models locally, and provides OpenAI-compatible APIs for integration with your applications.  
[Learn more](https://docs.docker.com/ai/model-runner/)

---

### 2. How do I enable Docker Model Runner?
- **Docker Desktop:** Go to **Settings > Beta features** and enable **Docker Model Runner**. On Windows with a supported NVIDIA GPU, you can also enable GPU-backed inference.
- **Docker Engine:** Install the `docker-model-plugin` package using your system’s package manager.  
[Enable Docker Model Runner](https://docs.docker.com/ai/model-runner/#enable-docker-model-runner)

---

### 3. What are the prerequisites for using Docker Model Runner?
- Docker Desktop 4.41+ (Windows) or 4.40+ (MacOS)
- Docker Compose v2.35 or later (for Compose integration)
- Docker Desktop for Mac with Apple Silicon or Windows with NVIDIA GPU  
[Prerequisites](https://docs.docker.com/compose/how-tos/model-runner/)

---

### 4. How do I pull a model?
Use the CLI command:
```bash
docker model pull <model-name>
```
For example:
```bash
docker model pull ai/qwen3
```
[Pull a model](https://hub.docker.com/r/ai/qwen3)

---

### 5. How do I run a model?
Use the CLI command:
```bash
docker model run <model-name>
```
For example:
```bash
docker model run ai/qwen3
```
[Run a model](https://hub.docker.com/r/ai/qwen3)

---

### 6. How do I check if Docker Model Runner is running?
Use:
```bash
docker model status
```
[CLI reference](https://docs.docker.com/reference/cli/docker/model/)

---

### 7. How do I list all models pulled to my environment?
Use:
```bash
docker model list
```
[CLI reference](https://docs.docker.com/reference/cli/docker/model/)

---

### 8. How do I remove a model?
Use:
```bash
docker model rm <model-name>
```
[CLI reference](https://docs.docker.com/reference/cli/docker/model/)

---

### 9. How do I view logs for Docker Model Runner?
Use:
```bash
docker model logs
```
[CLI reference](https://docs.docker.com/reference/cli/docker/model/)

---

### 10. How do I uninstall Docker Model Runner?
Use:
```bash
docker model uninstall-runner
```
Options:
- `--images` to remove model-runner images
- `--models` to remove model storage volume  
[Uninstall Model Runner](https://docs.docker.com/reference/cli/docker/model/uninstall-runner/)

---

### 11. How do I install Docker Model Runner manually?
Use:
```bash
docker model install-runner
```
Options:
- `--gpu` to specify GPU support (`none|auto|cuda`)
- `--port` to specify the container port  
[Install Model Runner](https://docs.docker.com/reference/cli/docker/model/install-runner/)

---

### 12. What models are available?
All available models are hosted in the [public Docker Hub namespace of `ai`](https://hub.docker.com/u/ai).  
[FAQs](https://docs.docker.com/ai/model-runner/#faqs)

---

### 13. What API endpoints are available?
Once enabled, Model Runner exposes OpenAI-compatible endpoints such as:
- `/engines/llama.cpp/v1/chat/completions`
- `/engines/llama.cpp/v1/completions`
- `/engines/llama.cpp/v1/embeddings`
- `/models/create`, `/models`, `/models/{namespace}/{name}`, `/models/{namespace}/{name}` (DELETE)  
[API Endpoints](https://docs.docker.com/ai/model-runner/#faqs)

---

### 14. How do I interact with Model Runner using the OpenAI API?
You can use `curl` to send requests to the Model Runner endpoints, either from within a container or from the host (using TCP or Unix socket).  
[OpenAI API interaction](https://docs.docker.com/ai/model-runner/#faqs)

---

### 15. How do I integrate Model Runner with Docker Compose?
Define a service with the `provider` type set to `model` in your `compose.yaml`. Example:
```yaml
services:
  chat:
    image: my-chat-app
    depends_on:
      - ai_runner
  ai_runner:
    provider:
      type: model
      options:
        model: ai/smollm2
```
[Compose integration](https://docs.docker.com/compose/how-tos/model-runner/)

---

### 16. How do I troubleshoot "docker: 'model' is not a docker command"?
Create a symlink so Docker can detect the plugin:
```bash
ln -s /Applications/Docker.app/Contents/Resources/cli-plugins/docker-model ~/.docker/cli-plugins/docker-model
```
[Known issues](https://docs.docker.com/ai/model-runner/#known-issues)

---

### 17. Are there safeguards for running oversized models?
No, currently Docker Model Runner does not prevent you from launching models that exceed your system’s available resources.  
[Known issues](https://docs.docker.com/ai/model-runner/#known-issues)

---

### 18. How do I check the version of Docker Model Runner?
Use:
```bash
docker model version
```
[CLI reference](https://docs.docker.com/reference/cli/docker/model/version/)

---

### 19. How do I provide feedback or report bugs?
Use the **Give feedback** link next to the **Enable Docker Model Runner** setting in Docker Desktop.  
[Share feedback](https://docs.docker.com/ai/model-runner/#share-feedback)

---

### 20. Where can I find more documentation or examples?
- [Docker Model Runner documentation](https://docs.docker.com/ai/model-runner/)
- [Sample GenAI app repository](https://github.com/docker/hello-genai)  
[Reference](https://docs.docker.com/compose/how-tos/provider-services/#reference)

