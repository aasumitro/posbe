import React, { useContext } from "react";
import { observer } from "mobx-react";
import { Line, Rect, Text } from "react-konva";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import { getWorkspaceColor } from "../../helpers/konvaUtil";

const Room = observer(({ isOccupied, shapeProps, isMultiSelected }) => {
    const context = useContext(ShapeContext);
    const { viewMode } = context;
    
    const roomProps = {
        width: 55,
        height: 55,
        stroke: isMultiSelected ? 'blue' : 'black',
        strokeWidth: isMultiSelected ? 1 : 0.5,
        fill: getWorkspaceColor(viewMode, isOccupied)
    };

    const lineProps = {
        points: [0, 40, 55, 40],
        stroke: 'black',
        strokeWidth: 0.5,
    };

    const labelProps = {
        y: 10,
        width: roomProps.width,
        wrap: 'char',
        align: 'center'
    };

    return (
        <>
            <Rect {...roomProps} />
            <Line 
                {...lineProps}
            />
            <Text 
                {...labelProps}
                text={shapeProps.floorLabel}
            />
        </>
    );
});

export default Room;