FROM node:16

EXPOSE 3000

ADD . /REACT_ARQUISOFT
WORKDIR /REACT_ARQUISOFT

COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
CMD ["npm", "run", "start"]