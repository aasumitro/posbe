import React, { useContext, useEffect, useState } from "react";
import { Transformer, Group } from "react-konva";
import { observer } from "mobx-react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import { ViewModes } from "../../models/enums/viewModes";
import { WorkspaceTypes } from "../../models/constants/itemTypes";
import Seat from "./Seat";
import Room from "./Room";

const Workspace = observer(({ shapeProps, onChange }) => {
    const context = useContext(ShapeContext);
	const shapeRef = React.useRef();
	const trRef = React.useRef<any>();
	const { viewMode, multipleSelections, onSelectShape, selectedShapeId, onShapeDragStart, toBookWorkspace } = context;
    const [ isMultiSelected, setIsMultiSelected ] = useState(false);
    const [ isOccupied, setIsOccupied ] = useState(shapeProps.occupied);

    const renderWorkspace = (isOccupied: boolean) => {
        return shapeProps.workspaceType === WorkspaceTypes.Seat 
            ? <Seat shapeProps={shapeProps} isOccupied={isOccupied} isMultiSelected={isMultiSelected} />
            : <Room shapeProps={shapeProps} isOccupied={isOccupied} isMultiSelected={isMultiSelected} />
    };

    useEffect(() => {
        if (toBookWorkspace && toBookWorkspace.id === shapeProps.id) {
            setIsOccupied(toBookWorkspace.isBooked);
        }
    }, [toBookWorkspace])

    useEffect(() => {
        setIsMultiSelected(multipleSelections.indexOf(shapeProps.id) !== -1);
    }, [multipleSelections])

	useEffect(() => {
		if (selectedShapeId === shapeProps.id && !isMultiSelected) {
			//attaching transformer manually
			trRef.current.setNode(shapeRef.current);
			trRef.current.getLayer().batchDraw();
		}
	}, [selectedShapeId]);

	return (
		<>
            <Group 
                {...shapeProps} 
                draggable={viewMode === ViewModes.Design}
                onDragStart={e => onShapeDragStart(shapeProps.id)}
                onDragEnd={e => {
                    onChange({
                        ...shapeProps,
                        x: e.target.x(),
                        y: e.target.y(),
                    });
                }}
                onClick={(e) => { onSelectShape(e, shapeProps.id) } }
                ref={shapeRef}>
                {renderWorkspace(isOccupied)}
            </Group>
			{selectedShapeId === shapeProps.id && !isMultiSelected && <Transformer ref={trRef} />}
		</>
	);
});

export default Workspace;