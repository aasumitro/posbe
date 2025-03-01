import React, { useContext } from "react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import { observer } from "mobx-react";
import Workspace from "./Workspace";

const WorkspaceList = observer(() => {
    const context = useContext(ShapeContext);
	const { workspaces, setWorkspaces, onShapeChange } = context;
    
    return (
        <>
            {workspaces.map((item, i) => {
                return(
                    <Workspace
                        key={i}
                        shapeProps={item}
                        onChange={newAttrs => {
                            onShapeChange(i, workspaces, newAttrs, setWorkspaces);
                        }}
                    />
                );
            })}
        </>
    );
});

export default WorkspaceList;