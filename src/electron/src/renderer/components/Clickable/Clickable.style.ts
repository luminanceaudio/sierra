import styled from 'styled-components';

export const StyledClickable = styled.button`
  &:hover {
    opacity: 0.68;
    cursor: pointer;
    transition: opacity 0.18s;
  }

  &[disabled] {
    opacity: 0.5;
    cursor: default;
  }
`;
