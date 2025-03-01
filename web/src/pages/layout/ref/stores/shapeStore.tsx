import { makeAutoObservable } from "mobx";
import { createContext } from "react";
import { IdPrefix } from "../models/constants/itemTypes";
import { ViewModes } from "../models/enums/viewModes";
import { IActiveWorkspace } from "../models/activeWorkspace";
import { IBookWorkspace } from "../models/bookWorkspace";
import { generateId } from "../helpers/konvaUtil";
import { cleanJsonData } from "../helpers/stringUtil";
import { IOccupiedWorkspace } from "../models/occupiedWorkspace";
import { IPcfInputs } from "../models/pcfInputs";
import { staticData } from "./staticData";

class ShapeStore {
	blocks = [];
	workspaces = [];
	clonedItems = [];
	shapes = [];
	selectedShapeId = null;
	multipleSelections: string[] = [];
	activeWorkspace: IActiveWorkspace | null = null;
	floorLayout = '';
	occupiedWorkspaces: IOccupiedWorkspace[] = [];
	viewMode = ViewModes.Design;
	toBookWorkspace: IBookWorkspace | null = null;

	pasteOffset = 100;
	canvasDimensions = {width: 1200, height: 600};

	isEditBoundary = false;
	points = [];
	flattenedPoints = [];
	position = [];
	isPolyComplete = false;
	isMouseOverPoint = false;
	boundaryItems: any;

	constructor() {
		makeAutoObservable(this);
	}

	resetStore = () => {
		this.selectedShapeId = null;
		this.toBookWorkspace = null;
		this.multipleSelections = [];
		this.occupiedWorkspaces = [];
		this.blocks = [];
		this.workspaces = [];
		this.clonedItems = [];
		this.shapes = [];
	}

	setMouseOverPoint = (value: boolean) => {
		this.isMouseOverPoint = value;
	}

	setPolyComplete = (value: boolean) => {
		this.isPolyComplete = value;
	}

	setPoints = (value: any[]) => {
		this.points = value;
	}

	setFlattenedPoints = (value: any[]) => {
		this.flattenedPoints = value;
	}

	setPosition = (value: any[]) => {
		this.position = value;
	}

	notifyOutput: () => void;

	setIsEditBoundary = (val: boolean) => {
		this.isEditBoundary = val;
	}

	setBoundaryItems = (val: any) => {
		this.boundaryItems = val;
	}

	clearBoundaryItems = () => {
		this.boundaryItems = null;
		this.points = [];
		this.flattenedPoints = [];
		this.points = [];
		this.isPolyComplete = false;
	}

	setNotifyOutput = (fn: () => void) => {
		this.notifyOutput = fn;
	}

	initializePcfInputs = (pcfInputs: IPcfInputs) => {
		if (!pcfInputs) {
			pcfInputs = {
				floorLayout: staticData,
				viewMode: '1',
				occupiedWorkspaces: null
			}
		};

		const floorLayoutInput = pcfInputs.floorLayout;
		if (floorLayoutInput && floorLayoutInput != '' && floorLayoutInput != 'undefined') {
			this.setFloorLayout(floorLayoutInput);
		}

		const occupiedWorkspacesInput = pcfInputs.occupiedWorkspaces;
		if (occupiedWorkspacesInput && occupiedWorkspacesInput != '' && occupiedWorkspacesInput != 'undefined') {
			this.extractOccupiedWorkspaceFromJson(occupiedWorkspacesInput);
		}

		const viewModeInput = pcfInputs.viewMode;
		if (viewModeInput && viewModeInput != '' && viewModeInput != 'undefined') {
			const mode = viewModeInput === '1' ? ViewModes.Design : ViewModes.Book;
			this.setViewMode(mode);
		}
	}

	extractOccupiedWorkspaceFromJson = (json: string) => {
		try {
			const data = JSON.parse(cleanJsonData(json));
			this.setOccupiedWorkspaces(data);
		} catch {
			console.log('cannot extract ocuupiedWorkspaces from invalid JSON')
		}
	}

