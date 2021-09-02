FROM gitpod/workspace-full

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/

RUN GO111MODULE=on go get golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go get -u github.com/uudashr/gopkgs/cmd/gopkgs \
    && go get -u github.com/ramya-rao-a/go-outline \
    && go get -u github.com/haya14busa/goplay/cmd/goplay \
    && GO111MODULE=on go get github.com/fatih/gomodifytags \
    && go get -u github.com/josharian/impl \
    && go get -u github.com/cweill/gotests/... \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && curl -sL install.mob.sh | sudo sh
