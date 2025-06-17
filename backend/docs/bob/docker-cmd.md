# 20 Most Used Docker Commands with Explanations

1. **docker run**  
   Run a new container from an image.  
   Example: `docker run -it ubuntu bash`  
   Starts an interactive bash shell in an Ubuntu container.

2. **docker ps**  
   List running containers.  
   Example: `docker ps`  
   Shows all currently running containers.

3. **docker ps -a**  
   List all containers (running and stopped).  
   Example: `docker ps -a`  
   Useful to see containers that have exited.

4. **docker images**  
   List all downloaded images.  
   Example: `docker images`  
   Shows images available locally.

5. **docker pull**  
   Download an image from a registry.  
   Example: `docker pull nginx`  
   Pulls the latest nginx image from Docker Hub.

6. **docker build**  
   Build an image from a Dockerfile.  
   Example: `docker build -t myapp .`  
   Creates an image named "myapp" from the current directory.

7. **docker exec**  
   Run a command inside a running container.  
   Example: `docker exec -it container_id bash`  
   Opens an interactive shell inside the container.

8. **docker stop**  
   Stop a running container.  
   Example: `docker stop container_id`  
   Gracefully stops the container.

9. **docker start**  
   Start a stopped container.  
   Example: `docker start container_id`  
   Restarts a previously stopped container.

10. **docker rm**  
    Remove one or more stopped containers.  
    Example: `docker rm container_id`  
    Deletes the container from the system.

11. **docker rmi**  
    Remove one or more images.  
    Example: `docker rmi image_id`  
    Deletes the image from local storage.

12. **docker logs**  
    Fetch logs from a container.  
    Example: `docker logs container_id`  
    Shows the output logs of the container.

13. **docker inspect**  
    Display detailed information about containers or images.  
    Example: `docker inspect container_id`  
    Outputs JSON with configuration and status.

14. **docker network ls**  
    List Docker networks.  
    Example: `docker network ls`  
    Shows all networks available.

15. **docker volume ls**  
    List Docker volumes.  
    Example: `docker volume ls`  
    Shows all volumes used for persistent data.

16. **docker-compose up**  
    Create and start containers defined in a docker-compose.yml file.  
    Example: `docker-compose up`  
    Starts all services defined in the compose file.

17. **docker-compose down**  
    Stop and remove containers, networks created by docker-compose.  
    Example: `docker-compose down`  
    Cleans up the environment.

18. **docker-compose logs**  
    View output logs from services in a compose project.  
    Example: `docker-compose logs`  
    Shows logs for all or specific services.

19. **docker tag**  
    Tag an image for a repository.  
    Example: `docker tag image_id username/repo:tag`  
    Prepares an image for pushing to a registry.

20. **docker push**  
    Push an image to a registry.  
    Example: `docker push username/repo:tag`  
    Uploads the tagged image to Docker Hub or other registries.

