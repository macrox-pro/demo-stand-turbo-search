/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string | number; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type IndexObject = {
  __typename?: 'IndexObject';
  ageRestriction?: Maybe<Scalars['String']['output']>;
  countries?: Maybe<Array<Scalars['String']['output']>>;
  description?: Maybe<Scalars['String']['output']>;
  genres?: Maybe<Array<Scalars['String']['output']>>;
  id: Scalars['ID']['output'];
  isActive: Scalars['Boolean']['output'];
  name?: Maybe<Scalars['String']['output']>;
  persons?: Maybe<Array<Scalars['String']['output']>>;
  picture?: Maybe<Scalars['String']['output']>;
  provider?: Maybe<Scalars['String']['output']>;
  score: Scalars['Float']['output'];
  service: Scalars['String']['output'];
  slug?: Maybe<Scalars['String']['output']>;
  title?: Maybe<Scalars['String']['output']>;
  type: Scalars['String']['output'];
  url?: Maybe<Scalars['String']['output']>;
  year?: Maybe<Scalars['String']['output']>;
  yearEnd?: Maybe<Scalars['String']['output']>;
  yearStart?: Maybe<Scalars['String']['output']>;
};

export type Query = {
  __typename?: 'Query';
  search: SearchResponse;
};


export type QuerySearchArgs = {
  query: Scalars['String']['input'];
  useNLP?: InputMaybe<Scalars['Boolean']['input']>;
  where?: InputMaybe<SearchWhereInput>;
};

export type SearchEntity = {
  __typename?: 'SearchEntity';
  end: Scalars['Int']['output'];
  normalValue?: Maybe<Scalars['String']['output']>;
  start: Scalars['Int']['output'];
  type: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type SearchIntent = {
  __typename?: 'SearchIntent';
  confidence: Scalars['Float']['output'];
  name: Scalars['String']['output'];
};

export type SearchResponse = {
  __typename?: 'SearchResponse';
  documents?: Maybe<Array<IndexObject>>;
  metadata: SearchResponseMetadata;
};

export type SearchResponseMetadata = {
  __typename?: 'SearchResponseMetadata';
  entities?: Maybe<Array<SearchEntity>>;
  intent?: Maybe<SearchIntent>;
  query: Scalars['String']['output'];
};

export type SearchWhereInput = {
  active?: InputMaybe<Scalars['Boolean']['input']>;
  service?: InputMaybe<Scalars['String']['input']>;
};

export type SearchQueryVariables = Exact<{
  query: Scalars['String']['input'];
}>;


export type SearchQuery = { __typename?: 'Query', search: { __typename?: 'SearchResponse', documents?: Array<{ __typename?: 'IndexObject', id: string, url?: string | null, year?: string | null, type: string, title?: string | null, score: number, genres?: Array<string> | null, service: string, persons?: Array<string> | null, picture?: string | null, description?: string | null, ageRestriction?: string | null }> | null, metadata: { __typename?: 'SearchResponseMetadata', query: string, intent?: { __typename?: 'SearchIntent', name: string, confidence: number } | null, entities?: Array<{ __typename?: 'SearchEntity', start: number, end: number, type: string, value: string, normalValue?: string | null }> | null } } };


export const SearchDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"search"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"query"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"search"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"query"},"value":{"kind":"Variable","name":{"kind":"Name","value":"query"}}},{"kind":"Argument","name":{"kind":"Name","value":"where"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"active"},"value":{"kind":"BooleanValue","value":true}}]}},{"kind":"Argument","name":{"kind":"Name","value":"useNLP"},"value":{"kind":"BooleanValue","value":true}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"documents"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"url"}},{"kind":"Field","name":{"kind":"Name","value":"year"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"score"}},{"kind":"Field","name":{"kind":"Name","value":"genres"}},{"kind":"Field","name":{"kind":"Name","value":"service"}},{"kind":"Field","name":{"kind":"Name","value":"persons"}},{"kind":"Field","name":{"kind":"Name","value":"picture"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"ageRestriction"}}]}},{"kind":"Field","name":{"kind":"Name","value":"metadata"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"query"}},{"kind":"Field","name":{"kind":"Name","value":"intent"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"confidence"}}]}},{"kind":"Field","name":{"kind":"Name","value":"entities"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"start"}},{"kind":"Field","name":{"kind":"Name","value":"end"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"normalValue"}}]}}]}}]}}]}}]} as unknown as DocumentNode<SearchQuery, SearchQueryVariables>;