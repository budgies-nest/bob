## Docker: Short Tutorial

### 1. Install Docker

First, install Docker Desktop for your operating system. Follow the instructions for your platform:
- [Mac](https://docs.docker.com/desktop/setup/install/mac-install)
- [Windows](https://docs.docker.com/desktop/setup/install/windows-install)
- [Linux](https://docs.docker.com/desktop/setup/install/linux/)

[Source: Get Docker Desktop](https://docs.docker.com/get-started/introduction/get-docker-desktop/)

---

### 2. Run Your First Container

Open your terminal and run:

```bash
docker run -d -p 8080:80 docker/welcome-to-docker
```

- This command downloads the `docker/welcome-to-docker` image (if not present), starts a container in detached mode (`-d`), and maps port 8080 on your machine to port 80 in the container.
- Visit [http://localhost:8080](http://localhost:8080) in your browser to see the running app.

[Source: Run your first container](https://docs.docker.com/get-started/introduction/get-docker-desktop/)

---

### 3. Build Your Own Image

Create a file named `Dockerfile` in your project directory. Example:

```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY . .
RUN yarn install
CMD ["yarn", "start"]
```

Build the image:

```bash
docker build -t my-app .
```

[Source: Build your first image](https://docs.docker.com/get-started/introduction/build-and-push-first-image/)

---

### 4. Run Your Custom Image

Start a container from your image:

```bash
docker run -p 3000:3000 my-app
```

---

### 5. Stop and Remove Containers

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

---

### 6. Next Steps

- Learn about images, containers, and registries.
- Explore Docker Compose for multi-container applications.
- Check out the [Docker workshop](https://docs.docker.com/get-started/workshop/) for a guided, hands-on experience.

[Source: Get started with Docker](https://docs.docker.com/get-started/introduction/)

---

This covers the basics to get you started with Docker!