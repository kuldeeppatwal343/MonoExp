
# Stage 1: Build the React application
FROM node:22-alpine AS build

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

COPY . .
RUN yarn build

# Stage 2: Serve the production files
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html