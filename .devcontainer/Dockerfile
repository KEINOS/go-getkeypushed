# =============================================================================
#  Pre Defined Variables
# =============================================================================
# User provided Go version of Docker image
ARG VARIANT
# Choices: https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go
ARG VARIANT_DEFAULT="latest"

# =============================================================================
#  Main
# =============================================================================
FROM ghcr.io/keinos/vscode-dev-container-go:${VARIANT:-$VARIANT_DEFAULT}
