FROM oven/bun

WORKDIR /app

ARG VITE_WEB3_AUTH_VERIFIER
ARG VITE_WEB3_AUTH_CLIENTID
ARG VITE_API_URL
ARG VITE_ETH_RPC_URL

ENV VITE_WEB3_AUTH_VERIFIER=${VITE_WEB3_AUTH_VERIFIER}
ENV VITE_WEB3_AUTH_CLIENTID=${VITE_WEB3_AUTH_CLIENTID}
ENV VITE_API_URL=${VITE_API_URL}
ENV VITE_ETH_RPC_URL=${VITE_ETH_RPC_URL}

ENV NODE_ENV="production"
ENV LC_ALL="C.UTF-8"
ENV LANG="C.UTF-8"

COPY ./package.json /app/package.json
COPY ./src /app/src
COPY ./index.html /app/index.html
COPY ./tsconfig.json /app/tsconfig.json
COPY ./tsconfig.node.json /app/tsconfig.node.json
COPY ./vite.config.ts /app/vite.config.ts

RUN bun install

RUN bun run build

CMD [ "bun", "run", "preview" ]
