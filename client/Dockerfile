FROM node:20-alpine AS builder

WORKDIR /app

# Install pnpm
RUN corepack enable && corepack prepare pnpm@10.8.1 --activate

# Copy dependencies
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the code
COPY . .

# Compile the application
RUN pnpm build

# Final image
FROM node:20-alpine

WORKDIR /app

# Install pnpm, for now, we don't need it
#RUN corepack enable && corepack prepare pnpm@10.8.1 --activate

# Copy the necessary files from the build stage
COPY --from=builder /app/.output ./.output
COPY --from=builder /app/package.json ./package.json

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["node", ".output/server/index.mjs"] 