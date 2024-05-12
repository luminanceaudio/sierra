import styled from 'styled-components';
import { Radio } from '@headlessui/react';

export const StyledRadioItem = styled(Radio)`
  display: flex;
  padding: 20px 10px;
  border: 1px solid #e1e1e1;
  border-radius: 7px;
  font-size: 14px;
  color: #454545;

  justify-content: space-between;
  align-items: center;

  .radio-icon {
    opacity: 0;
  }

  &[data-checked] .radio-icon {
    opacity: 1;
    transition: opacity 0.25s;
  }

  &[data-checked] {
    font-weight: 600;
  }

  &[data-disabled] {
    background-color: #e1e1e1;
  }
`;

export const StyledRadioIcon = styled.div`
  display: flex;
  flex-direction: row;
  width: 20px;
`;
