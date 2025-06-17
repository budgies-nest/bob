## Most Used Docker Bake Commands

### 1. `docker buildx bake`
Runs the default group or specified targets defined in your Bake file. This is the primary command to orchestrate builds using Bake.
```bash
docker buildx bake
```
[Source: Bake Introduction](https://docs.docker.com/build/bake/introduction/)

---

### 2. `docker buildx bake <target>`
Builds a specific target or multiple targets defined in your Bake file.
```bash
docker buildx bake webapp
docker buildx bake webapp api tests
```
[Source: Bake targets](https://docs.docker.com/build/bake/targets/)

---

### 3. `docker buildx bake --file <file>`
Specifies a custom Bake file (HCL, JSON, or Compose YAML) to use for the build.
```bash
docker buildx bake --file my-bake.hcl
```
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 4. `docker buildx bake --print`
Prints the resolved build configuration without actually building anything. Useful for debugging and verifying your Bake setup.
```bash
docker buildx bake --print
```
[Source: Bake file reference](https://docs.docker.com/build/bake/reference/)

---

### 5. `docker buildx bake --list`
Lists all available targets or variables in the Bake configuration, along with their descriptions if set.
```bash
docker buildx bake --list=targets
docker buildx bake --list=variables
```
[Source: List targets and variables](https://docs.docker.com/reference/cli/docker/buildx/bake/#list)

---

### 6. `docker buildx bake --set <targetpattern.key=value>`
Overrides a target value from the command line, such as build arguments or tags.
```bash
docker buildx bake --set webapp.args.VERSION=2.0
```
[Source: Bake CLI reference](https://docs.docker.com/reference/cli/docker/buildx/bake/)

---

### 7. `docker buildx bake --allow <entitlement>`
Grants the build process access to specified resources, such as privileged operations or file system paths.
```bash
docker buildx bake --allow fs.read=../src app
```
[Source: Allow extra privileged entitlement](https://docs.docker.com/reference/cli/docker/buildx/bake/#examples)

---

### 8. `docker buildx bake --no-cache`
Disables the build cache for the build, ensuring all steps are executed from scratch.
```bash
docker buildx bake --no-cache
```
[Source: Bake CLI reference](https://docs.docker.com/reference/cli/docker/buildx/bake/)

---

### 9. `docker buildx bake --push`
Pushes the resulting images to a registry after building.
```bash
docker buildx bake --push
```
[Source: Bake CLI reference](https://docs.docker.com/reference/cli/docker/buildx/bake/)

---

### 10. `docker buildx bake --progress <type>`
Sets the type of progress output (e.g., `auto`, `plain`, `tty`, `rawjson`). Use `plain` to show container output.
```bash
docker buildx bake --progress=plain
```
[Source: Bake CLI reference](https://docs.docker.com/reference/cli/docker/buildx/bake/)
