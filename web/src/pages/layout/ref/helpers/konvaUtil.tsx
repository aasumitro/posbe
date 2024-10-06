import { WorkspaceColors } from "../models/constants/itemTypes";
import { ViewModes } from "../models/enums/viewModes";

export const zoomStageByWheel = (stageRef: any, scaleBy: number, event: any) => {
    if (stageRef.current === null) return;

    const stage = stageRef.current;
    const oldScale = stage.scaleX();
    const { x: pointerX, y: pointerY } = stage.getPointerPosition();
    const mousePointTo = {
        x: (pointerX - stage.x()) / oldScale,
        y: (pointerY - stage.y()) / oldScale,
    };
    const newScale = event.evt.deltaY > 0 ? oldScale * scaleBy : oldScale / scaleBy;

    if (newScale < 0.4) return;

    stage.scale({ x: newScale, y: newScale });
    const newPos = {
        x: pointerX - mousePointTo.x * newScale,
        y: pointerY - mousePointTo.y * newScale,
    }
    stage.position(newPos);
    stage.batchDraw();
}

export const getMousePos = (e: any, layerRef: any) => {
    const transform = layerRef.current.getAbsoluteTransform().copy();
    // to detect relative position we need to invert transform
    transform.invert();
    // now we find relative point
    const pos = e.target.getStage().getPointerPosition();
    const relativePos = transform.point(pos);

    return [relativePos.x, relativePos.y];
};

export const generateId = (): string => {
    const uniqueId = Date.now().toString(36) + Math.random().toString(36).substring(2);
    return uniqueId;
}

export const getWorkspaceColor = (viewMode: ViewModes, isOccupied: boolean): string => {
    if (viewMode === ViewModes.Design) return WorkspaceColors.LightGrey;

    return isOccupied ? WorkspaceColors.LightRed : WorkspaceColors.LightGreen;
}