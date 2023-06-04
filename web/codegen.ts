import type {CodegenConfig} from '@graphql-codegen/cli';

const config: CodegenConfig = {
    overwrite: true,
    schema: process.env.GRAPHQL_API_URL ?? 'http://localhost:9000/query',
    documents: 'src/**/*.tsx',
    ignoreNoDocuments: true,
    generates: {
        'src/gql/': {
            preset: 'client',
            plugins: [],
        },
    },
};

export default config;
