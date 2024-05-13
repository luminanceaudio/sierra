import styled from 'styled-components';
import TransitionOnLoad from '../TransitionOnLoad/TransitionOnLoad';

export const StyledToast = styled(TransitionOnLoad)`
  display: flex;
  gap: 5px;
  align-items: center;

  z-index: 100;
  position: fixed;
  top: 0;
  right: 0;
  width: 300px;
  margin: 30px;
  border-radius: 7px;
  border: 1px solid #f1f1f1;
  padding: 15px 20px;
  font-size: 14px;
  background-color: #f1f1f1;
`;
