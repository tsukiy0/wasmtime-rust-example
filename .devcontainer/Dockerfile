FROM mcr.microsoft.com/vscode/devcontainers/go:1.19-bullseye

# node
ARG NODE_VERSION=lts/*
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# atlas cli
RUN curl -fL https://statlas.prod.atl-paas.net/atlas-cli/linux/atlas-latest-linux-amd64.tar.gz | tar -xzp atlas
RUN sudo mv atlas /usr/local/bin
RUN atlas plugin install \ 
    -n packages \
    -n micros \
    -n perfhammer

# rust
RUN apt-get update && \
    apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    build-essential \
    autoconf \
    automake \
    autotools-dev \
    libtool \
    xutils-dev
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH=/root/.cargo/bin:$PATH
RUN rustup target add wasm32-wasi
RUN cargo install cargo-wasi

# other
RUN apt-get update && \
    apt-get install --no-install-recommends -y \
    vim
