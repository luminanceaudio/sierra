import styled from 'styled-components';

export const StyledProgress = styled.div`
  position: absolute;
  height: 100%;
  background-color: #f1f1f1;
  z-index: -1;

  animation: my-animation 6s infinite;
  animation-timing-function: linear;

  @keyframes my-animation {
    from {
      width: 0;
    }
    to {
      width: 100%;
    }
  }
`;
