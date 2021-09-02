FROM gitpod/workspace-full

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/

RUN GO111MODULE=on sudo go get golang.org/x/tools/gopls@latest \
    && sudo go install github.com/go-delve/delve/cmd/dlv@latest \
    && sudo go get -u github.com/uudashr/gopkgs/cmd/gopkgs \
    && sudo go get -u github.com/ramya-rao-a/go-outline \
    && sudo go get -u github.com/haya14busa/goplay/cmd/goplay \
    && GO111MODULE=on sudo go get github.com/fatih/gomodifytags \
    && sudo go get -u github.com/josharian/impl \
    && sudo go get -u github.com/cweill/gotests/... \
    && sudo go install honnef.co/go/tools/cmd/staticcheck@latest \
    && curl -sL install.mob.sh | sudo sh