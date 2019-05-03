FROM node:8-alpine

# install
COPY package*.json ./

RUN npm install

# codes
COPY . .

# run
USER node
CMD [ "npm", "start" ]