	setToBookWorkspace = (id: string, label: string, isBooked = false) => {
		this.toBookWorkspace = {
			id: id,
			label: label,
			isBooked: isBooked
		}
	}

	setViewMode = (mode: ViewModes) => {
		this.viewMode = mode;
	}

	loadLayout = () => {
		try {
			if (!this.floorLayout) return;
			const data = JSON.parse(this.floorLayout);
			this.loadShapes(data.floorItems, 'Rect', IdPrefix.Block, this.setBlocks);
			this.loadShapes(data.floorItems, 'Group', IdPrefix.Workspace, this.setWorkspaces);
			this.loadBoundaryItems(data.boundaryItems)
			this.setBoundaryItems(data.boundaryItems);
			this.setCanvasDimensions(data.canvasDimensions.width, data.canvasDimensions.height);
		} catch(e) {
			console.log('cannot load json layout', e);
		}

	}

	loadBoundaryItems = (boundaryItems: any) => {
		if (boundaryItems && boundaryItems.points.length > 0 || boundaryItems.flattenedPoints.length > 0) {
			this.setBoundaryItems(boundaryItems);
		} else {
			this.clearBoundaryItems();
		}
	}

	loadShapes = (
		floorItems: any[],
		className: string, 
		idPrefix: string, 
		setFn: (items: any[]) => void) => {

		const nodes = floorItems.filter(i => 
			i.className === className
			&& i.attrs.class 
			&& i.attrs.class === 'floor-item' 
			&& i.attrs.id 
			&& i.attrs.id.indexOf(idPrefix) !== -1);

		const stateList = [];
		nodes.forEach(item => {
			item.attrs.occupied = this.occupiedWorkspaces && this.occupiedWorkspaces.length > 0 && this.occupiedWorkspaces.filter(occ => item.attrs.id === occ.floorItemId).length > 0;
			stateList.push({...item.attrs});
		});
		setFn(stateList);
		this.addShapes(stateList);
	}

	setOccupiedWorkspaces = (items: IOccupiedWorkspace[]) => {
		this.occupiedWorkspaces = items;
	}

	removeOccupied = (id: string) => {
		this.occupiedWorkspaces = this.occupiedWorkspaces.filter(item => item.floorItemId !== id);
	}

	addOccupied = (id: string) => {
		const item = {
			floorItemId: id,
			bookedBy: 'You'
		} as IOccupiedWorkspace;
		this.occupiedWorkspaces = this.occupiedWorkspaces.concat([item]);
	}

	getCommaSeparatedItems = (commaSeparatedItems: string) => {
		if (!commaSeparatedItems || commaSeparatedItems.length === 0) return [];
		
		return commaSeparatedItems.replace(' ', '').split(",");
	}

	setFloorLayout = (val: string) => {
		this.floorLayout = val;
	}

	setCanvasDimensions = (width: number, height: number) => {
		this.canvasDimensions = {width: width, height: height};
	}

	setActiveWorkspace = (id: string, label: string) => {
		this.activeWorkspace = {
			id: id,
			label: label
		};
	}

	setMultipleSelections = (ids: string[]) => {
		this.multipleSelections = ids;
	}

	addShapeIdsToMultipleSelections = (ids: string[]) => {
		const uniqueIds = ids.filter(id => !this.multipleSelections.includes(id));
		this.multipleSelections = this.multipleSelections.concat(uniqueIds);
	}

	addToOccupiedWorkspaces = (items: IOccupiedWorkspace[]) => {
		const bookedIds = this.occupiedWorkspaces.map(o => o.floorItemId);
		const uniqueItems = items.filter(item => !bookedIds.includes(item.floorItemId));
		this.occupiedWorkspaces = this.occupiedWorkspaces.concat(uniqueItems);
	}

	resetMultipleSelections = () => {
		this.multipleSelections = [];
	}

	addShapes = (shapes: any[]) => {
		this.shapes = this.shapes.concat(shapes);
	}

	setShapes = (shapes: any[]) => {
		this.shapes = shapes;
	}

	setClonedItems = (items: any[]) => {
		this.clonedItems = items;
	}

