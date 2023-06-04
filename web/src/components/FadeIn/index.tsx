import classNames from 'classnames';
import React, {useEffect, useState} from 'react';

export type FadeInProps = {
    children: React.ReactNode;
    className?: string;
};

export function FadeIn({className, children}: FadeInProps) {
    const [visible, setVisible] = useState(false);
    useEffect(() => {
        setVisible(true);
        return () => {
            setVisible(false);
        };
    }, []);
    return (
        <div className={classNames(className, visible ? 'opacity-100' : 'opacity-0', 'transition-opacity')}>
            {children}
        </div>
    );
}
