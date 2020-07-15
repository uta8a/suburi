FROM node:14.5-slim

WORKDIR /client
COPY . .
RUN npm install
RUN npm install -g serve
RUN npm run build
CMD ["serve","-l","3000", "-s", "build"]
EXPOSE 3000
