FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --frozen-lockfile

COPY . .

EXPOSE 3000

ENV DOCKER=true

CMD ["npm", "run", "dev", "--", "--host"]
