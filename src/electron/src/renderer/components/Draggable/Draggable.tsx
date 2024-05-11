import React from 'react';
import { useDraggable } from '@dnd-kit/core';

export type DraggableProps = {
  children?: React.ReactNode;
};

function Draggable({ children }: DraggableProps) {
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: 'draggable',
  });
  const style = transform
    ? {
        transform: `translate3d(${transform.x}px, ${transform.y}px, 0)`,
      }
    : undefined;

  return (
    <button
      type="button"
      ref={setNodeRef}
      style={{ backgroundColor: 'orange', ...style }}
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...listeners}
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...attributes}
    >
      {children}
    </button>
  );
}

Draggable.defaultProps = {
  children: null,
};

export default Draggable;
