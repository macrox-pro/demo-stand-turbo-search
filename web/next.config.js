/** @type {import('next').NextConfig} */
const nextConfig = {
    publicRuntimeConfig: {
        graphqlApiUrl: process.env.GRAPHQL_API_URL ?? 'http://localhost:9000/query',
    },
};

module.exports = nextConfig;
