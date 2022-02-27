import { useQuery } from 'react-query';

interface TreeFetcherProps {
    pkg: string
    children: (arg0: any) => any
}

export default function TreeFetcher({ pkg, children }: TreeFetcherProps) {
    const query = useQuery('tree', async () => {
        const response = await fetch(`/api/tree/${pkg}`);

        if (!response.ok) {
            throw new Error(`HTTP error: ${response.status}`);
        }

        const json = await response.json();

        return json;

    }, { refetchOnWindowFocus: false, cacheTime: 0 });

    if (query.isLoading) {
        return (
            <div className="container mx-auto flex align-items-center">
                <div className="spinner-border animate-spin inline-block w-8 h-8 border-4 rounded-full text-green-500 border-blue-500 border-r-white" role="status" />
                <div className="mx-3 py-1 flex align-items-center">
                    Building dependency graph for {pkg}...
                </div>
            </div>
        )
    }

    if (query.isError) {
        return <div>Oops :(</div>
    }

    return children({ dependencies: query.data })
}