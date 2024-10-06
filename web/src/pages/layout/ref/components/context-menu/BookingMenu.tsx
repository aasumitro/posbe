import { useContext, useEffect, useState } from "react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import { BookActions } from "../../models/constants/itemTypes";
import { IOccupiedWorkspace } from "../../models/occupiedWorkspace";
import { observer } from "mobx-react";

const BookingMenu = observer(({onClose}) => {
    const context = useContext(ShapeContext);
	const { setFdOutput, activeWorkspace, addOccupied, removeOccupied, notifyOutput, occupiedWorkspaces, setToBookWorkspace } = context;
    const [ actionLabel, setActionLabel ] = useState('');
    const [ occupiedDetails, setOccupiedDetails ] = useState<IOccupiedWorkspace>(null);

    const handleOk = (e) => {
        e.preventDefault();

        const isBookAction = actionLabel === BookActions.Book;
        if (isBookAction) {
            addOccupied(activeWorkspace.id);
        } else {
            removeOccupied(activeWorkspace.id);
        }

        setToBookWorkspace(activeWorkspace.id, activeWorkspace.label, isBookAction);
        setFdOutput(`${actionLabel}, ${activeWorkspace.id}`);
        notifyOutput();

        onClose();
    }

    useEffect(() => {
        if (!activeWorkspace) return;
        const occupiedSpace = occupiedWorkspaces.find(o => o.floorItemId === activeWorkspace.id);
        setOccupiedDetails(occupiedSpace);
        const label = occupiedSpace ? BookActions.Unbook : BookActions.Book;
        setActionLabel(label);
    }, [occupiedWorkspaces, activeWorkspace]);

    const renderOccupiedDetails = () => {
        return (
            <>
                <div className="py-3 d-flex flex-column align-items-start">
                    <div>Currently booked by:</div>
                    <div className="text-primary">{occupiedDetails.bookedBy}</div>
                </div>
            </>
        );
    }

    const renderBody = () => {
        return(
            <>
                { occupiedDetails && renderOccupiedDetails() }
                <div className="modal-body booking py-3">{actionLabel} workspace <b>{activeWorkspace.label}</b>?</div>
            </>
        );
    }

	return (
        <>
            <div className="modal-header">
                <b className="modal-title">Confirm Action</b>
                <button onClick={onClose} type="button" className="btn btn-outline close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            {activeWorkspace && renderBody()}
            <div className="modal-footer">
                <button onClick={handleOk} type="button" className="btn btn-outline">OK</button>
            </div>
        </>
	);
});
export default BookingMenu;