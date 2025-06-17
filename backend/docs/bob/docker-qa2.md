# 20 Frequently Asked Questions About Docker

### 1. What is Docker?
Docker is an open platform for developing, shipping, and running applications. It enables you to separate your applications from your infrastructure so you can deliver software quickly and consistently across environments.  
[Learn more](https://docs.docker.com/get-started/docker-overview/)

---

### 2. What is a Docker container?
A Docker container is a lightweight, standalone, executable package that includes everything needed to run a piece of software, including the code, runtime, system tools, libraries, and settings.  
[Source: What is Docker?](https://docs.docker.com/get-started/docker-overview/)

---

### 3. What is Docker Engine?
Docker Engine is an open source containerization technology for building and containerizing your applications. It consists of a server (the `dockerd` daemon), APIs, and a CLI client (`docker`).  
[Source: Docker Engine](https://docs.docker.com/engine/)

---

### 4. What is Docker Desktop?
Docker Desktop is a native application for Mac, Windows, and Linux that delivers all Docker tools to your computer, including Docker Engine, Docker CLI, Docker Compose, and more.  
[Source: Get Docker](https://docs.docker.com/get-started/get-docker/)

---

### 5. What is Docker Compose?
Docker Compose is a tool for defining and running multi-container Docker applications using a YAML file to configure your application’s services, networks, and volumes.  
[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

### 6. How do I install Docker?
You can install Docker Desktop for Mac, Windows, or Linux by following the instructions for your operating system:  
- [Mac](https://docs.docker.com/desktop/setup/install/mac-install/)  
- [Windows](https://docs.docker.com/desktop/setup/install/windows-install/)  
- [Linux](https://docs.docker.com/desktop/setup/install/linux/)

---

### 7. Is Docker free to use?
Docker offers a free Personal plan. For commercial use in larger enterprises (more than 250 employees OR more than $10 million USD in annual revenue), a paid subscription is required for Docker Desktop.  
[Source: Get Docker](https://docs.docker.com/get-started/get-docker/)

---

### 8. How do I run my first Docker container?
You can run your first container with:  
```bash
docker run -d -p 8080:80 docker/welcome-to-docker
```  
This starts a container and maps port 8080 on your machine to port 80 in the container.  
[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

### 9. How do I build a Docker image?
Create a `Dockerfile` and run:  
```bash
docker build -t my-app .
```  
This builds an image named `my-app` from your Dockerfile.  
[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

### 10. How do I stop and remove a Docker container?
List running containers:  
```bash
docker ps
```
Stop a container:  
```bash
docker stop <container_id>
```
Remove a container:  
```bash
docker rm <container_id>
```
[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

### 11. What is Docker Hub?
Docker Hub is a cloud-based registry service that allows you to share your Docker images with others.  
[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

### 12. What is the difference between Docker Engine and Docker Desktop?
Docker Engine is the core container runtime, while Docker Desktop is a bundled application that includes Docker Engine, Docker CLI, Docker Compose, and other tools for desktop environments.  
[Source: Docker Engine](https://docs.docker.com/engine/), [Get Docker](https://docs.docker.com/get-started/get-docker/)

---

### 13. How do I view logs from a running container?
You can view logs with:  
```bash
docker logs <container_id>
```
[Source: Docker Engine](https://docs.docker.com/engine/)

---

### 14. How do I remove unused Docker resources?
You can tidy up unused resources with:  
```bash
docker system prune
```
[Source: Docker Engine](https://docs.docker.com/engine/)

---

### 15. What is a Docker volume?
A Docker volume is a persistent data storage mechanism that allows data to be stored and shared among containers, independent of the container lifecycle.  
[Source: Docker Engine](https://docs.docker.com/engine/)

---

### 16. How do I back up and restore Docker data?
Instructions for backing up and restoring Docker data are available in the [Docker Desktop documentation](https://docs.docker.com/desktop/settings-and-maintenance/backup-and-restore/).

---

### 17. Where can I find troubleshooting help for Docker?
You can find troubleshooting guides and FAQs in the [Docker Desktop Troubleshooting section](https://docs.docker.com/desktop/troubleshoot-and-support/troubleshoot/).

---

### 18. How do I check the latest version of Docker?
The latest version information for Docker Engine and Docker Desktop can be found in their respective [release notes](https://docs.docker.com/desktop/release-notes/).  
[Source: Support QA]

---

### 19. What is rootless mode in Docker?
Rootless mode allows you to run Docker without root privileges, enhancing security.  
[Source: Docker Engine](https://docs.docker.com/engine/)

---

### 20. Where can I find more learning resources for Docker?
You can follow the [Docker workshop](https://docs.docker.com/get-started/workshop/) and explore the [Get Started guide](https://docs.docker.com/get-started/introduction/) for hands-on learning.  
[Source: Next steps for Docker Desktop](https://docs.docker.com/desktop/setup/install/linux/ubuntu/#next-steps)

