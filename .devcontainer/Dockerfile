FROM mcr.microsoft.com/devcontainers/go:1.22

ENV DOCKER_VERSION=20.10.7
# Save command line history 
RUN echo "export HISTFILE=/home/vscode/commandhistory/.bash_history" >> "/home/vscode/.bashrc" \
    && echo "export PROMPT_COMMAND='history -a'" >> "/home/vscode/.bashrc" \
    && mkdir -p /home/vscode/commandhistory \
    && touch /home/vscode/commandhistory/.bash_history

# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends netcat telnet wget

RUN wget https://gitlab.com/api/v4/projects/21018340/jobs/artifacts/main/raw/signaldctl?job=build%3Ax86 -O signalctrl && chmod +x signalctrl && mv signalctrl /usr/bin/signalctrl

# Install docker
ENV DOCKER_BUILDKIT=1
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz \
  && tar xzvf docker-${DOCKER_VERSION}.tgz --strip 1 \
                 -C /usr/local/bin docker/docker \
  && rm docker-${DOCKER_VERSION}.tgz

# Install Kubeval
RUN wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz \
    && tar xf kubeval-linux-amd64.tar.gz \
    && sudo cp kubeval /usr/local/bin
