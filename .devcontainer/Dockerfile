FROM pulumi/pulumi

ARG PULUMI_CTL_VERSION
RUN curl -L "https://github.com/pulumi/pulumictl/releases/download/v${PULUMI_CTL_VERSION}/pulumictl-v${PULUMI_CTL_VERSION}-linux-amd64.tar.gz" -o "pulumictl.tar.gz" && \
    tar xzfz pulumictl.tar.gz && \
    mv pulumictl /usr/local/bin && \
    chmod +x /usr/local/bin/pulumictl && \
    rm pulumictl.tar.gz && rm README.md && rm LICENSE