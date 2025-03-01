import { observer } from "mobx-react";
import { useContext, useState } from "react";
// import { Button, NavDropdown, Navbar } from "react-bootstrap";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import './CanvasSettings.css';
import { WorkspaceTypes } from "../../../models/constants/itemTypes";
import { ViewModes } from "../../../models/enums/viewModes";
import {Button} from "@/components/ui/button.tsx";

const CanvasSettings = observer((props) => {
    
    const { layerRef, stageRef } = props;
    const context = useContext(ShapeContext);
	const { setViewMode, viewMode, setFdOutput, clearBoundaryItems, addBlock, addShapes, addWorkspace, setCanvasDimensions, canvasDimensions, notifyOutput, isEditBoundary, setIsEditBoundary } = context;
    
	const [inputHeight, setInputHeight] = useState(canvasDimensions.height);
	const [inputWidth, setInputWidth] = useState(canvasDimensions.width);
    const [jsonLayout, setJsonLayout] = useState<string>("");
    
	const handleAddBlock = () => {
		const newBlock = addBlock();
		addShapes([newBlock]);
	}

	const handleAddWorkspace = (type: string) => {
		const newWorkspace = addWorkspace(type);
		addShapes([newWorkspace]);
	}
    
	const handleHeightChange = (event) => {
		setInputHeight(event.target.value)
	}
	
	const handleWidthChange = (event) => {
		setInputWidth(event.target.value)
	}

	const onChangeDimension = () => {
		setCanvasDimensions(inputWidth, inputHeight);
	}

    const saveLayout = () => {
		const konvaData = getKonvaData();
		const jsonData = JSON.stringify(konvaData);
		console.log('jsonData', jsonData);
		setFdOutput(jsonData);
        notifyOutput();
		alert('Layout saved!')
	}

	const getKonvaData = () => {
		const floorItems = getCanvasFloorItems();
		const boundaryItems = getBoundaryItems();

		return {
			floorItems: floorItems,
			boundaryItems: boundaryItems,
			canvasDimensions: canvasDimensions
		};
	}

	const getCanvasFloorItems = () => {
		const layer = layerRef.current as any;
		const data = JSON.parse(layer.toJSON());
		return data.children.filter(item => item.attrs && item.attrs.class === 'floor-item');
	}

	const getBoundaryItems = () => {
		const layer = layerRef.current as any;
		const data = JSON.parse(layer.toJSON());
		const polygon = data.children.filter(item => item.attrs && item.attrs.name === 'polygon')[0];
		const flattenedPoints = polygon.children.filter(item => item.className === 'Line')[0].attrs.points;
		const points = polygon.children.filter(item => item.className === 'Circle').map(point => {
			return [point.attrs.x, point.attrs.y]
		});

		return {
			points: points,
			flattenedPoints: flattenedPoints
		};
	}

	const resetView = () => {
		if (!stageRef.current) return;

		const origScale = 1;
		const stage = stageRef.current;
		const pos = {
			x: origScale,
			y: origScale,
		}

		stage.scale({ x: origScale, y: origScale });
		stage.position(pos);
		stage.batchDraw();
	}

	const handleIsEditBoundaryChange = () => {
		setIsEditBoundary(!isEditBoundary);
	}

	const removeBoundary = () => {
		clearBoundaryItems();
	}

	const onRadBookingChange = () => {
		setViewMode(ViewModes.Book);
	}

	const onRadDesignerChange = () => {
		setViewMode(ViewModes.Design);
	}

    return (
        <>
			<input value={jsonLayout} hidden id="txtJsonLayoutOutput"/>
            <div id="canvasNavBar" className="m-3 d-flex justify-content-center">
				<div title="App Mode" id="userMode">
					<div className="mode-wrapper">
						<div className="form-check">
							<input onChange={onRadDesignerChange} checked={viewMode === ViewModes.Design}
								className="form-check-input" type="radio" name="radMode" id="radDesigner" />
							<label className="form-check-label">
								Designer Mode
							</label>
						</div>
						<div className="form-check">
							<input onChange={onRadBookingChange} checked={viewMode === ViewModes.Book}
								className="form-check-input" type="radio" name="radMode" id="radBooking" />
							<label className="form-check-label">
								Booking Mode
							</label>
						</div>
					</div>
				</div>
				{
					viewMode === ViewModes.Design &&
					<div title="Canvas Settings" id="canvasSettings">
						<div className="p-3 w-100">
							<label>Add Object:</label>
							<div className='d-flex'>
								<Button variant="secondary" onClick={handleAddBlock}>
									Block
								</Button>
								<Button variant="secondary" onClick={() => handleAddWorkspace(WorkspaceTypes.Seat)}>
									Seat
								</Button>
								<Button variant="secondary" onClick={() => handleAddWorkspace(WorkspaceTypes.Room)}>
									Room
								</Button>
							</div>
						</div>
						<div className="p-3 w-100">
							<label>Canvas Dimensions:</label>
							<div className="d-flex mb-3">
								<div className="d-flex flex-column">
									<label>Width</label>
									<input defaultValue={canvasDimensions.width} type="text" onChange={handleWidthChange}/>
								</div>
								<div className="d-flex flex-column">
									<label>Height</label>
									<input defaultValue={canvasDimensions.height} type="text" onChange={handleHeightChange}/>
								</div>
							</div>
							<Button variant="secondary" onClick={onChangeDimension}>Apply Dimensions</Button>
						</div>
						<div className="p-3 w-100">
							<label>Actions:</label>
							<div className='d-flex'>
								<Button variant="secondary" onClick={saveLayout}>
									Save Layout
								</Button>
								<Button variant="secondary" onClick={resetView}>
									Reset View
								</Button>
							</div>
						</div>
						<div className="d-flex">
							<div className="d-flex p-3">
								<div>
									<input
									type="checkbox"
									checked={isEditBoundary}
									onChange={handleIsEditBoundaryChange}/>
								</div>
								<div>Edit Floor Boundary</div>
							</div>
							<div className='d-flex'>
								<Button variant="link" onClick={removeBoundary}>
									Remove Boundary
								</Button>
							</div>
						</div>
					</div>
				}
			</div>
        </>
    );
})

export default CanvasSettings;