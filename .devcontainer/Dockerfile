# syntax=docker/dockerfile:1
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-22.04

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl build-essential git cmake \
    gcc make unzip libssl-dev pkg-config clang && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# # Install Go
# ARG GO_VERSION=1.23.2
# RUN curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
#     tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
#     rm go${GO_VERSION}.linux-amd64.tar.gz

# ENV PATH="/usr/local/go/bin:${PATH}"

# Install TinyGo
#RUN curl -fsSL https://tinygo.org/get/tinygo | bash
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.33.0/tinygo_0.33.0_amd64.deb
RUN sudo dpkg -i tinygo_0.33.0_amd64.deb
RUN rm tinygo_0.33.0_amd64.deb


# Install Zig
ARG ZIG_VERSION=0.10.1
RUN curl -LO https://ziglang.org/download/${ZIG_VERSION}/zig-linux-x86_64-${ZIG_VERSION}.tar.xz && \
    tar -xf zig-linux-x86_64-${ZIG_VERSION}.tar.xz && \
    mv zig-linux-x86_64-${ZIG_VERSION} /opt/zig && \
    ln -s /opt/zig/zig /usr/local/bin/zig && \
    rm zig-linux-x86_64-${ZIG_VERSION}.tar.xz

# Set environment variables
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

USER vscode
WORKDIR /home/vscode/

# Install Rust
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH="/home/vscode/.cargo/bin:${PATH}"

# Install witgen
RUN cargo install wit-bindgen-cli

# Install wasm-tools
RUN cargo install wasm-tools
