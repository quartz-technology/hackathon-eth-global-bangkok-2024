FROM oven/bun

WORKDIR /app

ENV NODE_ENV="production"
ENV LC_ALL="C.UTF-8"
ENV LANG="C.UTF-8"

COPY ./package.json /app/package.json
COPY ./src /app/src
COPY ./tsconfig.json /app/tsconfig.json

RUN bun install

RUN bun run build

CMD [ "bun", "run", "start" ]
