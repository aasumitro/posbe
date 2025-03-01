import React, { useState } from "react";
import { Line, Circle, Group } from "react-konva";
import { dragBoundFunc, minMax } from "../../helpers/boundaryUtil";
/**
 *
 * @param {minMaxX} props
 * minMaxX[0]=>minX
 * minMaxX[1]=>maxX
 *
 */
const PolygonAnnotation = (props) => {
  const {
    points,
    flattenedPoints,
    isEnabled,
    isFinished,
    handlePointDragMove,
    handleGroupDragEnd,
    handleMouseOverStartPoint,
    handleMouseOutStartPoint,
  } = props;
  
  const vertexRadius = 2;

  const [stage, setStage] = useState();
  const handleGroupMouseOver = (e) => {
    if (!isFinished || !isEnabled) return;
    e.target.getStage().container().style.cursor = "move";
    setStage(e.target.getStage());
  };
  const handleGroupMouseOut = (e) => {
    if (!isEnabled) return;
    e.target.getStage().container().style.cursor = "default";
  };
  const [minMaxX, setMinMaxX] = useState([0, 0]); //min and max in x axis
  const [minMaxY, setMinMaxY] = useState([0, 0]); //min and max in y axis
  const handleGroupDragStart = (e) => {
    let arrX = points.map((p) => p[0]);
    let arrY = points.map((p) => p[1]);
    setMinMaxX(minMax(arrX));
    setMinMaxY(minMax(arrY));
  };
  
  return (
    <Group
      name="polygon"
      draggable={isEnabled && isFinished}
      onDragStart={handleGroupDragStart}
      onDragEnd={handleGroupDragEnd}
      onMouseOver={handleGroupMouseOver}
      onMouseOut={handleGroupMouseOut}
    >
      <Line
        points={flattenedPoints}
        stroke="darkgray"
        dash = {[5, 5]}
        strokeWidth={3}
        closed={isFinished}
      />
      {points.map((point, index) => {
        const x = point[0] - vertexRadius / 2;
        const y = point[1] - vertexRadius / 2;
        const startPointAttr =
          index === 0
            ? {
                hitStrokeWidth: 12,
                onMouseOver: handleMouseOverStartPoint,
                onMouseOut: handleMouseOutStartPoint,
              }
            : null;
        return (
          <Circle
            key={index}
            x={x}
            y={y}
            radius={vertexRadius}
            fill="black"
            stroke="black"
            strokeWidth={2}
            draggable={isEnabled}
            onDragMove={handlePointDragMove}
            dragBoundFunc={(pos) =>
              // @ts-ignore
              dragBoundFunc(stage.width(), stage.height(), vertexRadius, pos)
            }
            {...startPointAttr}
          />
        );
      })}
    </Group>
  );
};

export default PolygonAnnotation;
