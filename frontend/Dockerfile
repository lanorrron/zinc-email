#Node.js image
FROM node:18-alpine AS build
#Work dir
WORKDIR /frontend

#Copy file settings
COPY package.json package-lock.json ./

#Install dependencies
RUN npm install

#Copy source code
COPY . .
 
#Build the frontend
RUN npm run build

#Use a image of Nginx
FROM nginx:alpine
COPY --from=build /frontend/dist /usr/share/nginx/html

#Routes the Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

#Expose port nginx
EXPOSE 80

#Init comand
CMD ["nginx", "-g", "daemon off;"]