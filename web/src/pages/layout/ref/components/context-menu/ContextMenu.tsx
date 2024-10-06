import { useContext } from "react";
import './ContextMenu.css';
import DesignerMenu from "./DesignerMenu";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import { ViewModes } from "../../models/enums/viewModes";
import BookingMenu from "./BookingMenu";
import { observer } from "mobx-react";

const ContextMenu = observer(({menuRef}) => {
    const context = useContext(ShapeContext);
	const { viewMode } = context;

    const onClose = () => {
        const div = menuRef.current;
        div.style.display = 'none';
    }

    const renderBody = () => {
        return viewMode === ViewModes.Design ? <DesignerMenu onClose={onClose} /> : <BookingMenu onClose={onClose} />
    }
    
	return (
        <div ref={menuRef} id="contextMenu">
            <div className="modal-content p-3">
                {renderBody()}
            </div>
        </div>
	);
});
export default ContextMenu;