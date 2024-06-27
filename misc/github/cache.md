```yaml
name: Cache Docker Images

on:
  push:
    branches:
      - main

jobs:
  cache-docker-images:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v1

    - name: Log in to DockerHub
      uses: docker/login-action@v1
      with:
        username: ${{ secrets.DOCKER_USERNAME }}
        password: ${{ secrets.DOCKER_PASSWORD }}

    - name: Cache Docker layers
      uses: actions/cache@v3
      with:
        path: /tmp/.docker-cache
        key: ${{ runner.os }}-docker-${{ hashFiles('**/Dockerfile') }}
        restore-keys: |
          ${{ runner.os }}-docker-

    - name: Load cache
      if: steps.cache-docker-layers.outputs.cache-hit != 'true'
      run: |
        mkdir -p /tmp/.docker-cache
        docker save -o /tmp/.docker-cache/my_image.tar my_image:latest || true

    - name: Pull Docker image if not cached
      if: steps.cache-docker-layers.outputs.cache-hit != 'true'
      run: |
        docker pull my_image:latest

    - name: Save Docker image to cache
      if: steps.cache-docker-layers.outputs.cache-hit != 'true'
      run: |
        docker save -o /tmp/.docker-cache/my_image.tar my_image:latest

    - name: Load Docker image from cache
      if: steps.cache-docker-layers.outputs.cache-hit == 'true'
      run: |
        docker load -i /tmp/.docker-cache/my_image.tar
```