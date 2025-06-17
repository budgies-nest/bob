# 20 Frequently Asked Questions About Docker Bake

### 1. What is Docker Bake?
Docker Bake is a feature of Docker Buildx that lets you define your build configuration using a declarative file, as opposed to specifying a complex CLI expression. It also lets you run multiple builds concurrently with a single invocation.  
[Source: Bake Introduction](https://docs.docker.com/build/bake/introduction/)

---

### 2. What file formats does Bake support?
Bake files can be written in HCL, JSON, or YAML (Compose file) formats.  
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 3. How do I run a Bake build?
Use the command:
```bash
docker buildx bake
```
This executes the `default` group or specified targets in your Bake file.  
[Source: Bake Introduction](https://docs.docker.com/build/bake/introduction/)

---

### 4. What is a target in Bake?
A target in a Bake file represents a build invocation, holding all the information you would normally pass to a `docker build` command using flags.  
[Source: Bake targets](https://docs.docker.com/build/bake/targets/)

---

### 5. How do I define multiple builds to run concurrently?
Use the `group` block in your Bake file to define a group of targets that can be built concurrently.  
[Source: Bake](https://docs.docker.com/build/bake/)

---

### 6. How do I specify which Bake file to use?
Use the `-f` or `--file` flag to specify the Bake file:
```bash
docker buildx bake --file my-bake.hcl
```
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 7. What is the default file lookup order for Bake?
If you don't specify a file, Bake looks for files in this order:
1. `compose.yaml`
2. `compose.yml`
3. `docker-compose.yml`
4. `docker-compose.yaml`
5. `docker-bake.json`
6. `docker-bake.hcl`
7. `docker-bake.override.json`
8. `docker-bake.override.hcl`  
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 8. How does Bake handle multiple files?
If multiple Bake files are found, all files are loaded and merged into a single definition. Files are merged according to the lookup order, with later files overriding earlier ones for certain attributes.  
[Source: File overrides](https://docs.docker.com/build/bake/overrides/#file-overrides)

---

### 9. How do I override configuration in Bake?
You can override attributes using file overrides, CLI overrides, or environment variable overrides.  
[Source: Overriding configurations](https://docs.docker.com/build/bake/overrides/)

---

### 10. What attributes can be overridden in Bake?
Attributes such as `args`, `cache-from`, `cache-to`, `context`, `dockerfile`, `output`, `platform`, `tags`, and more can be overridden.  
[Source: Overriding configurations](https://docs.docker.com/build/bake/overrides/)

---

### 11. How do I use variables in Bake files?
Define variables using the `variable` block. Variables can have default values and can be interpolated into attribute values using `${}` syntax.  
[Source: Variables in Bake](https://docs.docker.com/build/bake/variables/)

---

### 12. Can I use expressions in Bake files?
Yes, HCL Bake files support expression evaluation, including arithmetic operations and ternary operators for conditional values.  
[Source: Expression evaluation in Bake](https://docs.docker.com/build/bake/expressions/)

---

### 13. How do I use a remote Bake file definition?
You can build Bake files directly from a remote Git repository or HTTPS URL:
```bash
docker buildx bake "https://github.com/docker/cli.git#v20.10.11"
```
[Source: Remote Bake file definition](https://docs.docker.com/build/bake/remote-definition/)

---

### 14. How do I use local files with a remote Bake definition?
Use the `cwd://` prefix to define contexts as relative to the directory where the Bake command is executed.  
[Source: Remote Bake file definition](https://docs.docker.com/build/bake/remote-definition/)

---

### 15. How do I specify which Bake definition to use in a remote repository?
Use the `--file` or `-f` flag to specify which Bake file to use from the remote repository.  
[Source: Remote Bake file definition](https://docs.docker.com/build/bake/remote-definition/)

---

### 16. How do I print the resolved Bake configuration without building?
Use the `--print` flag:
```bash
docker buildx bake --print
```
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 17. How do I build from a Compose file using Bake?
Bake supports the Compose file format and translates each service to a target.  
[Source: Building with Bake from a Compose file](https://docs.docker.com/build/bake/compose-file/)

---

### 18. How do I use environment variables in Bake?
You can declare default environment variables in a `.env` file, which will be loaded from the current working directory. System environment variables take precedence.  
[Source: Building with Bake from a Compose file](https://docs.docker.com/build/bake/compose-file/)

---

### 19. What are some common Bake CLI options?
Some common options include:
- `--allow` (grant access to resources)
- `--file` (specify Bake file)
- `--print` (print configuration)
- `--no-cache` (disable cache)
- `--push` (push images to registry)  
[Source: docker buildx bake CLI reference](https://docs.docker.com/reference/cli/docker/buildx/bake/)

---

### 20. Where can I find all the properties that can be set for a target?
See the [Bake file reference](https://docs.docker.com/build/bake/reference/#target) for a complete list of target attributes, such as `args`, `context`, `dockerfile`, `tags`, `platforms`, and more.
