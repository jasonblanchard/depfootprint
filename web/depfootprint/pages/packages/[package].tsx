import { useRouter } from 'next/router'
import dynamic from "next/dynamic";
import Link from 'next/link';

import TreeView from '../../components/TreeView';
import TreeFetcher from '../../components/TreeFetcher';

function PackagePage() {
    const router = useRouter();
    const { package: pkg } = router.query;

    if (!pkg) return null;
    return (
        <div>
            <div className="container mx-auto my-5">
                <Link href="/" passHref>
                    <a className="text-slate-500 hover:underline">home</a>
                </Link>
                <h2 className="font-bold text-lg">{pkg}</h2>
            </div>
            <TreeFetcher pkg={pkg as string}>
                {({ dependencies }) => {
                    return <TreeView dependencies={dependencies} />
                }}
            </TreeFetcher>
        </div>
    )
}

export default dynamic(() => Promise.resolve(PackagePage), {
    ssr: false,
});