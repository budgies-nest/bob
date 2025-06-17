## Most Used Docker Compose Commands

### 1. `docker compose up`
Creates and starts all the services defined in your `compose.yaml` file. This is the main command to launch your multi-container application.  
```bash
docker compose up
```
You can also run services in the background (detached mode) with the `-d` flag:  
```bash
docker compose up -d
```
[Quickstart: Build and run your app with Compose](https://docs.docker.com/compose/gettingstarted/#step-3-build-and-run-your-app-with-compose)

---

### 2. `docker compose down`
Stops and removes all containers, networks, and resources created by `docker compose up`.  
```bash
docker compose down
```
[How Compose works: Stop and remove services](https://docs.docker.com/compose/intro/compose-application-model/#cli)

---

### 3. `docker compose ps`
Lists all the containers and their current status for the services defined in your Compose file.  
```bash
docker compose ps
```
[Quickstart: See what is currently running](https://docs.docker.com/compose/gettingstarted/#step-8-experiment-with-some-other-commands)

---

### 4. `docker compose logs`
Streams the log output from all running services, which is useful for monitoring and debugging.  
```bash
docker compose logs
```
[How Compose works: View logs](https://docs.docker.com/compose/intro/compose-application-model/#cli)

---

### 5. `docker compose stop`
Stops running containers for services without removing them.  
```bash
docker compose stop
```
[Quickstart: Stop your services](https://docs.docker.com/compose/gettingstarted/#step-8-experiment-with-some-other-commands)

---

### 6. `docker compose start`
Starts existing containers for services that were previously stopped, without recreating them.  
```bash
docker compose start
```
[Compose FAQ: Difference between up, run, and start](https://docs.docker.com/compose/support-and-feedback/faq/)

---

### 7. `docker compose build`
Builds or rebuilds the images for the services defined in your Compose file.  
```bash
docker compose build
```
[Compose CLI reference: build](https://docs.docker.com/reference/cli/docker/compose/#subcommands)

---

### 8. `docker compose exec`
Executes a command in a running service container. For example, to open a shell in the `web` service:  
```bash
docker compose exec web sh
```
[Compose CLI reference: exec](https://docs.docker.com/reference/cli/docker/compose/exec/)

---

### 9. `docker compose run`
Runs a one-off command on a service, starting only the containers needed for that service and its dependencies.  
```bash
docker compose run <service> <command>
```
[Compose FAQ: Difference between up, run, and start](https://docs.docker.com/compose/support-and-feedback/faq/)

---

### 10. `docker compose images`
Lists images used by the created containers for your services.  
```bash
docker compose images
```
[Compose CLI reference: images](https://docs.docker.com/reference/cli/docker/compose/images/)

