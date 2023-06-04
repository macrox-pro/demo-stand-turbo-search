'use client';

import {Input, InputProps} from 'react-daisyui';
import {MagnifyingGlassCircleIcon} from '@heroicons/react/24/solid';
import {KeyboardEvent, useCallback} from 'react';

export type SearchInputProps = InputProps & {
    onEnter?: (e: KeyboardEvent<HTMLInputElement>) => void;
    loading?: boolean;
};

export function SearchInput({
    className = '',
    onKeyDown,
    disabled = false,
    loading = false,
    onEnter,
    id = 'search',
    ...props
}: SearchInputProps) {
    const onKeyDownCallback = useCallback(
        (e: KeyboardEvent<HTMLInputElement>) => {
            if (onKeyDown) {
                onKeyDown(e);
            }
            if (e.key === 'Enter' && onEnter) {
                onEnter(e);
            }
        },
        [onKeyDown, onEnter],
    );
    return (
        <label htmlFor={id} className={`${className} relative block`.trim()}>
            {loading ? (
                <span className="loading loading-spinner loading-sm absolute top-1/2 transform -translate-y-1/2 left-3" />
            ) : (
                <MagnifyingGlassCircleIcon className="pointer-events-none w-8 h-8 absolute top-1/2 transform -translate-y-1/2 left-3" />
            )}
            <Input
                {...props}
                disabled={loading || disabled}
                onKeyDown={onKeyDownCallback}
                className="w-full pl-14"
                id={id}
            />
        </label>
    );
}
