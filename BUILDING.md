1. Install go 1.13 or higher
2. Test

    go test ./...

3. Build binaries

    sh scripts/build.sh

4. Build images

    docker build -t mhealthvn/drone-runner-docker:latest -f docker/Dockerfile.linux.amd64 .
    docker build -t mhealthvn/drone-runner-docker:latest-linux-arm64 -f docker/Dockerfile.linux.arm64 .
    docker build -t mhealthvn/drone-runner-docker:latest-linux-arm   -f docker/Dockerfile.linux.arm   .

5. Run
    docker run -d   --name drone-runner   --restart always   -e DRONE_RPC_PROTO="https"   -e DRONE_RPC_HOST="x"   -e DRONE_RPC_SECRET="x"   -v /var/run/docker.sock:/var/run/docker.sock   mhealthvn/drone-runner-docker:latest