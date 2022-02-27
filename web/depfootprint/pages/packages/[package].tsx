import { useRouter } from 'next/router'
import dynamic from "next/dynamic";

import BaseLayout from '../../components/BaseLayout';
import TreeView from '../../components/TreeView';
import TreeFetcher from '../../components/TreeFetcher';

function PackagePage() {
    const router = useRouter();
    const { package: pkg } = router.query;

    if (!pkg) return null;
    return (
        <BaseLayout>
            <div className="container mx-auto my-5">
                <h2 className="font-bold text-lg">
                    <a className="hover:underline" href={`https://www.npmjs.com/package/${pkg}`} target="_blank" rel="noreferrer">{pkg}</a>
                </h2>
            </div>
            <TreeFetcher pkg={pkg as string}>
                {({ dependencies }) => {
                    return <TreeView dependencies={dependencies} />
                }}
            </TreeFetcher>
        </BaseLayout>
    )
}

export default dynamic(() => Promise.resolve(PackagePage), {
    ssr: false,
});