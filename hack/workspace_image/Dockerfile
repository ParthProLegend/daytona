# Copyright 2024 Daytona Platforms Inc.
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:22.04

RUN apt update -y && \
    apt install curl socat libgit2-1.1 -y

USER daytona

ENTRYPOINT ["sudo", "dockerd"]