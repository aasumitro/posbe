import React from "react";
import { Layer, Rect } from "react-konva";

const MainLayer = (props) => {
    const {height, width} = props;
    const rectConfig = {
		fill: 'white',
		x: 0,
		y: 0,
		stroke: '#ddd',
		strokeWidth: 0.5,
		width: width,
		height: height
	}

	return (
        <Layer>
            <Rect {...rectConfig}></Rect>
        </Layer>
	);
};
export default MainLayer;