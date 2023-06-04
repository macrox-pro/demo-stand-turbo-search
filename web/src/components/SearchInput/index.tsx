'use client';

import {Input, InputProps} from 'react-daisyui';
import {MagnifyingGlassCircleIcon} from '@heroicons/react/24/solid';
import {KeyboardEvent, useCallback} from 'react';

export type SearchInputProps = InputProps & {
    onEnter?: (e: KeyboardEvent<HTMLInputElement>) => void;
};

export function SearchInput({id = 'search', className = '', onEnter, onKeyDown, ...props}: SearchInputProps) {
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
        <label htmlFor={id} className={`${className} relative text-gray-400 focus-within:text-gray-600 block`.trim()}>
            <MagnifyingGlassCircleIcon className="pointer-events-none w-8 h-8 absolute top-1/2 transform -translate-y-1/2 left-3" />
            <Input {...props} onKeyDown={onKeyDownCallback} className="w-full pl-14" id={id} />
        </label>
    );
}
