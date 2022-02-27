import React from 'react';
import Link from 'next/link';

const BaseLayout: React.FC = ({ children }) => {
    return (
        <div className="text-zinc-700">
            <header className="container mx-auto my-5">
                <h1 className="font-bold text-2xl">
                    <Link href="/" passHref>
                        <a className="hover:underline">depfootprint</a>
                    </Link>
                </h1>
            </header>
            {children}
        </div>
    )
}

export default BaseLayout;