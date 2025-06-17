## Docker Compose: Short Tutorial

**Docker Compose** is a tool that helps you define and run multi-container applications. You describe your application’s services, networks, and volumes in a single YAML file (commonly named `compose.yaml`). With one command, you can start or stop your entire application stack.

### 1. Create a `compose.yaml` File

In your project directory, create a file named `compose.yaml`. Here’s an example for a Node.js app with a MySQL database:

```yaml
services:
  app:
    image: node:18-alpine
    command: sh -c "yarn install && yarn run dev"
    ports:
      - 127.0.0.1:3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos

  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos
    volumes:
      - todo-mysql-data:/var/lib/mysql

volumes:
  todo-mysql-data:
```
[Source: Use Docker Compose](https://docs.docker.com/get-started/workshop/08_using_compose/)

---

### 2. Start Your Application Stack

Open a terminal in your project directory and run:

```bash
docker compose up -d
```

This command will:
- Build and start all services defined in your `compose.yaml`
- Create a network for your app
- Create any defined volumes for persistent data

[Source: Run the application stack](https://docs.docker.com/get-started/workshop/08_using_compose/#run-the-application-stack)

---

### 3. View Logs

To see logs from all services:

```bash
docker compose logs -f
```

---

### 4. Stop and Remove Everything

To stop and remove all containers, networks, and resources created by Compose:

```bash
docker compose down
```

---

### Key Benefits

- Define your entire stack in a single file
- Start everything with one command
- Easily share and version-control your stack configuration

[Source: What is Docker Compose?](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-docker-compose/)

---

For more details, see the [Compose overview](https://docs.docker.com/compose/).