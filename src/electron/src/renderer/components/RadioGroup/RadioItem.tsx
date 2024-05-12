import React, { CSSProperties } from 'react';
import { CheckCircleIcon } from '@heroicons/react/24/solid';
import { StyledRadioIcon, StyledRadioItem } from './RadioItem.style';
import Clickable from '../Clickable/Clickable';

type RadioItemProps = {
  title: string;
  value: string;
  style?: CSSProperties;
  disabled?: boolean;
};

function RadioItem({ title, value, style, disabled }: RadioItemProps) {
  return (
    <Clickable onClick={() => {}} disabled={disabled}>
      <StyledRadioItem style={style} value={value} disabled={disabled}>
        {title}
        <StyledRadioIcon>
          <CheckCircleIcon className="radio-icon" />
        </StyledRadioIcon>
      </StyledRadioItem>
    </Clickable>
  );
}

RadioItem.defaultProps = {
  style: {},
  disabled: false,
};

export default RadioItem;
