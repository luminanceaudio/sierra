import { CSSProperties, forwardRef, KeyboardEvent } from 'react';
import Textarea from '../Textarea/Textarea';
import './InputListItem.css';

export type InputListItemProps = {
  value: string;
  onChange: (value: string) => void;
  onEnter?: (value: string) => void;
  onPrevious?: () => void;
  onNext?: () => void;
  onRight?: () => void;
  onDelete?: () => void;
  onBlur?: () => void;
  getRef: () => any;
  onFocus?: () => void;
  style?: CSSProperties;
};

const InputListItem = forwardRef(function InputListItem(
  {
    value,
    onChange,
    onEnter,
    onPrevious,
    onNext,
    onRight,
    onDelete,
    onBlur,
    getRef,
    onFocus,
    style,
  }: InputListItemProps,
  ref: any,
) {
  function tab(event: KeyboardEvent<HTMLTextAreaElement>) {
    const isTab =
      event.key === 'Tab' && !event.altKey && !event.ctrlKey && !event.metaKey;
    if (!isTab) {
      return;
    }

    const { selectionStart: eSelectionStart, selectionEnd: eSelectionEnd } =
      event.target as any;

    // Get start row
    let startRowValueIndex = value.slice(0, eSelectionStart).lastIndexOf('\n');
    if (startRowValueIndex === -1) {
      startRowValueIndex = 0;
    }

    // Get end row
    let endRowValueIndex = value.slice(0, eSelectionEnd).lastIndexOf('\n');
    if (endRowValueIndex === -1) {
      endRowValueIndex = 0;
    }

    if (!event.shiftKey) {
      // Indent
      if (startRowValueIndex === 0) {
        startRowValueIndex = -1;
      }

      const newValue = `${value.slice(
        0,
        startRowValueIndex + 1,
      )}\t${value.slice(startRowValueIndex + 1, value.length)}`;
      onChange(newValue);
    } else {
      // Undent
      if (startRowValueIndex === 0) {
        startRowValueIndex = -1;
      }

      if (value.at(startRowValueIndex + 1) === '\t') {
        const newValue = `${value.slice(
          0,
          startRowValueIndex + 1,
        )}${value.slice(startRowValueIndex + 2, value.length)}`;
        onChange(newValue);
      }
    }

    event.preventDefault();
  }

  function commandEnter(event: KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key !== 'Enter' || !event.altKey || !event.metaKey) {
      return;
    }

    onEnter?.(value);
    event.preventDefault();
  }

  function commandUp(event: KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key !== 'ArrowUp' || !event.altKey || !event.metaKey) {
      return;
    }

    // const { selectionStart: eSelectionStart } = event.target as any;

    // const firstLinebreakIndex = value.indexOf('\n');

    // if (firstLinebreakIndex === -1 || eSelectionStart <= firstLinebreakIndex) {
    onPrevious?.();
    event.preventDefault();
    // }
  }

  function commandRight(event: KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key !== 'ArrowRight' || !event.altKey || !event.metaKey) {
      return;
    }

    // const { selectionStart: eSelectionStart } = event.target as any;

    // const firstLinebreakIndex = value.indexOf('\n');

    // if (firstLinebreakIndex === -1 || eSelectionStart <= firstLinebreakIndex) {
    onRight?.();
    event.preventDefault();
    // }
  }

  function commandDown(event: KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key !== 'ArrowDown' || !event.altKey || !event.metaKey) {
      return;
    }

    // const { selectionStart: eSelectionStart } = event.target as any;

    // if (eSelectionStart >= value.lastIndexOf('\n')) {
    onNext?.();
    event.preventDefault();
    // }
  }

  function commandBackspace(event: KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key !== 'Backspace' || !event.altKey || !event.metaKey) {
      return;
    }

    // const { selectionStart: eSelectionStart } = event.target as any;

    // if (eSelectionStart >= value.lastIndexOf('\n')) {
    onDelete?.();
    event.preventDefault();
    // }
  }

  const keydown = (event: KeyboardEvent<HTMLTextAreaElement>) => {
    tab(event);
    commandEnter(event);
    commandUp(event);
    commandDown(event);
    commandBackspace(event);
    commandRight(event);
  };

  return (
    <div
      style={{
        display: 'flex',
        alignItems: 'start',
      }}
    >
      <input type="checkbox" style={{ marginTop: 10, zoom: 1.2 }} />

      <Textarea
        ref={ref}
        getRef={getRef}
        value={value}
        onChange={onChange}
        style={{
          flex: 1,
          backgroundColor: 'transparent',
          color: '#1c1c1c',
          marginLeft: 5,
          borderBottom: '1px solid #e1e1e1',
          paddingTop: 12,
          paddingBottom: 12,
          maxHeight: 200,
          overflow: 'scroll',
          ...style,
        }}
        onKeyDown={keydown}
        onBlur={onBlur}
        onFocus={onFocus}
        className="inputlistitem"
      />
    </div>
  );
});

InputListItem.defaultProps = {
  onEnter: () => {},
  onPrevious: () => {},
  onNext: () => {},
  onDelete: () => {},
  onBlur: () => {},
  onFocus: () => {},
  onRight: () => {},
  style: {},
};

export default InputListItem;
