function DragHeader() {
  return (
    <div
      style={{
        '-webkit-app-region': 'drag',
        height: 28,
        position: 'absolute',
        width: '100%',
      }}
    />
  );
}

export default DragHeader;
