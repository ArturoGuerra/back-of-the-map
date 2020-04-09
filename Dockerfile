FROM node:13-alpine AS BASE


FROM BASE as deps

COPY package-lock.json .
COPY package.json .
RUN npm install

FROM BASE as builder

COPY --from=deps node_modules node_modules
COPY . .
RUN npm install typescript -g
RUN npm run tsc

FROM BASE

WORKDIR /app
COPY package-lock.json .
COPY package.json .
COPY --from=deps node_modules node_modules
COPY --from=builder dist dist

CMD ["npm", "start"]