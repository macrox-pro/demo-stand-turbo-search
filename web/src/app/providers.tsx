'use client';

import React from 'react';
import client from '@/app/apollo-client';
import {ApolloProvider} from '@apollo/client';

export function Providers({children}: {children: React.ReactNode}) {
    return <ApolloProvider client={client}>{children}</ApolloProvider>;
}
