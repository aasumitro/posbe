import { observer } from "mobx-react";
import { useContext, useEffect, useState } from "react";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx";
import { Text } from "react-konva";

const ShapeCounter = observer((props) => { 
    const context = useContext(ShapeContext);
	const [selectedShapesCount, setSelectedShapesCount] = useState(0);
	const { multipleSelections, selectedShapeId } = context;

    const textProps = {
        x: 15,
        y: 20
    }

	useEffect(() => {
		if (multipleSelections.length > 0) {
			setSelectedShapesCount(multipleSelections.length);
			return;
		}
		if (selectedShapeId) {
			setSelectedShapesCount(1);
		} else {
			setSelectedShapesCount(0);
		}

	}, [multipleSelections, selectedShapeId])

    const renderCounter = () => {
        return <Text {...textProps} text={`(${selectedShapesCount}) selected`}></Text>;
    }

    return(selectedShapesCount === 0 ? <></> : renderCounter());
})

export default ShapeCounter