    setBlocks = (items: any[]) => {
		this.blocks = items;
    };

	setWorkspaces = (items: any[]) => {
		this.workspaces = items;
    };

	addBlocks = (items: any[]) => {
		this.blocks = this.blocks.concat(items);
	}

	addWorkspaces = (items: any[]) => {
		this.workspaces = this.workspaces.concat(items);
	}

    addBlock = () => {
		const id = generateId();
		const item = {
			x: 50,
			y: 50,
			width: 40,
			height: 40,
			stroke: 'darkgray',
			strokeWidth: 1,
			class: 'floor-item',
			id: `${IdPrefix.Block}${id}`,
		};
		const items = this.blocks.concat([item]);
		this.setBlocks(items);
		return item;
	};

	addWorkspace = (type: string) => {
		const id = generateId();
		let item = {
			x: 50,
			y: 50,
			class: 'floor-item',
			workspaceType: type,
			floorLabel: type,
			id: `${IdPrefix.Workspace}${id}`,
		};

		const items = this.workspaces.concat([item]);
		this.setWorkspaces(items);
		return item;
	};

	cloneShape = (node: any, idx: number) => {
		// will need to change this if we have 2+ itemTypes
		const {attrs} = node;
		const name = attrs.id.indexOf(IdPrefix.Block) !== -1 ? IdPrefix.Block : IdPrefix.Workspace;
		const id = generateId();
		return {
			...attrs,
			x: attrs.x + this.pasteOffset,
			y: attrs.y + this.pasteOffset,
			id: `${name}${id}`
		}
	}

	distributeClones = (cloneList: any[]) => {
		// TODO make this reusable
		const clonedBlocks = cloneList.filter(b => b.id.indexOf(IdPrefix.Block) !== -1);
		this.addBlocks(clonedBlocks);
		const clonedWorkspaces = cloneList.filter(b => b.id.indexOf(IdPrefix.Workspace) !== -1);
		this.addWorkspaces(clonedWorkspaces);
	}

	setSelectedShapeId = (id: string | null) => {
		this.selectedShapeId = id;
	}

	deleteShapes = (shapeIds: string[]) => {
		this.blocks = this.blocks.filter(i => !shapeIds.includes(i.id));
		this.workspaces = this.workspaces.filter(i => !shapeIds.includes(i.id));
		this.shapes = this.shapes.filter(i => !shapeIds.includes(i.id));
	}

	onSelectShape = (e, shapeId) => {
		if (this.viewMode !== ViewModes.Design) return;
		if (e.evt.ctrlKey && !e.evt.key) {
			this.addPreviouslySelectedShape();
			this.addShapeIdsToMultipleSelections([shapeId]);
			this.setSelectedShapeId(null);
		} else {
			this.resetMultipleSelections();
			this.setSelectedShapeId(shapeId);
		}
	}

	addPreviouslySelectedShape = () => {
		// add previously selected shape not clicked using ctrl+click
		if (this.selectedShapeId && this.multipleSelections.indexOf(this.selectedShapeId) === -1) {
			this.addShapeIdsToMultipleSelections([this.selectedShapeId])
		}
	}

	onShapeChange = (i: number, existingItems: any[], newAttrs: any, callbackFn: (items: any[]) => void) => {
		const items = existingItems.slice() as any;
		items[i] = newAttrs;
		callbackFn(items);
	}

	onShapeDragStart = (shapeId: string) => {
		const notSingleSelected = this.selectedShapeId && shapeId !== this.selectedShapeId;
		const notMultiSelected = this.multipleSelections.length > 0 && !this.multipleSelections.includes(shapeId);

		if (notSingleSelected) {
			this.setSelectedShapeId(shapeId);
		} else if (notMultiSelected) {
			this.setSelectedShapeId(shapeId);
			this.resetMultipleSelections();
		}
	}

	setFdOutput = (val: string) => {
		const fdOutputEl = document.getElementById('fdOutput') as any;
		fdOutputEl.value = val;
	}
}

export const shapeStore = new ShapeStore();
export const ShapeContext = createContext(shapeStore);