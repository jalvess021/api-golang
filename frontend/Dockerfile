FROM node:20.15.1 AS build

WORKDIR /vue/app

COPY ./ /vue/app/

RUN npm install

RUN npm run build

FROM nginx:alpine

COPY --from=build /vue/app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
