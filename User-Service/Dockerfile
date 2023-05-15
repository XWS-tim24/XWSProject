FROM node:19.8.1

WORKDIR /app

COPY ./User-Service/package*.json ./
RUN npm install -g nodemon && npm install && npm install -g ts-node && npm install -g mongoose
COPY ./User-Service/ .

EXPOSE 5000
CMD ["nodemon"]
