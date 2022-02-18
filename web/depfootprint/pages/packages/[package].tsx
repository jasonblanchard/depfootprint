import { useRouter } from 'next/router'
import dynamic from "next/dynamic";

import TreeView from '../../components/TreeView';
import TreeFetcher from '../../components/TreeFetcher';

function PackagePage() {
    const router = useRouter();
    const { package: pkg } = router.query;

    if (!pkg) return null;
    return (
        <TreeFetcher pkg={pkg as string}>
            {({ dependencies }) => {
                return <TreeView dependencies={dependencies} />
            }}
        </TreeFetcher>
    )
}

export default dynamic(() => Promise.resolve(PackagePage), {
    ssr: false,
});