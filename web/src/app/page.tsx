'use client';

import {Card} from 'react-daisyui';
import {graphql} from '@/gql';
import {useQuery} from '@apollo/client';
import {SearchInput} from '@/components/SearchInput';
import {KeyboardEvent, useCallback, useState} from 'react';

const searchQueryDocument = graphql(/* GraphQL */ `
    query search($query: String!) {
        search(query: $query, where: {active: true}, useNLP: true) {
            documents {
                id
                url
                year
                type
                title
                score
                genres
                service
                persons
                picture
                description
                ageRestriction
            }
            metadata {
                query
                intent {
                    name
                    confidence
                }
                entities {
                    start
                    end
                    type
                    value
                    normalValue
                }
            }
        }
    }
`);

export default function Home() {
    const [query, setQuery] = useState('');
    const onSearchEnter = useCallback(
        (e: KeyboardEvent<HTMLInputElement>) => {
            setQuery(((e.target as any).value || '').trim());
        },
        [setQuery],
    );
    const {data, loading} = useQuery(searchQueryDocument, {
        skip: query.length < 1,
        variables: {query},
    });
    return (
        <main className="flex min-h-screen flex-col gap-4 items-center p-8">
            <div className="w-full">
                <SearchInput
                    onEnter={onSearchEnter}
                    disabled={loading}
                    className="w-full"
                    placeholder="Введите текст поиска..."
                />
            </div>
            <div className="w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {data?.search?.documents?.map((document) => (
                    <Card key={document.id} className="w-full mb-4">
                        <Card.Body>
                            <Card.Title tag="div">
                                {document.picture ? (
                                    <div className="w-1/3 mr-2">
                                        {/* eslint-disable-next-line @next/next/no-img-element */}
                                        <img className="rounded" src={document.picture} alt={document.id} />
                                    </div>
                                ) : null}
                                <h2 className="items-start flex-col">
                                    {document.title}
                                    <p>
                                        {document.year ? <div className="badge mb-1 mr-1">{document.year}</div> : null}
                                        {document.genres?.map((genre, index) => (
                                            <div key={`genre-${index}`} className="badge badge-primary mb-1 mr-1">
                                                {genre}
                                            </div>
                                        ))}
                                        {document.ageRestriction ? (
                                            <div className="badge badge-secondary mb-1 mr-1">
                                                {document.ageRestriction}
                                            </div>
                                        ) : null}
                                    </p>
                                </h2>
                            </Card.Title>
                            <p>
                                {document.description}
                                {(document.persons?.length || 0) > 0 ? (
                                    <>
                                        <br />
                                        {document.persons?.map((person, index) => (
                                            <div key={`person-${index}`} className="badge badge-ghost mb-1 mr-1">
                                                {person}
                                            </div>
                                        ))}
                                    </>
                                ) : null}
                            </p>
                        </Card.Body>
                    </Card>
                ))}
            </div>
        </main>
    );
}
