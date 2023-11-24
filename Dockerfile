# Use the same image as in CircleCI configuration
FROM cimg/go:1.19.10-node

# Set environment variables
ENV RESULTS=/tmp/results
ENV ARTIFACTS=/tmp/artifacts
ENV TERM=xterm

# Set the working directory
WORKDIR /home/circleci/perun-node

# Install ganache-cli and aha
RUN sudo npm install -g ganache-cli \
    && sudo apt-get update \
    && sudo apt-get install -y aha \
    && go install github.com/vektra/mockery/v2@v2.14.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

# Copy your project files into the container
COPY . .

CMD ["/bin/bash"]
