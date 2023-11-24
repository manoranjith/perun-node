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
    && sudo apt-get install -y aha

# Copy your project files into the container
COPY . .

CMD ["/bin/bash"]
