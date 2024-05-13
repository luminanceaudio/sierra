import styled from 'styled-components';

export const StyledClickable = styled.button<{ onClick?: () => void }>`
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
