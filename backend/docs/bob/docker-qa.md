# Top 20 Most Asked Docker Questions with Answers

1. **What is Docker?**  
   Docker is a platform that enables developers to automate the deployment of applications inside lightweight, portable containers.

2. **What is a Docker container?**  
   A container is a lightweight, standalone, executable package that includes everything needed to run a piece of software, including code, runtime, system tools, libraries, and settings.

3. **What is a Docker image?**  
   A Docker image is a read-only template used to create containers. It contains the application and all dependencies.

4. **How do I install Docker?**  
   Docker can be installed on various platforms including Windows, Mac, and Linux. Follow the official installation guides for your OS on the Docker Docs site.

5. **What is Docker Hub?**  
   Docker Hub is a cloud-based registry service where you can find and share container images with your team.

6. **How do I create a Docker container?**  
   Use the command `docker run <image-name>` to create and start a container from an image.

7. **How do I list running containers?**  
   Use `docker ps` to list all running containers.

8. **How do I stop a running container?**  
   Use `docker stop <container-id>` to stop a running container.

9. **What is Docker Compose?**  
   Docker Compose is a tool for defining and running multi-container Docker applications using a YAML file.

10. **How do I build a Docker image?**  
    Use `docker build -t <image-name> .` in the directory containing your Dockerfile.

11. **What is a Dockerfile?**  
    A Dockerfile is a text file with instructions to build a Docker image.

12. **How do I share my Docker image?**  
    Push your image to a registry like Docker Hub using `docker push <image-name>`.

13. **How do I remove unused Docker images?**  
    Use `docker image prune` to remove dangling images or `docker system prune` to clean up unused data.

14. **How do I check Docker version?**  
    Run `docker version` to see the client and server version information.

15. **What is the difference between Docker and a virtual machine?**  
    Docker containers share the host OS kernel and are more lightweight, while VMs run full OS instances with their own kernel.

16. **How do I persist data in Docker containers?**  
    Use Docker volumes or bind mounts to persist data outside the container lifecycle.

17. **How do I debug a Docker container?**  
    Use `docker logs <container-id>` to view logs or `docker exec -it <container-id> /bin/bash` to get an interactive shell.

18. **What networking options does Docker provide?**  
    Docker supports bridge, host, overlay, and macvlan networks for container communication.

19. **How do I update a running container?**  
    You cannot update a running container directly; instead, update the image and redeploy a new container.

20. **How do I secure Docker containers?**  
    Follow best practices like running containers with least privilege, using trusted images, and enabling Docker security features.
