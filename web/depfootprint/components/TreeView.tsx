import { useCallback, useState } from 'react';
import { useRouter } from 'next/router'
import Tree from 'react-d3-tree';

interface Dependency {
    name: string
    version: string
    size: number
    healthScore: number
    children: Dependency[]
}

interface TreeViewProps {
    dependencies: Dependency
}

export const useCenteredTree = (defaultTranslate = { x: 0, y: 0 }) => {
    const [translate, setTranslate] = useState(defaultTranslate);
    const containerRef = useCallback((containerElem: any) => {
        if (containerElem !== null) {
            const { width, height } = containerElem.getBoundingClientRect();
            setTranslate({ x: width / 2, y: height / 5 });
        }
    }, []);
    return [translate, containerRef];
};

const containerStyles = {
    width: "100vw",
    height: "100vh"
};

const renderRectSvgNode = ({ nodeDatum, toggleNode }: any) => {
    let fill = "red";

    if (nodeDatum.healthScore > 70) {
        fill = "orange";
    }

    if (nodeDatum.healthScore > 80) {
        fill = "yellow";
    }

    if (nodeDatum.healthScore > 90) {
        fill = "green";
    }

    let rMultiplier = (nodeDatum.size / 10000) * .4;
    const lowerBound = .2;
    const upperBound = 1.5;
    
    // Bound it so the circles don't get too big or too small
    if (rMultiplier < lowerBound) {
        rMultiplier = lowerBound;
    }

    if (rMultiplier > upperBound) {
        rMultiplier = upperBound
    }

    return (
        <g>
            <circle fill={fill} r={40 * rMultiplier} onClick={toggleNode} />
            <text strokeWidth="1" x="20">
                {nodeDatum.name} ({nodeDatum.healthScore}, {nodeDatum.size}mb)
            </text>
        </g>
    );
}

export default function TreeView({ dependencies }: TreeViewProps) {
    const router = useRouter();
    const { package: pkg } = router.query;
    const [translate, containerRef] = useCenteredTree();

    return (
        // @ts-ignore
        <div style={containerStyles} ref={containerRef}>
            <Tree
                orientation="vertical"
                data={dependencies}
                renderCustomNodeElement={renderRectSvgNode}
                // @ts-ignore
                translate={translate}
                zoomable
            />
        </div>
    )
}