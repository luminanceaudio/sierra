import styled from 'styled-components';
import React from 'react';

export const StyledClickable = styled.button<{
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
}>`
  padding: inherit;
  &:hover {
    cursor: ${({ onClick }) => (onClick ? 'pointer' : 'inherit')};
    opacity: ${({ onClick }) => (onClick ? 0.68 : 1)};
    transition: opacity 0.18s;
  }

  &[disabled] {
    opacity: 0.5;
    cursor: default;
  }
`;
