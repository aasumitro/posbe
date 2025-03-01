import React, { useState, useEffect } from "react";
import { Line } from "react-konva";

const GridLines = (props) => {
    const {width, height} = props;
    const [gridLines, setGridLines] = useState([]);
    const blockSnapSize = 30;

    const addGridLines = () => {
        const padding = blockSnapSize;
        const lines = [];
        for (let i = 0; i < width / padding; i++) {
            const line = {
                points: [Math.round(i * padding) + 0.5, 0, Math.round(i * padding) + 0.5, height],
                stroke: '#ddd',
                strokeWidth: 1,
            };
            lines.push(line);
        }

        for (let j = 0; j < height / padding; j++) {
            const line = {
                points: [0, Math.round(j * padding), width, Math.round(j * padding)],
                stroke: '#ddd',
                strokeWidth: 0.5,
            };
            lines.push(line);
        }
        
        setGridLines(lines);
    }

    useEffect(() => {
        addGridLines();
    }, [width, height])

    return (
        <>
            {gridLines.map((line, i) => {
                return (
                    <Line key={i} {...line}></Line>
                );
            })}
        </>
    );
}

export default GridLines;