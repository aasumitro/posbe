import Canvas from "@/pages/layout/ref/components/canvas/Canvas.tsx";

export const LayoutBlueprintPage = () => {
  const notifyOutput = () => {
    console.log('notify output');
  }

  return (
    <Canvas notifyOutput={notifyOutput} />
  )
}