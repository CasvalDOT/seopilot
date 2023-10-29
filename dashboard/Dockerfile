FROM node:18 as base


# Development phase
FROM base AS dev

WORKDIR /app

EXPOSE 3000

CMD ["npm", "run", "dev", "--","--port","3000","--host"]

# Build the binary

FROM base AS build

WORKDIR /app

COPY . .

RUN npm ci && npm run build

COPY ./build .

EXPOSE 3000

CMD [ "node", "build" ]
