import styled from 'styled-components';

export const StyledAudioPlayerBar = styled.div`
  display: flex;

  padding: 20px;
  box-sizing: border-box;
  height: 60px;
  border-top: 1px solid #e1e1e1;
  align-items: center;
  position: fixed;
  background-color: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(8px);
  z-index: 1;
  width: 100%;
  bottom: 0;
`;
