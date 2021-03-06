<h1>ContainerSSH: An SSH server that launches containers</h1>

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/janoszen/containerssh/goreleaser?style=for-the-badge)](https://github.com/janoszen/containerssh/actions)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/janoszen/containerssh?sort=semver&style=for-the-badge)](https://github.com/janoszen/containerssh/releases)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/janoszen/containerssh?style=for-the-badge)](http://hub.docker.com/r/janoszen/containerssh)
[![Go Report Card](https://goreportcard.com/badge/github.com/janoszen/containerssh?style=for-the-badge)](https://goreportcard.com/report/github.com/janoszen/containerssh)
[![LGTM Alerts](https://img.shields.io/lgtm/alerts/github/janoszen/containerssh?style=for-the-badge)](https://lgtm.com/projects/g/janoszen/containerssh/)
[![GitHub](https://img.shields.io/github/license/janoszen/containerssh?style=for-the-badge)](https://github.com/janoszen/containerssh/blob/stable/LICENSE.md)

This is a Proof of Concept SSH server written in Go that sends any shell directly into a Docker container or Kubernetes
pod instead of launching it on a local machine. It uses an HTTP microservice as an authentication endpoint for SSH
connections.

## What is this?

This is an **SSH server that launches containers for every incoming connection**. You can run it on the host or in a
container. It needs two things: an authentication server and access to your container environment.

![Animation: SSH-ing into this SSH server lands you in a container where you can't access the network and you can't see any processes.](https://projects.pasztor.at/containerssh/images/ssh-in-action.gif)

## Use cases

- **Web hosting:** Imagine user A has access to site X and Y, user B has access to site Y and Z. You can use
  ContainerSSH to mount the appropriate sites for the SSH session.
- **Practicing environments:** Launch dummy containers for practice environment.
- **Honeypot:** Let attackers into an enclosed environment and observe them.
- **Security:** Grant your developers on-demand access to environments with changing credentials.

## How does it work?

```
+------+        +--------------+   2.   +-------------------+
|      |        |              | -----> |    Auth server    |
|      |        |              |        +-------------------+
|      |        |              |   
|      |   1.   |              |   3.   +-------------------+
| User | -----> | ContainerSSH | -----> |   Config server   |
|      |        |              |        +-------------------+
|      |        |              |   
|      |        |              |   4.   +-------------------+
|      |        |              | -----> | Container Backend |
+------+        +--------------+        +-------------------+
```

1. The user opens an SSH connection to ContainerSSH.
2. ContainerSSH calls the authentication server with the users username and password/pubkey to check if its valid.
3. ContainerSSH calls the config server to obtain backend location and configuration (if configured)
4. ContainerSSH calls the container backend to launch the container with the
   specified configuration. All input from the user is sent directly to the backend, output from the container is sent
   to the user.
   
> **Curious?** [Learn more about how this works in my blog post.](https://pasztor.at/blog/ssh-direct-to-docker)