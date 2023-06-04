'use client';

import {Card} from 'react-daisyui';
import {FadeIn} from '@/components/FadeIn';
import {graphql} from '@/gql';
import {useQuery} from '@apollo/client';
import {upperFirst} from 'lodash';
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
        fetchPolicy: 'network-only',
        variables: {query},
        skip: query.length < 1,
    });
    return (
        <main id="macrox-turbo-vss" className="flex min-h-screen flex-col gap-4 items-start p-8">
            <div className="prose w-full max-w-max">
                <div className="flex flex-col md:flex-row items-baseline">
                    <a href="https://macrox.pro/turbo">
                        <h1 className="mr-4 mb-0">
                            MACROX <i>Turbo</i>
                        </h1>
                    </a>
                    <h3 className="mb-0s">Демонстрационный стенд поиска фильмов и сериалов</h3>
                </div>
            </div>
            <div className="w-full">
                <SearchInput
                    onEnter={onSearchEnter}
                    loading={loading}
                    className="w-full"
                    placeholder="Введите текст поиска..."
                />
            </div>
            <div className="prose w-full max-w-max flex-col gap-4 pt-4 pl-4 pr-4">
                {!loading ? (
                    <>
                        {data?.search?.metadata ? (
                            <>
                                <h1 className="leading-none">
                                    <span
                                        dangerouslySetInnerHTML={{
                                            __html: (data.search.metadata.entities || []).reduce(
                                                (obj, entity) => {
                                                    const elem = `<span class="tooltip tooltip-bottom" data-tip="${
                                                        entity.type
                                                    } - ${entity.normalValue ?? entity.value}">${obj.text.substring(
                                                        obj.offset + entity.start,
                                                        obj.offset + entity.end,
                                                    )}</span>`;
                                                    obj.text = [
                                                        obj.text.substring(0, obj.offset + entity.start),
                                                        elem,
                                                        obj.text.substring(obj.offset + entity.end),
                                                    ].join('');
                                                    obj.offset += elem.length - entity.value.length;
                                                    return obj;
                                                },
                                                {text: upperFirst(data.search.metadata.query), offset: 0},
                                            ).text,
                                        }}
                                    />
                                    {data.search.metadata.intent ? (
                                        <>
                                            <br />
                                            <span className="badge badge-outline">
                                                {data.search.metadata.intent.name}
                                            </span>
                                        </>
                                    ) : null}
                                </h1>
                            </>
                        ) : null}
                        {query.length > 0 && (data?.search?.documents?.length || 0) < 1 ? (
                            <h3>Ничего не найдено</h3>
                        ) : null}
                    </>
                ) : (
                    <span className="loading loading-spinner loading-lg" />
                )}
            </div>
            <div className="w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {data?.search?.documents?.map((document, index) => (
                    <FadeIn key={document.id} className={`w-full mb-4 duration-500 delay-${Math.min(index * 75, 700)}`}>
                        <Card compact>
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
                                            {document.year ? (
                                                <div className="badge mb-1 mr-1">{document.year}</div>
                                            ) : null}
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
                                                <div key={`person-${index}`} className="badge badge-outline mb-1 mr-1">
                                                    {person}
                                                </div>
                                            ))}
                                        </>
                                    ) : null}
                                </p>
                            </Card.Body>
                        </Card>
                    </FadeIn>
                ))}
            </div>
        </main>
    );
}
