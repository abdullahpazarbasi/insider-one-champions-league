# Insider One Champions League

This repository contains a sample microservices stack with a Go backend (`back`), a Vue/TypeScript single-page app (`front`), and supporting Kubernetes manifests under `k8s`. It is configured to run locally with Minikube for development and testing.

## Getting started

1. **Clone the repository**:

    ```bash
    git clone git@github.com:abdullahpazarbasi/insider-one-champions-league.git
    ```

2. **Start the local setup** (provisions the image registry, TLS certs, builds images, and bootstraps Minikube):

    ```bash
    make setup
    ```

    This will install prerequisites if you are on a **macOS** or a debian-based **linux** host: **Docker**, **Minikube**, **cURL**, **mkcert**, and **bash**.

3. **Open the SPA in your browser**:

    [https://iocl.local/](https://iocl.local/)

4. **You can check the cluster status**:

    ```bash
    make status-checked
    ```

5. **Run tests as needed**:

    ```bash
    make service-test   # Go service tests
    make bff-test       # Go BFF tests
    make spa-test       # Frontend tests
    ```

> Use `make stop` to tear down local resources and `make cleaned-up` to remove everything after confirmation.
