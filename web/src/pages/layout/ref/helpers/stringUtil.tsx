export const cleanJsonData = (val: string) => {
    return val.replace(/\\"/g, '"');
}