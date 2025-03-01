import React, { useContext } from "react";
import { observer } from "mobx-react";
import { Circle, Text } from "react-konva";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import { getWorkspaceColor } from "../../helpers/konvaUtil";

const Seat = observer(({ isOccupied, isMultiSelected, shapeProps }) => {
    const context = useContext(ShapeContext);
    const { viewMode } = context;
    const circleProps = {
        offsetX: -20,
        radius: 20,
        stroke: isMultiSelected ? 'blue' : 'black',
        strokeWidth: isMultiSelected ? 1 : 0.5,
        fill: getWorkspaceColor(viewMode, isOccupied)
    };

    const labelProps = {
        y: -5,
        width: circleProps.radius * 2,
        wrap: 'char',
        align: 'center'
    };

    return (
        <>
            <Circle {...circleProps} />
            <Text 
                {...labelProps}
                text={shapeProps.floorLabel}
            />
        </>
    );
});

export default Seat;