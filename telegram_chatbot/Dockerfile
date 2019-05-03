FROM node:8

# install
COPY package*.json ./

RUN npm install

# codes
COPY . .

# run
USER node
CMD [ "npm", "start" ]