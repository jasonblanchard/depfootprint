import React from 'react';
import Link from 'next/link';

const BaseLayout: React.FC = ({ children }) => {
    return (
        <div className="text-zinc-700">
            <header className="container mx-auto my-5">
                <Link href="/" passHref>
                    <h1 className="font-bold text-2xl">
                        <a href="/" className="hover:underline">depfootprint</a>
                    </h1>
                </Link>
            </header>
            {children}
        </div>
    )
}

export default BaseLayout;