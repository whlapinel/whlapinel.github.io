+++
date = '2025-06-06T00:03:40-04:00'
draft = false
title = 'Technical Details'
+++

## Developer Environment Dependencies

- Go (see go.mod for go dependencies)
- Docker
  - Marp runs in container
  - Caddy runs in container
- Sqlc CLI
- Task CLI
- Templ CLI
- Goose (for db migrations)
- Tailwind
- Typescript
- Hugo stuff
  - [prettier plugin](https://github.com/NiklasPor/prettier-plugin-go-template)

## Production Environment Dependencies

- Docker
- Bash
- Task CLI

## VS Code extensions used

- [Task](https://marketplace.visualstudio.com/items?itemName=task.vscode-task)
- [Task](https://marketplace.visualstudio.com/items?itemName=paulvarache.vscode-taskfile)
- [Tailwind CSS Intellisense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [Tailwind Docs](https://marketplace.visualstudio.com/items?itemName=austenc.tailwind-docs)
- [Tailwind Fold](https://marketplace.visualstudio.com/items?itemName=stivo.tailwind-fold)
- [Templ](https://marketplace.visualstudio.com/items?itemName=a-h.templ)

## Workflows

### Run in development (local)

```bash
task run-dev
```

Builds app binary and mounts as volume, runs all services with docker compose:

- caddy
- marp
- echo app

After starting all containers with `task run-dev`, you can regenerate templates and rebuild go code during development by running:

```bash
task restart
```

To run migrations on development database (already included in production script):

```bash
task migrate
```

### Deploy

```bash
task deploy
```

- builds images
- pushes them to docker hub
- copies all scripts and env to remote server
- starts ssh session

### Run in production (remote, after deploy)

```bash
task run-prod
```

- pulls the latest images (caddy, marp, and app (includes migration))
- runs the migration binary as a separate container
- starts caddy, marp, and the app ('echo') servers.
