FROM node:16

EXPOSE 3000

ADD . /React_ArquiSoft
WORKDIR /React_ArquiSoft

COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
CMD ["npm", "run", "start"]