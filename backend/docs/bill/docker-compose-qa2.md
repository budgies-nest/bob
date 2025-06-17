# Docker Compose: Frequently Asked Questions

### 1. What is Docker Compose?
Docker Compose is a tool for defining and running multi-container applications. It lets you manage services, networks, and volumes in a single YAML configuration file, and start all services with a single command. Compose works in all environments: production, staging, development, testing, and CI workflows.  
[Learn more](https://docs.docker.com/compose/)  

---

### 2. What is the difference between `docker compose` and `docker-compose`?
- `docker-compose` (v1) was written in Python and invoked with a hyphen.
- `docker compose` (v2) is written in Go, invoked as a subcommand, and ignores the version top-level element in the compose.yaml file.  
[More details](https://docs.docker.com/compose/support-and-feedback/faq/#what-is-the-difference-between-docker-compose-and-docker-compose)

---

### 3. What is the difference between a Dockerfile and a Compose file?
A Dockerfile provides instructions to build a container image, while a Compose file defines your running containers and their configuration. Often, a Compose file references a Dockerfile to build an image for a service.  
[Source](https://docs.docker.com/guides/docker-compose/common-questions/)

---

### 4. Can I use JSON instead of YAML for my Compose file?
Yes. YAML is a superset of JSON, so any JSON file should be valid YAML. Specify the filename to use, for example:  
```bash
docker compose -f compose.json up
```  
[Source](https://docs.docker.com/compose/support-and-feedback/faq/#can-i-use-json-instead-of-yaml-for-my-compose-file)

---

### 5. Should I include my code with `COPY`/`ADD` or a volume?
- Use `COPY`/`ADD` in your Dockerfile to include code in the image for relocation (e.g., production, CI).
- Use a volume for live code changes during development.
- You can use both: `COPY` in the image and a volume in Compose to override during development.  
[Source](https://docs.docker.com/compose/support-and-feedback/faq/#should-i-include-my-code-with-copyadd-or-a-volume)

---

### 6. Do I need separate Compose files for development, testing, and staging?
No. You can define all services in a single Compose file and use profiles to group configurations for each environment. Activate the desired profile with:  
```bash
docker compose --profile dev up
```  
[Source](https://docs.docker.com/guides/docker-compose/common-questions/)

---

### 7. How can I enforce the database service to start before the frontend?
Use the `depends_on` property to control startup order. For more robust readiness, use health checks to ensure the database is ready before starting the frontend.  
[Source](https://docs.docker.com/guides/docker-compose/common-questions/)

---

### 8. Can I use Compose to build Docker images?
Yes. Specify the build context and Dockerfile in your Compose file. Use:  
```bash
docker compose up --build
```  
[Source](https://docs.docker.com/guides/docker-compose/common-questions/)

---

### 9. What is the difference between `docker compose up`, `run`, and `start`?
- `up`: Starts or restarts all services.
- `run`: Runs one-off or ad-hoc tasks for a service.
- `start`: Restarts stopped containers, never creates new ones.  
[Source](https://docs.docker.com/compose/support-and-feedback/faq/)

---

### 10. Why do my services take 10 seconds to recreate or stop?
`docker compose stop` sends a `SIGTERM` and waits 10 seconds before sending `SIGKILL`. If containers don't shut down, ensure your process handles signals properly or use an init system like `tini` or `dumb-init`.  
[Source](https://docs.docker.com/compose/support-and-feedback/faq/)

---

### 11. Can I scale services with Docker Compose?
Yes, you can scale services up or down within the multi-container setup, allowing efficient resource allocation.  
[Source](https://docs.docker.com/get-started/docker-concepts/running-containers/multi-container-applications/)

---

### 12. How do I define persistent data in Docker Compose?
Use the `volumes` key in your Compose file to define persistent data storage for services.  
[Source](https://docs.docker.com/compose/intro/compose-application-model/)

---

### 13. How do I set environment variables in Docker Compose?
Use the `environment` key under a service in your Compose file to define environment variables.  
[Source](https://docs.docker.com/get-started/workshop/08_using_compose/)

---

### 14. How do I specify a custom project name in Compose?
Use the top-level `name` attribute in your Compose file to set a custom project name, which groups resources and isolates them from other applications.  
[Source](https://docs.docker.com/compose/intro/compose-application-model/)

---

### 15. How do I use Compose with different environments (dev, test, prod)?
Use profiles in your Compose file to group service configurations for each environment and activate them as needed.  
[Source](https://docs.docker.com/guides/docker-compose/common-questions/)

---

### 16. Can I use Compose for single-container applications?
Yes, Compose can be used even for single-container applications to centralize configuration and simplify management.  
[Source](https://docs.docker.com/guides/docker-compose/)

---

### 17. How do I remove all containers created by Compose?
You can use the Docker Desktop Dashboard to delete the application stack, or use the CLI:  
```bash
docker compose down
```  
[Source](https://docs.docker.com/get-started/docker-concepts/running-containers/multi-container-applications/)

---

### 18. How do I view logs from all services?
Use:  
```bash
docker compose logs
```  
to stream log output from all running services.  
[Source](https://docs.docker.com/compose/)

---

### 19. How do I run a one-off command in a service?
Use:  
```bash
docker compose run <service> <command>
```  
to run a one-off or ad-hoc task in a service container.  
[Source](https://docs.docker.com/compose/support-and-feedback/faq/)

---

### 20. Where can I find the latest version of Docker Compose?
The latest version is listed on the [official release notes page](https://docs.docker.com/compose/release-notes/).  
[Source](https://docs.docker.com/compose/)

---

If you need more detailed answers or additional questions, please consult the [official Docker Compose documentation](https://docs.docker.com/compose/).