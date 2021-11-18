FROM python:3.9-slim

LABEL "repository"="https://github.com/pulumi/pulumi"
LABEL "homepage"="https://pulumi.com"
LABEL "maintainer"="Pulumi Team <team@pulumi.com>"

ENV GOLANG_VERSION 1.16.2
ENV GOLANG_SHA256 542e936b19542e62679766194364f45141fde55169db2d8d01046555ca9eb4b8

# Install deps all in one step
RUN apt-get update -y && \
    apt-get install -y \
    apt-transport-https \
    build-essential \
    ca-certificates \
    curl \
    git \
    gnupg \
    software-properties-common \
    wget \
    unzip && \
    # Get all of the signatures we need all at once.
    curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key  | apt-key add - && \
    curl -fsSL https://dl.yarnpkg.com/debian/pubkey.gpg              | apt-key add - && \
    curl -fsSL https://download.docker.com/linux/debian/gpg          | apt-key add - && \
    curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    curl -fsSL https://packages.microsoft.com/keys/microsoft.asc     | apt-key add - && \
    # IAM Authenticator for EKS
    curl -fsSLo /usr/bin/aws-iam-authenticator https://amazon-eks.s3-us-west-2.amazonaws.com/1.10.3/2018-07-26/bin/linux/amd64/aws-iam-authenticator && \
    chmod +x /usr/bin/aws-iam-authenticator && \
    # AWS v2 cli
    ARCH=`uname -m` && \
    curl "https://awscli.amazonaws.com/awscli-exe-linux-${ARCH}.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf aws && \
    # Add additional apt repos all at once
    echo "deb https://deb.nodesource.com/node_14.x $(lsb_release -cs) main"                         | tee /etc/apt/sources.list.d/node.list             && \
    echo "deb https://dl.yarnpkg.com/debian/ stable main"                                           | tee /etc/apt/sources.list.d/yarn.list             && \
    echo "deb [arch=$(dpkg --print-architecture)] https://download.docker.com/linux/debian $(lsb_release -cs) stable"      | tee /etc/apt/sources.list.d/docker.list           && \
    echo "deb http://packages.cloud.google.com/apt cloud-sdk-$(lsb_release -cs) main"               | tee /etc/apt/sources.list.d/google-cloud-sdk.list && \
    echo "deb http://apt.kubernetes.io/ kubernetes-xenial main"                                     | tee /etc/apt/sources.list.d/kubernetes.list       && \
    echo "deb [arch=$(dpkg --print-architecture)] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/azure.list            && \
    # Install second wave of dependencies
    apt-get update -y && \
    apt-get install -y \
    azure-cli \
    docker-ce \
    google-cloud-sdk \
    kubectl \
    nodejs \
    yarn && \
    # Clean up the lists work
    rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl -fsSLo /tmp/go.tgz https://golang.org/dl/go${GOLANG_VERSION}.linux-$(dpkg --print-architecture).tar.gz; \
    echo "${GOLANG_SHA256} /tmp/go.tgz" | sha256sum -c -; \
    tar -C /usr/local -xzf /tmp/go.tgz; \
    rm /tmp/go.tgz; \
    export PATH="/usr/local/go/bin:$PATH"; \
    go version
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH


# Install .NET Core SDK
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    libc6 \
    libgcc1 \
    libgssapi-krb5-2 \
    libicu67 \
    libssl1.1 \
    libstdc++6 \
    zlib1g \
    && rm -rf /var/lib/apt/lists/*

# Install .NET Core SDK
ARG DOTNET_SDK_VERSION=3.1.415
RUN curl -SL --output dotnet.tar.gz https://dotnetcli.azureedge.net/dotnet/Sdk/${DOTNET_SDK_VERSION}/dotnet-sdk-${DOTNET_SDK_VERSION}-linux-$(dpkg --print-architecture).tar.gz \
    && mkdir -p /usr/share/dotnet \
    && tar -ozxf dotnet.tar.gz -C /usr/share/dotnet \
    && rm dotnet.tar.gz \
    && ln -s /usr/share/dotnet/dotnet /usr/bin/dotnet \
    # Trigger first run experience by running arbitrary cmd
    && dotnet help

# Install Helm
# Explicitly set env variables that helm reads to their defaults, so that subsequent calls to
# helm will find the stable repo even if $HOME points to something other than /root
# (e.g. in GitHub actions where $HOME points to /github/home).
ENV XDG_CONFIG_HOME=/root/.config
ENV XDG_CACHE_HOME=/root/.cache
RUN curl -L https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash && \
    helm repo add stable https://charts.helm.sh/stable && \
    helm repo update

# Passing --build-arg PULUMI_VERSION=vX.Y.Z will use that version
# of the SDK. Otherwise, we use whatever get.pulumi.com thinks is
# the latest
ARG PULUMI_VERSION=latest

# Install the Pulumi SDK, including the CLI and language runtimes.
RUN if [ "$PULUMI_VERSION" = "latest" ]; then \
    curl -fsSL https://get.pulumi.com/ | bash; \
    else \
    curl -fsSL https://get.pulumi.com/ | bash -s -- --version $PULUMI_VERSION ; \
    fi && \
    mv ~/.pulumi/bin/* /usr/bin

ARG PULUMI_CTL_VERSION
RUN curl -L "https://github.com/pulumi/pulumictl/releases/download/v${PULUMI_CTL_VERSION}/pulumictl-v${PULUMI_CTL_VERSION}-linux-$(dpkg --print-architecture).tar.gz" -o "pulumictl.tar.gz" && \
    tar xzfz pulumictl.tar.gz && \
    mv pulumictl /usr/local/bin && \
    chmod +x /usr/local/bin/pulumictl && \
    rm pulumictl.tar.gz && rm README.md && rm LICENSE

# I think it's safe to say if we're using this mega image, we want pulumi
ENTRYPOINT ["pulumi"]