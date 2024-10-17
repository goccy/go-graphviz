FROM ghcr.io/webassembly/wasi-sdk:wasi-sdk-24

ARG GRAPHVIZ_VERSION

ENV PATH=/opt/wasi-sdk/bin:$PATH
ENV USE_CCACHE=1
ENV CCACHE_DIR=/ccache

RUN \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    --mount=type=cache,target=/var/cache/apt,sharing=locked \
    apt-get update && apt install -y pkg-config libexpat1-dev wget ccache

RUN wget https://github.com/WebAssembly/binaryen/releases/download/version_119/binaryen-version_119-x86_64-linux.tar.gz
RUN tar -zxvf binaryen-version_119-x86_64-linux.tar.gz

RUN wget https://github.com/libexpat/libexpat/releases/download/R_2_6_3/expat-2.6.3.tar.gz
RUN tar -zxvf expat-2.6.3.tar.gz && mv expat-2.6.3 expat

RUN wget https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/$GRAPHVIZ_VERSION/graphviz-$GRAPHVIZ_VERSION.tar.gz
RUN tar -zxvf graphviz-$GRAPHVIZ_VERSION.tar.gz && mv graphviz-$GRAPHVIZ_VERSION graphviz

RUN cd graphviz && ./configure --host=amd64 --enable-ltdl=no --with-ipsepcola=no
RUN cd expat && ./configure --host amd64

RUN rm /graphviz/lib/rbtree/test_red_black_tree.c # remove the file includes main function.

WORKDIR /work

ENV PATH=/binaryen-version_119/bin:$PATH

RUN \
    --mount=type=cache,target=/ccache,sharing=locked \
    --mount=type=bind,source=Makefile,target=Makefile \
    --mount=type=bind,source=patch.c,target=patch.c \
    --mount=type=bind,source=bind.c,target=bind.c \
    make build

RUN wasm-opt -g --strip --strip-producers -c -Os graphviz.wasm -o graphviz.wasm
