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

    }, { refetchOnWindowFocus: false });

    if (query.isLoading) {
        return <div>Loading...</div>
    }

    if (query.isError) {
        return <div>Oops :(</div>
    }

    return children({ dependencies: query.data })
}