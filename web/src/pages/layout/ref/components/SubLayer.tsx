import React, { useContext, useEffect } from "react";
import { Layer, Rect, Transformer } from "react-konva";
import GridLines from "./GridLines";
import BlockList from "./floor-items/BlockList";
import { observer } from "mobx-react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import ShapeCounter from "./canvas/shapeCounter/ShapeCounter";
import { ViewModes } from "../models/enums/viewModes";
import PolygonAnnotation from "./floor-boundary/PolygonAnnotation";
import { getMousePos } from "../helpers/konvaUtil";
import WorkspaceList from "./floor-items/WorkspaceList";

const SubLayer = observer((props) => {
    const { layerRef, height, width } = props;
    const context = useContext(ShapeContext);
	const { isEditBoundary, viewMode, deleteShapes, setSelectedShapeId, addShapes, cloneShape, multipleSelections, clonedItems, setClonedItems, distributeClones, setMultipleSelections } = context;
    const { boundaryItems, points, setPoints, flattenedPoints, setFlattenedPoints, position, setPosition, isPolyComplete, setPolyComplete, isMouseOverPoint, setMouseOverPoint } = context;
	const rectConfig = {
		fill: 'white',
		width: width,
		height: height
	}

	const trRef = React.useRef<any>();

	const handleRenderedClonedItems = () => {
		if (viewMode !== ViewModes.Design || clonedItems.length === 0) return;

		let toSelect = [];
		clonedItems.forEach(i => {
			const el = layerRef.current.children.find(layerElement => layerElement.attrs.id === i.id);
			if (el) toSelect.push(el);
		});

		if (toSelect.length === 0) return;

		trRef.current.nodes(toSelect);
		setClonedItems([]);

		if (toSelect.length === 1) {
			setSelectedShapeId(toSelect[0].attrs.id);
		} else {
			const ids = toSelect.map(s => s.attrs.id);
			setMultipleSelections(ids);
		}
	}

	const syncMultipleSelectionsAndTr = () => {
		if (multipleSelections.length === 0) {
			trRef.current.nodes([])
		} else {
			const elements = layerRef.current.children.filter(c => c.attrs.class === 'floor-item' && multipleSelections.indexOf(c.attrs.id) !== -1);
			trRef.current.nodes(elements);
		}
	}

	useEffect(() => {
		handleRenderedClonedItems();
	}, [clonedItems])

	// copy paste, delete events
	useEffect(() => {
		if (viewMode !== ViewModes.Design) return;

		syncMultipleSelectionsAndTr();

		let cloneList = [];
		const handleKeyDown = (event) => {
			// copy
			if (event.ctrlKey && event.key.toLowerCase() === 'c') {
				cloneList = [];
				const toCopyItems = trRef.current.nodes();
				if (toCopyItems.length === 1) {
					alert('Cannot copy single item. Please add more to your selection.');
					return;
				}
				toCopyItems.forEach((n, idx) => {
					const cloned = cloneShape(n, idx);
					cloneList.push(cloned);
				});
			}

			// paste
			if (event.ctrlKey && event.key.toLowerCase() === 'v' && cloneList.length > 0) {
				distributeClones(cloneList);
				addShapes(cloneList);
				setClonedItems(cloneList);
				cloneList = [];
			}

			// delete
			if (event.key === 'Delete' || event.keyCode === 46) {
				deleteShapes(multipleSelections);
				setSelectedShapeId(null);
				trRef.current.nodes([]);
			}
		};

		document.addEventListener('keydown', handleKeyDown);

		return () => {
			document.removeEventListener('keydown', handleKeyDown);
		};
	}, [multipleSelections]);

    //drawing begins when mousedown event fires.
    const handleMouseDown = (e) => {
		if (isEditBoundary) {
			if (isPolyComplete) return;
			const mousePos = getMousePos(e, layerRef);
			if (isMouseOverPoint && points.length >= 3) {
				setPolyComplete(true);
			} else {
				setPoints([...points, mousePos]);
			}
		}
    };

    const handleMouseMove = (e) => {
		if (isEditBoundary) {
			const mousePos = getMousePos(e, layerRef);
			setPosition(mousePos);
		}
    };

	useEffect(() => {
		if (boundaryItems && boundaryItems.points.length > 0 && boundaryItems.flattenedPoints.length > 0) {
			setPoints(boundaryItems.points);
			setFlattenedPoints(boundaryItems.flattenedPoints);
			setPolyComplete(true);
		}
	}, [boundaryItems]);
	
	useEffect(() => {
		setFlattenedPoints(
			points
			.concat(isPolyComplete ? [] : position)
			.reduce((a, b) => a.concat(b), [])
		);
	}, [points, isPolyComplete, position]);

	const renderBoundaryComponent = () => {		
		const handlePointDragMove = (e) => {
			const pos = getMousePos(e, layerRef);
			const stage = e.target.getStage();
			const index = e.target.index - 1;
			if (pos[0] < 0) pos[0] = 0;
			if (pos[1] < 0) pos[1] = 0;
			if (pos[0] > stage.width()) pos[0] = stage.width();
			if (pos[1] > stage.height()) pos[1] = stage.height();
			setPoints([...points.slice(0, index), pos, ...points.slice(index + 1)]);
		};
	
		const handleGroupDragEnd = (e) => {
			//drag end listens other children circles' drag end event
			//...that's, why 'name' attr is added, see in polygon annotation part
			if (e.target.name() === "polygon") {
				let result = [];
				let copyPoints = [...points];
				copyPoints.map((point) =>
				result.push([point[0] + e.target.x(), point[1] + e.target.y()])
				);
				e.target.position({ x: 0, y: 0 }); //needs for mouse position otherwise when click undo you will see that mouse click position is not normal:)
				setPoints(result);
			}
		};
	
		const handleMouseOverStartPoint = (e) => {
			if (isPolyComplete || points.length < 3) return;
			e.target.scale({ x: 3, y: 3 });
			setMouseOverPoint(true);
		};
	
		const handleMouseOutStartPoint = (e) => {
			e.target.scale({ x: 1, y: 1 });
			setMouseOverPoint(false);
		};

		return (
			<PolygonAnnotation
				points={points}
				flattenedPoints={flattenedPoints}
				handlePointDragMove={handlePointDragMove}
				handleGroupDragEnd={handleGroupDragEnd}
				handleMouseOverStartPoint={handleMouseOverStartPoint}
				handleMouseOutStartPoint={handleMouseOutStartPoint}
				isEnabled={isEditBoundary}
				isFinished={isPolyComplete}/>
		);
	}

	return (
		<Layer 
			ref={layerRef}
            onMouseMove={handleMouseMove}
            onMouseDown={handleMouseDown}>
			<Rect {...rectConfig}></Rect>
			{viewMode === ViewModes.Design && <GridLines width={width} height={height} />}
			{renderBoundaryComponent()}
			<BlockList />
			<WorkspaceList />
			<Transformer ref={trRef} resizeEnabled={false} />
			<ShapeCounter></ShapeCounter>
		</Layer>
	);
});
export default SubLayer;