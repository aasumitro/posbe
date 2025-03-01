import { useContext, useState } from "react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import { observer } from "mobx-react";

const DesignerMenu = observer(({onClose}) => {
    const context = useContext(ShapeContext);
	const { activeWorkspace, setWorkspaces, workspaces } = context;

    const [text, setText] = useState('');
    
    const handleInputChange = (e) => {
        setText(e.target.value);
    };

    const handleOk = (e) => {
        e.preventDefault();
        if (text.trim() === '') return;

        const currentItem = workspaces.find((s: any) => s.id === activeWorkspace.id) as any;
        currentItem.floorLabel = text;
        const newItems = workspaces.filter((s: any) => s.id !== activeWorkspace.id).concat([currentItem]);
        setWorkspaces(newItems);
        setText('')
        onClose();
    }

	return (
        <>
            <div className="modal-header">
                <b className="modal-title">Edit Workspace</b>
                <button onClick={onClose} type="button" className="btn btn-outline close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div className="modal-body py-3">
                <input type="text" value={text} onChange={handleInputChange} />
            </div>
            <div className="modal-footer">
                <button onClick={handleOk} type="button" className="btn btn-outline">OK</button>
            </div>
        </>
	);
});
export default DesignerMenu;