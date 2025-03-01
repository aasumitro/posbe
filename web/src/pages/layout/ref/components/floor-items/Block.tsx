import { observer } from "mobx-react";
import React, { useContext, useEffect, useState } from "react";
import { Rect, Transformer } from "react-konva";
import { ShapeContext } from "@/pages/layout/ref/stores/shapeStore.tsx"
import { ViewModes } from "../../models/enums/viewModes";

const Block = observer(({ shapeProps, onChange }) => {
  const shapeRef = React.useRef();
	const trRef = React.useRef<any>();
  const context = useContext(ShapeContext);
	const { viewMode, selectedShapeId, multipleSelections, onSelectShape } = context;
  const [ isMultiSelected, setIsMultiSelected ] = useState(false);

	useEffect(() => {
		if (selectedShapeId === shapeProps.id && !isMultiSelected) {
			//attaching transformer manually
			trRef.current.setNode(shapeRef.current);
			trRef.current.getLayer().batchDraw();
		}
	}, [selectedShapeId]);

  useEffect(() => {
    setIsMultiSelected(multipleSelections.indexOf(shapeProps.id) !== -1);
  }, [multipleSelections])

  return(
    <>
      <Rect
          onClick={(e) => { onSelectShape(e, shapeProps.id) } }
          ref={shapeRef}
          {...shapeProps}
          stroke={isMultiSelected ? 'blue' : 'black'}
          strokeWidth={isMultiSelected ? 1 : 0.5}
          draggable={viewMode === ViewModes.Design}
          onDragEnd={e => {
            onChange({
              ...shapeProps,
              x: e.target.x(),
              y: e.target.y(),
            });
          }}
          onTransformEnd={e => {
            //here the transform will change the scale
            const node = shapeRef.current as any;
            const scaleX = node.scaleX();
            const scaleY = node.scaleY();
            node.scaleX(1);
            node.scaleY(1);
            onChange({
              ...shapeProps,
              x: node.x(),
              y: node.y(),
              width: node.width() * scaleX,
              height: node.height() * scaleY,
            });
          }}
        />
      {selectedShapeId === shapeProps.id && !isMultiSelected && <Transformer ref={trRef} />}
    </>
  )
});
export default Block;