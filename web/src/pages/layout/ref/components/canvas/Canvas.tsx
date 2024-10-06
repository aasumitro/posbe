import React, { useContext, useEffect } from "react";
import { Stage } from "react-konva";
import { observer } from "mobx-react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import SubLayer from "../SubLayer";
import MainLayer from "../MainLayer";
import './Canvas.css';
import CanvasSettings from "./canvasSettings/CanvasSettings";
// import { ViewModes } from "@/pages/layout/ref/models/enums/viewModes.tsx";
import { zoomStageByWheel } from "../../helpers/konvaUtil";
import ContextMenu from "../context-menu/ContextMenu";

const Canvas = observer((props) => {
	const { notifyOutput, pcfInputs } = props;
    const context = useContext(ShapeContext);
	const { initializePcfInputs, extractOccupiedWorkspaceFromJson, setFloorLayout, resetStore, setSelectedShapeId, resetMultipleSelections, canvasDimensions, viewMode, setViewMode,
		occupiedWorkspaces, floorLayout, loadLayout, setNotifyOutput, toBookWorkspace, setActiveWorkspace } = context;
	const stageRef = React.useRef<any>();
	const layerRef = React.useRef<any>();
	const contextMenuRef = React.useRef<any>();
	const scaleBy = 1.10;

	setNotifyOutput(notifyOutput);

	useEffect(() => {
		// don't execute if there's a newly booked workspace
		if (toBookWorkspace && toBookWorkspace.isBooked) {
			return;
		}

		loadLayout();
	}, [occupiedWorkspaces, floorLayout]);

	useEffect(() => {
		resetStore();
		initializePcfInputs(pcfInputs);
	}, [])

	const handleZoomWheel = (event) => {
		event.evt.preventDefault();
		zoomStageByWheel(stageRef, scaleBy, event);
	}

	const deselectShapes = (e) => {
		const el = e.target;

		const isTransformer = el.attrs.name && el.attrs.name.indexOf('_anchor') !== -1;
		const isFloorChildItem = el.parent && el.parent.attrs.class && el.parent.attrs.class === 'floor-item';
		const isFloorItem = el.attrs.class && el.attrs.class === 'floor-item';

		if(!isFloorChildItem && !isFloorItem && !isTransformer) {
			setSelectedShapeId(null);
			resetMultipleSelections();
		}
	}

	const showContextMenu = (e: any) => {
		e.evt.preventDefault();
		const parent = e.target.parent;
		if (!parent || !parent.attrs || !parent.attrs.class || parent.attrs.class.indexOf('floor-item') < 0) return;
		
		setActiveWorkspace(parent.attrs.id, parent.attrs.floorLabel);
        const menuNode = contextMenuRef.current;
		const stage = stageRef.current;
		const containerRect = stage.container().getBoundingClientRect();
		const top = containerRect.top + stage.getPointerPosition().y - 100 + 'px';
		const left = containerRect.left + stage.getPointerPosition().x - 350 + 'px';
		menuNode.style.top = top;
		menuNode.style.left = left;
		menuNode.style.display = 'initial';
	}
	
	return (
		<div className="fd-canvas">
			<input hidden id="fdOutput" />
			<CanvasSettings stageRef={stageRef} layerRef={layerRef}></CanvasSettings>
			<div className="d-flex justify-content-center stage-wrapper">
				<input value={viewMode} hidden id="txtViewMode" />
				<Stage
					draggable={true}
					onWheel={handleZoomWheel}
					width={canvasDimensions.width}
					height={canvasDimensions.height}
					ref={stageRef}
					onContextMenu={(e) => {
						showContextMenu(e);
					}}
					onMouseDown={deselectShapes}>
					<MainLayer {...canvasDimensions}></MainLayer>
					<SubLayer {...canvasDimensions} layerRef={layerRef} />
				</Stage>
			</div>
			<ContextMenu menuRef={contextMenuRef} />
		</div>
	);
});

export default Canvas;


