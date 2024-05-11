import React, { CSSProperties } from 'react';
import { useDroppable } from '@dnd-kit/core';

export type DroppableProps = {
  id: string;
  children?: React.ReactNode;
  style?: CSSProperties;
};

function Droppable({ id, children, style }: DroppableProps) {
  const { isOver, setNodeRef } = useDroppable({
    id,
  });
  const extraStyle = {
    color: isOver ? 'green' : undefined,
  };

  return (
    <div
      ref={setNodeRef}
      style={{ backgroundColor: 'orange', flex: 1, ...style, ...extraStyle }}
    >
      {children}
    </div>
  );
}

Droppable.defaultProps = {
  children: null,
  style: {},
};

export default Droppable;
