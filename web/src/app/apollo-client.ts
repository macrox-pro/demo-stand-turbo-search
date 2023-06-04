import {ApolloClient, InMemoryCache} from '@apollo/client';

const client = new ApolloClient({
    uri: 'http://localhost:9000/query',
    cache: new InMemoryCache(),
});

export default client;
